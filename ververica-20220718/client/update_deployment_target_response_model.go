// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentTargetResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentTargetResponseBody) *UpdateDeploymentTargetResponse
	GetBody() *UpdateDeploymentTargetResponseBody
}

type UpdateDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentTargetResponse) GetBody() *UpdateDeploymentTargetResponseBody {
	return s.Body
}

func (s *UpdateDeploymentTargetResponse) SetHeaders(v map[string]*string) *UpdateDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentTargetResponse) SetStatusCode(v int32) *UpdateDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentTargetResponse) SetBody(v *UpdateDeploymentTargetResponseBody) *UpdateDeploymentTargetResponse {
	s.Body = v
	return s
}

func (s *UpdateDeploymentTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
