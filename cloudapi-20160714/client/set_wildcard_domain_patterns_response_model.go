// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWildcardDomainPatternsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetWildcardDomainPatternsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetWildcardDomainPatternsResponse
	GetStatusCode() *int32
	SetBody(v *SetWildcardDomainPatternsResponseBody) *SetWildcardDomainPatternsResponse
	GetBody() *SetWildcardDomainPatternsResponseBody
}

type SetWildcardDomainPatternsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetWildcardDomainPatternsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetWildcardDomainPatternsResponse) String() string {
	return dara.Prettify(s)
}

func (s SetWildcardDomainPatternsResponse) GoString() string {
	return s.String()
}

func (s *SetWildcardDomainPatternsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetWildcardDomainPatternsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetWildcardDomainPatternsResponse) GetBody() *SetWildcardDomainPatternsResponseBody {
	return s.Body
}

func (s *SetWildcardDomainPatternsResponse) SetHeaders(v map[string]*string) *SetWildcardDomainPatternsResponse {
	s.Headers = v
	return s
}

func (s *SetWildcardDomainPatternsResponse) SetStatusCode(v int32) *SetWildcardDomainPatternsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetWildcardDomainPatternsResponse) SetBody(v *SetWildcardDomainPatternsResponseBody) *SetWildcardDomainPatternsResponse {
	s.Body = v
	return s
}

func (s *SetWildcardDomainPatternsResponse) Validate() error {
	return dara.Validate(s)
}
