// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplementDagrunInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSupplementDagrunInstanceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSupplementDagrunInstanceResponseBody
	GetHttpStatusCode() *int32
	SetInstanceList(v []*GetSupplementDagrunInstanceResponseBodyInstanceList) *GetSupplementDagrunInstanceResponseBody
	GetInstanceList() []*GetSupplementDagrunInstanceResponseBodyInstanceList
	SetMessage(v string) *GetSupplementDagrunInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSupplementDagrunInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSupplementDagrunInstanceResponseBody
	GetSuccess() *bool
}

type GetSupplementDagrunInstanceResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	InstanceList   []*GetSupplementDagrunInstanceResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSupplementDagrunInstanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSupplementDagrunInstanceResponseBody) GetInstanceList() []*GetSupplementDagrunInstanceResponseBodyInstanceList {
	return s.InstanceList
}

func (s *GetSupplementDagrunInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSupplementDagrunInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupplementDagrunInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSupplementDagrunInstanceResponseBody) SetCode(v string) *GetSupplementDagrunInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetHttpStatusCode(v int32) *GetSupplementDagrunInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetInstanceList(v []*GetSupplementDagrunInstanceResponseBodyInstanceList) *GetSupplementDagrunInstanceResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetMessage(v string) *GetSupplementDagrunInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetRequestId(v string) *GetSupplementDagrunInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) SetSuccess(v bool) *GetSupplementDagrunInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunInstanceResponseBodyInstanceList struct {
	// example:
	//
	// 2024-04-01
	BizDate *int64 `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// 2024-04-02
	DueTime *int64 `json:"DueTime,omitempty" xml:"DueTime,omitempty"`
	// example:
	//
	// 60
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-04-12 00:25:02
	EndExecuteTime *int64 `json:"EndExecuteTime,omitempty" xml:"EndExecuteTime,omitempty"`
	// example:
	//
	// {"a":"b"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// t_239496_20210411_246982077481
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	Index    *int32                                                       `json:"Index,omitempty" xml:"Index,omitempty"`
	NodeInfo *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Struct"`
	// example:
	//
	// 2024-04-12 00:00:00
	StartExecuteTime *int64    `json:"StartExecuteTime,omitempty" xml:"StartExecuteTime,omitempty"`
	StatusList       []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// example:
	//
	// SUPPLEMENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetBizDate() *int64 {
	return s.BizDate
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetDueTime() *int64 {
	return s.DueTime
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetDuration() *string {
	return s.Duration
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetEndExecuteTime() *int64 {
	return s.EndExecuteTime
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetIndex() *int32 {
	return s.Index
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetNodeInfo() *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	return s.NodeInfo
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetStartExecuteTime() *int64 {
	return s.StartExecuteTime
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetStatusList() []*string {
	return s.StatusList
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) GetType() *string {
	return s.Type
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetBizDate(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.BizDate = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetDueTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.DueTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetDuration(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Duration = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetEndExecuteTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.EndExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetExtendInfo(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.ExtendInfo = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetIndex(v int32) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Index = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetNodeInfo(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.NodeInfo = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetStartExecuteTime(v int64) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.StartExecuteTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetStatusList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.StatusList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) SetType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceList {
	s.Type = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo struct {
	// example:
	//
	// xx测试
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-01-30 10:08:49
	CreateTime *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// xx测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// DATA_PROCESS
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// true
	HasDev *bool `json:"HasDev,omitempty" xml:"HasDev,omitempty"`
	// example:
	//
	// true
	HasProd *bool `json:"HasProd,omitempty" xml:"HasProd,omitempty"`
	// example:
	//
	// n_239496
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-01-30 10:08:49
	LastModifiedTime *string                                                              `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	// example:
	//
	// xx测试
	Name              *string                                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList         []*GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList      []*string                                                               `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ResourceGroupList []*string                                                               `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	SchedulePaused     *bool     `json:"SchedulePaused,omitempty" xml:"SchedulePaused,omitempty"`
	SchedulePeriodList []*string `json:"SchedulePeriodList,omitempty" xml:"SchedulePeriodList,omitempty" type:"Repeated"`
	// example:
	//
	// SHELL
	SubDetailType *string `json:"SubDetailType,omitempty" xml:"SubDetailType,omitempty"`
	// example:
	//
	// DATA_PROCESS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetCreator() *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator {
	return s.Creator
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetDescription() *string {
	return s.Description
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetDryRun() *bool {
	return s.DryRun
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetFrom() *string {
	return s.From
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetHasDev() *bool {
	return s.HasDev
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetHasProd() *bool {
	return s.HasProd
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetModifier() *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier {
	return s.Modifier
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetName() *string {
	return s.Name
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetOwnerList() []*GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList {
	return s.OwnerList
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetResourceGroupList() []*string {
	return s.ResourceGroupList
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetSubDetailType() *string {
	return s.SubDetailType
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) GetType() *string {
	return s.Type
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetBizUnitName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.BizUnitName = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetCreateTime(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.CreateTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetCreator(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Creator = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetDescription(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Description = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetDryRun(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.DryRun = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetFrom(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.From = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetHasDev(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.HasDev = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetHasProd(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.HasProd = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetLastModifiedTime(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.LastModifiedTime = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetModifier(v *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Modifier = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Name = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetOwnerList(v []*GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.OwnerList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetPriorityList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.PriorityList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetResourceGroupList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.ResourceGroupList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSchedulePaused(v bool) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SchedulePaused = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSchedulePeriodList(v []*string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SchedulePeriodList = v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetSubDetailType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.SubDetailType = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) SetType(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo {
	s.Type = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfo) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) GetName() *string {
	return s.Name
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator {
	s.Name = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoCreator) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) GetName() *string {
	return s.Name
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier {
	s.Name = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoModifier) Validate() error {
	return dara.Validate(s)
}

type GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList struct {
	// example:
	//
	// 1001012
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// xx测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) String() string {
	return dara.Prettify(s)
}

func (s GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) GoString() string {
	return s.String()
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) GetId() *string {
	return s.Id
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) GetName() *string {
	return s.Name
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) SetId(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList {
	s.Id = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) SetName(v string) *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList {
	s.Name = &v
	return s
}

func (s *GetSupplementDagrunInstanceResponseBodyInstanceListNodeInfoOwnerList) Validate() error {
	return dara.Validate(s)
}
