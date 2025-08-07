// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckKMSAuthorizedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckKMSAuthorizedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckKMSAuthorizedResponse
	GetStatusCode() *int32
	SetBody(v *CheckKMSAuthorizedResponseBody) *CheckKMSAuthorizedResponse
	GetBody() *CheckKMSAuthorizedResponseBody
}

type CheckKMSAuthorizedResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckKMSAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckKMSAuthorizedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckKMSAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *CheckKMSAuthorizedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckKMSAuthorizedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckKMSAuthorizedResponse) GetBody() *CheckKMSAuthorizedResponseBody {
	return s.Body
}

func (s *CheckKMSAuthorizedResponse) SetHeaders(v map[string]*string) *CheckKMSAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *CheckKMSAuthorizedResponse) SetStatusCode(v int32) *CheckKMSAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckKMSAuthorizedResponse) SetBody(v *CheckKMSAuthorizedResponseBody) *CheckKMSAuthorizedResponse {
	s.Body = v
	return s
}

func (s *CheckKMSAuthorizedResponse) Validate() error {
	return dara.Validate(s)
}
