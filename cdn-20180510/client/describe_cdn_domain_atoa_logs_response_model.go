// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainAtoaLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainAtoaLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainAtoaLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainAtoaLogsResponseBody) *DescribeCdnDomainAtoaLogsResponse
	GetBody() *DescribeCdnDomainAtoaLogsResponseBody
}

type DescribeCdnDomainAtoaLogsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainAtoaLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainAtoaLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainAtoaLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainAtoaLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainAtoaLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainAtoaLogsResponse) GetBody() *DescribeCdnDomainAtoaLogsResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainAtoaLogsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainAtoaLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainAtoaLogsResponse) SetStatusCode(v int32) *DescribeCdnDomainAtoaLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainAtoaLogsResponse) SetBody(v *DescribeCdnDomainAtoaLogsResponseBody) *DescribeCdnDomainAtoaLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainAtoaLogsResponse) Validate() error {
	return dara.Validate(s)
}
