// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRegionAndIspResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnRegionAndIspResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnRegionAndIspResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnRegionAndIspResponseBody) *DescribeDcdnRegionAndIspResponse
	GetBody() *DescribeDcdnRegionAndIspResponseBody
}

type DescribeDcdnRegionAndIspResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnRegionAndIspResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnRegionAndIspResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRegionAndIspResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRegionAndIspResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnRegionAndIspResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnRegionAndIspResponse) GetBody() *DescribeDcdnRegionAndIspResponseBody {
	return s.Body
}

func (s *DescribeDcdnRegionAndIspResponse) SetHeaders(v map[string]*string) *DescribeDcdnRegionAndIspResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponse) SetStatusCode(v int32) *DescribeDcdnRegionAndIspResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnRegionAndIspResponse) SetBody(v *DescribeDcdnRegionAndIspResponseBody) *DescribeDcdnRegionAndIspResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnRegionAndIspResponse) Validate() error {
	return dara.Validate(s)
}
