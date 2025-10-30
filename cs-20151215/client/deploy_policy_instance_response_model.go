// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPolicyInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployPolicyInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployPolicyInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeployPolicyInstanceResponseBody) *DeployPolicyInstanceResponse
	GetBody() *DeployPolicyInstanceResponseBody
}

type DeployPolicyInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployPolicyInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployPolicyInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployPolicyInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeployPolicyInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployPolicyInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployPolicyInstanceResponse) GetBody() *DeployPolicyInstanceResponseBody {
	return s.Body
}

func (s *DeployPolicyInstanceResponse) SetHeaders(v map[string]*string) *DeployPolicyInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeployPolicyInstanceResponse) SetStatusCode(v int32) *DeployPolicyInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployPolicyInstanceResponse) SetBody(v *DeployPolicyInstanceResponseBody) *DeployPolicyInstanceResponse {
	s.Body = v
	return s
}

func (s *DeployPolicyInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
