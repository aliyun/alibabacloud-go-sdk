// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnFullDomainsBlockIPHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnFullDomainsBlockIPHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnFullDomainsBlockIPHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnFullDomainsBlockIPHistoryResponseBody) *DescribeCdnFullDomainsBlockIPHistoryResponse
	GetBody() *DescribeCdnFullDomainsBlockIPHistoryResponseBody
}

type DescribeCdnFullDomainsBlockIPHistoryResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnFullDomainsBlockIPHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnFullDomainsBlockIPHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) GetBody() *DescribeCdnFullDomainsBlockIPHistoryResponseBody {
	return s.Body
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) SetHeaders(v map[string]*string) *DescribeCdnFullDomainsBlockIPHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) SetStatusCode(v int32) *DescribeCdnFullDomainsBlockIPHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) SetBody(v *DescribeCdnFullDomainsBlockIPHistoryResponseBody) *DescribeCdnFullDomainsBlockIPHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnFullDomainsBlockIPHistoryResponse) Validate() error {
	return dara.Validate(s)
}
