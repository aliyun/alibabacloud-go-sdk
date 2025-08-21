// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApmResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApmResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApmResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApmResponseBody) *DescribeApmResponse
	GetBody() *DescribeApmResponseBody
}

type DescribeApmResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApmResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApmResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApmResponse) GoString() string {
	return s.String()
}

func (s *DescribeApmResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApmResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApmResponse) GetBody() *DescribeApmResponseBody {
	return s.Body
}

func (s *DescribeApmResponse) SetHeaders(v map[string]*string) *DescribeApmResponse {
	s.Headers = v
	return s
}

func (s *DescribeApmResponse) SetStatusCode(v int32) *DescribeApmResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApmResponse) SetBody(v *DescribeApmResponseBody) *DescribeApmResponse {
	s.Body = v
	return s
}

func (s *DescribeApmResponse) Validate() error {
	return dara.Validate(s)
}
