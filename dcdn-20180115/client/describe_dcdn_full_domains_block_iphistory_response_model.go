// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnFullDomainsBlockIPHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnFullDomainsBlockIPHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnFullDomainsBlockIPHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) *DescribeDcdnFullDomainsBlockIPHistoryResponse
	GetBody() *DescribeDcdnFullDomainsBlockIPHistoryResponseBody
}

type DescribeDcdnFullDomainsBlockIPHistoryResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnFullDomainsBlockIPHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnFullDomainsBlockIPHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) GetBody() *DescribeDcdnFullDomainsBlockIPHistoryResponseBody {
	return s.Body
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) SetHeaders(v map[string]*string) *DescribeDcdnFullDomainsBlockIPHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) SetStatusCode(v int32) *DescribeDcdnFullDomainsBlockIPHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) SetBody(v *DescribeDcdnFullDomainsBlockIPHistoryResponseBody) *DescribeDcdnFullDomainsBlockIPHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnFullDomainsBlockIPHistoryResponse) Validate() error {
	return dara.Validate(s)
}
