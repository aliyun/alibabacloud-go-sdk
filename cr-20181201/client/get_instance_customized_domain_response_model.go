// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCustomizedDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceCustomizedDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceCustomizedDomainResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceCustomizedDomainResponseBody) *GetInstanceCustomizedDomainResponse
	GetBody() *GetInstanceCustomizedDomainResponseBody
}

type GetInstanceCustomizedDomainResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceCustomizedDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceCustomizedDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCustomizedDomainResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceCustomizedDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceCustomizedDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceCustomizedDomainResponse) GetBody() *GetInstanceCustomizedDomainResponseBody {
	return s.Body
}

func (s *GetInstanceCustomizedDomainResponse) SetHeaders(v map[string]*string) *GetInstanceCustomizedDomainResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceCustomizedDomainResponse) SetStatusCode(v int32) *GetInstanceCustomizedDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponse) SetBody(v *GetInstanceCustomizedDomainResponseBody) *GetInstanceCustomizedDomainResponse {
	s.Body = v
	return s
}

func (s *GetInstanceCustomizedDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
