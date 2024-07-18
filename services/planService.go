package services

import (
	"fmt"

	"github.com/citiaps/GOLANG-Malla-backend/config"
	"github.com/citiaps/GOLANG-Malla-backend/models"
)

const (
	CollectionNamePlan = "Plan"
)

func GetAllPlanesService() ([]models.Plan, error) {
	var planes []models.Plan
	db := config.Database.Debug()
	query := db.
		Preload("CarreraTotal.AreaCarrera").
		Preload("CarreraTotal.PropietarioCarrera").
		Preload("CarreraTotal.PropietarioCarrera", "cod_propietario = ?", "1").
		Preload("CarreraTotal.Carrera").
		Preload("CarreraTotal.Centro").
		Preload("CarreraTotal.Especialidad").
		Preload("CarreraTotal.Mencion").
		Preload("CarreraTotal.NivelCarrera").
		Preload("NivelGlobalCarrera").
		Preload("TipoPlan").
		Preload("Regimen").
		Where("vigencia IN (?)", 1).

		//	Where("cod_nivel_global != ?", 2).
		Find(&planes).
		Statement
	if query.Error != nil {
		fmt.Println("Error:", query.Error)
		return nil, query.Error
	}
	sql := query.SQL.String()
	fmt.Println(sql)
	return planes, nil
}
