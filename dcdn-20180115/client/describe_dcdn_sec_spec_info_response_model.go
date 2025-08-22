// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSecSpecInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSecSpecInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSecSpecInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSecSpecInfoResponseBody) *DescribeDcdnSecSpecInfoResponse
	GetBody() *DescribeDcdnSecSpecInfoResponseBody
}

type DescribeDcdnSecSpecInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSecSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSecSpecInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecSpecInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSecSpecInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSecSpecInfoResponse) GetBody() *DescribeDcdnSecSpecInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnSecSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnSecSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponse) SetStatusCode(v int32) *DescribeDcdnSecSpecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponse) SetBody(v *DescribeDcdnSecSpecInfoResponseBody) *DescribeDcdnSecSpecInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSecSpecInfoResponse) Validate() error {
	return dara.Validate(s)
}
