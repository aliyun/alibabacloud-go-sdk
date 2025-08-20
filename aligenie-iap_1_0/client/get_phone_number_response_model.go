// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneNumberResponseBody) *GetPhoneNumberResponse
	GetBody() *GetPhoneNumberResponseBody
}

type GetPhoneNumberResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneNumberResponse) GetBody() *GetPhoneNumberResponseBody {
	return s.Body
}

func (s *GetPhoneNumberResponse) SetHeaders(v map[string]*string) *GetPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberResponse) SetStatusCode(v int32) *GetPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberResponse) SetBody(v *GetPhoneNumberResponseBody) *GetPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *GetPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
