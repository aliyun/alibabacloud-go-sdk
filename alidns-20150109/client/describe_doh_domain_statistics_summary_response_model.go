// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohDomainStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohDomainStatisticsSummaryResponseBody) *DescribeDohDomainStatisticsSummaryResponse
	GetBody() *DescribeDohDomainStatisticsSummaryResponseBody
}

type DescribeDohDomainStatisticsSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohDomainStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohDomainStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohDomainStatisticsSummaryResponse) GetBody() *DescribeDohDomainStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeDohDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeDohDomainStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponse) SetBody(v *DescribeDohDomainStatisticsSummaryResponseBody) *DescribeDohDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDohDomainStatisticsSummaryResponse) Validate() error {
	return dara.Validate(s)
}
