// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationTimelineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v interface{}) *GetAgentConversationTimelineResponseBody
	GetItems() interface{}
	SetReplayCursor(v int64) *GetAgentConversationTimelineResponseBody
	GetReplayCursor() *int64
	SetResponseFeedback(v interface{}) *GetAgentConversationTimelineResponseBody
	GetResponseFeedback() interface{}
	SetRuns(v interface{}) *GetAgentConversationTimelineResponseBody
	GetRuns() interface{}
	SetSource(v string) *GetAgentConversationTimelineResponseBody
	GetSource() *string
}

type GetAgentConversationTimelineResponseBody struct {
	// example:
	//
	// [{"seq":1,"type":"user"}]
	Items interface{} `json:"items,omitempty" xml:"items,omitempty"`
	// example:
	//
	// 42
	ReplayCursor *int64 `json:"replayCursor,omitempty" xml:"replayCursor,omitempty"`
	// example:
	//
	// {"run_01":"up"}
	ResponseFeedback interface{} `json:"responseFeedback,omitempty" xml:"responseFeedback,omitempty"`
	// example:
	//
	// [{"id":"run_01","modelId":"qwen3-coder"}]
	Runs interface{} `json:"runs,omitempty" xml:"runs,omitempty"`
	// example:
	//
	// events
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s GetAgentConversationTimelineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentConversationTimelineResponseBody) GetItems() interface{} {
	return s.Items
}

func (s *GetAgentConversationTimelineResponseBody) GetReplayCursor() *int64 {
	return s.ReplayCursor
}

func (s *GetAgentConversationTimelineResponseBody) GetResponseFeedback() interface{} {
	return s.ResponseFeedback
}

func (s *GetAgentConversationTimelineResponseBody) GetRuns() interface{} {
	return s.Runs
}

func (s *GetAgentConversationTimelineResponseBody) GetSource() *string {
	return s.Source
}

func (s *GetAgentConversationTimelineResponseBody) SetItems(v interface{}) *GetAgentConversationTimelineResponseBody {
	s.Items = v
	return s
}

func (s *GetAgentConversationTimelineResponseBody) SetReplayCursor(v int64) *GetAgentConversationTimelineResponseBody {
	s.ReplayCursor = &v
	return s
}

func (s *GetAgentConversationTimelineResponseBody) SetResponseFeedback(v interface{}) *GetAgentConversationTimelineResponseBody {
	s.ResponseFeedback = v
	return s
}

func (s *GetAgentConversationTimelineResponseBody) SetRuns(v interface{}) *GetAgentConversationTimelineResponseBody {
	s.Runs = v
	return s
}

func (s *GetAgentConversationTimelineResponseBody) SetSource(v string) *GetAgentConversationTimelineResponseBody {
	s.Source = &v
	return s
}

func (s *GetAgentConversationTimelineResponseBody) Validate() error {
	return dara.Validate(s)
}
