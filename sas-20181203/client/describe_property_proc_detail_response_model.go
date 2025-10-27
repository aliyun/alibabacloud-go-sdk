// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyProcDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyProcDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyProcDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyProcDetailResponseBody) *DescribePropertyProcDetailResponse
	GetBody() *DescribePropertyProcDetailResponseBody
}

type DescribePropertyProcDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyProcDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyProcDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyProcDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyProcDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyProcDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyProcDetailResponse) GetBody() *DescribePropertyProcDetailResponseBody {
	return s.Body
}

func (s *DescribePropertyProcDetailResponse) SetHeaders(v map[string]*string) *DescribePropertyProcDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyProcDetailResponse) SetStatusCode(v int32) *DescribePropertyProcDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyProcDetailResponse) SetBody(v *DescribePropertyProcDetailResponseBody) *DescribePropertyProcDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyProcDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
