// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceDeployableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceDeployableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceDeployableResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceDeployableResponseBody) *CheckServiceDeployableResponse
	GetBody() *CheckServiceDeployableResponseBody
}

type CheckServiceDeployableResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceDeployableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceDeployableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceDeployableResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceDeployableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceDeployableResponse) GetBody() *CheckServiceDeployableResponseBody {
	return s.Body
}

func (s *CheckServiceDeployableResponse) SetHeaders(v map[string]*string) *CheckServiceDeployableResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceDeployableResponse) SetStatusCode(v int32) *CheckServiceDeployableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceDeployableResponse) SetBody(v *CheckServiceDeployableResponseBody) *CheckServiceDeployableResponse {
	s.Body = v
	return s
}

func (s *CheckServiceDeployableResponse) Validate() error {
	return dara.Validate(s)
}
