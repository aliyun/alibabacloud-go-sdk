// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentByNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentByNameResponseBody) *UpdateDeploymentByNameResponse
	GetBody() *UpdateDeploymentByNameResponseBody
}

type UpdateDeploymentByNameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentByNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentByNameResponse) GetBody() *UpdateDeploymentByNameResponseBody {
	return s.Body
}

func (s *UpdateDeploymentByNameResponse) SetHeaders(v map[string]*string) *UpdateDeploymentByNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentByNameResponse) SetStatusCode(v int32) *UpdateDeploymentByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentByNameResponse) SetBody(v *UpdateDeploymentByNameResponseBody) *UpdateDeploymentByNameResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
