package wrapper

import (
	"errors"

	"github.com/andig/evcc/api"
)

type ChargeProgress struct {
	charger api.Charger
	vehicle api.Vehicle
}

func NewChargeProgress(charger api.Charger, vehicle api.Vehicle) *ChargeProgress {
	return &ChargeProgress{
		charger: charger,
		vehicle: vehicle,
	}
}

func (cp *ChargeProgress) Vehicle(vehicle api.Vehicle) {
	cp.vehicle = vehicle
}

func (cp *ChargeProgress) SoC() (res float64, err error) {
	err = api.ErrNotAvailable
	if charger, ok := cp.charger.(api.Battery); ok {
		res, err = charger.SoC()
	}

	if vehicle, ok := cp.vehicle.(api.Battery); errors.Is(err, api.ErrNotAvailable) && ok {
		res, err = vehicle.SoC()
	}

	return
}
