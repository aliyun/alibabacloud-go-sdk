// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneNumberVerificationStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneNumberVerificationStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneNumberVerificationStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneNumberVerificationStatusResponseBody) *GetPhoneNumberVerificationStatusResponse
	GetBody() *GetPhoneNumberVerificationStatusResponseBody
}

type GetPhoneNumberVerificationStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneNumberVerificationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneNumberVerificationStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneNumberVerificationStatusResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneNumberVerificationStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneNumberVerificationStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneNumberVerificationStatusResponse) GetBody() *GetPhoneNumberVerificationStatusResponseBody {
	return s.Body
}

func (s *GetPhoneNumberVerificationStatusResponse) SetHeaders(v map[string]*string) *GetPhoneNumberVerificationStatusResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponse) SetStatusCode(v int32) *GetPhoneNumberVerificationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponse) SetBody(v *GetPhoneNumberVerificationStatusResponseBody) *GetPhoneNumberVerificationStatusResponse {
	s.Body = v
	return s
}

func (s *GetPhoneNumberVerificationStatusResponse) Validate() error {
	return dara.Validate(s)
}
