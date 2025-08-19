// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendPhoneVerificationForMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendPhoneVerificationForMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendPhoneVerificationForMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *SendPhoneVerificationForMessageContactResponseBody) *SendPhoneVerificationForMessageContactResponse
	GetBody() *SendPhoneVerificationForMessageContactResponseBody
}

type SendPhoneVerificationForMessageContactResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendPhoneVerificationForMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendPhoneVerificationForMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s SendPhoneVerificationForMessageContactResponse) GoString() string {
	return s.String()
}

func (s *SendPhoneVerificationForMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendPhoneVerificationForMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendPhoneVerificationForMessageContactResponse) GetBody() *SendPhoneVerificationForMessageContactResponseBody {
	return s.Body
}

func (s *SendPhoneVerificationForMessageContactResponse) SetHeaders(v map[string]*string) *SendPhoneVerificationForMessageContactResponse {
	s.Headers = v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponse) SetStatusCode(v int32) *SendPhoneVerificationForMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponse) SetBody(v *SendPhoneVerificationForMessageContactResponseBody) *SendPhoneVerificationForMessageContactResponse {
	s.Body = v
	return s
}

func (s *SendPhoneVerificationForMessageContactResponse) Validate() error {
	return dara.Validate(s)
}
