// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCloudResourceAuthorizedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckCloudResourceAuthorizedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckCloudResourceAuthorizedResponse
	GetStatusCode() *int32
	SetBody(v *CheckCloudResourceAuthorizedResponseBody) *CheckCloudResourceAuthorizedResponse
	GetBody() *CheckCloudResourceAuthorizedResponseBody
}

type CheckCloudResourceAuthorizedResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckCloudResourceAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckCloudResourceAuthorizedResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckCloudResourceAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *CheckCloudResourceAuthorizedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckCloudResourceAuthorizedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckCloudResourceAuthorizedResponse) GetBody() *CheckCloudResourceAuthorizedResponseBody {
	return s.Body
}

func (s *CheckCloudResourceAuthorizedResponse) SetHeaders(v map[string]*string) *CheckCloudResourceAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *CheckCloudResourceAuthorizedResponse) SetStatusCode(v int32) *CheckCloudResourceAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckCloudResourceAuthorizedResponse) SetBody(v *CheckCloudResourceAuthorizedResponseBody) *CheckCloudResourceAuthorizedResponse {
	s.Body = v
	return s
}

func (s *CheckCloudResourceAuthorizedResponse) Validate() error {
	return dara.Validate(s)
}
