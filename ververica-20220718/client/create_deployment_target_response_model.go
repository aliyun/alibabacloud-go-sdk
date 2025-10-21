// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentTargetResponseBody) *CreateDeploymentTargetResponse
	GetBody() *CreateDeploymentTargetResponseBody
}

type CreateDeploymentTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentTargetResponse) GetBody() *CreateDeploymentTargetResponseBody {
	return s.Body
}

func (s *CreateDeploymentTargetResponse) SetHeaders(v map[string]*string) *CreateDeploymentTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateDeploymentTargetResponse) SetStatusCode(v int32) *CreateDeploymentTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentTargetResponse) SetBody(v *CreateDeploymentTargetResponseBody) *CreateDeploymentTargetResponse {
	s.Body = v
	return s
}

func (s *CreateDeploymentTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
