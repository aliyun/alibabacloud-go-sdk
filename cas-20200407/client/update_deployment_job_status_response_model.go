// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentJobStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentJobStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentJobStatusResponseBody) *UpdateDeploymentJobStatusResponse
	GetBody() *UpdateDeploymentJobStatusResponseBody
}

type UpdateDeploymentJobStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentJobStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentJobStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentJobStatusResponse) GetBody() *UpdateDeploymentJobStatusResponseBody {
	return s.Body
}

func (s *UpdateDeploymentJobStatusResponse) SetHeaders(v map[string]*string) *UpdateDeploymentJobStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentJobStatusResponse) SetStatusCode(v int32) *UpdateDeploymentJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentJobStatusResponse) SetBody(v *UpdateDeploymentJobStatusResponseBody) *UpdateDeploymentJobStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentJobStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
