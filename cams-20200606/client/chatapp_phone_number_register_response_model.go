// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberRegisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChatappPhoneNumberRegisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChatappPhoneNumberRegisterResponse
	GetStatusCode() *int32
	SetBody(v *ChatappPhoneNumberRegisterResponseBody) *ChatappPhoneNumberRegisterResponse
	GetBody() *ChatappPhoneNumberRegisterResponseBody
}

type ChatappPhoneNumberRegisterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChatappPhoneNumberRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChatappPhoneNumberRegisterResponse) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberRegisterResponse) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberRegisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChatappPhoneNumberRegisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChatappPhoneNumberRegisterResponse) GetBody() *ChatappPhoneNumberRegisterResponseBody {
	return s.Body
}

func (s *ChatappPhoneNumberRegisterResponse) SetHeaders(v map[string]*string) *ChatappPhoneNumberRegisterResponse {
	s.Headers = v
	return s
}

func (s *ChatappPhoneNumberRegisterResponse) SetStatusCode(v int32) *ChatappPhoneNumberRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *ChatappPhoneNumberRegisterResponse) SetBody(v *ChatappPhoneNumberRegisterResponseBody) *ChatappPhoneNumberRegisterResponse {
	s.Body = v
	return s
}

func (s *ChatappPhoneNumberRegisterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
