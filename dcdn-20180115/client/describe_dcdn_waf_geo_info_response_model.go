// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGeoInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafGeoInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafGeoInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafGeoInfoResponseBody) *DescribeDcdnWafGeoInfoResponse
	GetBody() *DescribeDcdnWafGeoInfoResponseBody
}

type DescribeDcdnWafGeoInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafGeoInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafGeoInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafGeoInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafGeoInfoResponse) GetBody() *DescribeDcdnWafGeoInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafGeoInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafGeoInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponse) SetStatusCode(v int32) *DescribeDcdnWafGeoInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponse) SetBody(v *DescribeDcdnWafGeoInfoResponseBody) *DescribeDcdnWafGeoInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponse) Validate() error {
	return dara.Validate(s)
}
