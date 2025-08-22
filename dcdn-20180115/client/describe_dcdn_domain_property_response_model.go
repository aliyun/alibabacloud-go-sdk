// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainPropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainPropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainPropertyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainPropertyResponseBody) *DescribeDcdnDomainPropertyResponse
	GetBody() *DescribeDcdnDomainPropertyResponseBody
}

type DescribeDcdnDomainPropertyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainPropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainPropertyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainPropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainPropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainPropertyResponse) GetBody() *DescribeDcdnDomainPropertyResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainPropertyResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainPropertyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainPropertyResponse) SetStatusCode(v int32) *DescribeDcdnDomainPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainPropertyResponse) SetBody(v *DescribeDcdnDomainPropertyResponseBody) *DescribeDcdnDomainPropertyResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainPropertyResponse) Validate() error {
	return dara.Validate(s)
}
