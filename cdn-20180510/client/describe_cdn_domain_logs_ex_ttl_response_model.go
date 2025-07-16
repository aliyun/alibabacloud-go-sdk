// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainLogsExTtlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainLogsExTtlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainLogsExTtlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainLogsExTtlResponseBody) *DescribeCdnDomainLogsExTtlResponse
	GetBody() *DescribeCdnDomainLogsExTtlResponseBody
}

type DescribeCdnDomainLogsExTtlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainLogsExTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainLogsExTtlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainLogsExTtlResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainLogsExTtlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainLogsExTtlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainLogsExTtlResponse) GetBody() *DescribeCdnDomainLogsExTtlResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainLogsExTtlResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainLogsExTtlResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponse) SetStatusCode(v int32) *DescribeCdnDomainLogsExTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponse) SetBody(v *DescribeCdnDomainLogsExTtlResponseBody) *DescribeCdnDomainLogsExTtlResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainLogsExTtlResponse) Validate() error {
	return dara.Validate(s)
}
