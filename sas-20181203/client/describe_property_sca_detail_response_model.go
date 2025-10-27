// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyScaDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyScaDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyScaDetailResponseBody) *DescribePropertyScaDetailResponse
	GetBody() *DescribePropertyScaDetailResponseBody
}

type DescribePropertyScaDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyScaDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyScaDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyScaDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyScaDetailResponse) GetBody() *DescribePropertyScaDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyScaDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyScaDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyScaDetailResponse) SetStatusCode(v int32) *DescribePropertyScaDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyScaDetailResponse) SetBody(v *DescribePropertyScaDetailResponseBody) *DescribePropertyScaDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyScaDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
