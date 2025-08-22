// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafServiceResponseBody) *DescribeDcdnWafServiceResponse
	GetBody() *DescribeDcdnWafServiceResponseBody
}

type DescribeDcdnWafServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafServiceResponse) GetBody() *DescribeDcdnWafServiceResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafServiceResponse) SetStatusCode(v int32) *DescribeDcdnWafServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafServiceResponse) SetBody(v *DescribeDcdnWafServiceResponseBody) *DescribeDcdnWafServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafServiceResponse) Validate() error {
	return dara.Validate(s)
}
