// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAsyncChatMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendAsyncChatMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendAsyncChatMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendAsyncChatMessageResponseBody) *SendAsyncChatMessageResponse
	GetBody() *SendAsyncChatMessageResponseBody
}

type SendAsyncChatMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendAsyncChatMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendAsyncChatMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendAsyncChatMessageResponse) GoString() string {
	return s.String()
}

func (s *SendAsyncChatMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendAsyncChatMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendAsyncChatMessageResponse) GetBody() *SendAsyncChatMessageResponseBody {
	return s.Body
}

func (s *SendAsyncChatMessageResponse) SetHeaders(v map[string]*string) *SendAsyncChatMessageResponse {
	s.Headers = v
	return s
}

func (s *SendAsyncChatMessageResponse) SetStatusCode(v int32) *SendAsyncChatMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendAsyncChatMessageResponse) SetBody(v *SendAsyncChatMessageResponseBody) *SendAsyncChatMessageResponse {
	s.Body = v
	return s
}

func (s *SendAsyncChatMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
