// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAgentResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAgentResponseBody
	GetRequestId() *string
}

type DeleteAgentResponseBody struct {
	// example:
	//
	// AgentNotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Agent with name \\"xxx\\" not found for account 1186xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B896B484-F16D-59DE-9E23-DD0E5C361108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgentResponseBody) SetCode(v string) *DeleteAgentResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAgentResponseBody) SetMessage(v string) *DeleteAgentResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAgentResponseBody) SetRequestId(v string) *DeleteAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
