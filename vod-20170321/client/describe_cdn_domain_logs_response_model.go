// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainLogsResponseBody) *DescribeCdnDomainLogsResponse
	GetBody() *DescribeCdnDomainLogsResponseBody
}

type DescribeCdnDomainLogsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainLogsResponse) GetBody() *DescribeCdnDomainLogsResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainLogsResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainLogsResponse) SetStatusCode(v int32) *DescribeCdnDomainLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainLogsResponse) SetBody(v *DescribeCdnDomainLogsResponseBody) *DescribeCdnDomainLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainLogsResponse) Validate() error {
	return dara.Validate(s)
}
