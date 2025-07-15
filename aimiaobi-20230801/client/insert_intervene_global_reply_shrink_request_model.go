// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertInterveneGlobalReplyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *InsertInterveneGlobalReplyShrinkRequest
	GetAgentKey() *string
	SetReplyMessagListShrink(v string) *InsertInterveneGlobalReplyShrinkRequest
	GetReplyMessagListShrink() *string
}

type InsertInterveneGlobalReplyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxx_efm
	AgentKey              *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	ReplyMessagListShrink *string `json:"ReplyMessagList,omitempty" xml:"ReplyMessagList,omitempty"`
}

func (s InsertInterveneGlobalReplyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertInterveneGlobalReplyShrinkRequest) GoString() string {
	return s.String()
}

func (s *InsertInterveneGlobalReplyShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *InsertInterveneGlobalReplyShrinkRequest) GetReplyMessagListShrink() *string {
	return s.ReplyMessagListShrink
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetAgentKey(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *InsertInterveneGlobalReplyShrinkRequest) SetReplyMessagListShrink(v string) *InsertInterveneGlobalReplyShrinkRequest {
	s.ReplyMessagListShrink = &v
	return s
}

func (s *InsertInterveneGlobalReplyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
