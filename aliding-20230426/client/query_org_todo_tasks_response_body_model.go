// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *QueryOrgTodoTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryOrgTodoTasksResponseBody
	GetRequestId() *string
	SetTodoCards(v []*QueryOrgTodoTasksResponseBodyTodoCards) *QueryOrgTodoTasksResponseBody
	GetTodoCards() []*QueryOrgTodoTasksResponseBodyTodoCards
}

type QueryOrgTodoTasksResponseBody struct {
	// example:
	//
	// 15
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TodoCards []*QueryOrgTodoTasksResponseBodyTodoCards `json:"todoCards,omitempty" xml:"todoCards,omitempty" type:"Repeated"`
}

func (s QueryOrgTodoTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgTodoTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryOrgTodoTasksResponseBody) GetTodoCards() []*QueryOrgTodoTasksResponseBodyTodoCards {
	return s.TodoCards
}

func (s *QueryOrgTodoTasksResponseBody) SetNextToken(v string) *QueryOrgTodoTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) SetRequestId(v string) *QueryOrgTodoTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) SetTodoCards(v []*QueryOrgTodoTasksResponseBodyTodoCards) *QueryOrgTodoTasksResponseBody {
	s.TodoCards = v
	return s
}

func (s *QueryOrgTodoTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryOrgTodoTasksResponseBodyTodoCards struct {
	// example:
	//
	// isv_dingtalkTodo
	BizTag *string `json:"bizTag,omitempty" xml:"bizTag,omitempty"`
	// example:
	//
	// 1617675000000
	CreatedTime *int64 `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// xxxx
	CreatorId *string                                          `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	DetailUrl *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl `json:"detailUrl,omitempty" xml:"detailUrl,omitempty" type:"Struct"`
	// example:
	//
	// 1617675000000
	DueTime *int64 `json:"dueTime,omitempty" xml:"dueTime,omitempty"`
	// example:
	//
	// true
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// 1617675000000
	ModifiedTime *int64 `json:"modifiedTime,omitempty" xml:"modifiedTime,omitempty"`
	// example:
	//
	// 10
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// example:
	//
	// isv_dingtalkTodo1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// 接入钉钉待办
	Subject *string `json:"subject,omitempty" xml:"subject,omitempty"`
	// example:
	//
	// taskOPJpwtwPVNGIFKURjrzd
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCards) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetBizTag() *string {
	return s.BizTag
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetCreatorId() *string {
	return s.CreatorId
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetDetailUrl() *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	return s.DetailUrl
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetDueTime() *int64 {
	return s.DueTime
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetIsDone() *bool {
	return s.IsDone
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetPriority() *int32 {
	return s.Priority
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetSourceId() *string {
	return s.SourceId
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetSubject() *string {
	return s.Subject
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetBizTag(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.BizTag = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetCreatorId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.CreatorId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDetailUrl(v *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DetailUrl = v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetDueTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.DueTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetIsDone(v bool) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetModifiedTime(v int64) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.ModifiedTime = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetPriority(v int32) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Priority = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSourceId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.SourceId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetSubject(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.Subject = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) SetTaskId(v string) *QueryOrgTodoTasksResponseBodyTodoCards {
	s.TaskId = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCards) Validate() error {
	return dara.Validate(s)
}

type QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl struct {
	// example:
	//
	// https://www.dingtalk.com
	AppUrl *string `json:"appUrl,omitempty" xml:"appUrl,omitempty"`
	// example:
	//
	// https://www.dingtalk.com
	PcUrl *string `json:"pcUrl,omitempty" xml:"pcUrl,omitempty"`
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) GetAppUrl() *string {
	return s.AppUrl
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) GetPcUrl() *string {
	return s.PcUrl
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetAppUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.AppUrl = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) SetPcUrl(v string) *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl {
	s.PcUrl = &v
	return s
}

func (s *QueryOrgTodoTasksResponseBodyTodoCardsDetailUrl) Validate() error {
	return dara.Validate(s)
}
