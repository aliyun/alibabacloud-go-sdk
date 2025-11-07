// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyControlStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlStrategyList(v []*ModifyControlStrategyRequestControlStrategyList) *ModifyControlStrategyRequest
	GetControlStrategyList() []*ModifyControlStrategyRequestControlStrategyList
	SetProductType(v string) *ModifyControlStrategyRequest
	GetProductType() *string
	SetRegionId(v string) *ModifyControlStrategyRequest
	GetRegionId() *string
}

type ModifyControlStrategyRequest struct {
	// List of security alarm rules.
	ControlStrategyList []*ModifyControlStrategyRequestControlStrategyList `json:"ControlStrategyList,omitempty" xml:"ControlStrategyList,omitempty" type:"Repeated"`
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade Real Person), all others are phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Region ID of the intelligent access gateway instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyControlStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyControlStrategyRequest) GetControlStrategyList() []*ModifyControlStrategyRequestControlStrategyList {
	return s.ControlStrategyList
}

func (s *ModifyControlStrategyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyControlStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyControlStrategyRequest) SetControlStrategyList(v []*ModifyControlStrategyRequestControlStrategyList) *ModifyControlStrategyRequest {
	s.ControlStrategyList = v
	return s
}

func (s *ModifyControlStrategyRequest) SetProductType(v string) *ModifyControlStrategyRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyControlStrategyRequest) SetRegionId(v string) *ModifyControlStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyControlStrategyRequest) Validate() error {
	if s.ControlStrategyList != nil {
		for _, item := range s.ControlStrategyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyControlStrategyRequestControlStrategyList struct {
	// API name, same as **ProductCode**.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Rule configuration type:
	//
	// - **QPS**: QPS greater than
	//
	// - **SUCCESS_RATE_5_MIN**: Success rate in the last 5 minutes less than
	//
	// - **RESP_TIME_5_MIN**: Average response time in the last 5 minutes greater than
	//
	// - **AMOUNT_RISE**: Call volume growth ratio greater than
	//
	// - **AMOUNT_FALL**: Call volume decline ratio less than
	//
	// - **PASSED_RATE_1_HOUR**: Verification consistency rate in the last hour less than
	//
	// - **PARAM_ERROR_RATE_1_HOUR**: Parameter error rate in the last hour greater than
	//
	// example:
	//
	// SUCCESS_RATE_5_MIN
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Status:
	//
	// - **disabled**: Disabled
	//
	// - **normal**: Enabled
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Alarm threshold for the rule.
	//
	// example:
	//
	// 0.9
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s ModifyControlStrategyRequestControlStrategyList) String() string {
	return dara.Prettify(s)
}

func (s ModifyControlStrategyRequestControlStrategyList) GoString() string {
	return s.String()
}

func (s *ModifyControlStrategyRequestControlStrategyList) GetApiName() *string {
	return s.ApiName
}

func (s *ModifyControlStrategyRequestControlStrategyList) GetBizType() *string {
	return s.BizType
}

func (s *ModifyControlStrategyRequestControlStrategyList) GetId() *int64 {
	return s.Id
}

func (s *ModifyControlStrategyRequestControlStrategyList) GetStatus() *string {
	return s.Status
}

func (s *ModifyControlStrategyRequestControlStrategyList) GetThreshold() *float64 {
	return s.Threshold
}

func (s *ModifyControlStrategyRequestControlStrategyList) SetApiName(v string) *ModifyControlStrategyRequestControlStrategyList {
	s.ApiName = &v
	return s
}

func (s *ModifyControlStrategyRequestControlStrategyList) SetBizType(v string) *ModifyControlStrategyRequestControlStrategyList {
	s.BizType = &v
	return s
}

func (s *ModifyControlStrategyRequestControlStrategyList) SetId(v int64) *ModifyControlStrategyRequestControlStrategyList {
	s.Id = &v
	return s
}

func (s *ModifyControlStrategyRequestControlStrategyList) SetStatus(v string) *ModifyControlStrategyRequestControlStrategyList {
	s.Status = &v
	return s
}

func (s *ModifyControlStrategyRequestControlStrategyList) SetThreshold(v float64) *ModifyControlStrategyRequestControlStrategyList {
	s.Threshold = &v
	return s
}

func (s *ModifyControlStrategyRequestControlStrategyList) Validate() error {
	return dara.Validate(s)
}
