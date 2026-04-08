// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployTrafficControlTaskCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployTrafficControlTaskCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployTrafficControlTaskCodeResponse
	GetStatusCode() *int32
	SetBody(v *DeployTrafficControlTaskCodeResponseBody) *DeployTrafficControlTaskCodeResponse
	GetBody() *DeployTrafficControlTaskCodeResponseBody
}

type DeployTrafficControlTaskCodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployTrafficControlTaskCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployTrafficControlTaskCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployTrafficControlTaskCodeResponse) GoString() string {
	return s.String()
}

func (s *DeployTrafficControlTaskCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployTrafficControlTaskCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployTrafficControlTaskCodeResponse) GetBody() *DeployTrafficControlTaskCodeResponseBody {
	return s.Body
}

func (s *DeployTrafficControlTaskCodeResponse) SetHeaders(v map[string]*string) *DeployTrafficControlTaskCodeResponse {
	s.Headers = v
	return s
}

func (s *DeployTrafficControlTaskCodeResponse) SetStatusCode(v int32) *DeployTrafficControlTaskCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployTrafficControlTaskCodeResponse) SetBody(v *DeployTrafficControlTaskCodeResponseBody) *DeployTrafficControlTaskCodeResponse {
	s.Body = v
	return s
}

func (s *DeployTrafficControlTaskCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
