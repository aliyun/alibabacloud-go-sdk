// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanDeductableCommodityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetSavingPlanDeductableCommodityResponseBodyData) *GetSavingPlanDeductableCommodityResponseBody
	GetData() []*GetSavingPlanDeductableCommodityResponseBodyData
	SetRequestId(v string) *GetSavingPlanDeductableCommodityResponseBody
	GetRequestId() *string
}

type GetSavingPlanDeductableCommodityResponseBody struct {
	Data      []*GetSavingPlanDeductableCommodityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBody) GetData() []*GetSavingPlanDeductableCommodityResponseBodyData {
	return s.Data
}

func (s *GetSavingPlanDeductableCommodityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSavingPlanDeductableCommodityResponseBody) SetData(v []*GetSavingPlanDeductableCommodityResponseBodyData) *GetSavingPlanDeductableCommodityResponseBody {
	s.Data = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBody) SetRequestId(v string) *GetSavingPlanDeductableCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBody) Validate() error {
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

type GetSavingPlanDeductableCommodityResponseBodyData struct {
	ActivityId            *int64                                                            `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	CommodityCode         *string                                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityId           *int64                                                            `json:"CommodityId,omitempty" xml:"CommodityId,omitempty"`
	CommodityName         *string                                                           `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	CycleList             []*GetSavingPlanDeductableCommodityResponseBodyDataCycleList      `json:"CycleList,omitempty" xml:"CycleList,omitempty" type:"Repeated"`
	FilterModules         []*GetSavingPlanDeductableCommodityResponseBodyDataFilterModules  `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ItemCode              *string                                                           `json:"ItemCode,omitempty" xml:"ItemCode,omitempty"`
	ItemId                *int64                                                            `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	ItemName              *string                                                           `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	ModuleMapList         []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList  `json:"ModuleMapList,omitempty" xml:"ModuleMapList,omitempty" type:"Repeated"`
	PayModeList           []*GetSavingPlanDeductableCommodityResponseBodyDataPayModeList    `json:"PayModeList,omitempty" xml:"PayModeList,omitempty" type:"Repeated"`
	PricingModules        []*GetSavingPlanDeductableCommodityResponseBodyDataPricingModules `json:"PricingModules,omitempty" xml:"PricingModules,omitempty" type:"Repeated"`
	SpnCommodityCode      *string                                                           `json:"SpnCommodityCode,omitempty" xml:"SpnCommodityCode,omitempty"`
	SpnCommodityName      *string                                                           `json:"SpnCommodityName,omitempty" xml:"SpnCommodityName,omitempty"`
	SpnDiscountConfigType *string                                                           `json:"SpnDiscountConfigType,omitempty" xml:"SpnDiscountConfigType,omitempty"`
	StepPriceMap          map[string][]*DataStepPriceMapValue                               `json:"StepPriceMap,omitempty" xml:"StepPriceMap,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetCommodityId() *int64 {
	return s.CommodityId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetCycleList() []*GetSavingPlanDeductableCommodityResponseBodyDataCycleList {
	return s.CycleList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetFilterModules() []*GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	return s.FilterModules
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetItemCode() *string {
	return s.ItemCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetItemId() *int64 {
	return s.ItemId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetItemName() *string {
	return s.ItemName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetModuleMapList() []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	return s.ModuleMapList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetPayModeList() []*GetSavingPlanDeductableCommodityResponseBodyDataPayModeList {
	return s.PayModeList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetPricingModules() []*GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	return s.PricingModules
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetSpnCommodityCode() *string {
	return s.SpnCommodityCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetSpnCommodityName() *string {
	return s.SpnCommodityName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetSpnDiscountConfigType() *string {
	return s.SpnDiscountConfigType
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) GetStepPriceMap() map[string][]*DataStepPriceMapValue {
	return s.StepPriceMap
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetActivityId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ActivityId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCommodityName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetCycleList(v []*GetSavingPlanDeductableCommodityResponseBodyDataCycleList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.CycleList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetFilterModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.FilterModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemId(v int64) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetItemName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ItemName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetModuleMapList(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.ModuleMapList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetPayModeList(v []*GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.PayModeList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetPricingModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.PricingModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnCommodityCode(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnCommodityCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnCommodityName(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnCommodityName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetSpnDiscountConfigType(v string) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.SpnDiscountConfigType = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) SetStepPriceMap(v map[string][]*DataStepPriceMapValue) *GetSavingPlanDeductableCommodityResponseBodyData {
	s.StepPriceMap = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyData) Validate() error {
	if s.CycleList != nil {
		for _, item := range s.CycleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterModules != nil {
		for _, item := range s.FilterModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModuleMapList != nil {
		for _, item := range s.ModuleMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PayModeList != nil {
		for _, item := range s.PayModeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PricingModules != nil {
		for _, item := range s.PricingModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSavingPlanDeductableCommodityResponseBodyDataCycleList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataCycleList) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataCycleList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) GetCode() *string {
	return s.Code
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) GetName() *string {
	return s.Name
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataCycleList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataCycleList {
	s.Name = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataCycleList) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataFilterModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataFilterModules) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList struct {
	FilterModules   []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules   `json:"FilterModules,omitempty" xml:"FilterModules,omitempty" type:"Repeated"`
	ModuleCode      *string                                                                         `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId        *int64                                                                          `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName      *string                                                                         `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	ShowModules     []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules     `json:"ShowModules,omitempty" xml:"ShowModules,omitempty" type:"Repeated"`
	SpnTypeList     []*string                                                                       `json:"SpnTypeList,omitempty" xml:"SpnTypeList,omitempty" type:"Repeated"`
	SpnTypeMapList  []map[string]*DataModuleMapListSpnTypeMapListValue                              `json:"SpnTypeMapList,omitempty" xml:"SpnTypeMapList,omitempty" type:"Repeated"`
	SpnTypeNameList []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList `json:"SpnTypeNameList,omitempty" xml:"SpnTypeNameList,omitempty" type:"Repeated"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetFilterModules() []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	return s.FilterModules
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetShowModules() []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	return s.ShowModules
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetSpnTypeList() []*string {
	return s.SpnTypeList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetSpnTypeMapList() []map[string]*DataModuleMapListSpnTypeMapListValue {
	return s.SpnTypeMapList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) GetSpnTypeNameList() []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList {
	return s.SpnTypeNameList
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetFilterModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.FilterModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetShowModules(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.ShowModules = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeList(v []*string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeMapList(v []map[string]*DataModuleMapListSpnTypeMapListValue) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeMapList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) SetSpnTypeNameList(v []*GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList {
	s.SpnTypeNameList = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapList) Validate() error {
	if s.FilterModules != nil {
		for _, item := range s.FilterModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ShowModules != nil {
		for _, item := range s.ShowModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SpnTypeNameList != nil {
		for _, item := range s.SpnTypeNameList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListFilterModules) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListShowModules) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) GetCode() *string {
	return s.Code
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) GetName() *string {
	return s.Name
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList {
	s.Name = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataModuleMapListSpnTypeNameList) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataPayModeList struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) GetCode() *string {
	return s.Code
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) GetName() *string {
	return s.Name
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) SetCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList {
	s.Code = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) SetName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList {
	s.Name = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPayModeList) Validate() error {
	return dara.Validate(s)
}

type GetSavingPlanDeductableCommodityResponseBodyDataPricingModules struct {
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	ModuleId   *int64  `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleCode(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleId(v int64) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleId = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) SetModuleName(v string) *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules {
	s.ModuleName = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponseBodyDataPricingModules) Validate() error {
	return dara.Validate(s)
}
