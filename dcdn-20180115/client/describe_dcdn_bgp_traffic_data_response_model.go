// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnBgpTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnBgpTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnBgpTrafficDataResponseBody) *DescribeDcdnBgpTrafficDataResponse
	GetBody() *DescribeDcdnBgpTrafficDataResponseBody
}

type DescribeDcdnBgpTrafficDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnBgpTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnBgpTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnBgpTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnBgpTrafficDataResponse) GetBody() *DescribeDcdnBgpTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnBgpTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnBgpTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponse) SetStatusCode(v int32) *DescribeDcdnBgpTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponse) SetBody(v *DescribeDcdnBgpTrafficDataResponseBody) *DescribeDcdnBgpTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnBgpTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
