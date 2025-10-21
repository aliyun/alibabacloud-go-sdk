// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse
	GetBody() *CreateDeploymentResponseBody
}

type CreateDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentResponse) GetBody() *CreateDeploymentResponseBody {
	return s.Body
}

func (s *CreateDeploymentResponse) SetHeaders(v map[string]*string) *CreateDeploymentResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentResponse) SetStatusCode(v int32) *CreateDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentResponse) SetBody(v *CreateDeploymentResponseBody) *CreateDeploymentResponse {
	s.Body = v
	return s
}

func (s *CreateDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
