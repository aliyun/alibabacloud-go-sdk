// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageCopyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendMessageCopyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendMessageCopyResponse
	GetStatusCode() *int32
	SetBody(v *SendMessageCopyResponseBody) *SendMessageCopyResponse
	GetBody() *SendMessageCopyResponseBody
}

type SendMessageCopyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendMessageCopyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendMessageCopyResponse) String() string {
	return dara.Prettify(s)
}

func (s SendMessageCopyResponse) GoString() string {
	return s.String()
}

func (s *SendMessageCopyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendMessageCopyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendMessageCopyResponse) GetBody() *SendMessageCopyResponseBody {
	return s.Body
}

func (s *SendMessageCopyResponse) SetHeaders(v map[string]*string) *SendMessageCopyResponse {
	s.Headers = v
	return s
}

func (s *SendMessageCopyResponse) SetStatusCode(v int32) *SendMessageCopyResponse {
	s.StatusCode = &v
	return s
}

func (s *SendMessageCopyResponse) SetBody(v *SendMessageCopyResponseBody) *SendMessageCopyResponse {
	s.Body = v
	return s
}

func (s *SendMessageCopyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
