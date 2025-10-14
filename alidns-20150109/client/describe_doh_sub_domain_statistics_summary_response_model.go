// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohSubDomainStatisticsSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohSubDomainStatisticsSummaryResponseBody) *DescribeDohSubDomainStatisticsSummaryResponse
	GetBody() *DescribeDohSubDomainStatisticsSummaryResponseBody
}

type DescribeDohSubDomainStatisticsSummaryResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohSubDomainStatisticsSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohSubDomainStatisticsSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) GetBody() *DescribeDohSubDomainStatisticsSummaryResponseBody {
	return s.Body
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) SetStatusCode(v int32) *DescribeDohSubDomainStatisticsSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) SetBody(v *DescribeDohSubDomainStatisticsSummaryResponseBody) *DescribeDohSubDomainStatisticsSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeDohSubDomainStatisticsSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
