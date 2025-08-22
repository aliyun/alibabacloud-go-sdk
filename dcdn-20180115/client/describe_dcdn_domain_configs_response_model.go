// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainConfigsResponseBody) *DescribeDcdnDomainConfigsResponse
	GetBody() *DescribeDcdnDomainConfigsResponseBody
}

type DescribeDcdnDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainConfigsResponse) GetBody() *DescribeDcdnDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponse) SetStatusCode(v int32) *DescribeDcdnDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainConfigsResponse) SetBody(v *DescribeDcdnDomainConfigsResponseBody) *DescribeDcdnDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
