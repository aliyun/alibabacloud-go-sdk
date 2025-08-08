// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployEnvironmentResponse
	GetStatusCode() *int32
	SetBody(v *EnvironmentDeployment) *DeployEnvironmentResponse
	GetBody() *EnvironmentDeployment
}

type DeployEnvironmentResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnvironmentDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeployEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployEnvironmentResponse) GetBody() *EnvironmentDeployment {
	return s.Body
}

func (s *DeployEnvironmentResponse) SetHeaders(v map[string]*string) *DeployEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeployEnvironmentResponse) SetStatusCode(v int32) *DeployEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployEnvironmentResponse) SetBody(v *EnvironmentDeployment) *DeployEnvironmentResponse {
	s.Body = v
	return s
}

func (s *DeployEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
