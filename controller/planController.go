package controller

import (
	"net/http"

	"github.com/citiaps/GOLANG-Malla-backend/services"

	"github.com/gin-gonic/gin"
)

const (
	CollectionNamePlan = "Plan"
)

// GetAllPlanes godoc
// @Summary        Obtiene todos los planes de estudio
// @Description    Esta función obtiene todos los planes de estudio.
// @Tags           Plan
// @Accept         json
// @Product        json
// @Method         GET
// @Security       Bearer
// @Success		   200 {array} models.Plan "Consulta realizada con éxito"
// @Failure        404 {object} ErrorResponse "No se encontró datos."
// @Router         /mallas/plan/ [get]
func GetAllPlanes(ctx *gin.Context) {

	resultsPlanes, err := services.GetAllPlanesService()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resultsPlanes)
}
