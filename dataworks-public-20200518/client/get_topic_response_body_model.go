// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTopicResponseBodyData) *GetTopicResponseBody
	GetData() *GetTopicResponseBodyData
	SetErrorCode(v string) *GetTopicResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTopicResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetTopicResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTopicResponseBody
	GetSuccess() *bool
}

type GetTopicResponseBody struct {
	// The details of the event.
	Data *GetTopicResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFGH-IJKLMNOPQ
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBody) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBody) GetData() *GetTopicResponseBodyData {
	return s.Data
}

func (s *GetTopicResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTopicResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTopicResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTopicResponseBody) SetData(v *GetTopicResponseBodyData) *GetTopicResponseBody {
	s.Data = v
	return s
}

func (s *GetTopicResponseBody) SetErrorCode(v string) *GetTopicResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTopicResponseBody) SetErrorMessage(v string) *GetTopicResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTopicResponseBody) SetHttpStatusCode(v int32) *GetTopicResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTopicResponseBody) SetRequestId(v string) *GetTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTopicResponseBody) SetSuccess(v bool) *GetTopicResponseBody {
	s.Success = &v
	return s
}

func (s *GetTopicResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTopicResponseBodyData struct {
	// The timestamp when the event was found.
	//
	// example:
	//
	// 1553524393000
	AddTime *int64 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The timestamp when the first alert was reported.
	//
	// example:
	//
	// 1553524393000
	AlertTime *int64 `json:"AlertTime,omitempty" xml:"AlertTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the alert recipient.
	//
	// example:
	//
	// 952795****
	Assigner *string `json:"Assigner,omitempty" xml:"Assigner,omitempty"`
	// The margin of the worst baseline instance. Unit: seconds.
	//
	// example:
	//
	// 3600
	BaselineBuffer *int64 `json:"BaselineBuffer,omitempty" xml:"BaselineBuffer,omitempty"`
	// The ID of the baseline to which the worst baseline instance belongs.
	//
	// example:
	//
	// 1234
	BaselineId *int64 `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The ID of the cycle of the worst baseline instance.
	//
	// example:
	//
	// 1
	BaselineInGroupId *int32 `json:"BaselineInGroupId,omitempty" xml:"BaselineInGroupId,omitempty"`
	// The name of the baseline to which the worst baseline instance belongs.
	//
	// example:
	//
	// Baseline name
	BaselineName *string `json:"BaselineName,omitempty" xml:"BaselineName,omitempty"`
	// The status of the baseline. Valid values: ERROR, SAFE, DANGROUS, and OVER. The value ERROR indicates that no nodes are associated with the baseline, or all nodes associated with the baseline are suspended. The value SAFE indicates that nodes are run before the alert duration begins. The value DANGROUS indicates that nodes are still running after the alert duration ends but the committed completion time does not arrive. The value OVER indicates that nodes are still running after the committed completion time.
	//
	// example:
	//
	// SAFE
	BaselineStatus *string `json:"BaselineStatus,omitempty" xml:"BaselineStatus,omitempty"`
	// The margin of the event. Unit: seconds.
	//
	// example:
	//
	// 1200
	Buffer *int64 `json:"Buffer,omitempty" xml:"Buffer,omitempty"`
	// The timestamp when the event was last processed.
	//
	// example:
	//
	// 1553524393000
	DealTime *int64 `json:"DealTime,omitempty" xml:"DealTime,omitempty"`
	// The ID of the Alibaba Cloud account used by the user who last processed the event.
	//
	// example:
	//
	// 952795****
	DealUser *string `json:"DealUser,omitempty" xml:"DealUser,omitempty"`
	// The timestamp when the event was processed.
	//
	// example:
	//
	// 1553524393000
	FixTime *int64 `json:"FixTime,omitempty" xml:"FixTime,omitempty"`
	// The timestamp when the event occurred. A time difference may exist between the time when the event occurred and the time when the event was found.
	//
	// example:
	//
	// 1553524393000
	HappenTime *int64 `json:"HappenTime,omitempty" xml:"HappenTime,omitempty"`
	// The ID of the instance that triggered the event.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The timestamp when the system reports the next alert.
	//
	// example:
	//
	// 1553524393000
	NextAlertTime *int64 `json:"NextAlertTime,omitempty" xml:"NextAlertTime,omitempty"`
	// The ID of the node that triggered the event.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node that triggered the event.
	//
	// example:
	//
	// Node name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account used by the event owner.
	//
	// example:
	//
	// 952795****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace to which the node that triggered the event belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// 1234 error
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The status of the event. Valid values: IGNORE, NEW, FIXING, and RECOVER.
	//
	// example:
	//
	// FIXING
	TopicStatus *string `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty"`
	// The type of the event. Valid values: SLOW and ERROR. The value SLOW indicates that the duration of the task is significantly longer than the average duration of the task in previous cycles. The value ERROR indicates that the task fails to run.
	//
	// example:
	//
	// ERROR
	TopicType *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s GetTopicResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTopicResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTopicResponseBodyData) GetAddTime() *int64 {
	return s.AddTime
}

func (s *GetTopicResponseBodyData) GetAlertTime() *int64 {
	return s.AlertTime
}

func (s *GetTopicResponseBodyData) GetAssigner() *string {
	return s.Assigner
}

func (s *GetTopicResponseBodyData) GetBaselineBuffer() *int64 {
	return s.BaselineBuffer
}

func (s *GetTopicResponseBodyData) GetBaselineId() *int64 {
	return s.BaselineId
}

func (s *GetTopicResponseBodyData) GetBaselineInGroupId() *int32 {
	return s.BaselineInGroupId
}

func (s *GetTopicResponseBodyData) GetBaselineName() *string {
	return s.BaselineName
}

func (s *GetTopicResponseBodyData) GetBaselineStatus() *string {
	return s.BaselineStatus
}

func (s *GetTopicResponseBodyData) GetBuffer() *int64 {
	return s.Buffer
}

func (s *GetTopicResponseBodyData) GetDealTime() *int64 {
	return s.DealTime
}

func (s *GetTopicResponseBodyData) GetDealUser() *string {
	return s.DealUser
}

func (s *GetTopicResponseBodyData) GetFixTime() *int64 {
	return s.FixTime
}

func (s *GetTopicResponseBodyData) GetHappenTime() *int64 {
	return s.HappenTime
}

func (s *GetTopicResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetTopicResponseBodyData) GetNextAlertTime() *int64 {
	return s.NextAlertTime
}

func (s *GetTopicResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetTopicResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetTopicResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetTopicResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTopicResponseBodyData) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetTopicResponseBodyData) GetTopicName() *string {
	return s.TopicName
}

func (s *GetTopicResponseBodyData) GetTopicStatus() *string {
	return s.TopicStatus
}

func (s *GetTopicResponseBodyData) GetTopicType() *string {
	return s.TopicType
}

func (s *GetTopicResponseBodyData) SetAddTime(v int64) *GetTopicResponseBodyData {
	s.AddTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetAlertTime(v int64) *GetTopicResponseBodyData {
	s.AlertTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetAssigner(v string) *GetTopicResponseBodyData {
	s.Assigner = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBaselineBuffer(v int64) *GetTopicResponseBodyData {
	s.BaselineBuffer = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBaselineId(v int64) *GetTopicResponseBodyData {
	s.BaselineId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBaselineInGroupId(v int32) *GetTopicResponseBodyData {
	s.BaselineInGroupId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBaselineName(v string) *GetTopicResponseBodyData {
	s.BaselineName = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBaselineStatus(v string) *GetTopicResponseBodyData {
	s.BaselineStatus = &v
	return s
}

func (s *GetTopicResponseBodyData) SetBuffer(v int64) *GetTopicResponseBodyData {
	s.Buffer = &v
	return s
}

func (s *GetTopicResponseBodyData) SetDealTime(v int64) *GetTopicResponseBodyData {
	s.DealTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetDealUser(v string) *GetTopicResponseBodyData {
	s.DealUser = &v
	return s
}

func (s *GetTopicResponseBodyData) SetFixTime(v int64) *GetTopicResponseBodyData {
	s.FixTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetHappenTime(v int64) *GetTopicResponseBodyData {
	s.HappenTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetInstanceId(v int64) *GetTopicResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetNextAlertTime(v int64) *GetTopicResponseBodyData {
	s.NextAlertTime = &v
	return s
}

func (s *GetTopicResponseBodyData) SetNodeId(v int64) *GetTopicResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetNodeName(v string) *GetTopicResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetTopicResponseBodyData) SetOwner(v string) *GetTopicResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetTopicResponseBodyData) SetProjectId(v int64) *GetTopicResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicId(v int64) *GetTopicResponseBodyData {
	s.TopicId = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicName(v string) *GetTopicResponseBodyData {
	s.TopicName = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicStatus(v string) *GetTopicResponseBodyData {
	s.TopicStatus = &v
	return s
}

func (s *GetTopicResponseBodyData) SetTopicType(v string) *GetTopicResponseBodyData {
	s.TopicType = &v
	return s
}

func (s *GetTopicResponseBodyData) Validate() error {
	return dara.Validate(s)
}
