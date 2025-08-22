// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainDetailResponseBody) *DescribeDcdnDomainDetailResponse
	GetBody() *DescribeDcdnDomainDetailResponseBody
}

type DescribeDcdnDomainDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainDetailResponse) GetBody() *DescribeDcdnDomainDetailResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainDetailResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainDetailResponse) SetStatusCode(v int32) *DescribeDcdnDomainDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainDetailResponse) SetBody(v *DescribeDcdnDomainDetailResponseBody) *DescribeDcdnDomainDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainDetailResponse) Validate() error {
	return dara.Validate(s)
}
