// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescDomainResponseBody) *DescDomainResponse
	GetBody() *DescDomainResponseBody
}

type DescDomainResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescDomainResponse) GoString() string {
	return s.String()
}

func (s *DescDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescDomainResponse) GetBody() *DescDomainResponseBody {
	return s.Body
}

func (s *DescDomainResponse) SetHeaders(v map[string]*string) *DescDomainResponse {
	s.Headers = v
	return s
}

func (s *DescDomainResponse) SetStatusCode(v int32) *DescDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescDomainResponse) SetBody(v *DescDomainResponseBody) *DescDomainResponse {
	s.Body = v
	return s
}

func (s *DescDomainResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
