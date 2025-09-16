// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListClusterTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListClusterTasksResponseBodyResult) *ListClusterTasksResponseBody
	GetResult() []*ListClusterTasksResponseBodyResult
}

type ListClusterTasksResponseBody struct {
	// id of request
	//
	// example:
	//
	// CC5EC8FA-5C0D-56AF-BEF4-6FCCEABD0511
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	Result []*ListClusterTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListClusterTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterTasksResponseBody) GetResult() []*ListClusterTasksResponseBodyResult {
	return s.Result
}

func (s *ListClusterTasksResponseBody) SetRequestId(v string) *ListClusterTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterTasksResponseBody) SetResult(v []*ListClusterTasksResponseBodyResult) *ListClusterTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListClusterTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterTasksResponseBodyResult struct {
	// The additional attributes of the card.
	//
	// example:
	//
	// " "
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	// The field3 field that was passed when the FSM was created.
	//
	// example:
	//
	// " "
	Field3 *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// The ID of the finite state machine (FSM).
	//
	// example:
	//
	// tisplus_opensearch@datasource_flow_fsm@1865410598556969-ha-cn-zvp2ljiwe01_api2@bj_vpc_domain_1@null@MANUAL-ha-cn-zvp2ljiwe01_api2@1649729867698@028315
	FsmId *string `json:"fsmId,omitempty" xml:"fsmId,omitempty"`
	// The change group type.
	//
	// example:
	//
	// " "
	GroupType *string `json:"groupType,omitempty" xml:"groupType,omitempty"`
	// The card name.
	//
	// example:
	//
	// ha-cn-pl32rf0js04_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The FSM status.
	//
	// example:
	//
	// onlyPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags of the progress bar.
	Tags []*ListClusterTasksResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The task information.
	TaskNodes []*ListClusterTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	// The timestamp of the card.
	//
	// example:
	//
	// 1657610520
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The card type.
	//
	// example:
	//
	// qrs
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The user who triggered the generation of the FSM process.
	//
	// example:
	//
	// " "
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListClusterTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResult) GetExtraAttribute() *string {
	return s.ExtraAttribute
}

func (s *ListClusterTasksResponseBodyResult) GetField3() *string {
	return s.Field3
}

func (s *ListClusterTasksResponseBodyResult) GetFsmId() *string {
	return s.FsmId
}

func (s *ListClusterTasksResponseBodyResult) GetGroupType() *string {
	return s.GroupType
}

func (s *ListClusterTasksResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListClusterTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListClusterTasksResponseBodyResult) GetTags() []*ListClusterTasksResponseBodyResultTags {
	return s.Tags
}

func (s *ListClusterTasksResponseBodyResult) GetTaskNodes() []*ListClusterTasksResponseBodyResultTaskNodes {
	return s.TaskNodes
}

func (s *ListClusterTasksResponseBodyResult) GetTime() *string {
	return s.Time
}

func (s *ListClusterTasksResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListClusterTasksResponseBodyResult) GetUser() *string {
	return s.User
}

func (s *ListClusterTasksResponseBodyResult) SetExtraAttribute(v string) *ListClusterTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetField3(v string) *ListClusterTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetFsmId(v string) *ListClusterTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetGroupType(v string) *ListClusterTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetName(v string) *ListClusterTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetStatus(v string) *ListClusterTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTags(v []*ListClusterTasksResponseBodyResultTags) *ListClusterTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTaskNodes(v []*ListClusterTasksResponseBodyResultTaskNodes) *ListClusterTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetTime(v string) *ListClusterTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetType(v string) *ListClusterTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) SetUser(v string) *ListClusterTasksResponseBodyResult {
	s.User = &v
	return s
}

func (s *ListClusterTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListClusterTasksResponseBodyResultTags struct {
	// The tag content.
	//
	// example:
	//
	// succeed in handling request
	Msg *string `json:"msg,omitempty" xml:"msg,omitempty"`
	// The tag level.
	//
	// example:
	//
	// " "
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTags) GetMsg() *string {
	return s.Msg
}

func (s *ListClusterTasksResponseBodyResultTags) GetTagLevel() *string {
	return s.TagLevel
}

func (s *ListClusterTasksResponseBodyResultTags) SetMsg(v string) *ListClusterTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTags) SetTagLevel(v string) *ListClusterTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

type ListClusterTasksResponseBodyResultTaskNodes struct {
	// The time when the task was complete.
	//
	// example:
	//
	// " "
	FinishDate *string `json:"finishDate,omitempty" xml:"finishDate,omitempty"`
	// The ordinal number of the task.
	//
	// example:
	//
	// 100
	Index *int64 `json:"index,omitempty" xml:"index,omitempty"`
	// The task name.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task status.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListClusterTasksResponseBodyResultTaskNodes) String() string {
	return dara.Prettify(s)
}

func (s ListClusterTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) GetFinishDate() *string {
	return s.FinishDate
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) GetIndex() *int64 {
	return s.Index
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) GetName() *string {
	return s.Name
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) GetStatus() *string {
	return s.Status
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetName(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListClusterTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

func (s *ListClusterTasksResponseBodyResultTaskNodes) Validate() error {
	return dara.Validate(s)
}
