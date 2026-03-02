// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertPdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevertPdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevertPdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpServiceDeployment) *RevertPdpServiceResponse
	GetBody() *PdpServiceDeployment
}

type RevertPdpServiceResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpServiceDeployment `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevertPdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RevertPdpServiceResponse) GoString() string {
	return s.String()
}

func (s *RevertPdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevertPdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevertPdpServiceResponse) GetBody() *PdpServiceDeployment {
	return s.Body
}

func (s *RevertPdpServiceResponse) SetHeaders(v map[string]*string) *RevertPdpServiceResponse {
	s.Headers = v
	return s
}

func (s *RevertPdpServiceResponse) SetStatusCode(v int32) *RevertPdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertPdpServiceResponse) SetBody(v *PdpServiceDeployment) *RevertPdpServiceResponse {
	s.Body = v
	return s
}

func (s *RevertPdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
