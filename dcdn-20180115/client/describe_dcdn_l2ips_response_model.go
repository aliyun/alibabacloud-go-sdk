// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnL2IpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnL2IpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnL2IpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnL2IpsResponseBody) *DescribeDcdnL2IpsResponse
	GetBody() *DescribeDcdnL2IpsResponseBody
}

type DescribeDcdnL2IpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnL2IpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnL2IpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnL2IpsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnL2IpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnL2IpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnL2IpsResponse) GetBody() *DescribeDcdnL2IpsResponseBody {
	return s.Body
}

func (s *DescribeDcdnL2IpsResponse) SetHeaders(v map[string]*string) *DescribeDcdnL2IpsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnL2IpsResponse) SetStatusCode(v int32) *DescribeDcdnL2IpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnL2IpsResponse) SetBody(v *DescribeDcdnL2IpsResponseBody) *DescribeDcdnL2IpsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnL2IpsResponse) Validate() error {
	return dara.Validate(s)
}
