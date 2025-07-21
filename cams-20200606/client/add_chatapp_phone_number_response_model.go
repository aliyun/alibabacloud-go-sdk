// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatappPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddChatappPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddChatappPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *AddChatappPhoneNumberResponseBody) *AddChatappPhoneNumberResponse
	GetBody() *AddChatappPhoneNumberResponseBody
}

type AddChatappPhoneNumberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddChatappPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddChatappPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddChatappPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *AddChatappPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddChatappPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddChatappPhoneNumberResponse) GetBody() *AddChatappPhoneNumberResponseBody {
	return s.Body
}

func (s *AddChatappPhoneNumberResponse) SetHeaders(v map[string]*string) *AddChatappPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *AddChatappPhoneNumberResponse) SetStatusCode(v int32) *AddChatappPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddChatappPhoneNumberResponse) SetBody(v *AddChatappPhoneNumberResponseBody) *AddChatappPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *AddChatappPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
