package firstpresentent

import "fmt"

type FPFactory struct {
}

const (
	_AR = "AR"
	_BR = "BR"
)

func (factory FPFactory) InitFirstPresentment(country string) (FirstPresentmentInterface, error) {
	switch country {
	case _AR:
		return NewARFirstPresentment(), nil
	case _BR:
		return NewBRFirstPresentment(), nil
	}
	return ARFirstPresentment{}, fmt.Errorf("Couldn't init First Presentment")

}
