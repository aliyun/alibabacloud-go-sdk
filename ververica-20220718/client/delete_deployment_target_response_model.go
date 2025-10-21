// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeploymentTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeploymentTargetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeploymentTargetResponseBody) *DeleteDeploymentTargetResponse
	GetBody() *DeleteDeploymentTargetResponseBody
}

type DeleteDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeploymentTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeploymentTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeploymentTargetResponse) GetBody() *DeleteDeploymentTargetResponseBody {
	return s.Body
}

func (s *DeleteDeploymentTargetResponse) SetHeaders(v map[string]*string) *DeleteDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeploymentTargetResponse) SetStatusCode(v int32) *DeleteDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeploymentTargetResponse) SetBody(v *DeleteDeploymentTargetResponseBody) *DeleteDeploymentTargetResponse {
	s.Body = v
	return s
}

func (s *DeleteDeploymentTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
