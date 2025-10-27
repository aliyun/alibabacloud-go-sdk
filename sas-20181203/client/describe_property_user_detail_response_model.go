// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyUserDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyUserDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyUserDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyUserDetailResponseBody) *DescribePropertyUserDetailResponse
	GetBody() *DescribePropertyUserDetailResponseBody
}

type DescribePropertyUserDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyUserDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyUserDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyUserDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyUserDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyUserDetailResponse) GetBody() *DescribePropertyUserDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyUserDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyUserDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyUserDetailResponse) SetStatusCode(v int32) *DescribePropertyUserDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyUserDetailResponse) SetBody(v *DescribePropertyUserDetailResponseBody) *DescribePropertyUserDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyUserDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
