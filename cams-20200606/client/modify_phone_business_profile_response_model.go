// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPhoneBusinessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPhoneBusinessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPhoneBusinessProfileResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPhoneBusinessProfileResponseBody) *ModifyPhoneBusinessProfileResponse
	GetBody() *ModifyPhoneBusinessProfileResponseBody
}

type ModifyPhoneBusinessProfileResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPhoneBusinessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPhoneBusinessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPhoneBusinessProfileResponse) GoString() string {
	return s.String()
}

func (s *ModifyPhoneBusinessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPhoneBusinessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPhoneBusinessProfileResponse) GetBody() *ModifyPhoneBusinessProfileResponseBody {
	return s.Body
}

func (s *ModifyPhoneBusinessProfileResponse) SetHeaders(v map[string]*string) *ModifyPhoneBusinessProfileResponse {
	s.Headers = v
	return s
}

func (s *ModifyPhoneBusinessProfileResponse) SetStatusCode(v int32) *ModifyPhoneBusinessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPhoneBusinessProfileResponse) SetBody(v *ModifyPhoneBusinessProfileResponseBody) *ModifyPhoneBusinessProfileResponse {
	s.Body = v
	return s
}

func (s *ModifyPhoneBusinessProfileResponse) Validate() error {
	return dara.Validate(s)
}
