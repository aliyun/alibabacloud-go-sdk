// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFileDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckFileDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckFileDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *CheckFileDeploymentResponseBody) *CheckFileDeploymentResponse
	GetBody() *CheckFileDeploymentResponseBody
}

type CheckFileDeploymentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckFileDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckFileDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckFileDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CheckFileDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckFileDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckFileDeploymentResponse) GetBody() *CheckFileDeploymentResponseBody {
	return s.Body
}

func (s *CheckFileDeploymentResponse) SetHeaders(v map[string]*string) *CheckFileDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CheckFileDeploymentResponse) SetStatusCode(v int32) *CheckFileDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckFileDeploymentResponse) SetBody(v *CheckFileDeploymentResponseBody) *CheckFileDeploymentResponse {
	s.Body = v
	return s
}

func (s *CheckFileDeploymentResponse) Validate() error {
	return dara.Validate(s)
}
