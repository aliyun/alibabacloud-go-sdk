// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTaskFlowInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDAGInstanceList(v *ListTaskFlowInstanceResponseBodyDAGInstanceList) *ListTaskFlowInstanceResponseBody
	GetDAGInstanceList() *ListTaskFlowInstanceResponseBodyDAGInstanceList
	SetErrorCode(v string) *ListTaskFlowInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTaskFlowInstanceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListTaskFlowInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTaskFlowInstanceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListTaskFlowInstanceResponseBody
	GetTotalCount() *int32
}

type ListTaskFlowInstanceResponseBody struct {
	// The information about the execution records returned.
	DAGInstanceList *ListTaskFlowInstanceResponseBodyDAGInstanceList `json:"DAGInstanceList,omitempty" xml:"DAGInstanceList,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8CFF2295-8249-5287-B888-DBD4F0D76CB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of execution records returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTaskFlowInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskFlowInstanceResponseBody) GetDAGInstanceList() *ListTaskFlowInstanceResponseBodyDAGInstanceList {
	return s.DAGInstanceList
}

func (s *ListTaskFlowInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTaskFlowInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTaskFlowInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTaskFlowInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTaskFlowInstanceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTaskFlowInstanceResponseBody) SetDAGInstanceList(v *ListTaskFlowInstanceResponseBodyDAGInstanceList) *ListTaskFlowInstanceResponseBody {
	s.DAGInstanceList = v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) SetErrorCode(v string) *ListTaskFlowInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) SetErrorMessage(v string) *ListTaskFlowInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) SetRequestId(v string) *ListTaskFlowInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) SetSuccess(v bool) *ListTaskFlowInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) SetTotalCount(v int32) *ListTaskFlowInstanceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBody) Validate() error {
	if s.DAGInstanceList != nil {
		if err := s.DAGInstanceList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTaskFlowInstanceResponseBodyDAGInstanceList struct {
	DAGInstance []*ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance `json:"DAGInstance,omitempty" xml:"DAGInstance,omitempty" type:"Repeated"`
}

func (s ListTaskFlowInstanceResponseBodyDAGInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowInstanceResponseBodyDAGInstanceList) GoString() string {
	return s.String()
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceList) GetDAGInstance() []*ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	return s.DAGInstance
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceList) SetDAGInstance(v []*ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) *ListTaskFlowInstanceResponseBodyDAGInstanceList {
	s.DAGInstance = v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceList) Validate() error {
	if s.DAGInstance != nil {
		for _, item := range s.DAGInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance struct {
	// The business time of the task flow. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-11-10 14:37:26
	BusinessTime *string `json:"BusinessTime,omitempty" xml:"BusinessTime,omitempty"`
	// The ID of the task flow.
	//
	// example:
	//
	// 7***
	DagId *string `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The name of the task flow.
	//
	// example:
	//
	// Spark_SQL_test
	DagName *string `json:"DagName,omitempty" xml:"DagName,omitempty"`
	// The version of the task flow.
	//
	// example:
	//
	// []
	DagVersion *string `json:"DagVersion,omitempty" xml:"DagVersion,omitempty"`
	// The time when the execution of the task flow was complete. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-11-11 14:38:57
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the previously published version of the task flow.
	//
	// example:
	//
	// 2****
	HistoryDagId *int64 `json:"HistoryDagId,omitempty" xml:"HistoryDagId,omitempty"`
	// The ID of the execution record.
	//
	// example:
	//
	// 9234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the task flow owner.
	//
	// example:
	//
	// test_name
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The status of the task flow. Valid values:
	//
	// 	- **0**: The task flow is waiting to be scheduled.
	//
	// 	- **1**: The task flow is being executed.
	//
	// 	- **2**: The task flow is paused.
	//
	// 	- **3**: The task flow failed.
	//
	// 	- **4**: The task flow is executed.
	//
	// 	- **5**: The task flow is complete.
	//
	// example:
	//
	// 4
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode in which the task flow is triggered. Valid values:
	//
	// 	- **0**: The task flow is automatically triggered based on periodic scheduling.
	//
	// 	- **1**: The task flow is manually triggered.
	//
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The time when the execution of the task flow was start. The time is displayed in the yyyy-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2021-11-11 14:35:57
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) String() string {
	return dara.Prettify(s)
}

func (s ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GoString() string {
	return s.String()
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetBusinessTime() *string {
	return s.BusinessTime
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetDagId() *string {
	return s.DagId
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetDagName() *string {
	return s.DagName
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetDagVersion() *string {
	return s.DagVersion
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetHistoryDagId() *int64 {
	return s.HistoryDagId
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetId() *int64 {
	return s.Id
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetMessage() *string {
	return s.Message
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetStatus() *int32 {
	return s.Status
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) GetStartTime() *string {
	return s.StartTime
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetBusinessTime(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.BusinessTime = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetDagId(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.DagId = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetDagName(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.DagName = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetDagVersion(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.DagVersion = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetEndTime(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.EndTime = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetHistoryDagId(v int64) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.HistoryDagId = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetId(v int64) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.Id = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetMessage(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.Message = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetOwnerName(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.OwnerName = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetStatus(v int32) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.Status = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetTriggerType(v int32) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.TriggerType = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) SetStartTime(v string) *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance {
	s.StartTime = &v
	return s
}

func (s *ListTaskFlowInstanceResponseBodyDAGInstanceListDAGInstance) Validate() error {
	return dara.Validate(s)
}
