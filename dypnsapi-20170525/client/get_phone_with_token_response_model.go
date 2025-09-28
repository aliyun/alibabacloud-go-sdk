// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhoneWithTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhoneWithTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhoneWithTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetPhoneWithTokenResponseBody) *GetPhoneWithTokenResponse
	GetBody() *GetPhoneWithTokenResponseBody
}

type GetPhoneWithTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhoneWithTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhoneWithTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhoneWithTokenResponse) GoString() string {
	return s.String()
}

func (s *GetPhoneWithTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhoneWithTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhoneWithTokenResponse) GetBody() *GetPhoneWithTokenResponseBody {
	return s.Body
}

func (s *GetPhoneWithTokenResponse) SetHeaders(v map[string]*string) *GetPhoneWithTokenResponse {
	s.Headers = v
	return s
}

func (s *GetPhoneWithTokenResponse) SetStatusCode(v int32) *GetPhoneWithTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhoneWithTokenResponse) SetBody(v *GetPhoneWithTokenResponseBody) *GetPhoneWithTokenResponse {
	s.Body = v
	return s
}

func (s *GetPhoneWithTokenResponse) Validate() error {
	return dara.Validate(s)
}
