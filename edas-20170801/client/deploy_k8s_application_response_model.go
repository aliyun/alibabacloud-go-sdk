// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployK8sApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployK8sApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployK8sApplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeployK8sApplicationResponseBody) *DeployK8sApplicationResponse
	GetBody() *DeployK8sApplicationResponseBody
}

type DeployK8sApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployK8sApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployK8sApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployK8sApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployK8sApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployK8sApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployK8sApplicationResponse) GetBody() *DeployK8sApplicationResponseBody {
	return s.Body
}

func (s *DeployK8sApplicationResponse) SetHeaders(v map[string]*string) *DeployK8sApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployK8sApplicationResponse) SetStatusCode(v int32) *DeployK8sApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployK8sApplicationResponse) SetBody(v *DeployK8sApplicationResponseBody) *DeployK8sApplicationResponse {
	s.Body = v
	return s
}

func (s *DeployK8sApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
