// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSaasServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PauseSaasServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PauseSaasServiceResponse
	GetStatusCode() *int32
	SetBody(v *PauseSaasServiceResponseBody) *PauseSaasServiceResponse
	GetBody() *PauseSaasServiceResponseBody
}

type PauseSaasServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PauseSaasServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PauseSaasServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PauseSaasServiceResponse) GoString() string {
	return s.String()
}

func (s *PauseSaasServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PauseSaasServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PauseSaasServiceResponse) GetBody() *PauseSaasServiceResponseBody {
	return s.Body
}

func (s *PauseSaasServiceResponse) SetHeaders(v map[string]*string) *PauseSaasServiceResponse {
	s.Headers = v
	return s
}

func (s *PauseSaasServiceResponse) SetStatusCode(v int32) *PauseSaasServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PauseSaasServiceResponse) SetBody(v *PauseSaasServiceResponseBody) *PauseSaasServiceResponse {
	s.Body = v
	return s
}

func (s *PauseSaasServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
