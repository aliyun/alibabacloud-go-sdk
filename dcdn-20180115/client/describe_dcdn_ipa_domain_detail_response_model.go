// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpaDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpaDomainDetailResponseBody) *DescribeDcdnIpaDomainDetailResponse
	GetBody() *DescribeDcdnIpaDomainDetailResponseBody
}

type DescribeDcdnIpaDomainDetailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpaDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpaDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpaDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpaDomainDetailResponse) GetBody() *DescribeDcdnIpaDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpaDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponse) SetStatusCode(v int32) *DescribeDcdnIpaDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponse) SetBody(v *DescribeDcdnIpaDomainDetailResponseBody) *DescribeDcdnIpaDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpaDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
