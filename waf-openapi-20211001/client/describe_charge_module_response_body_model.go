// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChargeModuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChargeModules(v []*DescribeChargeModuleResponseBodyChargeModules) *DescribeChargeModuleResponseBody
	GetChargeModules() []*DescribeChargeModuleResponseBodyChargeModules
	SetRequestId(v string) *DescribeChargeModuleResponseBody
	GetRequestId() *string
}

type DescribeChargeModuleResponseBody struct {
	ChargeModules []*DescribeChargeModuleResponseBodyChargeModules `json:"ChargeModules,omitempty" xml:"ChargeModules,omitempty" type:"Repeated"`
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChargeModuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleResponseBody) GetChargeModules() []*DescribeChargeModuleResponseBodyChargeModules {
	return s.ChargeModules
}

func (s *DescribeChargeModuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChargeModuleResponseBody) SetChargeModules(v []*DescribeChargeModuleResponseBodyChargeModules) *DescribeChargeModuleResponseBody {
	s.ChargeModules = v
	return s
}

func (s *DescribeChargeModuleResponseBody) SetRequestId(v string) *DescribeChargeModuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChargeModuleResponseBody) Validate() error {
	if s.ChargeModules != nil {
		for _, item := range s.ChargeModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChargeModuleResponseBodyChargeModules struct {
	// example:
	//
	// NORMAL_PRICE
	ChargeMode        *string   `json:"ChargeMode,omitempty" xml:"ChargeMode,omitempty"`
	ChargeModeDetails []*string `json:"ChargeModeDetails,omitempty" xml:"ChargeModeDetails,omitempty" type:"Repeated"`
	// example:
	//
	// domainCount
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// example:
	//
	// Hour
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// example:
	//
	// domain
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
	// example:
	//
	// 1
	UsageUnitFactor *int32 `json:"UsageUnitFactor,omitempty" xml:"UsageUnitFactor,omitempty"`
}

func (s DescribeChargeModuleResponseBodyChargeModules) String() string {
	return dara.Prettify(s)
}

func (s DescribeChargeModuleResponseBodyChargeModules) GoString() string {
	return s.String()
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetChargeMode() *string {
	return s.ChargeMode
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetChargeModeDetails() []*string {
	return s.ChargeModeDetails
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeChargeModuleResponseBodyChargeModules) GetUsageUnitFactor() *int32 {
	return s.UsageUnitFactor
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetChargeMode(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ChargeMode = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetChargeModeDetails(v []*string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ChargeModeDetails = v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetModuleCode(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.ModuleCode = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetPeriodType(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.PeriodType = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetUsageType(v string) *DescribeChargeModuleResponseBodyChargeModules {
	s.UsageType = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) SetUsageUnitFactor(v int32) *DescribeChargeModuleResponseBodyChargeModules {
	s.UsageUnitFactor = &v
	return s
}

func (s *DescribeChargeModuleResponseBodyChargeModules) Validate() error {
	return dara.Validate(s)
}
