// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafDomainDetailResponseBody) *DescribeDcdnWafDomainDetailResponse
	GetBody() *DescribeDcdnWafDomainDetailResponseBody
}

type DescribeDcdnWafDomainDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafDomainDetailResponse) GetBody() *DescribeDcdnWafDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponse) SetStatusCode(v int32) *DescribeDcdnWafDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponse) SetBody(v *DescribeDcdnWafDomainDetailResponseBody) *DescribeDcdnWafDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
