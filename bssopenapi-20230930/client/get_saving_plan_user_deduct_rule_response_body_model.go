// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanUserDeductRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSavingPlanUserDeductRuleResponseBodyData) *GetSavingPlanUserDeductRuleResponseBody
	GetData() []*GetSavingPlanUserDeductRuleResponseBodyData
	SetRequestId(v string) *GetSavingPlanUserDeductRuleResponseBody
	GetRequestId() *string
}

type GetSavingPlanUserDeductRuleResponseBody struct {
	Data      []*GetSavingPlanUserDeductRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponseBody) GetData() []*GetSavingPlanUserDeductRuleResponseBodyData {
	return s.Data
}

func (s *GetSavingPlanUserDeductRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSavingPlanUserDeductRuleResponseBody) SetData(v []*GetSavingPlanUserDeductRuleResponseBodyData) *GetSavingPlanUserDeductRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBody) SetRequestId(v string) *GetSavingPlanUserDeductRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSavingPlanUserDeductRuleResponseBodyData struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	ModuleCode    *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleName    *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	SkipDeduct    *bool   `json:"SkipDeduct,omitempty" xml:"SkipDeduct,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) GetSkipDeduct() *bool {
	return s.SkipDeduct
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetCommodityCode(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetCommodityName(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetModuleCode(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetModuleName(v string) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) SetSkipDeduct(v bool) *GetSavingPlanUserDeductRuleResponseBodyData {
	s.SkipDeduct = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
