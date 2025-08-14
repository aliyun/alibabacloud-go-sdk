// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentDataChannelMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAIAgentDataChannelMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAIAgentDataChannelMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendAIAgentDataChannelMessageResponseBody) *SendAIAgentDataChannelMessageResponse
	GetBody() *SendAIAgentDataChannelMessageResponseBody
}

type SendAIAgentDataChannelMessageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAIAgentDataChannelMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAIAgentDataChannelMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentDataChannelMessageResponse) GoString() string {
	return s.String()
}

func (s *SendAIAgentDataChannelMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAIAgentDataChannelMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAIAgentDataChannelMessageResponse) GetBody() *SendAIAgentDataChannelMessageResponseBody {
	return s.Body
}

func (s *SendAIAgentDataChannelMessageResponse) SetHeaders(v map[string]*string) *SendAIAgentDataChannelMessageResponse {
	s.Headers = v
	return s
}

func (s *SendAIAgentDataChannelMessageResponse) SetStatusCode(v int32) *SendAIAgentDataChannelMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAIAgentDataChannelMessageResponse) SetBody(v *SendAIAgentDataChannelMessageResponseBody) *SendAIAgentDataChannelMessageResponse {
	s.Body = v
	return s
}

func (s *SendAIAgentDataChannelMessageResponse) Validate() error {
	return dara.Validate(s)
}
