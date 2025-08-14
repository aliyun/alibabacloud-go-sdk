// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForwardAIAgentCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ForwardAIAgentCallResponseBody
	GetRequestId() *string
}

type ForwardAIAgentCallResponseBody struct {
	// example:
	//
	// 550e8400********55440000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ForwardAIAgentCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ForwardAIAgentCallResponseBody) GoString() string {
	return s.String()
}

func (s *ForwardAIAgentCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ForwardAIAgentCallResponseBody) SetRequestId(v string) *ForwardAIAgentCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *ForwardAIAgentCallResponseBody) Validate() error {
	return dara.Validate(s)
}
