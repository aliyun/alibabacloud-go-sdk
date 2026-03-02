// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceDeployment) *TriggerDeploymentResponse
	GetBody() *PdpServiceDeployment
}

type TriggerDeploymentResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerDeploymentResponse) GoString() string {
	return s.String()
}

func (s *TriggerDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerDeploymentResponse) GetBody() *PdpServiceDeployment {
	return s.Body
}

func (s *TriggerDeploymentResponse) SetHeaders(v map[string]*string) *TriggerDeploymentResponse {
	s.Headers = v
	return s
}

func (s *TriggerDeploymentResponse) SetStatusCode(v int32) *TriggerDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerDeploymentResponse) SetBody(v *PdpServiceDeployment) *TriggerDeploymentResponse {
	s.Body = v
	return s
}

func (s *TriggerDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
