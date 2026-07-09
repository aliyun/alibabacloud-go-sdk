// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceCustomizedDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceCustomizedDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceCustomizedDomainResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceCustomizedDomainResponseBody) *CreateInstanceCustomizedDomainResponse
	GetBody() *CreateInstanceCustomizedDomainResponseBody
}

type CreateInstanceCustomizedDomainResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceCustomizedDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceCustomizedDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceCustomizedDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceCustomizedDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceCustomizedDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceCustomizedDomainResponse) GetBody() *CreateInstanceCustomizedDomainResponseBody {
	return s.Body
}

func (s *CreateInstanceCustomizedDomainResponse) SetHeaders(v map[string]*string) *CreateInstanceCustomizedDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceCustomizedDomainResponse) SetStatusCode(v int32) *CreateInstanceCustomizedDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceCustomizedDomainResponse) SetBody(v *CreateInstanceCustomizedDomainResponseBody) *CreateInstanceCustomizedDomainResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceCustomizedDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
