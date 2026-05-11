// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExportProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExportProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExportProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExportProgressResponseBody) *DescribeExportProgressResponse
	GetBody() *DescribeExportProgressResponseBody
}

type DescribeExportProgressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExportProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExportProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExportProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeExportProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExportProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExportProgressResponse) GetBody() *DescribeExportProgressResponseBody {
	return s.Body
}

func (s *DescribeExportProgressResponse) SetHeaders(v map[string]*string) *DescribeExportProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeExportProgressResponse) SetStatusCode(v int32) *DescribeExportProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExportProgressResponse) SetBody(v *DescribeExportProgressResponseBody) *DescribeExportProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeExportProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
