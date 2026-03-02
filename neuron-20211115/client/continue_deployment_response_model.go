// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinueDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinueDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceDeployment) *ContinueDeploymentResponse
	GetBody() *PdpServiceDeployment
}

type ContinueDeploymentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinueDeploymentResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinueDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinueDeploymentResponse) GetBody() *PdpServiceDeployment {
	return s.Body
}

func (s *ContinueDeploymentResponse) SetHeaders(v map[string]*string) *ContinueDeploymentResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeploymentResponse) SetStatusCode(v int32) *ContinueDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeploymentResponse) SetBody(v *PdpServiceDeployment) *ContinueDeploymentResponse {
	s.Body = v
	return s
}

func (s *ContinueDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
