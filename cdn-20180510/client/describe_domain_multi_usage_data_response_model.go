// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMultiUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainMultiUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainMultiUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainMultiUsageDataResponseBody) *DescribeDomainMultiUsageDataResponse
	GetBody() *DescribeDomainMultiUsageDataResponseBody
}

type DescribeDomainMultiUsageDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainMultiUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainMultiUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMultiUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainMultiUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainMultiUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainMultiUsageDataResponse) GetBody() *DescribeDomainMultiUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDomainMultiUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDomainMultiUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponse) SetStatusCode(v int32) *DescribeDomainMultiUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainMultiUsageDataResponse) SetBody(v *DescribeDomainMultiUsageDataResponseBody) *DescribeDomainMultiUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainMultiUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
