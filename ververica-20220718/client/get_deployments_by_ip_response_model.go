// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeploymentsByIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeploymentsByIpResponse
	GetStatusCode() *int32
	SetBody(v *GetDeploymentsByIpResponseBody) *GetDeploymentsByIpResponse
	GetBody() *GetDeploymentsByIpResponseBody
}

type GetDeploymentsByIpResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeploymentsByIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeploymentsByIpResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByIpResponse) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeploymentsByIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeploymentsByIpResponse) GetBody() *GetDeploymentsByIpResponseBody {
	return s.Body
}

func (s *GetDeploymentsByIpResponse) SetHeaders(v map[string]*string) *GetDeploymentsByIpResponse {
	s.Headers = v
	return s
}

func (s *GetDeploymentsByIpResponse) SetStatusCode(v int32) *GetDeploymentsByIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeploymentsByIpResponse) SetBody(v *GetDeploymentsByIpResponseBody) *GetDeploymentsByIpResponse {
	s.Body = v
	return s
}

func (s *GetDeploymentsByIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
