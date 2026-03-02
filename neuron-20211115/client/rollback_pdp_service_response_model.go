// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackPdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackPdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackPdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceDeployment) *RollbackPdpServiceResponse
	GetBody() *PdpServiceDeployment
}

type RollbackPdpServiceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackPdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackPdpServiceResponse) GoString() string {
	return s.String()
}

func (s *RollbackPdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackPdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackPdpServiceResponse) GetBody() *PdpServiceDeployment {
	return s.Body
}

func (s *RollbackPdpServiceResponse) SetHeaders(v map[string]*string) *RollbackPdpServiceResponse {
	s.Headers = v
	return s
}

func (s *RollbackPdpServiceResponse) SetStatusCode(v int32) *RollbackPdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackPdpServiceResponse) SetBody(v *PdpServiceDeployment) *RollbackPdpServiceResponse {
	s.Body = v
	return s
}

func (s *RollbackPdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
