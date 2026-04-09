// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDomainResponse
	GetStatusCode() *int32
	SetBody(v *AddDomainResponseBody) *AddDomainResponse
	GetBody() *AddDomainResponseBody
}

type AddDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDomainResponse) GoString() string {
	return s.String()
}

func (s *AddDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDomainResponse) GetBody() *AddDomainResponseBody {
	return s.Body
}

func (s *AddDomainResponse) SetHeaders(v map[string]*string) *AddDomainResponse {
	s.Headers = v
	return s
}

func (s *AddDomainResponse) SetStatusCode(v int32) *AddDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDomainResponse) SetBody(v *AddDomainResponseBody) *AddDomainResponse {
	s.Body = v
	return s
}

func (s *AddDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
