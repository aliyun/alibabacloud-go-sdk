// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLastDeploymentConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLastDeploymentConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLastDeploymentConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeployConfigInfo) *GetLastDeploymentConfigResponse
	GetBody() *DeployConfigInfo
}

type GetLastDeploymentConfigResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployConfigInfo  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLastDeploymentConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLastDeploymentConfigResponse) GoString() string {
	return s.String()
}

func (s *GetLastDeploymentConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLastDeploymentConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLastDeploymentConfigResponse) GetBody() *DeployConfigInfo {
	return s.Body
}

func (s *GetLastDeploymentConfigResponse) SetHeaders(v map[string]*string) *GetLastDeploymentConfigResponse {
	s.Headers = v
	return s
}

func (s *GetLastDeploymentConfigResponse) SetStatusCode(v int32) *GetLastDeploymentConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLastDeploymentConfigResponse) SetBody(v *DeployConfigInfo) *GetLastDeploymentConfigResponse {
	s.Body = v
	return s
}

func (s *GetLastDeploymentConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
