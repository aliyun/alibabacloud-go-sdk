// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpService) *CreatePdpServiceResponse
	GetBody() *PdpService
}

type CreatePdpServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpService        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePdpServiceResponse) GoString() string {
	return s.String()
}

func (s *CreatePdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePdpServiceResponse) GetBody() *PdpService {
	return s.Body
}

func (s *CreatePdpServiceResponse) SetHeaders(v map[string]*string) *CreatePdpServiceResponse {
	s.Headers = v
	return s
}

func (s *CreatePdpServiceResponse) SetStatusCode(v int32) *CreatePdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePdpServiceResponse) SetBody(v *PdpService) *CreatePdpServiceResponse {
	s.Body = v
	return s
}

func (s *CreatePdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
