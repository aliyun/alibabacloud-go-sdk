// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBaselineStatusResponseBodyData) *GetBaselineStatusResponseBody
	GetData() *GetBaselineStatusResponseBodyData
	SetErrorCode(v string) *GetBaselineStatusResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBaselineStatusResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBaselineStatusResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBaselineStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaselineStatusResponseBody
	GetSuccess() *bool
}

type GetBaselineStatusResponseBody struct {
	// The details of the baseline instance.
	Data *GetBaselineStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusResponseBody) GetData() *GetBaselineStatusResponseBodyData {
	return s.Data
}

func (s *GetBaselineStatusResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBaselineStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBaselineStatusResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaselineStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaselineStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaselineStatusResponseBody) SetData(v *GetBaselineStatusResponseBodyData) *GetBaselineStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineStatusResponseBody) SetErrorCode(v string) *GetBaselineStatusResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBaselineStatusResponseBody) SetErrorMessage(v string) *GetBaselineStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBaselineStatusResponseBody) SetHttpStatusCode(v int32) *GetBaselineStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineStatusResponseBody) SetRequestId(v string) *GetBaselineStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineStatusResponseBody) SetSuccess(v bool) *GetBaselineStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaselineStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBaselineStatusResponseBodyData struct {
	// The ID of the baseline.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The data timestamp of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The information about the key instance.
	BlockInstance *GetBaselineStatusResponseBodyDataBlockInstance `json:"BlockInstance,omitempty" xml:"BlockInstance,omitempty" type:"Struct"`
	// The margin of the baseline instance. Unit: seconds.
	//
	// example:
	//
	// 1200
	Buffer *float32 `json:"Buffer,omitempty" xml:"Buffer,omitempty"`
	// The timestamp of the predicted time when the baseline instance finished running.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp of the alerting time of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	ExpTime *int64 `json:"ExpTime,omitempty" xml:"ExpTime,omitempty"`
	// The status of the baseline instance. Valid values: UNFINISH and FINISH. The value UNFINISH indicates that the baseline instance is still running. The value FINISH indicates that the baseline instance finishes running.
	//
	// example:
	//
	// UNFINISH
	FinishStatus *string `json:"FinishStatus,omitempty" xml:"FinishStatus,omitempty"`
	// The timestamp of the actual time when the baseline instance finished running. This parameter is returned if the value of the FinishStatus parameter is FINISH.
	//
	// example:
	//
	// 1553443200000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the scheduling cycle of the baseline instance. For a baseline instance that is scheduled by day, the value of this parameter is 1. For a baseline instance that is scheduled by hour, the value of this parameter ranges from 1 to 24.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The information about the last generated instance.
	LastInstance *GetBaselineStatusResponseBodyDataLastInstance `json:"LastInstance,omitempty" xml:"LastInstance,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account used by the baseline owner. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 9527952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The priority of the baseline. Valid values: 1, 2, 5, 7, and 8.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the workspace to which the baseline belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The timestamp of the committed completion time of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	SlaTime *int64 `json:"SlaTime,omitempty" xml:"SlaTime,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGEROUS, and OVER. The value ERROR indicates that no nodes are associated with the baseline, or all nodes associated with the baseline are suspended. The value SAFE indicates that nodes finish running before the alerting time. The value DANGEROUS indicates that nodes are still running after the alerting time but before the committed completion time. The value OVER indicates that nodes are still running after the committed completion time.
	//
	// example:
	//
	// SAFE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBaselineStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetBaselineStatusResponseBodyData) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetBaselineStatusResponseBodyData) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetBaselineStatusResponseBodyData) GetBlockInstance() *GetBaselineStatusResponseBodyDataBlockInstance {
	return s.BlockInstance
}

func (s *GetBaselineStatusResponseBodyData) GetBuffer() *float32 {
	return s.Buffer
}

func (s *GetBaselineStatusResponseBodyData) GetEndCast() *int64 {
	return s.EndCast
}

func (s *GetBaselineStatusResponseBodyData) GetExpTime() *int64 {
	return s.ExpTime
}

func (s *GetBaselineStatusResponseBodyData) GetFinishStatus() *string {
	return s.FinishStatus
}

func (s *GetBaselineStatusResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetBaselineStatusResponseBodyData) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetBaselineStatusResponseBodyData) GetLastInstance() *GetBaselineStatusResponseBodyDataLastInstance {
	return s.LastInstance
}

func (s *GetBaselineStatusResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineStatusResponseBodyData) GetPriority() *int32 {
	return s.Priority
}

func (s *GetBaselineStatusResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineStatusResponseBodyData) GetSlaTime() *int64 {
	return s.SlaTime
}

func (s *GetBaselineStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetBaselineStatusResponseBodyData) SetBaselineId(v int64) *GetBaselineStatusResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetBaselineName(v string) *GetBaselineStatusResponseBodyData {
	s.BaselineName = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetBizdate(v int64) *GetBaselineStatusResponseBodyData {
	s.Bizdate = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetBlockInstance(v *GetBaselineStatusResponseBodyDataBlockInstance) *GetBaselineStatusResponseBodyData {
	s.BlockInstance = v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetBuffer(v float32) *GetBaselineStatusResponseBodyData {
	s.Buffer = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetEndCast(v int64) *GetBaselineStatusResponseBodyData {
	s.EndCast = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetExpTime(v int64) *GetBaselineStatusResponseBodyData {
	s.ExpTime = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetFinishStatus(v string) *GetBaselineStatusResponseBodyData {
	s.FinishStatus = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetFinishTime(v int64) *GetBaselineStatusResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetInGroupId(v int32) *GetBaselineStatusResponseBodyData {
	s.InGroupId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetLastInstance(v *GetBaselineStatusResponseBodyDataLastInstance) *GetBaselineStatusResponseBodyData {
	s.LastInstance = v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetOwner(v string) *GetBaselineStatusResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetPriority(v int32) *GetBaselineStatusResponseBodyData {
	s.Priority = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetProjectId(v int64) *GetBaselineStatusResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetSlaTime(v int64) *GetBaselineStatusResponseBodyData {
	s.SlaTime = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) SetStatus(v string) *GetBaselineStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetBaselineStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetBaselineStatusResponseBodyDataBlockInstance struct {
	// The timestamp of the predicted time when the instance finished running.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp of the actual time when the instance finished running.
	//
	// example:
	//
	// 1553443200000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 9527952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS. The value NOT_RUN indicates that the instance is not run. The value WAIT_TIME indicates that the instance is waiting to be run. The value WAIT_RESOURCE indicates that the instance is waiting for resources. The value RUNNING indicates that the instance is running. The value CHECKING indicates that data quality is being checked for the instance. The value CHECKING_CONDITION indicates that branch conditions are being checked for the instance. The value FAILURE indicates that the instance fails to run. The value SUCCESS indicates that the instance is run.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBaselineStatusResponseBodyDataBlockInstance) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusResponseBodyDataBlockInstance) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetEndCast() *int64 {
	return s.EndCast
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetNodeName() *string {
	return s.NodeName
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) GetStatus() *string {
	return s.Status
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetEndCast(v int64) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.EndCast = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetFinishTime(v int64) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.FinishTime = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetInstanceId(v int64) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.InstanceId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetNodeId(v int64) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.NodeId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetNodeName(v string) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.NodeName = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetOwner(v string) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.Owner = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetProjectId(v int64) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) SetStatus(v string) *GetBaselineStatusResponseBodyDataBlockInstance {
	s.Status = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataBlockInstance) Validate() error {
	return dara.Validate(s)
}

type GetBaselineStatusResponseBodyDataLastInstance struct {
	// The timestamp of the predicted time when the instance finished running.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp of the actual time when the instance finished running.
	//
	// example:
	//
	// 1553443200000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the node owner.
	//
	// example:
	//
	// 9527952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS. The value NOT_RUN indicates that the instance is not run. The value WAIT_TIME indicates that the instance is waiting to be run. The value WAIT_RESOURCE indicates that the instance is waiting for resources. The value RUNNING indicates that the instance is running. The value CHECKING indicates that data quality is being checked for the instance. The value CHECKING_CONDITION indicates that branch conditions are being checked for the instance. The value FAILURE indicates that the instance fails to run. The value SUCCESS indicates that the instance is run.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBaselineStatusResponseBodyDataLastInstance) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineStatusResponseBodyDataLastInstance) GoString() string {
	return s.String()
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetEndCast() *int64 {
	return s.EndCast
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetNodeName() *string {
	return s.NodeName
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) GetStatus() *string {
	return s.Status
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetEndCast(v int64) *GetBaselineStatusResponseBodyDataLastInstance {
	s.EndCast = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetFinishTime(v int64) *GetBaselineStatusResponseBodyDataLastInstance {
	s.FinishTime = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetInstanceId(v int64) *GetBaselineStatusResponseBodyDataLastInstance {
	s.InstanceId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetNodeId(v int64) *GetBaselineStatusResponseBodyDataLastInstance {
	s.NodeId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetNodeName(v string) *GetBaselineStatusResponseBodyDataLastInstance {
	s.NodeName = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetOwner(v string) *GetBaselineStatusResponseBodyDataLastInstance {
	s.Owner = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetProjectId(v int64) *GetBaselineStatusResponseBodyDataLastInstance {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) SetStatus(v string) *GetBaselineStatusResponseBodyDataLastInstance {
	s.Status = &v
	return s
}

func (s *GetBaselineStatusResponseBodyDataLastInstance) Validate() error {
	return dara.Validate(s)
}
