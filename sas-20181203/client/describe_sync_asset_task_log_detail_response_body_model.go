// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncAssetTaskLogDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) *DescribeSyncAssetTaskLogDetailResponseBody
	GetPageInfo() *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo
	SetRequestId(v string) *DescribeSyncAssetTaskLogDetailResponseBody
	GetRequestId() *string
	SetTaskRecordDetails(v []*DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) *DescribeSyncAssetTaskLogDetailResponseBody
	GetTaskRecordDetails() []*DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails
}

type DescribeSyncAssetTaskLogDetailResponseBody struct {
	// The pagination information.
	PageInfo *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the tasks.
	TaskRecordDetails []*DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails `json:"TaskRecordDetails,omitempty" xml:"TaskRecordDetails,omitempty" type:"Repeated"`
}

func (s DescribeSyncAssetTaskLogDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskLogDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) GetPageInfo() *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) GetTaskRecordDetails() []*DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	return s.TaskRecordDetails
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) SetPageInfo(v *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) *DescribeSyncAssetTaskLogDetailResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) SetRequestId(v string) *DescribeSyncAssetTaskLogDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) SetTaskRecordDetails(v []*DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) *DescribeSyncAssetTaskLogDetailResponseBody {
	s.TaskRecordDetails = v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TaskRecordDetails != nil {
		for _, item := range s.TaskRecordDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyncAssetTaskLogDetailResponseBodyPageInfo struct {
	// The number of entries on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) SetCount(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) SetPageSize(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails struct {
	// The total number of assets.
	//
	// example:
	//
	// 5
	AssetCount *int32 `json:"AssetCount,omitempty" xml:"AssetCount,omitempty"`
	// The region of the server in the data center.
	//
	// example:
	//
	// cn-shanghai
	IdcRegion *string `json:"IdcRegion,omitempty" xml:"IdcRegion,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 6c4e4c36ffc3e5919120b405c2b3****
	LeafTaskId *string `json:"LeafTaskId,omitempty" xml:"LeafTaskId,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **INIT**: The task is not started.
	//
	// 	- **START**: The task is started.
	//
	// 	- **MESSAGE_SEND**: The command is sent.
	//
	// 	- **SUCCESS**: The task is complete.
	//
	// 	- **FAIL**: The task failed.
	//
	// 	- **TIMEOUT**: The task timed out.
	//
	// example:
	//
	// INIT
	LeafTaskStatus *string `json:"LeafTaskStatus,omitempty" xml:"LeafTaskStatus,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// unknown reason
	TaskMsg *string `json:"TaskMsg,omitempty" xml:"TaskMsg,omitempty"`
	// The timestamp when the task results were reported.
	//
	// example:
	//
	// 1671614217000
	TaskReportTime *int64 `json:"TaskReportTime,omitempty" xml:"TaskReportTime,omitempty"`
	// The number of unprotected assets.
	//
	// example:
	//
	// 0
	UnprotectedAssetCount *int32 `json:"UnprotectedAssetCount,omitempty" xml:"UnprotectedAssetCount,omitempty"`
}

func (s DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GoString() string {
	return s.String()
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetAssetCount() *int32 {
	return s.AssetCount
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetIdcRegion() *string {
	return s.IdcRegion
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetLeafTaskId() *string {
	return s.LeafTaskId
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetLeafTaskStatus() *string {
	return s.LeafTaskStatus
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetTaskMsg() *string {
	return s.TaskMsg
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetTaskReportTime() *int64 {
	return s.TaskReportTime
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) GetUnprotectedAssetCount() *int32 {
	return s.UnprotectedAssetCount
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetAssetCount(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.AssetCount = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetIdcRegion(v string) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.IdcRegion = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetLeafTaskId(v string) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.LeafTaskId = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetLeafTaskStatus(v string) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.LeafTaskStatus = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetTaskMsg(v string) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.TaskMsg = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetTaskReportTime(v int64) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.TaskReportTime = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) SetUnprotectedAssetCount(v int32) *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails {
	s.UnprotectedAssetCount = &v
	return s
}

func (s *DescribeSyncAssetTaskLogDetailResponseBodyTaskRecordDetails) Validate() error {
	return dara.Validate(s)
}
