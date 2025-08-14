// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentDataChannelMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendAIAgentDataChannelMessageResponseBody
	GetRequestId() *string
}

type SendAIAgentDataChannelMessageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendAIAgentDataChannelMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentDataChannelMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendAIAgentDataChannelMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAIAgentDataChannelMessageResponseBody) SetRequestId(v string) *SendAIAgentDataChannelMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAIAgentDataChannelMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
