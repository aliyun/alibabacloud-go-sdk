// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceCustomizedDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceCustomizedDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceCustomizedDomainResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceCustomizedDomainResponseBody) *DeleteInstanceCustomizedDomainResponse
	GetBody() *DeleteInstanceCustomizedDomainResponseBody
}

type DeleteInstanceCustomizedDomainResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceCustomizedDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceCustomizedDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceCustomizedDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceCustomizedDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceCustomizedDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceCustomizedDomainResponse) GetBody() *DeleteInstanceCustomizedDomainResponseBody {
	return s.Body
}

func (s *DeleteInstanceCustomizedDomainResponse) SetHeaders(v map[string]*string) *DeleteInstanceCustomizedDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponse) SetStatusCode(v int32) *DeleteInstanceCustomizedDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponse) SetBody(v *DeleteInstanceCustomizedDomainResponseBody) *DeleteInstanceCustomizedDomainResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceCustomizedDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
