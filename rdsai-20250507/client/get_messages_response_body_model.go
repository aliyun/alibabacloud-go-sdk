// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetMessagesResponseBodyData) *GetMessagesResponseBody
	GetData() []*GetMessagesResponseBodyData
	SetHasMore(v bool) *GetMessagesResponseBody
	GetHasMore() *bool
	SetLimit(v int64) *GetMessagesResponseBody
	GetLimit() *int64
	SetRequestId(v string) *GetMessagesResponseBody
	GetRequestId() *string
}

type GetMessagesResponseBody struct {
	// The returned results.
	Data []*GetMessagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether the current page is followed by a page.
	//
	// example:
	//
	// true
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetMessagesResponseBody) GetData() []*GetMessagesResponseBodyData {
	return s.Data
}

func (s *GetMessagesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetMessagesResponseBody) GetLimit() *int64 {
	return s.Limit
}

func (s *GetMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMessagesResponseBody) SetData(v []*GetMessagesResponseBodyData) *GetMessagesResponseBody {
	s.Data = v
	return s
}

func (s *GetMessagesResponseBody) SetHasMore(v bool) *GetMessagesResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetMessagesResponseBody) SetLimit(v int64) *GetMessagesResponseBody {
	s.Limit = &v
	return s
}

func (s *GetMessagesResponseBody) SetRequestId(v string) *GetMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMessagesResponseBody) Validate() error {
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

type GetMessagesResponseBodyData struct {
	// The response to the query.
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The ID of the conversation.
	//
	// example:
	//
	// 9cbbe885-b240-4803-9d15-6781a3fd****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The creation time of the conversation.
	//
	// example:
	//
	// 1763986004
	CreatedAt *string                              `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Events    []*GetMessagesResponseBodyDataEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The feedback.
	//
	// example:
	//
	// like
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// The message ID.
	//
	// example:
	//
	// 84dc9f9b-424a-404d-9c36-35e9d000****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The query statement.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The retriever resources.
	RetrieverResources []interface{} `json:"RetrieverResources,omitempty" xml:"RetrieverResources,omitempty" type:"Repeated"`
}

func (s GetMessagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMessagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMessagesResponseBodyData) GetAnswer() *string {
	return s.Answer
}

func (s *GetMessagesResponseBodyData) GetConversationId() *string {
	return s.ConversationId
}

func (s *GetMessagesResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMessagesResponseBodyData) GetEvents() []*GetMessagesResponseBodyDataEvents {
	return s.Events
}

func (s *GetMessagesResponseBodyData) GetFeedback() *string {
	return s.Feedback
}

func (s *GetMessagesResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetMessagesResponseBodyData) GetQuery() *string {
	return s.Query
}

func (s *GetMessagesResponseBodyData) GetRetrieverResources() []interface{} {
	return s.RetrieverResources
}

func (s *GetMessagesResponseBodyData) SetAnswer(v string) *GetMessagesResponseBodyData {
	s.Answer = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetConversationId(v string) *GetMessagesResponseBodyData {
	s.ConversationId = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetCreatedAt(v string) *GetMessagesResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetEvents(v []*GetMessagesResponseBodyDataEvents) *GetMessagesResponseBodyData {
	s.Events = v
	return s
}

func (s *GetMessagesResponseBodyData) SetFeedback(v string) *GetMessagesResponseBodyData {
	s.Feedback = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetId(v string) *GetMessagesResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetQuery(v string) *GetMessagesResponseBodyData {
	s.Query = &v
	return s
}

func (s *GetMessagesResponseBodyData) SetRetrieverResources(v []interface{}) *GetMessagesResponseBodyData {
	s.RetrieverResources = v
	return s
}

func (s *GetMessagesResponseBodyData) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMessagesResponseBodyDataEvents struct {
	Answer *string `json:"answer,omitempty" xml:"answer,omitempty"`
	Event  *string `json:"event,omitempty" xml:"event,omitempty"`
}

func (s GetMessagesResponseBodyDataEvents) String() string {
	return dara.Prettify(s)
}

func (s GetMessagesResponseBodyDataEvents) GoString() string {
	return s.String()
}

func (s *GetMessagesResponseBodyDataEvents) GetAnswer() *string {
	return s.Answer
}

func (s *GetMessagesResponseBodyDataEvents) GetEvent() *string {
	return s.Event
}

func (s *GetMessagesResponseBodyDataEvents) SetAnswer(v string) *GetMessagesResponseBodyDataEvents {
	s.Answer = &v
	return s
}

func (s *GetMessagesResponseBodyDataEvents) SetEvent(v string) *GetMessagesResponseBodyDataEvents {
	s.Event = &v
	return s
}

func (s *GetMessagesResponseBodyDataEvents) Validate() error {
	return dara.Validate(s)
}
