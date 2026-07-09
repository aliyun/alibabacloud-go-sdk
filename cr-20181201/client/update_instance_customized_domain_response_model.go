// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceCustomizedDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceCustomizedDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceCustomizedDomainResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceCustomizedDomainResponseBody) *UpdateInstanceCustomizedDomainResponse
	GetBody() *UpdateInstanceCustomizedDomainResponseBody
}

type UpdateInstanceCustomizedDomainResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceCustomizedDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceCustomizedDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceCustomizedDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceCustomizedDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceCustomizedDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceCustomizedDomainResponse) GetBody() *UpdateInstanceCustomizedDomainResponseBody {
	return s.Body
}

func (s *UpdateInstanceCustomizedDomainResponse) SetHeaders(v map[string]*string) *UpdateInstanceCustomizedDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponse) SetStatusCode(v int32) *UpdateInstanceCustomizedDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponse) SetBody(v *UpdateInstanceCustomizedDomainResponseBody) *UpdateInstanceCustomizedDomainResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceCustomizedDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
