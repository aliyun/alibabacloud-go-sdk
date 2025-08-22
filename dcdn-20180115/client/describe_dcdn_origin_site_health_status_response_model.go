// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnOriginSiteHealthStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnOriginSiteHealthStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnOriginSiteHealthStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnOriginSiteHealthStatusResponseBody) *DescribeDcdnOriginSiteHealthStatusResponse
	GetBody() *DescribeDcdnOriginSiteHealthStatusResponseBody
}

type DescribeDcdnOriginSiteHealthStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnOriginSiteHealthStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnOriginSiteHealthStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnOriginSiteHealthStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) GetBody() *DescribeDcdnOriginSiteHealthStatusResponseBody {
	return s.Body
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) SetHeaders(v map[string]*string) *DescribeDcdnOriginSiteHealthStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) SetStatusCode(v int32) *DescribeDcdnOriginSiteHealthStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) SetBody(v *DescribeDcdnOriginSiteHealthStatusResponseBody) *DescribeDcdnOriginSiteHealthStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponse) Validate() error {
	return dara.Validate(s)
}
