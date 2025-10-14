// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarVersionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColumnarVersionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColumnarVersionListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColumnarVersionListResponseBody) *DescribeColumnarVersionListResponse
	GetBody() *DescribeColumnarVersionListResponseBody
}

type DescribeColumnarVersionListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnarVersionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnarVersionListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarVersionListResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnarVersionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColumnarVersionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColumnarVersionListResponse) GetBody() *DescribeColumnarVersionListResponseBody {
	return s.Body
}

func (s *DescribeColumnarVersionListResponse) SetHeaders(v map[string]*string) *DescribeColumnarVersionListResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnarVersionListResponse) SetStatusCode(v int32) *DescribeColumnarVersionListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnarVersionListResponse) SetBody(v *DescribeColumnarVersionListResponseBody) *DescribeColumnarVersionListResponse {
	s.Body = v
	return s
}

func (s *DescribeColumnarVersionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
