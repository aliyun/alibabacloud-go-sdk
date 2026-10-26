// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentConversationTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v int64) *GetAgentConversationTimelineRequest
	GetAfter() *int64
	SetIncludeImageData(v bool) *GetAgentConversationTimelineRequest
	GetIncludeImageData() *bool
}

type GetAgentConversationTimelineRequest struct {
	// example:
	//
	// 42
	After *int64 `json:"after,omitempty" xml:"after,omitempty"`
	// example:
	//
	// false
	IncludeImageData *bool `json:"includeImageData,omitempty" xml:"includeImageData,omitempty"`
}

func (s GetAgentConversationTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentConversationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetAgentConversationTimelineRequest) GetAfter() *int64 {
	return s.After
}

func (s *GetAgentConversationTimelineRequest) GetIncludeImageData() *bool {
	return s.IncludeImageData
}

func (s *GetAgentConversationTimelineRequest) SetAfter(v int64) *GetAgentConversationTimelineRequest {
	s.After = &v
	return s
}

func (s *GetAgentConversationTimelineRequest) SetIncludeImageData(v bool) *GetAgentConversationTimelineRequest {
	s.IncludeImageData = &v
	return s
}

func (s *GetAgentConversationTimelineRequest) Validate() error {
	return dara.Validate(s)
}
