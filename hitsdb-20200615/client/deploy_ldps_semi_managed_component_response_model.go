// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployLdpsSemiManagedComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployLdpsSemiManagedComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployLdpsSemiManagedComponentResponse
	GetStatusCode() *int32
	SetBody(v *DeployLdpsSemiManagedComponentResponseBody) *DeployLdpsSemiManagedComponentResponse
	GetBody() *DeployLdpsSemiManagedComponentResponseBody
}

type DeployLdpsSemiManagedComponentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployLdpsSemiManagedComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployLdpsSemiManagedComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentResponse) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployLdpsSemiManagedComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployLdpsSemiManagedComponentResponse) GetBody() *DeployLdpsSemiManagedComponentResponseBody {
	return s.Body
}

func (s *DeployLdpsSemiManagedComponentResponse) SetHeaders(v map[string]*string) *DeployLdpsSemiManagedComponentResponse {
	s.Headers = v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponse) SetStatusCode(v int32) *DeployLdpsSemiManagedComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponse) SetBody(v *DeployLdpsSemiManagedComponentResponseBody) *DeployLdpsSemiManagedComponentResponse {
	s.Body = v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponse) Validate() error {
	return dara.Validate(s)
}
