// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpInfoResponseBody) *DescribeDcdnIpInfoResponse
	GetBody() *DescribeDcdnIpInfoResponseBody
}

type DescribeDcdnIpInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpInfoResponse) GetBody() *DescribeDcdnIpInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpInfoResponse) SetStatusCode(v int32) *DescribeDcdnIpInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpInfoResponse) SetBody(v *DescribeDcdnIpInfoResponseBody) *DescribeDcdnIpInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpInfoResponse) Validate() error {
	return dara.Validate(s)
}
