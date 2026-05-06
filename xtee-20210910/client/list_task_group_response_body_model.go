// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTaskGroupResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListTaskGroupResponseBody
	GetCurrentPage() *int32
	SetHttpStatusCode(v string) *ListTaskGroupResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ListTaskGroupResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListTaskGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTaskGroupResponseBody
	GetRequestId() *string
	SetResultObject(v []*ListTaskGroupResponseBodyResultObject) *ListTaskGroupResponseBody
	GetResultObject() []*ListTaskGroupResponseBodyResultObject
	SetTotalItem(v int32) *ListTaskGroupResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *ListTaskGroupResponseBody
	GetTotalPage() *int32
}

type ListTaskGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject []*ListTaskGroupResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalItem *int32 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListTaskGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTaskGroupResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTaskGroupResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListTaskGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTaskGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTaskGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskGroupResponseBody) GetResultObject() []*ListTaskGroupResponseBodyResultObject {
	return s.ResultObject
}

func (s *ListTaskGroupResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *ListTaskGroupResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListTaskGroupResponseBody) SetCode(v string) *ListTaskGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetCurrentPage(v int32) *ListTaskGroupResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetHttpStatusCode(v string) *ListTaskGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetMessage(v string) *ListTaskGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetPageSize(v int32) *ListTaskGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetRequestId(v string) *ListTaskGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetResultObject(v []*ListTaskGroupResponseBodyResultObject) *ListTaskGroupResponseBody {
	s.ResultObject = v
	return s
}

func (s *ListTaskGroupResponseBody) SetTotalItem(v int32) *ListTaskGroupResponseBody {
	s.TotalItem = &v
	return s
}

func (s *ListTaskGroupResponseBody) SetTotalPage(v int32) *ListTaskGroupResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListTaskGroupResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskGroupResponseBodyResultObject struct {
	// example:
	//
	// 1588820785212
	CreateTime    *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatorUserId *int64 `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// RUNNING
	GroupStatus *string   `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	SampleNames []*string `json:"SampleNames,omitempty" xml:"SampleNames,omitempty" type:"Repeated"`
	// example:
	//
	// 4
	SubTaskCount *int32                                              `json:"SubTaskCount,omitempty" xml:"SubTaskCount,omitempty"`
	SubTaskList  []*ListTaskGroupResponseBodyResultObjectSubTaskList `json:"SubTaskList,omitempty" xml:"SubTaskList,omitempty" type:"Repeated"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// g-uf62fwvw2f8dx88xo2lt
	TaskGroupId *int32 `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// example:
	//
	// GroupTest
	TaskGroupName *string `json:"TaskGroupName,omitempty" xml:"TaskGroupName,omitempty"`
}

func (s ListTaskGroupResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ListTaskGroupResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ListTaskGroupResponseBodyResultObject) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTaskGroupResponseBodyResultObject) GetCreatorUserId() *int64 {
	return s.CreatorUserId
}

func (s *ListTaskGroupResponseBodyResultObject) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListTaskGroupResponseBodyResultObject) GetSampleNames() []*string {
	return s.SampleNames
}

func (s *ListTaskGroupResponseBodyResultObject) GetSubTaskCount() *int32 {
	return s.SubTaskCount
}

func (s *ListTaskGroupResponseBodyResultObject) GetSubTaskList() []*ListTaskGroupResponseBodyResultObjectSubTaskList {
	return s.SubTaskList
}

func (s *ListTaskGroupResponseBodyResultObject) GetTab() *string {
	return s.Tab
}

func (s *ListTaskGroupResponseBodyResultObject) GetTaskGroupId() *int32 {
	return s.TaskGroupId
}

func (s *ListTaskGroupResponseBodyResultObject) GetTaskGroupName() *string {
	return s.TaskGroupName
}

func (s *ListTaskGroupResponseBodyResultObject) SetCreateTime(v int64) *ListTaskGroupResponseBodyResultObject {
	s.CreateTime = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetCreatorUserId(v int64) *ListTaskGroupResponseBodyResultObject {
	s.CreatorUserId = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetGroupStatus(v string) *ListTaskGroupResponseBodyResultObject {
	s.GroupStatus = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetSampleNames(v []*string) *ListTaskGroupResponseBodyResultObject {
	s.SampleNames = v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetSubTaskCount(v int32) *ListTaskGroupResponseBodyResultObject {
	s.SubTaskCount = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetSubTaskList(v []*ListTaskGroupResponseBodyResultObjectSubTaskList) *ListTaskGroupResponseBodyResultObject {
	s.SubTaskList = v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetTab(v string) *ListTaskGroupResponseBodyResultObject {
	s.Tab = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetTaskGroupId(v int32) *ListTaskGroupResponseBodyResultObject {
	s.TaskGroupId = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) SetTaskGroupName(v string) *ListTaskGroupResponseBodyResultObject {
	s.TaskGroupName = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObject) Validate() error {
	if s.SubTaskList != nil {
		for _, item := range s.SubTaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskGroupResponseBodyResultObjectSubTaskList struct {
	// example:
	//
	// 2026-01-12 15:47:23
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// example:
	//
	// pts-demo
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// true
	HideViewResultButton *bool `json:"HideViewResultButton,omitempty" xml:"HideViewResultButton,omitempty"`
	// example:
	//
	// true
	IsCharge *bool `json:"IsCharge,omitempty" xml:"IsCharge,omitempty"`
	// example:
	//
	// rfs
	ModelScene *string `json:"ModelScene,omitempty" xml:"ModelScene,omitempty"`
	// example:
	//
	// 4
	SampleId *int32 `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
	// example:
	//
	// fs
	SampleName *string `json:"SampleName,omitempty" xml:"SampleName,omitempty"`
	// example:
	//
	// anti_fraud_riskscreening
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// 3
	SubTaskId *int32 `json:"SubTaskId,omitempty" xml:"SubTaskId,omitempty"`
	// example:
	//
	// true
	SupportCancel *bool `json:"SupportCancel,omitempty" xml:"SupportCancel,omitempty"`
	// example:
	//
	// true
	SupportDownload *bool `json:"SupportDownload,omitempty" xml:"SupportDownload,omitempty"`
	// example:
	//
	// false
	SupportMergeSelect *bool `json:"SupportMergeSelect,omitempty" xml:"SupportMergeSelect,omitempty"`
	// example:
	//
	// true
	SupportView *bool `json:"SupportView,omitempty" xml:"SupportView,omitempty"`
	// example:
	//
	// FINANCE
	Tab *string `json:"Tab,omitempty" xml:"Tab,omitempty"`
	// example:
	//
	// 4
	TaskGroupId *int32 `json:"TaskGroupId,omitempty" xml:"TaskGroupId,omitempty"`
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s ListTaskGroupResponseBodyResultObjectSubTaskList) String() string {
	return dara.Prettify(s)
}

func (s ListTaskGroupResponseBodyResultObjectSubTaskList) GoString() string {
	return s.String()
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetHideViewResultButton() *bool {
	return s.HideViewResultButton
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetIsCharge() *bool {
	return s.IsCharge
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetModelScene() *string {
	return s.ModelScene
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSampleId() *int32 {
	return s.SampleId
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSampleName() *string {
	return s.SampleName
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSubTaskId() *int32 {
	return s.SubTaskId
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSupportCancel() *bool {
	return s.SupportCancel
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSupportDownload() *bool {
	return s.SupportDownload
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSupportMergeSelect() *bool {
	return s.SupportMergeSelect
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetSupportView() *bool {
	return s.SupportView
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetTab() *string {
	return s.Tab
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetTaskGroupId() *int32 {
	return s.TaskGroupId
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetFinishTime(v int64) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.FinishTime = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetGroupName(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.GroupName = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetHideViewResultButton(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.HideViewResultButton = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetIsCharge(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.IsCharge = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetModelScene(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.ModelScene = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSampleId(v int32) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SampleId = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSampleName(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SampleName = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetServiceCode(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.ServiceCode = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetServiceName(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.ServiceName = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSubTaskId(v int32) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SubTaskId = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSupportCancel(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SupportCancel = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSupportDownload(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SupportDownload = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSupportMergeSelect(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SupportMergeSelect = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetSupportView(v bool) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.SupportView = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetTab(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.Tab = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetTaskGroupId(v int32) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.TaskGroupId = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) SetTaskStatus(v string) *ListTaskGroupResponseBodyResultObjectSubTaskList {
	s.TaskStatus = &v
	return s
}

func (s *ListTaskGroupResponseBodyResultObjectSubTaskList) Validate() error {
	return dara.Validate(s)
}
