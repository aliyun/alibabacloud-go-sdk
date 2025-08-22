// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainCnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainCnameResponseBody) *DescribeDcdnDomainCnameResponse
	GetBody() *DescribeDcdnDomainCnameResponseBody
}

type DescribeDcdnDomainCnameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainCnameResponse) GetBody() *DescribeDcdnDomainCnameResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainCnameResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainCnameResponse) SetStatusCode(v int32) *DescribeDcdnDomainCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainCnameResponse) SetBody(v *DescribeDcdnDomainCnameResponseBody) *DescribeDcdnDomainCnameResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainCnameResponse) Validate() error {
	return dara.Validate(s)
}
