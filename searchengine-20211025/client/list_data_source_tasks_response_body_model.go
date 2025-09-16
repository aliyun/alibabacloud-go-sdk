// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataSourceTasksResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataSourceTasksResponseBodyResult) *ListDataSourceTasksResponseBody
	GetResult() []*ListDataSourceTasksResponseBodyResult
}

type ListDataSourceTasksResponseBody struct {
	// id of request
	//
	// example:
	//
	// CC5EC8FA-5C0D-56AF-BEF4-6FCCEABD0511
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	Result []*ListDataSourceTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourceTasksResponseBody) GetResult() []*ListDataSourceTasksResponseBodyResult {
	return s.Result
}

func (s *ListDataSourceTasksResponseBody) SetRequestId(v string) *ListDataSourceTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTasksResponseBody) SetResult(v []*ListDataSourceTasksResponseBodyResult) *ListDataSourceTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListDataSourceTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceTasksResponseBodyResult struct {
	// The additional attributes of the card.
	//
	// example:
	//
	// ""
	ExtraAttribute *string `json:"extraAttribute,omitempty" xml:"extraAttribute,omitempty"`
	// The field3 field that was passed when the FSM was created.
	//
	// example:
	//
	// ""
	Field3 *string `json:"field3,omitempty" xml:"field3,omitempty"`
	// The ID of the finite state machine (FSM).
	//
	// example:
	//
	// tisplus_opensearch@datasource_flow_fsm@1062017779051424-ha-cn-2r42ostoc01_ecom_table@vpc_hz_domain_1@null@MANUAL-ha-cn-2r42ostoc01_ecom_table@1655974525756@006754
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
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The FSM status.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags of the progress bar.
	Tags []*ListDataSourceTasksResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The task information.
	TaskNodes []*ListDataSourceTasksResponseBodyResultTaskNodes `json:"taskNodes,omitempty" xml:"taskNodes,omitempty" type:"Repeated"`
	// The timestamp of the card.
	//
	// example:
	//
	// 1646279473
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// The card type.
	//
	// example:
	//
	// search
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The user who triggered the generation of the FSM process.
	//
	// example:
	//
	// ""
	User *string `json:"user,omitempty" xml:"user,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResult) GetExtraAttribute() *string {
	return s.ExtraAttribute
}

func (s *ListDataSourceTasksResponseBodyResult) GetField3() *string {
	return s.Field3
}

func (s *ListDataSourceTasksResponseBodyResult) GetFsmId() *string {
	return s.FsmId
}

func (s *ListDataSourceTasksResponseBodyResult) GetGroupType() *string {
	return s.GroupType
}

func (s *ListDataSourceTasksResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDataSourceTasksResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListDataSourceTasksResponseBodyResult) GetTags() []*ListDataSourceTasksResponseBodyResultTags {
	return s.Tags
}

func (s *ListDataSourceTasksResponseBodyResult) GetTaskNodes() []*ListDataSourceTasksResponseBodyResultTaskNodes {
	return s.TaskNodes
}

func (s *ListDataSourceTasksResponseBodyResult) GetTime() *string {
	return s.Time
}

func (s *ListDataSourceTasksResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDataSourceTasksResponseBodyResult) GetUser() *string {
	return s.User
}

func (s *ListDataSourceTasksResponseBodyResult) SetExtraAttribute(v string) *ListDataSourceTasksResponseBodyResult {
	s.ExtraAttribute = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetField3(v string) *ListDataSourceTasksResponseBodyResult {
	s.Field3 = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetFsmId(v string) *ListDataSourceTasksResponseBodyResult {
	s.FsmId = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetGroupType(v string) *ListDataSourceTasksResponseBodyResult {
	s.GroupType = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetName(v string) *ListDataSourceTasksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetStatus(v string) *ListDataSourceTasksResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTags(v []*ListDataSourceTasksResponseBodyResultTags) *ListDataSourceTasksResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTaskNodes(v []*ListDataSourceTasksResponseBodyResultTaskNodes) *ListDataSourceTasksResponseBodyResult {
	s.TaskNodes = v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetTime(v string) *ListDataSourceTasksResponseBodyResult {
	s.Time = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetType(v string) *ListDataSourceTasksResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) SetUser(v string) *ListDataSourceTasksResponseBodyResult {
	s.User = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceTasksResponseBodyResultTags struct {
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
	// ""
	TagLevel *string `json:"tagLevel,omitempty" xml:"tagLevel,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTags) GetMsg() *string {
	return s.Msg
}

func (s *ListDataSourceTasksResponseBodyResultTags) GetTagLevel() *string {
	return s.TagLevel
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetMsg(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.Msg = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTags) SetTagLevel(v string) *ListDataSourceTasksResponseBodyResultTags {
	s.TagLevel = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}

type ListDataSourceTasksResponseBodyResultTaskNodes struct {
	// The time when the task was complete.
	//
	// example:
	//
	// ""
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
	// ha-cn-7pp2ngv4s02_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task status.
	//
	// example:
	//
	// onlyPublished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTasksResponseBodyResultTaskNodes) GoString() string {
	return s.String()
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) GetFinishDate() *string {
	return s.FinishDate
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) GetIndex() *int64 {
	return s.Index
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) GetName() *string {
	return s.Name
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) GetStatus() *string {
	return s.Status
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetFinishDate(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.FinishDate = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetIndex(v int64) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Index = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetName(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Name = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) SetStatus(v string) *ListDataSourceTasksResponseBodyResultTaskNodes {
	s.Status = &v
	return s
}

func (s *ListDataSourceTasksResponseBodyResultTaskNodes) Validate() error {
	return dara.Validate(s)
}
