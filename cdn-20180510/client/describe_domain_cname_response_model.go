// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCnameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCnameResponseBody) *DescribeDomainCnameResponse
	GetBody() *DescribeDomainCnameResponseBody
}

type DescribeDomainCnameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCnameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCnameResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCnameResponse) GetBody() *DescribeDomainCnameResponseBody {
	return s.Body
}

func (s *DescribeDomainCnameResponse) SetHeaders(v map[string]*string) *DescribeDomainCnameResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCnameResponse) SetStatusCode(v int32) *DescribeDomainCnameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCnameResponse) SetBody(v *DescribeDomainCnameResponseBody) *DescribeDomainCnameResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCnameResponse) Validate() error {
	return dara.Validate(s)
}
