// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentJobResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentJobResponseBody) *UpdateDeploymentJobResponse
	GetBody() *UpdateDeploymentJobResponseBody
}

type UpdateDeploymentJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentJobResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentJobResponse) GetBody() *UpdateDeploymentJobResponseBody {
	return s.Body
}

func (s *UpdateDeploymentJobResponse) SetHeaders(v map[string]*string) *UpdateDeploymentJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentJobResponse) SetStatusCode(v int32) *UpdateDeploymentJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentJobResponse) SetBody(v *UpdateDeploymentJobResponseBody) *UpdateDeploymentJobResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentJobResponse) Validate() error {
	return dara.Validate(s)
}
