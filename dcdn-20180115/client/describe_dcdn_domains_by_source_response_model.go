// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainsBySourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainsBySourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainsBySourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainsBySourceResponseBody) *DescribeDcdnDomainsBySourceResponse
	GetBody() *DescribeDcdnDomainsBySourceResponseBody
}

type DescribeDcdnDomainsBySourceResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainsBySourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainsBySourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainsBySourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainsBySourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainsBySourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainsBySourceResponse) GetBody() *DescribeDcdnDomainsBySourceResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainsBySourceResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainsBySourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponse) SetStatusCode(v int32) *DescribeDcdnDomainsBySourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponse) SetBody(v *DescribeDcdnDomainsBySourceResponseBody) *DescribeDcdnDomainsBySourceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainsBySourceResponse) Validate() error {
	return dara.Validate(s)
}
