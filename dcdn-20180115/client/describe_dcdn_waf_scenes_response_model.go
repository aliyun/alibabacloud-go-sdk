// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafScenesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnWafScenesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnWafScenesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnWafScenesResponseBody) *DescribeDcdnWafScenesResponse
	GetBody() *DescribeDcdnWafScenesResponseBody
}

type DescribeDcdnWafScenesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnWafScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnWafScenesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafScenesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafScenesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnWafScenesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnWafScenesResponse) GetBody() *DescribeDcdnWafScenesResponseBody {
	return s.Body
}

func (s *DescribeDcdnWafScenesResponse) SetHeaders(v map[string]*string) *DescribeDcdnWafScenesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnWafScenesResponse) SetStatusCode(v int32) *DescribeDcdnWafScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnWafScenesResponse) SetBody(v *DescribeDcdnWafScenesResponseBody) *DescribeDcdnWafScenesResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnWafScenesResponse) Validate() error {
	return dara.Validate(s)
}
