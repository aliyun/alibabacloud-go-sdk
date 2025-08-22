// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpaServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnIpaServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnIpaServiceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnIpaServiceResponseBody) *DescribeDcdnIpaServiceResponse
	GetBody() *DescribeDcdnIpaServiceResponseBody
}

type DescribeDcdnIpaServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnIpaServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnIpaServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpaServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpaServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnIpaServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnIpaServiceResponse) GetBody() *DescribeDcdnIpaServiceResponseBody {
	return s.Body
}

func (s *DescribeDcdnIpaServiceResponse) SetHeaders(v map[string]*string) *DescribeDcdnIpaServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnIpaServiceResponse) SetStatusCode(v int32) *DescribeDcdnIpaServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnIpaServiceResponse) SetBody(v *DescribeDcdnIpaServiceResponseBody) *DescribeDcdnIpaServiceResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnIpaServiceResponse) Validate() error {
	return dara.Validate(s)
}
