// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSrcTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSrcTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSrcTrafficDataResponseBody) *DescribeDomainSrcTrafficDataResponse
	GetBody() *DescribeDomainSrcTrafficDataResponseBody
}

type DescribeDomainSrcTrafficDataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSrcTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSrcTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSrcTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSrcTrafficDataResponse) GetBody() *DescribeDomainSrcTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDomainSrcTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponse) SetStatusCode(v int32) *DescribeDomainSrcTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponse) SetBody(v *DescribeDomainSrcTrafficDataResponseBody) *DescribeDomainSrcTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSrcTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
