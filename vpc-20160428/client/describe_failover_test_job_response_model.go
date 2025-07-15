// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFailoverTestJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFailoverTestJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFailoverTestJobResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFailoverTestJobResponseBody) *DescribeFailoverTestJobResponse
	GetBody() *DescribeFailoverTestJobResponseBody
}

type DescribeFailoverTestJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFailoverTestJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFailoverTestJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFailoverTestJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeFailoverTestJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFailoverTestJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFailoverTestJobResponse) GetBody() *DescribeFailoverTestJobResponseBody {
	return s.Body
}

func (s *DescribeFailoverTestJobResponse) SetHeaders(v map[string]*string) *DescribeFailoverTestJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeFailoverTestJobResponse) SetStatusCode(v int32) *DescribeFailoverTestJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFailoverTestJobResponse) SetBody(v *DescribeFailoverTestJobResponseBody) *DescribeFailoverTestJobResponse {
	s.Body = v
	return s
}

func (s *DescribeFailoverTestJobResponse) Validate() error {
	return dara.Validate(s)
}
