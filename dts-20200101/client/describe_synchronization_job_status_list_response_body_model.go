// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSynchronizationJobStatusListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DescribeSynchronizationJobStatusListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSynchronizationJobStatusListResponseBody
	GetErrMessage() *string
	SetPageNumber(v int32) *DescribeSynchronizationJobStatusListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeSynchronizationJobStatusListResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeSynchronizationJobStatusListResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSynchronizationJobStatusListResponseBody
	GetSuccess() *string
	SetSynchronizationJobListStatusList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) *DescribeSynchronizationJobStatusListResponseBody
	GetSynchronizationJobListStatusList() []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList
	SetTotalRecordCount(v int64) *DescribeSynchronizationJobStatusListResponseBody
	GetTotalRecordCount() *int64
}

type DescribeSynchronizationJobStatusListResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of data synchronization instances displayed on one page.
	//
	// example:
	//
	// 2
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1413460B-138A-48D1-836C-B24EDDC1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The status of the data synchronization tasks.
	SynchronizationJobListStatusList []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList `json:"SynchronizationJobListStatusList,omitempty" xml:"SynchronizationJobListStatusList,omitempty" type:"Repeated"`
	// The total number of data synchronization instances.
	//
	// example:
	//
	// 2
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetSynchronizationJobListStatusList() []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList {
	return s.SynchronizationJobListStatusList
}

func (s *DescribeSynchronizationJobStatusListResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSynchronizationJobListStatusList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) *DescribeSynchronizationJobStatusListResponseBody {
	s.SynchronizationJobListStatusList = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobStatusListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) Validate() error {
	if s.SynchronizationJobListStatusList != nil {
		for _, item := range s.SynchronizationJobListStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList struct {
	// The details of data synchronization tasks in each direction.
	SynchronizationDirectionInfoList []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList `json:"SynchronizationDirectionInfoList,omitempty" xml:"SynchronizationDirectionInfoList,omitempty" type:"Repeated"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtsexjk1alb116****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) GetSynchronizationDirectionInfoList() []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	return s.SynchronizationDirectionInfoList
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) SetSynchronizationDirectionInfoList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList {
	s.SynchronizationDirectionInfoList = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) Validate() error {
	if s.SynchronizationDirectionInfoList != nil {
		for _, item := range s.SynchronizationDirectionInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList struct {
	// The UNIX timestamp generated when the latest data record was synchronized.
	//
	// >  You can use a search engine to obtain a UNIX timestamp converter.
	//
	// example:
	//
	// 1610524452
	Checkpoint *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	// The status of the data synchronization task in this direction. Valid values:
	//
	// 	- **NotStarted**: The task is not started.
	//
	// 	- **Prechecking**: The task is being prechecked.
	//
	// 	- **PrecheckFailed**: The task failed to pass the precheck.
	//
	// 	- **Initializing**: The task is performing initial synchronization.
	//
	// 	- **InitializeFailed**: Initial synchronization failed.
	//
	// 	- **Synchronizing**: The task is synchronizing data.
	//
	// 	- **Failed**: The task failed to synchronize data.
	//
	// 	- **Suspending**: The task is paused.
	//
	// 	- **Modifying**: The objects in the task are being modified.
	//
	// 	- **Finished**: The task is completed.
	//
	// example:
	//
	// InitializeFailed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- **Forward**
	//
	// 	- **Reverse**
	//
	// example:
	//
	// Forward
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) GetCheckpoint() *string {
	return s.Checkpoint
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) GetStatus() *string {
	return s.Status
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) GetSynchronizationDirection() *string {
	return s.SynchronizationDirection
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetCheckpoint(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetStatus(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) Validate() error {
	return dara.Validate(s)
}
