// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentByNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentByNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentByNameResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentByNameResponseBody) *DeleteDeploymentByNameResponse
	GetBody() *DeleteDeploymentByNameResponseBody
}

type DeleteDeploymentByNameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentByNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentByNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentByNameResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentByNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentByNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentByNameResponse) GetBody() *DeleteDeploymentByNameResponseBody {
	return s.Body
}

func (s *DeleteDeploymentByNameResponse) SetHeaders(v map[string]*string) *DeleteDeploymentByNameResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentByNameResponse) SetStatusCode(v int32) *DeleteDeploymentByNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentByNameResponse) SetBody(v *DeleteDeploymentByNameResponseBody) *DeleteDeploymentByNameResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentByNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
