// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeletedDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDeletedDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDeletedDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDeletedDomainsResponseBody) *DescribeCdnDeletedDomainsResponse
	GetBody() *DescribeCdnDeletedDomainsResponseBody
}

type DescribeCdnDeletedDomainsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDeletedDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDeletedDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeletedDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeletedDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDeletedDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDeletedDomainsResponse) GetBody() *DescribeCdnDeletedDomainsResponseBody {
	return s.Body
}

func (s *DescribeCdnDeletedDomainsResponse) SetHeaders(v map[string]*string) *DescribeCdnDeletedDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDeletedDomainsResponse) SetStatusCode(v int32) *DescribeCdnDeletedDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDeletedDomainsResponse) SetBody(v *DescribeCdnDeletedDomainsResponseBody) *DescribeCdnDeletedDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDeletedDomainsResponse) Validate() error {
	return dara.Validate(s)
}
