// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafUsageDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafUsageDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafUsageDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafUsageDataResponseBody) *DescribeDcdnWafUsageDataResponse
	GetBody() *DescribeDcdnWafUsageDataResponseBody
}

type DescribeDcdnWafUsageDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafUsageDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafUsageDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafUsageDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafUsageDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafUsageDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafUsageDataResponse) GetBody() *DescribeDcdnWafUsageDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafUsageDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafUsageDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafUsageDataResponse) SetStatusCode(v int32) *DescribeDcdnWafUsageDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafUsageDataResponse) SetBody(v *DescribeDcdnWafUsageDataResponseBody) *DescribeDcdnWafUsageDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafUsageDataResponse) Validate() error {
	return dara.Validate(s)
}
