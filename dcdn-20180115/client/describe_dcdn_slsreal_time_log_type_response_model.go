// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSLSRealTimeLogTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSLSRealTimeLogTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSLSRealTimeLogTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSLSRealTimeLogTypeResponseBody) *DescribeDcdnSLSRealTimeLogTypeResponse
	GetBody() *DescribeDcdnSLSRealTimeLogTypeResponseBody
}

type DescribeDcdnSLSRealTimeLogTypeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSLSRealTimeLogTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSLSRealTimeLogTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSLSRealTimeLogTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) GetBody() *DescribeDcdnSLSRealTimeLogTypeResponseBody {
	return s.Body
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) SetHeaders(v map[string]*string) *DescribeDcdnSLSRealTimeLogTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) SetStatusCode(v int32) *DescribeDcdnSLSRealTimeLogTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) SetBody(v *DescribeDcdnSLSRealTimeLogTypeResponseBody) *DescribeDcdnSLSRealTimeLogTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSLSRealTimeLogTypeResponse) Validate() error {
	return dara.Validate(s)
}
