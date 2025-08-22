// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafFilterInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafFilterInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafFilterInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafFilterInfoResponseBody) *DescribeDcdnWafFilterInfoResponse
	GetBody() *DescribeDcdnWafFilterInfoResponseBody
}

type DescribeDcdnWafFilterInfoResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafFilterInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafFilterInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafFilterInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafFilterInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafFilterInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafFilterInfoResponse) GetBody() *DescribeDcdnWafFilterInfoResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafFilterInfoResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafFilterInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponse) SetStatusCode(v int32) *DescribeDcdnWafFilterInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponse) SetBody(v *DescribeDcdnWafFilterInfoResponseBody) *DescribeDcdnWafFilterInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafFilterInfoResponse) Validate() error {
	return dara.Validate(s)
}
