// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainStatisticsSummaryResponseBody) *DescribeDomainStatisticsSummaryResponse
	GetBody() *DescribeDomainStatisticsSummaryResponseBody
}

type DescribeDomainStatisticsSummaryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainStatisticsSummaryResponse) GetBody() *DescribeDomainStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeDomainStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponse) SetBody(v *DescribeDomainStatisticsSummaryResponseBody) *DescribeDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainStatisticsSummaryResponse) Validate() error {
	return dara.Validate(s)
}
