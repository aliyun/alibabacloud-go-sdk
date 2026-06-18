// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataImportTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeDataImportTaskInfoResponseBody
	GetCode() *int64
	SetData(v *DescribeDataImportTaskInfoResponseBodyData) *DescribeDataImportTaskInfoResponseBody
	GetData() *DescribeDataImportTaskInfoResponseBodyData
	SetMessage(v string) *DescribeDataImportTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDataImportTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDataImportTaskInfoResponseBody
	GetSuccess() *string
}

type DescribeDataImportTaskInfoResponseBody struct {
	// The return code. This parameter is empty when the request succeeds. When the request fails, an exception message such as an error code is returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result set.
	Data *DescribeDataImportTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message. This parameter has a value only when the task status is success. Otherwise, an empty value is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B2F3840-5C98-****-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDataImportTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeDataImportTaskInfoResponseBody) GetData() *DescribeDataImportTaskInfoResponseBodyData {
	return s.Data
}

func (s *DescribeDataImportTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDataImportTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataImportTaskInfoResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDataImportTaskInfoResponseBody) SetCode(v int64) *DescribeDataImportTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBody) SetData(v *DescribeDataImportTaskInfoResponseBodyData) *DescribeDataImportTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBody) SetMessage(v string) *DescribeDataImportTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBody) SetRequestId(v string) *DescribeDataImportTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBody) SetSuccess(v string) *DescribeDataImportTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataImportTaskInfoResponseBodyData struct {
	// The task details.
	DataImportTaskDetailInfo *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo `json:"DataImportTaskDetailInfo,omitempty" xml:"DataImportTaskDetailInfo,omitempty" type:"Struct"`
}

func (s DescribeDataImportTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponseBodyData) GetDataImportTaskDetailInfo() *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo {
	return s.DataImportTaskDetailInfo
}

func (s *DescribeDataImportTaskInfoResponseBodyData) SetDataImportTaskDetailInfo(v *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) *DescribeDataImportTaskInfoResponseBodyData {
	s.DataImportTaskDetailInfo = v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyData) Validate() error {
	if s.DataImportTaskDetailInfo != nil {
		if err := s.DataImportTaskDetailInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo struct {
	// The progress ID.
	//
	// example:
	//
	// 1
	FsmId *int64 `json:"FsmId,omitempty" xml:"FsmId,omitempty"`
	// The state identifier in a data migration or synchronization task.
	//
	// example:
	//
	// RECON_FINISHED_CATCH_UP
	FsmState *string `json:"FsmState,omitempty" xml:"FsmState,omitempty"`
	// The status in a data migration, import, or synchronization system.
	//
	// example:
	//
	// IMPORT_NOT_BEGIN
	FsmStatus *string `json:"FsmStatus,omitempty" xml:"FsmStatus,omitempty"`
	// The data import task details.
	ServiceDetailList []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList `json:"ServiceDetailList,omitempty" xml:"ServiceDetailList,omitempty" type:"Repeated"`
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) GetFsmId() *int64 {
	return s.FsmId
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) GetFsmState() *string {
	return s.FsmState
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) GetFsmStatus() *string {
	return s.FsmStatus
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) GetServiceDetailList() []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList {
	return s.ServiceDetailList
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) SetFsmId(v int64) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo {
	s.FsmId = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) SetFsmState(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo {
	s.FsmState = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) SetFsmStatus(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo {
	s.FsmStatus = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) SetServiceDetailList(v []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo {
	s.ServiceDetailList = v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfo) Validate() error {
	if s.ServiceDetailList != nil {
		for _, item := range s.ServiceDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList struct {
	// The service detail ID.
	//
	// example:
	//
	// 1991609
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task execution status.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task details.
	TaskDetailList []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList `json:"TaskDetailList,omitempty" xml:"TaskDetailList,omitempty" type:"Repeated"`
	// Valid values:
	//
	// - FULL_COPY: full replication.
	//
	// - INC_COPY: incremental replication.
	//
	// example:
	//
	// FULL_COPY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) GetTaskDetailList() []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	return s.TaskDetailList
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) GetType() *string {
	return s.Type
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) SetId(v int64) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList {
	s.Id = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) SetStatus(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList {
	s.Status = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) SetTaskDetailList(v []*DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList {
	s.TaskDetailList = v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) SetType(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList {
	s.Type = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailList) Validate() error {
	if s.TaskDetailList != nil {
		for _, item := range s.TaskDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList struct {
	// The delay time.
	//
	// example:
	//
	// 58329
	Delay *int64 `json:"Delay,omitempty" xml:"Delay,omitempty"`
	// The most recent error message.
	//
	// example:
	//
	// null
	LastError *string `json:"LastError,omitempty" xml:"LastError,omitempty"`
	// The physical database name.
	//
	// example:
	//
	// drds_test
	PhysicalDbName *string `json:"PhysicalDbName,omitempty" xml:"PhysicalDbName,omitempty"`
	// The data import progress.
	//
	// example:
	//
	// 0
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The performance and runtime metrics collected during the execution of the data migration or import task.
	//
	// example:
	//
	// {
	//
	//   "applyCount": 0,
	//
	//   "cpuUseRatio": 41,
	//
	//   "fsmId": 1,
	//
	//   "fullGcCount":
	//
	// }
	Statistics *string `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The task status.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task ID.
	//
	// example:
	//
	// fc3b1568-ad96-****-adca-dfe018b38077
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// example:
	//
	// FULL_COPY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetDelay() *int64 {
	return s.Delay
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetLastError() *string {
	return s.LastError
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetPhysicalDbName() *string {
	return s.PhysicalDbName
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetProgress() *int64 {
	return s.Progress
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetStatistics() *string {
	return s.Statistics
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) GetType() *string {
	return s.Type
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetDelay(v int64) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.Delay = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetLastError(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.LastError = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetPhysicalDbName(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.PhysicalDbName = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetProgress(v int64) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.Progress = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetStatistics(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.Statistics = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetStatus(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.Status = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetTaskId(v int64) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.TaskId = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) SetType(v string) *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList {
	s.Type = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponseBodyDataDataImportTaskDetailInfoServiceDetailListTaskDetailList) Validate() error {
	return dara.Validate(s)
}
