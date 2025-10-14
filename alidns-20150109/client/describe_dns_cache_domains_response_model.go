// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsCacheDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDnsCacheDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDnsCacheDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDnsCacheDomainsResponseBody) *DescribeDnsCacheDomainsResponse
	GetBody() *DescribeDnsCacheDomainsResponseBody
}

type DescribeDnsCacheDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDnsCacheDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDnsCacheDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsCacheDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDnsCacheDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDnsCacheDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDnsCacheDomainsResponse) GetBody() *DescribeDnsCacheDomainsResponseBody {
	return s.Body
}

func (s *DescribeDnsCacheDomainsResponse) SetHeaders(v map[string]*string) *DescribeDnsCacheDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDnsCacheDomainsResponse) SetStatusCode(v int32) *DescribeDnsCacheDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDnsCacheDomainsResponse) SetBody(v *DescribeDnsCacheDomainsResponseBody) *DescribeDnsCacheDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDnsCacheDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
