// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafSpecInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafSpecInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafSpecInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafSpecInfoResponseBody) *DescribeDcdnWafSpecInfoResponse
	GetBody() *DescribeDcdnWafSpecInfoResponseBody
}

type DescribeDcdnWafSpecInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafSpecInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafSpecInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafSpecInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafSpecInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafSpecInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafSpecInfoResponse) GetBody() *DescribeDcdnWafSpecInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafSpecInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafSpecInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponse) SetStatusCode(v int32) *DescribeDcdnWafSpecInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponse) SetBody(v *DescribeDcdnWafSpecInfoResponseBody) *DescribeDcdnWafSpecInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafSpecInfoResponse) Validate() error {
	return dara.Validate(s)
}
