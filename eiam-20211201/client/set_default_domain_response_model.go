// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultDomainResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultDomainResponseBody) *SetDefaultDomainResponse
	GetBody() *SetDefaultDomainResponseBody
}

type SetDefaultDomainResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultDomainResponse) GetBody() *SetDefaultDomainResponseBody {
	return s.Body
}

func (s *SetDefaultDomainResponse) SetHeaders(v map[string]*string) *SetDefaultDomainResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultDomainResponse) SetStatusCode(v int32) *SetDefaultDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultDomainResponse) SetBody(v *SetDefaultDomainResponseBody) *SetDefaultDomainResponse {
	s.Body = v
	return s
}

func (s *SetDefaultDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
