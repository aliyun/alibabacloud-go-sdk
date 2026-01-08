// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberDeregisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappPhoneNumberDeregisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappPhoneNumberDeregisterResponse
	GetStatusCode() *int32
	SetBody(v *ChatappPhoneNumberDeregisterResponseBody) *ChatappPhoneNumberDeregisterResponse
	GetBody() *ChatappPhoneNumberDeregisterResponseBody
}

type ChatappPhoneNumberDeregisterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappPhoneNumberDeregisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappPhoneNumberDeregisterResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappPhoneNumberDeregisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappPhoneNumberDeregisterResponse) GetBody() *ChatappPhoneNumberDeregisterResponseBody {
	return s.Body
}

func (s *ChatappPhoneNumberDeregisterResponse) SetHeaders(v map[string]*string) *ChatappPhoneNumberDeregisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponse) SetStatusCode(v int32) *ChatappPhoneNumberDeregisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponse) SetBody(v *ChatappPhoneNumberDeregisterResponseBody) *ChatappPhoneNumberDeregisterResponse {
	s.Body = v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
