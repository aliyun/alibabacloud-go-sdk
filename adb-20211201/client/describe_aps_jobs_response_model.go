// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsJobsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsJobsResponseBody) *DescribeApsJobsResponse
	GetBody() *DescribeApsJobsResponseBody
}

type DescribeApsJobsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsJobsResponse) GetBody() *DescribeApsJobsResponseBody {
	return s.Body
}

func (s *DescribeApsJobsResponse) SetHeaders(v map[string]*string) *DescribeApsJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsJobsResponse) SetStatusCode(v int32) *DescribeApsJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsJobsResponse) SetBody(v *DescribeApsJobsResponseBody) *DescribeApsJobsResponse {
	s.Body = v
	return s
}

func (s *DescribeApsJobsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
