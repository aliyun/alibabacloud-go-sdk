// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpService) *UpdatePdpServiceResponse
	GetBody() *PdpService
}

type UpdatePdpServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpService        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePdpServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdatePdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePdpServiceResponse) GetBody() *PdpService {
	return s.Body
}

func (s *UpdatePdpServiceResponse) SetHeaders(v map[string]*string) *UpdatePdpServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdatePdpServiceResponse) SetStatusCode(v int32) *UpdatePdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePdpServiceResponse) SetBody(v *PdpService) *UpdatePdpServiceResponse {
	s.Body = v
	return s
}

func (s *UpdatePdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
