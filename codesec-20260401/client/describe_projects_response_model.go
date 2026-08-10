// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectsResponseBody) *DescribeProjectsResponse
	GetBody() *DescribeProjectsResponseBody
}

type DescribeProjectsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectsResponse) GetBody() *DescribeProjectsResponseBody {
	return s.Body
}

func (s *DescribeProjectsResponse) SetHeaders(v map[string]*string) *DescribeProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectsResponse) SetStatusCode(v int32) *DescribeProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectsResponse) SetBody(v *DescribeProjectsResponseBody) *DescribeProjectsResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
