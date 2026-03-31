// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModuleDetails(v []*DescribeChargeResultResponseBodyModuleDetails) *DescribeChargeResultResponseBody
	GetModuleDetails() []*DescribeChargeResultResponseBodyModuleDetails
	SetRequestId(v string) *DescribeChargeResultResponseBody
	GetRequestId() *string
	SetTotalSeCu(v float64) *DescribeChargeResultResponseBody
	GetTotalSeCu() *float64
}

type DescribeChargeResultResponseBody struct {
	ModuleDetails []*DescribeChargeResultResponseBodyModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1080
	TotalSeCu *float64 `json:"TotalSeCu,omitempty" xml:"TotalSeCu,omitempty"`
}

func (s DescribeChargeResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChargeResultResponseBody) GetModuleDetails() []*DescribeChargeResultResponseBodyModuleDetails {
	return s.ModuleDetails
}

func (s *DescribeChargeResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChargeResultResponseBody) GetTotalSeCu() *float64 {
	return s.TotalSeCu
}

func (s *DescribeChargeResultResponseBody) SetModuleDetails(v []*DescribeChargeResultResponseBodyModuleDetails) *DescribeChargeResultResponseBody {
	s.ModuleDetails = v
	return s
}

func (s *DescribeChargeResultResponseBody) SetRequestId(v string) *DescribeChargeResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChargeResultResponseBody) SetTotalSeCu(v float64) *DescribeChargeResultResponseBody {
	s.TotalSeCu = &v
	return s
}

func (s *DescribeChargeResultResponseBody) Validate() error {
	if s.ModuleDetails != nil {
		for _, item := range s.ModuleDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChargeResultResponseBodyModuleDetails struct {
	// example:
	//
	// domainCount
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// example:
	//
	// 1080
	SeCu *float64 `json:"SeCu,omitempty" xml:"SeCu,omitempty"`
}

func (s DescribeChargeResultResponseBodyModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeResultResponseBodyModuleDetails) GoString() string {
	return s.String()
}

func (s *DescribeChargeResultResponseBodyModuleDetails) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeChargeResultResponseBodyModuleDetails) GetSeCu() *float64 {
	return s.SeCu
}

func (s *DescribeChargeResultResponseBodyModuleDetails) SetModuleCode(v string) *DescribeChargeResultResponseBodyModuleDetails {
	s.ModuleCode = &v
	return s
}

func (s *DescribeChargeResultResponseBodyModuleDetails) SetSeCu(v float64) *DescribeChargeResultResponseBodyModuleDetails {
	s.SeCu = &v
	return s
}

func (s *DescribeChargeResultResponseBodyModuleDetails) Validate() error {
	return dara.Validate(s)
}
