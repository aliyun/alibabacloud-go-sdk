// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListInstancesResponseBody
	GetMessage() *string
	SetPageResult(v *ListInstancesResponseBodyPageResult) *ListInstancesResponseBody
	GetPageResult() *ListInstancesResponseBodyPageResult
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstancesResponseBody
	GetSuccess() *bool
}

type ListInstancesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListInstancesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstancesResponseBody) GetPageResult() *ListInstancesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstancesResponseBody) SetCode(v string) *ListInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstancesResponseBody) SetHttpStatusCode(v int32) *ListInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListInstancesResponseBody) SetMessage(v string) *ListInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstancesResponseBody) SetPageResult(v *ListInstancesResponseBodyPageResult) *ListInstancesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetSuccess(v bool) *ListInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyPageResult struct {
	Data []*ListInstancesResponseBodyPageResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 107
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResult) GetData() []*ListInstancesResponseBodyPageResultData {
	return s.Data
}

func (s *ListInstancesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBodyPageResult) SetData(v []*ListInstancesResponseBodyPageResultData) *ListInstancesResponseBodyPageResult {
	s.Data = v
	return s
}

func (s *ListInstancesResponseBodyPageResult) SetTotalCount(v int32) *ListInstancesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBodyPageResult) Validate() error {
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

type ListInstancesResponseBodyPageResultData struct {
	// example:
	//
	// 2024-05-30
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	DueTime *string `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 60s
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// xx
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// t_23231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                           `json:"Index,omitempty" xml:"Index,omitempty"`
	NodeInfo *ListInstancesResponseBodyPageResultDataNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2024-05-30 16:46:13
	StartExecuteTime *int64    `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	StatusList       []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyPageResultData) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultData) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultData) GetBizDate() *string {
	return s.BizDate
}

func (s *ListInstancesResponseBodyPageResultData) GetDueTime() *string {
	return s.DueTime
}

func (s *ListInstancesResponseBodyPageResultData) GetDuration() *string {
	return s.Duration
}

func (s *ListInstancesResponseBodyPageResultData) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *ListInstancesResponseBodyPageResultData) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListInstancesResponseBodyPageResultData) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyPageResultData) GetIndex() *int32 {
	return s.Index
}

func (s *ListInstancesResponseBodyPageResultData) GetNodeInfo() *ListInstancesResponseBodyPageResultDataNodeInfo {
	return s.NodeInfo
}

func (s *ListInstancesResponseBodyPageResultData) GetStartExecuteTime() *int64 {
	return s.StartExecuteTime
}

func (s *ListInstancesResponseBodyPageResultData) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListInstancesResponseBodyPageResultData) SetBizDate(v string) *ListInstancesResponseBodyPageResultData {
	s.BizDate = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetDueTime(v string) *ListInstancesResponseBodyPageResultData {
	s.DueTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetDuration(v string) *ListInstancesResponseBodyPageResultData {
	s.Duration = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetEndExecuteTime(v int64) *ListInstancesResponseBodyPageResultData {
	s.EndExecuteTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetExtendInfo(v string) *ListInstancesResponseBodyPageResultData {
	s.ExtendInfo = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetId(v string) *ListInstancesResponseBodyPageResultData {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetIndex(v int32) *ListInstancesResponseBodyPageResultData {
	s.Index = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetNodeInfo(v *ListInstancesResponseBodyPageResultDataNodeInfo) *ListInstancesResponseBodyPageResultData {
	s.NodeInfo = v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetStartExecuteTime(v int64) *ListInstancesResponseBodyPageResultData {
	s.StartExecuteTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) SetStatusList(v []*string) *ListInstancesResponseBodyPageResultData {
	s.StatusList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultData) Validate() error {
	if s.NodeInfo != nil {
		if err := s.NodeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesResponseBodyPageResultDataNodeInfo struct {
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	CreateTime  *string                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator     *ListInstancesResponseBodyPageResultDataNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	Description *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DATA_PROCES
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// false
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_132331
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	LastModifiedTime  *string                                                     `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier          *ListInstancesResponseBodyPageResultDataNodeInfoModifier    `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name              *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList         []*ListInstancesResponseBodyPageResultDataNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList      []*string                                                   `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ResourceGroupList []*string                                                   `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// false
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCES
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfo) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetCreator() *ListInstancesResponseBodyPageResultDataNodeInfoCreator {
	return s.Creator
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetFrom() *string {
	return s.From
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetHasDev() *bool {
	return s.HasDev
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetHasProd() *bool {
	return s.HasProd
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetModifier() *ListInstancesResponseBodyPageResultDataNodeInfoModifier {
	return s.Modifier
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetOwnerList() []*ListInstancesResponseBodyPageResultDataNodeInfoOwnerList {
	return s.OwnerList
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetResourceGroupList() []*string {
	return s.ResourceGroupList
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetSubDetailType() *string {
	return s.SubDetailType
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetBizUnitName(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetCreateTime(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetCreator(v *ListInstancesResponseBodyPageResultDataNodeInfoCreator) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Creator = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetDescription(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetDryRun(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.DryRun = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetFrom(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.From = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetHasDev(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.HasDev = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetHasProd(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.HasProd = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetLastModifiedTime(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetModifier(v *ListInstancesResponseBodyPageResultDataNodeInfoModifier) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Modifier = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetOwnerList(v []*ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.OwnerList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetPriorityList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.PriorityList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetResourceGroupList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSchedulePaused(v bool) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSchedulePeriodList(v []*string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetSubDetailType(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) SetType(v string) *ListInstancesResponseBodyPageResultDataNodeInfo {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfo) Validate() error {
	if s.Creator != nil {
		if err := s.Creator.Validate(); err != nil {
			return err
		}
	}
	if s.Modifier != nil {
		if err := s.Modifier.Validate(); err != nil {
			return err
		}
	}
	if s.OwnerList != nil {
		for _, item := range s.OwnerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyPageResultDataNodeInfoCreator struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoCreator {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoCreator) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyPageResultDataNodeInfoModifier struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoModifier) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoModifier {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoModifier) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyPageResultDataNodeInfoOwnerList struct {
	// example:
	//
	// 21313112
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) GetId() *string {
	return s.Id
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) GetName() *string {
	return s.Name
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) SetId(v string) *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) SetName(v string) *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList {
	s.Name = &v
	return s
}

func (s *ListInstancesResponseBodyPageResultDataNodeInfoOwnerList) Validate() error {
	return dara.Validate(s)
}
