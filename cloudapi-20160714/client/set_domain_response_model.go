// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDomainResponse
	GetStatusCode() *int32
	SetBody(v *SetDomainResponseBody) *SetDomainResponse
	GetBody() *SetDomainResponseBody
}

type SetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDomainResponse) GetBody() *SetDomainResponseBody {
	return s.Body
}

func (s *SetDomainResponse) SetHeaders(v map[string]*string) *SetDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDomainResponse) SetStatusCode(v int32) *SetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDomainResponse) SetBody(v *SetDomainResponseBody) *SetDomainResponse {
	s.Body = v
	return s
}

func (s *SetDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
