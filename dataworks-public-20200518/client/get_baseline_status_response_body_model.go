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
	// The error code.
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique ID of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// The business date timestamp.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The information about the critical instance.
	BlockInstance *GetBaselineStatusResponseBodyDataBlockInstance `json:"BlockInstance,omitempty" xml:"BlockInstance,omitempty" type:"Struct"`
	// The buffer time of the baseline instance, in seconds.
	//
	// example:
	//
	// 1200
	Buffer *float32 `json:"Buffer,omitempty" xml:"Buffer,omitempty"`
	// The estimated completion timestamp of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The warning timestamp of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	ExpTime *int64 `json:"ExpTime,omitempty" xml:"ExpTime,omitempty"`
	// Indicates whether the baseline instance is completed. Valid values: UNFINISH and FINISH.
	//
	// example:
	//
	// UNFINISH
	FinishStatus *string `json:"FinishStatus,omitempty" xml:"FinishStatus,omitempty"`
	// The completion timestamp of the baseline instance. This parameter is returned only when FinishStatus is FINISH.
	//
	// example:
	//
	// 1553443200000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The cycle number of the baseline instance. The value is 1 for daily baselines. The value ranges from [1,24\\] for hourly baselines.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The information about the latest instance.
	LastInstance *GetBaselineStatusResponseBodyDataLastInstance `json:"LastInstance,omitempty" xml:"LastInstance,omitempty" type:"Struct"`
	// The Alibaba Cloud UID of the baseline owner. Multiple owners are separated by commas (,).
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
	// The committed completion timestamp of the baseline instance.
	//
	// example:
	//
	// 1553443200000
	SlaTime *int64 `json:"SlaTime,omitempty" xml:"SlaTime,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGROUS (warning), and OVER (exceeded).
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
	if s.BlockInstance != nil {
		if err := s.BlockInstance.Validate(); err != nil {
			return err
		}
	}
	if s.LastInstance != nil {
		if err := s.LastInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBaselineStatusResponseBodyDataBlockInstance struct {
	// The estimated completion timestamp of the instance.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The actual completion timestamp of the instance.
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
	// NodeName
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The Alibaba Cloud UID of the node owner.
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
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS.
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
	// The estimated completion timestamp of the instance.
	//
	// example:
	//
	// 1553443200000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The actual completion timestamp of the instance.
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
	// The Alibaba Cloud UID of the node owner.
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
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS.
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
