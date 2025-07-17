// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIMRobotsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIMRobotsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIMRobotsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIMRobotsResponseBody) *DescribeIMRobotsResponse
	GetBody() *DescribeIMRobotsResponseBody
}

type DescribeIMRobotsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIMRobotsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIMRobotsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIMRobotsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIMRobotsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIMRobotsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIMRobotsResponse) GetBody() *DescribeIMRobotsResponseBody {
	return s.Body
}

func (s *DescribeIMRobotsResponse) SetHeaders(v map[string]*string) *DescribeIMRobotsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIMRobotsResponse) SetStatusCode(v int32) *DescribeIMRobotsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIMRobotsResponse) SetBody(v *DescribeIMRobotsResponseBody) *DescribeIMRobotsResponse {
	s.Body = v
	return s
}

func (s *DescribeIMRobotsResponse) Validate() error {
	return dara.Validate(s)
}
