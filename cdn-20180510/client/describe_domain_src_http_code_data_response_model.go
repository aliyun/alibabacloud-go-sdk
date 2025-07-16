// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSrcHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSrcHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSrcHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSrcHttpCodeDataResponseBody) *DescribeDomainSrcHttpCodeDataResponse
	GetBody() *DescribeDomainSrcHttpCodeDataResponseBody
}

type DescribeDomainSrcHttpCodeDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSrcHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSrcHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSrcHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSrcHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSrcHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSrcHttpCodeDataResponse) GetBody() *DescribeDomainSrcHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDomainSrcHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainSrcHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDomainSrcHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponse) SetBody(v *DescribeDomainSrcHttpCodeDataResponseBody) *DescribeDomainSrcHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSrcHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
