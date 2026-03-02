// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPdpServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPdpServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPdpServiceResponse
	GetStatusCode() *int32
	SetBody(v *PdpService) *GetPdpServiceResponse
	GetBody() *PdpService
}

type GetPdpServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PdpService        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPdpServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPdpServiceResponse) GoString() string {
	return s.String()
}

func (s *GetPdpServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPdpServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPdpServiceResponse) GetBody() *PdpService {
	return s.Body
}

func (s *GetPdpServiceResponse) SetHeaders(v map[string]*string) *GetPdpServiceResponse {
	s.Headers = v
	return s
}

func (s *GetPdpServiceResponse) SetStatusCode(v int32) *GetPdpServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPdpServiceResponse) SetBody(v *PdpService) *GetPdpServiceResponse {
	s.Body = v
	return s
}

func (s *GetPdpServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
