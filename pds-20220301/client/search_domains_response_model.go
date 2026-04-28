// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchDomainsResponse
	GetStatusCode() *int32
	SetBody(v *SearchDomainsResponseBody) *SearchDomainsResponse
	GetBody() *SearchDomainsResponseBody
}

type SearchDomainsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchDomainsResponse) GoString() string {
	return s.String()
}

func (s *SearchDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchDomainsResponse) GetBody() *SearchDomainsResponseBody {
	return s.Body
}

func (s *SearchDomainsResponse) SetHeaders(v map[string]*string) *SearchDomainsResponse {
	s.Headers = v
	return s
}

func (s *SearchDomainsResponse) SetStatusCode(v int32) *SearchDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDomainsResponse) SetBody(v *SearchDomainsResponseBody) *SearchDomainsResponse {
	s.Body = v
	return s
}

func (s *SearchDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
