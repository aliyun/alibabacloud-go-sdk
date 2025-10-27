// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyPortDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyPortDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyPortDetailResponseBody) *DescribePropertyPortDetailResponse
	GetBody() *DescribePropertyPortDetailResponseBody
}

type DescribePropertyPortDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyPortDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyPortDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyPortDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyPortDetailResponse) GetBody() *DescribePropertyPortDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyPortDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyPortDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyPortDetailResponse) SetStatusCode(v int32) *DescribePropertyPortDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyPortDetailResponse) SetBody(v *DescribePropertyPortDetailResponseBody) *DescribePropertyPortDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyPortDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
