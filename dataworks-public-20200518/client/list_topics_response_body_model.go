// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody
	GetData() *ListTopicsResponseBodyData
	SetErrorCode(v string) *ListTopicsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListTopicsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListTopicsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListTopicsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTopicsResponseBody
	GetSuccess() *bool
}

type ListTopicsResponseBody struct {
	// The information about the events returned.
	Data *ListTopicsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListTopicsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBody) GetData() *ListTopicsResponseBodyData {
	return s.Data
}

func (s *ListTopicsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListTopicsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListTopicsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTopicsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTopicsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTopicsResponseBody) SetData(v *ListTopicsResponseBodyData) *ListTopicsResponseBody {
	s.Data = v
	return s
}

func (s *ListTopicsResponseBody) SetErrorCode(v string) *ListTopicsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetErrorMessage(v string) *ListTopicsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListTopicsResponseBody) SetHttpStatusCode(v int32) *ListTopicsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTopicsResponseBody) SetRequestId(v string) *ListTopicsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTopicsResponseBody) SetSuccess(v bool) *ListTopicsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTopicsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTopicsResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The events returned.
	Topics []*ListTopicsResponseBodyDataTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
	// The total number of the events returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTopicsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTopicsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTopicsResponseBodyData) GetTopics() []*ListTopicsResponseBodyDataTopics {
	return s.Topics
}

func (s *ListTopicsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTopicsResponseBodyData) SetPageNumber(v int32) *ListTopicsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetPageSize(v int32) *ListTopicsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListTopicsResponseBodyData) SetTopics(v []*ListTopicsResponseBodyDataTopics) *ListTopicsResponseBodyData {
	s.Topics = v
	return s
}

func (s *ListTopicsResponseBodyData) SetTotalCount(v int32) *ListTopicsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListTopicsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTopicsResponseBodyDataTopics struct {
	// The timestamp when the event was found.
	//
	// example:
	//
	// 1553524393000
	AddTime *int64 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The timestamp when the event was processed.
	//
	// example:
	//
	// 1553508465000
	FixTime *int64 `json:"FixTime,omitempty" xml:"FixTime,omitempty"`
	// The timestamp when the event occurred. A time difference may exist between the time when the event occurred and the time when the event was found.
	//
	// example:
	//
	// 1553508465000
	HappenTime *int64 `json:"HappenTime,omitempty" xml:"HappenTime,omitempty"`
	// The ID of the node instance that triggers the event.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node that triggers the event.
	//
	// example:
	//
	// 1234
	NodeId *int64 `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// Node Name
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The ID of the Alibaba Cloud account that is used by the node owner.
	//
	// example:
	//
	// 952795****
	NodeOwner *string `json:"NodeOwner,omitempty" xml:"NodeOwner,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 1234
	TopicId *int64 `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
	// The name of the event.
	//
	// example:
	//
	// 1234 Error
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
	// The status of the event. Valid values: IGNORE, NEW, FIXING, and RECOVER. The value IGNORE indicates that the event is ignored. The value NEW indicates that the event is a new event. The value FIXING indicates that the event is being processed. The value RECOVER indicates that the event is processed.
	//
	// example:
	//
	// NEW
	TopicStatus *string `json:"TopicStatus,omitempty" xml:"TopicStatus,omitempty"`
	// The type of the event. Valid values: SLOW and ERROR. The value SLOW indicates that the running duration of the node in the current scheduling cycle is significantly longer than the average running duration of the node in previous scheduling cycles. The value ERROR indicates that the node fails to run.
	//
	// example:
	//
	// ERROR
	TopicType *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s ListTopicsResponseBodyDataTopics) String() string {
	return dara.Prettify(s)
}

func (s ListTopicsResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *ListTopicsResponseBodyDataTopics) GetAddTime() *int64 {
	return s.AddTime
}

func (s *ListTopicsResponseBodyDataTopics) GetFixTime() *int64 {
	return s.FixTime
}

func (s *ListTopicsResponseBodyDataTopics) GetHappenTime() *int64 {
	return s.HappenTime
}

func (s *ListTopicsResponseBodyDataTopics) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListTopicsResponseBodyDataTopics) GetNodeId() *int64 {
	return s.NodeId
}

func (s *ListTopicsResponseBodyDataTopics) GetNodeName() *string {
	return s.NodeName
}

func (s *ListTopicsResponseBodyDataTopics) GetNodeOwner() *string {
	return s.NodeOwner
}

func (s *ListTopicsResponseBodyDataTopics) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTopicsResponseBodyDataTopics) GetTopicId() *int64 {
	return s.TopicId
}

func (s *ListTopicsResponseBodyDataTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *ListTopicsResponseBodyDataTopics) GetTopicStatus() *string {
	return s.TopicStatus
}

func (s *ListTopicsResponseBodyDataTopics) GetTopicType() *string {
	return s.TopicType
}

func (s *ListTopicsResponseBodyDataTopics) SetAddTime(v int64) *ListTopicsResponseBodyDataTopics {
	s.AddTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetFixTime(v int64) *ListTopicsResponseBodyDataTopics {
	s.FixTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetHappenTime(v int64) *ListTopicsResponseBodyDataTopics {
	s.HappenTime = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetInstanceId(v int64) *ListTopicsResponseBodyDataTopics {
	s.InstanceId = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetNodeId(v int64) *ListTopicsResponseBodyDataTopics {
	s.NodeId = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetNodeName(v string) *ListTopicsResponseBodyDataTopics {
	s.NodeName = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetNodeOwner(v string) *ListTopicsResponseBodyDataTopics {
	s.NodeOwner = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetProjectId(v int64) *ListTopicsResponseBodyDataTopics {
	s.ProjectId = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetTopicId(v int64) *ListTopicsResponseBodyDataTopics {
	s.TopicId = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetTopicName(v string) *ListTopicsResponseBodyDataTopics {
	s.TopicName = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetTopicStatus(v string) *ListTopicsResponseBodyDataTopics {
	s.TopicStatus = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) SetTopicType(v string) *ListTopicsResponseBodyDataTopics {
	s.TopicType = &v
	return s
}

func (s *ListTopicsResponseBodyDataTopics) Validate() error {
	return dara.Validate(s)
}
