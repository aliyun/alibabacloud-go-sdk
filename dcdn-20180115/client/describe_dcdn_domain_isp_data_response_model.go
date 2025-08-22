// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainIspDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainIspDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainIspDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainIspDataResponseBody) *DescribeDcdnDomainIspDataResponse
	GetBody() *DescribeDcdnDomainIspDataResponseBody
}

type DescribeDcdnDomainIspDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainIspDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainIspDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainIspDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainIspDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainIspDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainIspDataResponse) GetBody() *DescribeDcdnDomainIspDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainIspDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainIspDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainIspDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainIspDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainIspDataResponse) SetBody(v *DescribeDcdnDomainIspDataResponseBody) *DescribeDcdnDomainIspDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainIspDataResponse) Validate() error {
	return dara.Validate(s)
}
