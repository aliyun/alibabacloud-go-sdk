// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBgpPackByIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBgpPackByIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBgpPackByIpResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBgpPackByIpResponseBody) *DescribeBgpPackByIpResponse
	GetBody() *DescribeBgpPackByIpResponseBody
}

type DescribeBgpPackByIpResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBgpPackByIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBgpPackByIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBgpPackByIpResponse) GoString() string {
	return s.String()
}

func (s *DescribeBgpPackByIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBgpPackByIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBgpPackByIpResponse) GetBody() *DescribeBgpPackByIpResponseBody {
	return s.Body
}

func (s *DescribeBgpPackByIpResponse) SetHeaders(v map[string]*string) *DescribeBgpPackByIpResponse {
	s.Headers = v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetStatusCode(v int32) *DescribeBgpPackByIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBgpPackByIpResponse) SetBody(v *DescribeBgpPackByIpResponseBody) *DescribeBgpPackByIpResponse {
	s.Body = v
	return s
}

func (s *DescribeBgpPackByIpResponse) Validate() error {
	return dara.Validate(s)
}
