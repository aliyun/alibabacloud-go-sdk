// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBgpBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnBgpBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnBgpBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnBgpBpsDataResponseBody) *DescribeDcdnBgpBpsDataResponse
	GetBody() *DescribeDcdnBgpBpsDataResponseBody
}

type DescribeDcdnBgpBpsDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnBgpBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnBgpBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBgpBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBgpBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnBgpBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnBgpBpsDataResponse) GetBody() *DescribeDcdnBgpBpsDataResponseBody {
	return s.Body
}

func (s *DescribeDcdnBgpBpsDataResponse) SetHeaders(v map[string]*string) *DescribeDcdnBgpBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponse) SetStatusCode(v int32) *DescribeDcdnBgpBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponse) SetBody(v *DescribeDcdnBgpBpsDataResponseBody) *DescribeDcdnBgpBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnBgpBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
