// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainLogExTtlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainLogExTtlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainLogExTtlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainLogExTtlResponseBody) *DescribeDcdnDomainLogExTtlResponse
	GetBody() *DescribeDcdnDomainLogExTtlResponseBody
}

type DescribeDcdnDomainLogExTtlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainLogExTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainLogExTtlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainLogExTtlResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainLogExTtlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainLogExTtlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainLogExTtlResponse) GetBody() *DescribeDcdnDomainLogExTtlResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainLogExTtlResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainLogExTtlResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponse) SetStatusCode(v int32) *DescribeDcdnDomainLogExTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponse) SetBody(v *DescribeDcdnDomainLogExTtlResponseBody) *DescribeDcdnDomainLogExTtlResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainLogExTtlResponse) Validate() error {
	return dara.Validate(s)
}
