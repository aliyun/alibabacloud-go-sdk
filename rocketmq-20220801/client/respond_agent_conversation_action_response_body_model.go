// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondAgentConversationActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccepted(v bool) *RespondAgentConversationActionResponseBody
	GetAccepted() *bool
	SetContinuationCreated(v bool) *RespondAgentConversationActionResponseBody
	GetContinuationCreated() *bool
	SetRunId(v string) *RespondAgentConversationActionResponseBody
	GetRunId() *string
}

type RespondAgentConversationActionResponseBody struct {
	// example:
	//
	// true
	Accepted *bool `json:"accepted,omitempty" xml:"accepted,omitempty"`
	// example:
	//
	// true
	ContinuationCreated *bool `json:"continuationCreated,omitempty" xml:"continuationCreated,omitempty"`
	// example:
	//
	// 0198c8a91f4d7f01a3c21e98
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
}

func (s RespondAgentConversationActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RespondAgentConversationActionResponseBody) GoString() string {
	return s.String()
}

func (s *RespondAgentConversationActionResponseBody) GetAccepted() *bool {
	return s.Accepted
}

func (s *RespondAgentConversationActionResponseBody) GetContinuationCreated() *bool {
	return s.ContinuationCreated
}

func (s *RespondAgentConversationActionResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *RespondAgentConversationActionResponseBody) SetAccepted(v bool) *RespondAgentConversationActionResponseBody {
	s.Accepted = &v
	return s
}

func (s *RespondAgentConversationActionResponseBody) SetContinuationCreated(v bool) *RespondAgentConversationActionResponseBody {
	s.ContinuationCreated = &v
	return s
}

func (s *RespondAgentConversationActionResponseBody) SetRunId(v string) *RespondAgentConversationActionResponseBody {
	s.RunId = &v
	return s
}

func (s *RespondAgentConversationActionResponseBody) Validate() error {
	return dara.Validate(s)
}
