// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnL2VipsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnL2VipsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnL2VipsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnL2VipsResponseBody) *DescribeDcdnL2VipsResponse
	GetBody() *DescribeDcdnL2VipsResponseBody
}

type DescribeDcdnL2VipsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnL2VipsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnL2VipsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnL2VipsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnL2VipsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnL2VipsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnL2VipsResponse) GetBody() *DescribeDcdnL2VipsResponseBody {
	return s.Body
}

func (s *DescribeDcdnL2VipsResponse) SetHeaders(v map[string]*string) *DescribeDcdnL2VipsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnL2VipsResponse) SetStatusCode(v int32) *DescribeDcdnL2VipsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnL2VipsResponse) SetBody(v *DescribeDcdnL2VipsResponseBody) *DescribeDcdnL2VipsResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnL2VipsResponse) Validate() error {
	return dara.Validate(s)
}
