// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainPvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainPvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainPvDataResponseBody) *DescribeDcdnDomainPvDataResponse
	GetBody() *DescribeDcdnDomainPvDataResponseBody
}

type DescribeDcdnDomainPvDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainPvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainPvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainPvDataResponse) GetBody() *DescribeDcdnDomainPvDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponse) SetStatusCode(v int32) *DescribeDcdnDomainPvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainPvDataResponse) SetBody(v *DescribeDcdnDomainPvDataResponseBody) *DescribeDcdnDomainPvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainPvDataResponse) Validate() error {
	return dara.Validate(s)
}
