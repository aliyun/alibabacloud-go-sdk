// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentResponseBody) *DeleteDeploymentResponse
	GetBody() *DeleteDeploymentResponseBody
}

type DeleteDeploymentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentResponse) GetBody() *DeleteDeploymentResponseBody {
	return s.Body
}

func (s *DeleteDeploymentResponse) SetHeaders(v map[string]*string) *DeleteDeploymentResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentResponse) SetStatusCode(v int32) *DeleteDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentResponse) SetBody(v *DeleteDeploymentResponseBody) *DeleteDeploymentResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
