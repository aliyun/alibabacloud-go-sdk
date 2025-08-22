// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSecFuncInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnSecFuncInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnSecFuncInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnSecFuncInfoResponseBody) *DescribeDcdnSecFuncInfoResponse
	GetBody() *DescribeDcdnSecFuncInfoResponseBody
}

type DescribeDcdnSecFuncInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnSecFuncInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnSecFuncInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSecFuncInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSecFuncInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnSecFuncInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnSecFuncInfoResponse) GetBody() *DescribeDcdnSecFuncInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnSecFuncInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnSecFuncInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponse) SetStatusCode(v int32) *DescribeDcdnSecFuncInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponse) SetBody(v *DescribeDcdnSecFuncInfoResponseBody) *DescribeDcdnSecFuncInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnSecFuncInfoResponse) Validate() error {
	return dara.Validate(s)
}
