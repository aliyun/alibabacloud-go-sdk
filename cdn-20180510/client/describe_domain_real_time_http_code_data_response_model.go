// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeHttpCodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainRealTimeHttpCodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainRealTimeHttpCodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainRealTimeHttpCodeDataResponseBody) *DescribeDomainRealTimeHttpCodeDataResponse
	GetBody() *DescribeDomainRealTimeHttpCodeDataResponseBody
}

type DescribeDomainRealTimeHttpCodeDataResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainRealTimeHttpCodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainRealTimeHttpCodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeHttpCodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) GetBody() *DescribeDomainRealTimeHttpCodeDataResponseBody {
	return s.Body
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) SetHeaders(v map[string]*string) *DescribeDomainRealTimeHttpCodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) SetStatusCode(v int32) *DescribeDomainRealTimeHttpCodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) SetBody(v *DescribeDomainRealTimeHttpCodeDataResponseBody) *DescribeDomainRealTimeHttpCodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainRealTimeHttpCodeDataResponse) Validate() error {
	return dara.Validate(s)
}
