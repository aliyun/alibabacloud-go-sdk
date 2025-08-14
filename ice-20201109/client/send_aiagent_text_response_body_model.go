// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendAIAgentTextResponseBody
	GetRequestId() *string
}

type SendAIAgentTextResponseBody struct {
	// example:
	//
	// DB488837-3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendAIAgentTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentTextResponseBody) GoString() string {
	return s.String()
}

func (s *SendAIAgentTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAIAgentTextResponseBody) SetRequestId(v string) *SendAIAgentTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAIAgentTextResponseBody) Validate() error {
	return dara.Validate(s)
}
