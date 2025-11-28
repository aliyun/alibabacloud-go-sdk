// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployPrivateRAGServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeployPrivateRAGServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeployPrivateRAGServiceResponse
	GetStatusCode() *int32
	SetBody(v *DeployPrivateRAGServiceResponseBody) *DeployPrivateRAGServiceResponse
	GetBody() *DeployPrivateRAGServiceResponseBody
}

type DeployPrivateRAGServiceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployPrivateRAGServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployPrivateRAGServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeployPrivateRAGServiceResponse) GoString() string {
	return s.String()
}

func (s *DeployPrivateRAGServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeployPrivateRAGServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeployPrivateRAGServiceResponse) GetBody() *DeployPrivateRAGServiceResponseBody {
	return s.Body
}

func (s *DeployPrivateRAGServiceResponse) SetHeaders(v map[string]*string) *DeployPrivateRAGServiceResponse {
	s.Headers = v
	return s
}

func (s *DeployPrivateRAGServiceResponse) SetStatusCode(v int32) *DeployPrivateRAGServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployPrivateRAGServiceResponse) SetBody(v *DeployPrivateRAGServiceResponseBody) *DeployPrivateRAGServiceResponse {
	s.Body = v
	return s
}

func (s *DeployPrivateRAGServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
