// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOk(v bool) *DeleteAgentConversationResponseBody
	GetOk() *bool
}

type DeleteAgentConversationResponseBody struct {
	// example:
	//
	// true
	Ok *bool `json:"ok,omitempty" xml:"ok,omitempty"`
}

func (s DeleteAgentConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentConversationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentConversationResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *DeleteAgentConversationResponseBody) SetOk(v bool) *DeleteAgentConversationResponseBody {
	s.Ok = &v
	return s
}

func (s *DeleteAgentConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
