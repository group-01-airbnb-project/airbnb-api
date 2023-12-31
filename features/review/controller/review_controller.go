package controller

import (
	"be-api/app/middlewares"
	"be-api/features"
	"be-api/features/review"
	"be-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ReviewControll struct {
	reviewControll review.ReviewServiceInterface
}

func New(review review.ReviewServiceInterface) *ReviewControll {
	return &ReviewControll{
		reviewControll: review,
	}
}

func (control *ReviewControll) AddReview(c echo.Context) error {
	inputReview := features.ReviewEntity{}
	id_Costumer := middlewares.ExtractTokenUserId(c)

	err := c.Bind(&inputReview)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind review data", nil))
	}

	id_Review, errReview := control.reviewControll.AddReview(inputReview, uint(id_Costumer))
	if errReview != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}
	review, errReview := control.reviewControll.GetId(id_Review)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed, data tidak ditemukan", nil))
	}

	data := EntityToResponse(review)

	return c.JSON(http.StatusOK, utils.SuccessResponse("review created successfully", data))
}

func (control *ReviewControll) DeleteReview(c echo.Context) error {
	Id := c.Param("review_id")
	idConv, errConv := strconv.Atoi(Id)
	if errConv != nil {
		if strings.Contains(errConv.Error(), "bind failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(errConv.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to bind data. "+errConv.Error(), nil))
		}
	}

	err := control.reviewControll.DeleteReview(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("delete Fail to Delete akun User", nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessWhitoutResponse("Success delete akun User"))
}

func (control *ReviewControll) GetAllReview(c echo.Context) error {
	Id := c.Param("homestay_id")
	idConv, errConv := strconv.Atoi(Id)
	if errConv != nil {
		if strings.Contains(errConv.Error(), "bind failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(errConv.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to bind data. "+errConv.Error(), nil))
		}
	}

	allReviews, errGetAll := control.reviewControll.GetAll(uint(idConv))
	if errGetAll != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(errGetAll.Error(), nil))
	}
	var mapAllReview []ResponseReviews
	for _, reviews := range allReviews {
		mapreview := EntityToResponse(reviews)
		mapAllReview = append(mapAllReview, mapreview)
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("review read review successfully", mapAllReview))

}
