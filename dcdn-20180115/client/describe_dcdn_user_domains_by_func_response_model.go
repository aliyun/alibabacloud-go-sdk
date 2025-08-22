// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserDomainsByFuncResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsByFuncResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnUserDomainsByFuncResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnUserDomainsByFuncResponseBody) *DescribeDcdnUserDomainsByFuncResponse
	GetBody() *DescribeDcdnUserDomainsByFuncResponseBody
}

type DescribeDcdnUserDomainsByFuncResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnUserDomainsByFuncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnUserDomainsByFuncResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserDomainsByFuncResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserDomainsByFuncResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnUserDomainsByFuncResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnUserDomainsByFuncResponse) GetBody() *DescribeDcdnUserDomainsByFuncResponseBody {
	return s.Body
}

func (s *DescribeDcdnUserDomainsByFuncResponse) SetHeaders(v map[string]*string) *DescribeDcdnUserDomainsByFuncResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponse) SetStatusCode(v int32) *DescribeDcdnUserDomainsByFuncResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponse) SetBody(v *DescribeDcdnUserDomainsByFuncResponseBody) *DescribeDcdnUserDomainsByFuncResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnUserDomainsByFuncResponse) Validate() error {
	return dara.Validate(s)
}
