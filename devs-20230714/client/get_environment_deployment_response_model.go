// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEnvironmentDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEnvironmentDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEnvironmentDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *EnvironmentDeployment) *GetEnvironmentDeploymentResponse
	GetBody() *EnvironmentDeployment
}

type GetEnvironmentDeploymentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvironmentDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEnvironmentDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEnvironmentDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetEnvironmentDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEnvironmentDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEnvironmentDeploymentResponse) GetBody() *EnvironmentDeployment {
	return s.Body
}

func (s *GetEnvironmentDeploymentResponse) SetHeaders(v map[string]*string) *GetEnvironmentDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetEnvironmentDeploymentResponse) SetStatusCode(v int32) *GetEnvironmentDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEnvironmentDeploymentResponse) SetBody(v *EnvironmentDeployment) *GetEnvironmentDeploymentResponse {
	s.Body = v
	return s
}

func (s *GetEnvironmentDeploymentResponse) Validate() error {
	return dara.Validate(s)
}
