// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListNodesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListNodesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListNodesResponseBody
	GetMessage() *string
	SetPageResult(v *ListNodesResponseBodyPageResult) *ListNodesResponseBody
	GetPageResult() *ListNodesResponseBodyPageResult
	SetRequestId(v string) *ListNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListNodesResponseBody
	GetSuccess() *bool
}

type ListNodesResponseBody struct {
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
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListNodesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListNodesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListNodesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListNodesResponseBody) GetPageResult() *ListNodesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListNodesResponseBody) SetCode(v string) *ListNodesResponseBody {
	s.Code = &v
	return s
}

func (s *ListNodesResponseBody) SetHttpStatusCode(v int32) *ListNodesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListNodesResponseBody) SetMessage(v string) *ListNodesResponseBody {
	s.Message = &v
	return s
}

func (s *ListNodesResponseBody) SetPageResult(v *ListNodesResponseBodyPageResult) *ListNodesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetSuccess(v bool) *ListNodesResponseBody {
	s.Success = &v
	return s
}

func (s *ListNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResult struct {
	NodeList []*ListNodesResponseBodyPageResultNodeList `json:"NodeList,omitempty" xml:"NodeList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNodesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResult) GetNodeList() []*ListNodesResponseBodyPageResultNodeList {
	return s.NodeList
}

func (s *ListNodesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNodesResponseBodyPageResult) SetNodeList(v []*ListNodesResponseBodyPageResultNodeList) *ListNodesResponseBodyPageResult {
	s.NodeList = v
	return s
}

func (s *ListNodesResponseBodyPageResult) SetTotalCount(v int32) *ListNodesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBodyPageResult) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResultNodeList struct {
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	CreateTime *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Creator    *ListNodesResponseBodyPageResultNodeListCreator `json:"Creator,omitempty" xml:"Creator,omitempty" type:"Struct"`
	// example:
	//
	// xx test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// {"xx":"zz"}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
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
	// n_31111
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-05-30 16:47:13
	LastModifiedTime *string                                             `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	Modifier         *ListNodesResponseBodyPageResultNodeListModifier    `json:"Modifier,omitempty" xml:"Modifier,omitempty" type:"Struct"`
	Name             *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerList        []*ListNodesResponseBodyPageResultNodeListOwnerList `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	PriorityList     []*string                                           `json:"PriorityList,omitempty" xml:"PriorityList,omitempty" type:"Repeated"`
	ProjectInfo      *ListNodesResponseBodyPageResultNodeListProjectInfo `json:"ProjectInfo,omitempty" xml:"ProjectInfo,omitempty" type:"Struct"`
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

func (s ListNodesResponseBodyPageResultNodeList) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeList) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListNodesResponseBodyPageResultNodeList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNodesResponseBodyPageResultNodeList) GetCreator() *ListNodesResponseBodyPageResultNodeListCreator {
	return s.Creator
}

func (s *ListNodesResponseBodyPageResultNodeList) GetDescription() *string {
	return s.Description
}

func (s *ListNodesResponseBodyPageResultNodeList) GetDryRun() *bool {
	return s.DryRun
}

func (s *ListNodesResponseBodyPageResultNodeList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListNodesResponseBodyPageResultNodeList) GetFrom() *string {
	return s.From
}

func (s *ListNodesResponseBodyPageResultNodeList) GetHasDev() *bool {
	return s.HasDev
}

func (s *ListNodesResponseBodyPageResultNodeList) GetHasProd() *bool {
	return s.HasProd
}

func (s *ListNodesResponseBodyPageResultNodeList) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyPageResultNodeList) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *ListNodesResponseBodyPageResultNodeList) GetModifier() *ListNodesResponseBodyPageResultNodeListModifier {
	return s.Modifier
}

func (s *ListNodesResponseBodyPageResultNodeList) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPageResultNodeList) GetOwnerList() []*ListNodesResponseBodyPageResultNodeListOwnerList {
	return s.OwnerList
}

func (s *ListNodesResponseBodyPageResultNodeList) GetPriorityList() []*string {
	return s.PriorityList
}

func (s *ListNodesResponseBodyPageResultNodeList) GetProjectInfo() *ListNodesResponseBodyPageResultNodeListProjectInfo {
	return s.ProjectInfo
}

func (s *ListNodesResponseBodyPageResultNodeList) GetSchedulePaused() *bool {
	return s.SchedulePaused
}

func (s *ListNodesResponseBodyPageResultNodeList) GetSchedulePeriodList() []*string {
	return s.SchedulePeriodList
}

func (s *ListNodesResponseBodyPageResultNodeList) GetSubDetailType() *string {
	return s.SubDetailType
}

func (s *ListNodesResponseBodyPageResultNodeList) GetType() *string {
	return s.Type
}

func (s *ListNodesResponseBodyPageResultNodeList) SetBizUnitName(v string) *ListNodesResponseBodyPageResultNodeList {
	s.BizUnitName = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetCreateTime(v string) *ListNodesResponseBodyPageResultNodeList {
	s.CreateTime = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetCreator(v *ListNodesResponseBodyPageResultNodeListCreator) *ListNodesResponseBodyPageResultNodeList {
	s.Creator = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetDescription(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Description = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetDryRun(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.DryRun = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetExtendInfo(v string) *ListNodesResponseBodyPageResultNodeList {
	s.ExtendInfo = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetFrom(v string) *ListNodesResponseBodyPageResultNodeList {
	s.From = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetHasDev(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.HasDev = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetHasProd(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.HasProd = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetId(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetLastModifiedTime(v string) *ListNodesResponseBodyPageResultNodeList {
	s.LastModifiedTime = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetModifier(v *ListNodesResponseBodyPageResultNodeListModifier) *ListNodesResponseBodyPageResultNodeList {
	s.Modifier = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetName(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetOwnerList(v []*ListNodesResponseBodyPageResultNodeListOwnerList) *ListNodesResponseBodyPageResultNodeList {
	s.OwnerList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetPriorityList(v []*string) *ListNodesResponseBodyPageResultNodeList {
	s.PriorityList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetProjectInfo(v *ListNodesResponseBodyPageResultNodeListProjectInfo) *ListNodesResponseBodyPageResultNodeList {
	s.ProjectInfo = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSchedulePaused(v bool) *ListNodesResponseBodyPageResultNodeList {
	s.SchedulePaused = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSchedulePeriodList(v []*string) *ListNodesResponseBodyPageResultNodeList {
	s.SchedulePeriodList = v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetSubDetailType(v string) *ListNodesResponseBodyPageResultNodeList {
	s.SubDetailType = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) SetType(v string) *ListNodesResponseBodyPageResultNodeList {
	s.Type = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeList) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResultNodeListCreator struct {
	// example:
	//
	// 23222
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListCreator) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListCreator) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) SetId(v string) *ListNodesResponseBodyPageResultNodeListCreator {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) SetName(v string) *ListNodesResponseBodyPageResultNodeListCreator {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListCreator) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResultNodeListModifier struct {
	// example:
	//
	// 311131
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListModifier) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListModifier) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) SetId(v string) *ListNodesResponseBodyPageResultNodeListModifier {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) SetName(v string) *ListNodesResponseBodyPageResultNodeListModifier {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListModifier) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResultNodeListOwnerList struct {
	// example:
	//
	// 23222
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListOwnerList) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListOwnerList) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) SetId(v string) *ListNodesResponseBodyPageResultNodeListOwnerList {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) SetName(v string) *ListNodesResponseBodyPageResultNodeListOwnerList {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListOwnerList) Validate() error {
	return dara.Validate(s)
}

type ListNodesResponseBodyPageResultNodeListProjectInfo struct {
	// example:
	//
	// 1121321
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListNodesResponseBodyPageResultNodeListProjectInfo) String() string {
	return dara.Prettify(s)
}

func (s ListNodesResponseBodyPageResultNodeListProjectInfo) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) GetId() *string {
	return s.Id
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) GetName() *string {
	return s.Name
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) SetId(v string) *ListNodesResponseBodyPageResultNodeListProjectInfo {
	s.Id = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) SetName(v string) *ListNodesResponseBodyPageResultNodeListProjectInfo {
	s.Name = &v
	return s
}

func (s *ListNodesResponseBodyPageResultNodeListProjectInfo) Validate() error {
	return dara.Validate(s)
}
