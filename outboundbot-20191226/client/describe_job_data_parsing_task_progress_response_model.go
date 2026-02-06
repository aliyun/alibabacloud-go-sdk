// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobDataParsingTaskProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeJobDataParsingTaskProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeJobDataParsingTaskProgressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeJobDataParsingTaskProgressResponseBody) *DescribeJobDataParsingTaskProgressResponse
	GetBody() *DescribeJobDataParsingTaskProgressResponseBody
}

type DescribeJobDataParsingTaskProgressResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeJobDataParsingTaskProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeJobDataParsingTaskProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobDataParsingTaskProgressResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobDataParsingTaskProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeJobDataParsingTaskProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeJobDataParsingTaskProgressResponse) GetBody() *DescribeJobDataParsingTaskProgressResponseBody {
	return s.Body
}

func (s *DescribeJobDataParsingTaskProgressResponse) SetHeaders(v map[string]*string) *DescribeJobDataParsingTaskProgressResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponse) SetStatusCode(v int32) *DescribeJobDataParsingTaskProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponse) SetBody(v *DescribeJobDataParsingTaskProgressResponseBody) *DescribeJobDataParsingTaskProgressResponse {
	s.Body = v
	return s
}

func (s *DescribeJobDataParsingTaskProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
