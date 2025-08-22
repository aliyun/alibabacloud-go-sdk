// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafDomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafDomainResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafDomainResponseBody) *DescribeDcdnWafDomainResponse
	GetBody() *DescribeDcdnWafDomainResponseBody
}

type DescribeDcdnWafDomainResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafDomainResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafDomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafDomainResponse) GetBody() *DescribeDcdnWafDomainResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafDomainResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafDomainResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafDomainResponse) SetStatusCode(v int32) *DescribeDcdnWafDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafDomainResponse) SetBody(v *DescribeDcdnWafDomainResponseBody) *DescribeDcdnWafDomainResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafDomainResponse) Validate() error {
	return dara.Validate(s)
}
