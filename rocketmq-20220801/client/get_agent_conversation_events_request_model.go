// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v int64) *GetAgentConversationEventsRequest
	GetAfter() *int64
}

type GetAgentConversationEventsRequest struct {
	// example:
	//
	// 42
	After *int64 `json:"after,omitempty" xml:"after,omitempty"`
}

func (s GetAgentConversationEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationEventsRequest) GoString() string {
	return s.String()
}

func (s *GetAgentConversationEventsRequest) GetAfter() *int64 {
	return s.After
}

func (s *GetAgentConversationEventsRequest) SetAfter(v int64) *GetAgentConversationEventsRequest {
	s.After = &v
	return s
}

func (s *GetAgentConversationEventsRequest) Validate() error {
	return dara.Validate(s)
}
