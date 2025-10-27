// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoghubDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoghubDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoghubDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoghubDetailResponseBody) *DescribeLoghubDetailResponse
	GetBody() *DescribeLoghubDetailResponseBody
}

type DescribeLoghubDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoghubDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoghubDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoghubDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoghubDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoghubDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoghubDetailResponse) GetBody() *DescribeLoghubDetailResponseBody {
	return s.Body
}

func (s *DescribeLoghubDetailResponse) SetHeaders(v map[string]*string) *DescribeLoghubDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoghubDetailResponse) SetStatusCode(v int32) *DescribeLoghubDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoghubDetailResponse) SetBody(v *DescribeLoghubDetailResponseBody) *DescribeLoghubDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeLoghubDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
