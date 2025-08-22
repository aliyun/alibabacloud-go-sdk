// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainMultiUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainMultiUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainMultiUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainMultiUsageDataResponseBody) *DescribeDcdnDomainMultiUsageDataResponse
	GetBody() *DescribeDcdnDomainMultiUsageDataResponseBody
}

type DescribeDcdnDomainMultiUsageDataResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainMultiUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainMultiUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainMultiUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) GetBody() *DescribeDcdnDomainMultiUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainMultiUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainMultiUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) SetBody(v *DescribeDcdnDomainMultiUsageDataResponseBody) *DescribeDcdnDomainMultiUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainMultiUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
