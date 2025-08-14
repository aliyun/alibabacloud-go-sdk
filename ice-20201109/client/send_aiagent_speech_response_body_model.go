// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentSpeechResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SendAIAgentSpeechResponseBody
	GetRequestId() *string
}

type SendAIAgentSpeechResponseBody struct {
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendAIAgentSpeechResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentSpeechResponseBody) GoString() string {
	return s.String()
}

func (s *SendAIAgentSpeechResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendAIAgentSpeechResponseBody) SetRequestId(v string) *SendAIAgentSpeechResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendAIAgentSpeechResponseBody) Validate() error {
	return dara.Validate(s)
}
