// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOnceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOnceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOnceTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOnceTaskResponseBody) *DescribeOnceTaskResponse
	GetBody() *DescribeOnceTaskResponseBody
}

type DescribeOnceTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOnceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOnceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOnceTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeOnceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOnceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOnceTaskResponse) GetBody() *DescribeOnceTaskResponseBody {
	return s.Body
}

func (s *DescribeOnceTaskResponse) SetHeaders(v map[string]*string) *DescribeOnceTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeOnceTaskResponse) SetStatusCode(v int32) *DescribeOnceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOnceTaskResponse) SetBody(v *DescribeOnceTaskResponseBody) *DescribeOnceTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeOnceTaskResponse) Validate() error {
	return dara.Validate(s)
}
