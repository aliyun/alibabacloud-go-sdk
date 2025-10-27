// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationProcessDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *ListOperationProcessDetailResponseBodyPageInfo) *ListOperationProcessDetailResponseBody
	GetPageInfo() *ListOperationProcessDetailResponseBodyPageInfo
	SetProcessDetails(v []*ListOperationProcessDetailResponseBodyProcessDetails) *ListOperationProcessDetailResponseBody
	GetProcessDetails() []*ListOperationProcessDetailResponseBodyProcessDetails
	SetRequestId(v string) *ListOperationProcessDetailResponseBody
	GetRequestId() *string
}

type ListOperationProcessDetailResponseBody struct {
	// The pagination information.
	PageInfo *ListOperationProcessDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The information about the operation subtasks.
	ProcessDetails []*ListOperationProcessDetailResponseBodyProcessDetails `json:"ProcessDetails,omitempty" xml:"ProcessDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CE500770-42D3-442E-9DDD-156E0F9F3***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListOperationProcessDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailResponseBody) GetPageInfo() *ListOperationProcessDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListOperationProcessDetailResponseBody) GetProcessDetails() []*ListOperationProcessDetailResponseBodyProcessDetails {
	return s.ProcessDetails
}

func (s *ListOperationProcessDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOperationProcessDetailResponseBody) SetPageInfo(v *ListOperationProcessDetailResponseBodyPageInfo) *ListOperationProcessDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListOperationProcessDetailResponseBody) SetProcessDetails(v []*ListOperationProcessDetailResponseBodyProcessDetails) *ListOperationProcessDetailResponseBody {
	s.ProcessDetails = v
	return s
}

func (s *ListOperationProcessDetailResponseBody) SetRequestId(v string) *ListOperationProcessDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOperationProcessDetailResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ProcessDetails != nil {
		for _, item := range s.ProcessDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationProcessDetailResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOperationProcessDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) SetCount(v int32) *ListOperationProcessDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) SetCurrentPage(v int32) *ListOperationProcessDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) SetPageSize(v int32) *ListOperationProcessDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) SetTotalCount(v int32) *ListOperationProcessDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListOperationProcessDetailResponseBodyProcessDetails struct {
	// The subtype of the asset associated with the operation subtask.
	//
	// example:
	//
	// 1
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the asset associated with the operation subtask.
	//
	// example:
	//
	// 8
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The vendor of the asset associated with the operation subtask.
	//
	// example:
	//
	// 0
	AssetVendor *int32 `json:"AssetVendor,omitempty" xml:"AssetVendor,omitempty"`
	// The check items associated with the operation subtask.
	Checks []*ListOperationProcessDetailResponseBodyProcessDetailsChecks `json:"Checks,omitempty" xml:"Checks,omitempty" type:"Repeated"`
	// The timestamp when the task was created. Unit: milliseconds.
	//
	// example:
	//
	// 1706544199000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the operation subtask.
	//
	// example:
	//
	// fb4bcd41-a916-46bc-ab1a-65fd383be***
	DetailTaskId *string `json:"DetailTaskId,omitempty" xml:"DetailTaskId,omitempty"`
	// The end timestamp of the operation subtask. Unit: milliseconds.
	//
	// example:
	//
	// 1706544199000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start timestamp of the operation subtask. Unit: milliseconds.
	//
	// example:
	//
	// 1730335622000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The subtask status code. Enumerated values:
	//
	// 	- 0: not started.
	//
	// 	- 1: running.
	//
	// 	- 2: successful.
	//
	// 	- 3: times out.
	//
	// 	- 4: failed.
	//
	// example:
	//
	// 0
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The ID of the operation subtask.
	//
	// example:
	//
	// v34578b8-e567-47ec-2345-3e5b077ca***
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListOperationProcessDetailResponseBodyProcessDetails) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailResponseBodyProcessDetails) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetAssetVendor() *int32 {
	return s.AssetVendor
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetChecks() []*ListOperationProcessDetailResponseBodyProcessDetailsChecks {
	return s.Checks
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetDetailTaskId() *string {
	return s.DetailTaskId
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetAssetSubType(v int32) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.AssetSubType = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetAssetType(v int32) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.AssetType = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetAssetVendor(v int32) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.AssetVendor = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetChecks(v []*ListOperationProcessDetailResponseBodyProcessDetailsChecks) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.Checks = v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetCreateTime(v int64) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.CreateTime = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetDetailTaskId(v string) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.DetailTaskId = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetEndTime(v int64) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.EndTime = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetStartTime(v int64) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.StartTime = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetStatusCode(v int32) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.StatusCode = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) SetTaskId(v string) *ListOperationProcessDetailResponseBodyProcessDetails {
	s.TaskId = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetails) Validate() error {
	if s.Checks != nil {
		for _, item := range s.Checks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOperationProcessDetailResponseBodyProcessDetailsChecks struct {
	// The ID of the check item associated with the operation subtask.
	//
	// example:
	//
	// 133
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The name of the check item associated with the operation subtask.
	//
	// example:
	//
	// Check for Security Center Agent Status
	CheckShowName *string `json:"CheckShowName,omitempty" xml:"CheckShowName,omitempty"`
}

func (s ListOperationProcessDetailResponseBodyProcessDetailsChecks) String() string {
	return dara.Prettify(s)
}

func (s ListOperationProcessDetailResponseBodyProcessDetailsChecks) GoString() string {
	return s.String()
}

func (s *ListOperationProcessDetailResponseBodyProcessDetailsChecks) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListOperationProcessDetailResponseBodyProcessDetailsChecks) GetCheckShowName() *string {
	return s.CheckShowName
}

func (s *ListOperationProcessDetailResponseBodyProcessDetailsChecks) SetCheckId(v int64) *ListOperationProcessDetailResponseBodyProcessDetailsChecks {
	s.CheckId = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetailsChecks) SetCheckShowName(v string) *ListOperationProcessDetailResponseBodyProcessDetailsChecks {
	s.CheckShowName = &v
	return s
}

func (s *ListOperationProcessDetailResponseBodyProcessDetailsChecks) Validate() error {
	return dara.Validate(s)
}
