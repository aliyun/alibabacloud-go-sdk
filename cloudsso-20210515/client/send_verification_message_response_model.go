// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendVerificationMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendVerificationMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendVerificationMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendVerificationMessageResponseBody) *SendVerificationMessageResponse
	GetBody() *SendVerificationMessageResponseBody
}

type SendVerificationMessageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendVerificationMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendVerificationMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendVerificationMessageResponse) GoString() string {
	return s.String()
}

func (s *SendVerificationMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendVerificationMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendVerificationMessageResponse) GetBody() *SendVerificationMessageResponseBody {
	return s.Body
}

func (s *SendVerificationMessageResponse) SetHeaders(v map[string]*string) *SendVerificationMessageResponse {
	s.Headers = v
	return s
}

func (s *SendVerificationMessageResponse) SetStatusCode(v int32) *SendVerificationMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendVerificationMessageResponse) SetBody(v *SendVerificationMessageResponseBody) *SendVerificationMessageResponse {
	s.Body = v
	return s
}

func (s *SendVerificationMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
