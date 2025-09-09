// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse
	GetBody() *DeployServiceInstanceResponseBody
}

type DeployServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployServiceInstanceResponse) GetBody() *DeployServiceInstanceResponseBody {
	return s.Body
}

func (s *DeployServiceInstanceResponse) SetHeaders(v map[string]*string) *DeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployServiceInstanceResponse) SetStatusCode(v int32) *DeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployServiceInstanceResponse) SetBody(v *DeployServiceInstanceResponseBody) *DeployServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *DeployServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
