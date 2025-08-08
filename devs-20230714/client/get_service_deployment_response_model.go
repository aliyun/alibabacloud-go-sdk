// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *ServiceDeployment) *GetServiceDeploymentResponse
	GetBody() *ServiceDeployment
}

type GetServiceDeploymentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServiceDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetServiceDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceDeploymentResponse) GetBody() *ServiceDeployment {
	return s.Body
}

func (s *GetServiceDeploymentResponse) SetHeaders(v map[string]*string) *GetServiceDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetServiceDeploymentResponse) SetStatusCode(v int32) *GetServiceDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceDeploymentResponse) SetBody(v *ServiceDeployment) *GetServiceDeploymentResponse {
	s.Body = v
	return s
}

func (s *GetServiceDeploymentResponse) Validate() error {
	return dara.Validate(s)
}
