package firstpresentent

import "fmt"

type FirstPresentmentInterface interface {
	AddDES()
	AddPDS()
}

type commonFirstPresentment struct {
}

type ARFirstPresentment struct {
	commonFirstPresentment commonFirstPresentment
}

type BRFirstPresentment struct {
	commonFirstPresentment commonFirstPresentment
}

func NewARFirstPresentment() ARFirstPresentment {
	return ARFirstPresentment{}
}

func NewBRFirstPresentment() BRFirstPresentment {
	return BRFirstPresentment{}
}

func (fp ARFirstPresentment) AddDES() {
	fp.commonFirstPresentment.AddCommonDES()
	fmt.Print("DE66 added\n")
	fmt.Print("DE77 added\n")

}

func (fp ARFirstPresentment) AddPDS() {
	fp.commonFirstPresentment.AddCommonPDS()
	fmt.Print("PDS1001 added\n")
	fmt.Print("PDS1002 added\n")
}

func (fp BRFirstPresentment) AddDES() {
	fp.commonFirstPresentment.AddCommonDES()
	fmt.Print("DE22 added\n")
	fmt.Print("DE24 added\n")
}

func (fp BRFirstPresentment) AddPDS() {
	fp.commonFirstPresentment.AddCommonPDS()
	fmt.Print("PDS181 added\n")
	fmt.Print("PDS182 added\n")
}

func (fp commonFirstPresentment) AddCommonDES() {
	fmt.Print("DE01 added\n")
	fmt.Print("DE02 added\n")
	fmt.Print("DE03 added\n")
	fmt.Print("DE04 added\n")
}

func (fp commonFirstPresentment) AddCommonPDS() {
	fmt.Printf("PDS 663 added\n")
}
