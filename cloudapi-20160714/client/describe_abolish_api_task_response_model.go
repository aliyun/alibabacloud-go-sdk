// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAbolishApiTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAbolishApiTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAbolishApiTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAbolishApiTaskResponseBody) *DescribeAbolishApiTaskResponse
	GetBody() *DescribeAbolishApiTaskResponseBody
}

type DescribeAbolishApiTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAbolishApiTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAbolishApiTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAbolishApiTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeAbolishApiTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAbolishApiTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAbolishApiTaskResponse) GetBody() *DescribeAbolishApiTaskResponseBody {
	return s.Body
}

func (s *DescribeAbolishApiTaskResponse) SetHeaders(v map[string]*string) *DescribeAbolishApiTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeAbolishApiTaskResponse) SetStatusCode(v int32) *DescribeAbolishApiTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAbolishApiTaskResponse) SetBody(v *DescribeAbolishApiTaskResponseBody) *DescribeAbolishApiTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeAbolishApiTaskResponse) Validate() error {
	return dara.Validate(s)
}
