package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	validator "gopkg.in/go-playground/validator.v9"
	"microservice/domain/entitie"
	_interface "microservice/domain/interface"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TagHandler struct {
	TagService _interface.TagService
}

func NewTagHandler(echoInstance *echo.Echo, service _interface.TagService) {
	handler := &TagHandler{
		TagService: service,
	}
	echoInstance.PUT("/tag", handler.Update)
	echoInstance.POST("/tag", handler.Store)
	echoInstance.GET("/tag/:id", handler.GetByID)

}

func (a *TagHandler) GetByID(echoContext echo.Context) error {

	idP, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		return echoContext.JSON(http.StatusNotFound, _interface.ErrNotFound.Error())
	}

	id := uint(idP)
	ctx := echoContext.Request().Context()

	art, err := a.TagService.GetByID(ctx, id)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusOK, art)
}

func (a *TagHandler) Store(echoContext echo.Context) (err error) {
	var ent entitie.Tag
	err = echoContext.Bind(&ent)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&ent); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := echoContext.Request().Context()
	err = a.TagService.Store(ctx, &ent)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusCreated, ent)
}

func (a *TagHandler) Update(echoContext echo.Context) (err error) {
	var ent entitie.Tag
	err = echoContext.Bind(&ent)
	if err != nil {
		return echoContext.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&ent); !ok {
		return echoContext.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := echoContext.Request().Context()
	err = a.TagService.Update(ctx, &ent)
	if err != nil {
		return echoContext.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return echoContext.JSON(http.StatusCreated, ent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	zap.L().Error("error", zap.Error(err))

	switch err {
	case _interface.ErrInternalServerError:
		return http.StatusInternalServerError
	case _interface.ErrNotFound:
		return http.StatusNotFound
	case _interface.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func isRequestValid(m *entitie.Tag) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
