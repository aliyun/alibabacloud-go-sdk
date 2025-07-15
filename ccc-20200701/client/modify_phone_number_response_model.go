// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneNumberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPhoneNumberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPhoneNumberResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPhoneNumberResponseBody) *ModifyPhoneNumberResponse
	GetBody() *ModifyPhoneNumberResponseBody
}

type ModifyPhoneNumberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPhoneNumberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPhoneNumberResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneNumberResponse) GoString() string {
	return s.String()
}

func (s *ModifyPhoneNumberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPhoneNumberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPhoneNumberResponse) GetBody() *ModifyPhoneNumberResponseBody {
	return s.Body
}

func (s *ModifyPhoneNumberResponse) SetHeaders(v map[string]*string) *ModifyPhoneNumberResponse {
	s.Headers = v
	return s
}

func (s *ModifyPhoneNumberResponse) SetStatusCode(v int32) *ModifyPhoneNumberResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPhoneNumberResponse) SetBody(v *ModifyPhoneNumberResponseBody) *ModifyPhoneNumberResponse {
	s.Body = v
	return s
}

func (s *ModifyPhoneNumberResponse) Validate() error {
	return dara.Validate(s)
}
