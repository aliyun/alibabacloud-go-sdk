// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBaselineKeyPathResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetBaselineKeyPathResponseBodyData) *GetBaselineKeyPathResponseBody
	GetData() []*GetBaselineKeyPathResponseBodyData
	SetErrorCode(v string) *GetBaselineKeyPathResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBaselineKeyPathResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBaselineKeyPathResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBaselineKeyPathResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBaselineKeyPathResponseBody
	GetSuccess() *bool
}

type GetBaselineKeyPathResponseBody struct {
	// The information about the key path.
	Data []*GetBaselineKeyPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Error code
	//
	// example:
	//
	// 1031203110005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The timestamp when the event was found.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique ID of the call. After an error occurs, you can troubleshoot the problem based on the ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBaselineKeyPathResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathResponseBody) GetData() []*GetBaselineKeyPathResponseBodyData {
	return s.Data
}

func (s *GetBaselineKeyPathResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBaselineKeyPathResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBaselineKeyPathResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBaselineKeyPathResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBaselineKeyPathResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBaselineKeyPathResponseBody) SetData(v []*GetBaselineKeyPathResponseBodyData) *GetBaselineKeyPathResponseBody {
	s.Data = v
	return s
}

func (s *GetBaselineKeyPathResponseBody) SetErrorCode(v string) *GetBaselineKeyPathResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBaselineKeyPathResponseBody) SetErrorMessage(v string) *GetBaselineKeyPathResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBaselineKeyPathResponseBody) SetHttpStatusCode(v int32) *GetBaselineKeyPathResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBaselineKeyPathResponseBody) SetRequestId(v string) *GetBaselineKeyPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBody) SetSuccess(v bool) *GetBaselineKeyPathResponseBody {
	s.Success = &v
	return s
}

func (s *GetBaselineKeyPathResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetBaselineKeyPathResponseBodyData struct {
	// The data timestamp of the instance.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The ID of the scheduling cycle of the instance. Valid values: 1 to 288.
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 12345
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The node ID.
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
	// 9527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The type of the node. Valid values: 23, 10, 6, and 99. The value 23 indicates that the node is a Data Integration node. The value 10 indicates that the node is a MaxCompute SQL node. The value 6 indicates that the node is a Shell node. The value 99 indicates that the node is a zero load node.
	//
	// example:
	//
	// 10
	PrgType *int32 `json:"PrgType,omitempty" xml:"PrgType,omitempty"`
	// The ID of the workspace to which the node belongs.
	//
	// example:
	//
	// 1234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The running records of the instance.
	Runs []*GetBaselineKeyPathResponseBodyDataRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The information about the events that are associated with the instance.
	Topics []*GetBaselineKeyPathResponseBodyDataTopics `json:"Topics,omitempty" xml:"Topics,omitempty" type:"Repeated"`
}

func (s GetBaselineKeyPathResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathResponseBodyData) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetBaselineKeyPathResponseBodyData) GetInGroupId() *int32 {
	return s.InGroupId
}

func (s *GetBaselineKeyPathResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetBaselineKeyPathResponseBodyData) GetNodeId() *int64 {
	return s.NodeId
}

func (s *GetBaselineKeyPathResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *GetBaselineKeyPathResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBaselineKeyPathResponseBodyData) GetPrgType() *int32 {
	return s.PrgType
}

func (s *GetBaselineKeyPathResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBaselineKeyPathResponseBodyData) GetRuns() []*GetBaselineKeyPathResponseBodyDataRuns {
	return s.Runs
}

func (s *GetBaselineKeyPathResponseBodyData) GetTopics() []*GetBaselineKeyPathResponseBodyDataTopics {
	return s.Topics
}

func (s *GetBaselineKeyPathResponseBodyData) SetBizdate(v int64) *GetBaselineKeyPathResponseBodyData {
	s.Bizdate = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetInGroupId(v int32) *GetBaselineKeyPathResponseBodyData {
	s.InGroupId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetInstanceId(v int64) *GetBaselineKeyPathResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetNodeId(v int64) *GetBaselineKeyPathResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetNodeName(v string) *GetBaselineKeyPathResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetOwner(v string) *GetBaselineKeyPathResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetPrgType(v int32) *GetBaselineKeyPathResponseBodyData {
	s.PrgType = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetProjectId(v int64) *GetBaselineKeyPathResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetRuns(v []*GetBaselineKeyPathResponseBodyDataRuns) *GetBaselineKeyPathResponseBodyData {
	s.Runs = v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) SetTopics(v []*GetBaselineKeyPathResponseBodyDataTopics) *GetBaselineKeyPathResponseBodyData {
	s.Topics = v
	return s
}

func (s *GetBaselineKeyPathResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetBaselineKeyPathResponseBodyDataRuns struct {
	// The timestamp obtained by adding the predicted time when the instance started to run to the historical average running duration of the instance.
	//
	// example:
	//
	// 1553531402000
	AbsTime *int64 `json:"AbsTime,omitempty" xml:"AbsTime,omitempty"`
	// The timestamp of the predicted time when the instance started to run.
	//
	// example:
	//
	// 1553531686000
	BeginCast *int64 `json:"BeginCast,omitempty" xml:"BeginCast,omitempty"`
	// The timestamp of the actual time when the instance started to run.
	//
	// example:
	//
	// 1553531401000
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The timestamp when the instance started to wait for resources.
	//
	// example:
	//
	// 1553531401000
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The timestamp when the instance started to wait for the scheduling time.
	//
	// example:
	//
	// 1553531400000
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The timestamp of the predicted time when the instance finished running.
	//
	// example:
	//
	// 1553531687000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp of the actual time when the instance finished running.
	//
	// example:
	//
	// 1553531401000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The status of the instance. Valid values: NOT_RUN, WAIT_TIME, WAIT_RESOURCE, RUNNING, CHECKING, CHECKING_CONDITION, FAILURE, and SUCCESS. The value NOT_RUN indicates that the instance is not run. The value WAIT_TIME indicates that the instance is waiting to be run. The value WAIT_RESOURCE indicates that the instance is waiting for resources. The value RUNNING indicates that the instance is running. The value CHECKING indicates that data quality is being checked for the instance. The value CHECKING_CONDITION indicates that branch conditions are being checked for the instance. The value FAILURE indicates that the instance fails to run. The value SUCCESS indicates that the instance is run.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBaselineKeyPathResponseBodyDataRuns) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathResponseBodyDataRuns) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetAbsTime() *int64 {
	return s.AbsTime
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetBeginCast() *int64 {
	return s.BeginCast
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetBeginRunningTime() *int64 {
	return s.BeginRunningTime
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetBeginWaitResTime() *int64 {
	return s.BeginWaitResTime
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetBeginWaitTimeTime() *int64 {
	return s.BeginWaitTimeTime
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetEndCast() *int64 {
	return s.EndCast
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) GetStatus() *string {
	return s.Status
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetAbsTime(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.AbsTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetBeginCast(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.BeginCast = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetBeginRunningTime(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.BeginRunningTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetBeginWaitResTime(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.BeginWaitResTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetBeginWaitTimeTime(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.BeginWaitTimeTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetEndCast(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.EndCast = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetFinishTime(v int64) *GetBaselineKeyPathResponseBodyDataRuns {
	s.FinishTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) SetStatus(v string) *GetBaselineKeyPathResponseBodyDataRuns {
	s.Status = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataRuns) Validate() error {
	return dara.Validate(s)
}

type GetBaselineKeyPathResponseBodyDataTopics struct {
	// The timestamp when the event was found.
	//
	// example:
	//
	// 1553531401000
	AddTime *int64 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// 1234
	TopicName *string `json:"TopicName,omitempty" xml:"TopicName,omitempty"`
}

func (s GetBaselineKeyPathResponseBodyDataTopics) String() string {
	return dara.Prettify(s)
}

func (s GetBaselineKeyPathResponseBodyDataTopics) GoString() string {
	return s.String()
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) GetAddTime() *int64 {
	return s.AddTime
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) GetTopicId() *int64 {
	return s.TopicId
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) GetTopicName() *string {
	return s.TopicName
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) SetAddTime(v int64) *GetBaselineKeyPathResponseBodyDataTopics {
	s.AddTime = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) SetInstanceId(v int64) *GetBaselineKeyPathResponseBodyDataTopics {
	s.InstanceId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) SetTopicId(v int64) *GetBaselineKeyPathResponseBodyDataTopics {
	s.TopicId = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) SetTopicName(v string) *GetBaselineKeyPathResponseBodyDataTopics {
	s.TopicName = &v
	return s
}

func (s *GetBaselineKeyPathResponseBodyDataTopics) Validate() error {
	return dara.Validate(s)
}
