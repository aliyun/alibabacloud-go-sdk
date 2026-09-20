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
	// The critical path information.
	Data []*GetBaselineKeyPathResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// Indicates whether the request was successful.
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
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBaselineKeyPathResponseBodyData struct {
	// The timestamp of the business date of the instance.
	//
	// example:
	//
	// 1553443200000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The cycle number of the instance. Valid values: [1,288\\].
	//
	// example:
	//
	// 1
	InGroupId *int32 `json:"InGroupId,omitempty" xml:"InGroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 123456
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
	// 9527952****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The node type. Common node types include Data Integration (23), MaxCompute SQL (10), Shell (6), and virtual node (99).
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
	// The run records of the instance.
	Runs []*GetBaselineKeyPathResponseBodyDataRuns `json:"Runs,omitempty" xml:"Runs,omitempty" type:"Repeated"`
	// The event information associated with the instance.
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
	if s.Runs != nil {
		for _, item := range s.Runs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Topics != nil {
		for _, item := range s.Topics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBaselineKeyPathResponseBodyDataRuns struct {
	// The timestamp calculated by adding the historical average run duration to the estimated start time of the instance.
	//
	// example:
	//
	// 1553531402000
	AbsTime *int64 `json:"AbsTime,omitempty" xml:"AbsTime,omitempty"`
	// The estimated start time of the instance.
	//
	// example:
	//
	// 1553531686000
	BeginCast *int64 `json:"BeginCast,omitempty" xml:"BeginCast,omitempty"`
	// The timestamp when the instance actually started running.
	//
	// example:
	//
	// 1553531401000
	BeginRunningTime *int64 `json:"BeginRunningTime,omitempty" xml:"BeginRunningTime,omitempty"`
	// The timestamp when the instance entered the waiting-for-resources state.
	//
	// example:
	//
	// 1553531401000
	BeginWaitResTime *int64 `json:"BeginWaitResTime,omitempty" xml:"BeginWaitResTime,omitempty"`
	// The timestamp when the instance entered the waiting-for-time state.
	//
	// example:
	//
	// 1553531400000
	BeginWaitTimeTime *int64 `json:"BeginWaitTimeTime,omitempty" xml:"BeginWaitTimeTime,omitempty"`
	// The estimated end time of the instance.
	//
	// example:
	//
	// 1553531687000
	EndCast *int64 `json:"EndCast,omitempty" xml:"EndCast,omitempty"`
	// The timestamp when the instance actually finished running.
	//
	// example:
	//
	// 1553531401000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The status of the instance. Valid values:
	//
	// - NOT_RUN: not run.
	//
	// - WAIT_TIME: waiting for the scheduled time.
	//
	// - WAIT_RESOURCE: waiting for resources.
	//
	// - RUNNING: running.
	//
	// - CHECKING: checking.
	//
	// - CHECKING_CONDITION: checking conditions.
	//
	// - FAILURE: failed.
	//
	// - SUCCESS: succeeded.
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
	// The timestamp when the event was detected.
	//
	// example:
	//
	// 1553531401000
	AddTime *int64 `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
