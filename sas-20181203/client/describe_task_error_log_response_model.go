// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskErrorLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskErrorLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskErrorLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskErrorLogResponseBody) *DescribeTaskErrorLogResponse
	GetBody() *DescribeTaskErrorLogResponseBody
}

type DescribeTaskErrorLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskErrorLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskErrorLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskErrorLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskErrorLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskErrorLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskErrorLogResponse) GetBody() *DescribeTaskErrorLogResponseBody {
	return s.Body
}

func (s *DescribeTaskErrorLogResponse) SetHeaders(v map[string]*string) *DescribeTaskErrorLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskErrorLogResponse) SetStatusCode(v int32) *DescribeTaskErrorLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskErrorLogResponse) SetBody(v *DescribeTaskErrorLogResponseBody) *DescribeTaskErrorLogResponse {
	s.Body = v
	return s
}

func (s *DescribeTaskErrorLogResponse) Validate() error {
	return dara.Validate(s)
}
