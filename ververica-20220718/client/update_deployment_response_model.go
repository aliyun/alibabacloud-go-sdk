// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentResponseBody) *UpdateDeploymentResponse
	GetBody() *UpdateDeploymentResponseBody
}

type UpdateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentResponse) GetBody() *UpdateDeploymentResponseBody {
	return s.Body
}

func (s *UpdateDeploymentResponse) SetHeaders(v map[string]*string) *UpdateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentResponse) SetStatusCode(v int32) *UpdateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentResponse) SetBody(v *UpdateDeploymentResponseBody) *UpdateDeploymentResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
