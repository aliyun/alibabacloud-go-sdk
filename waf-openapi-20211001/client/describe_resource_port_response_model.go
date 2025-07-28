// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePortResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourcePortResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourcePortResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourcePortResponseBody) *DescribeResourcePortResponse
	GetBody() *DescribeResourcePortResponseBody
}

type DescribeResourcePortResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcePortResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcePortResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePortResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourcePortResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourcePortResponse) GetBody() *DescribeResourcePortResponseBody {
	return s.Body
}

func (s *DescribeResourcePortResponse) SetHeaders(v map[string]*string) *DescribeResourcePortResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcePortResponse) SetStatusCode(v int32) *DescribeResourcePortResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcePortResponse) SetBody(v *DescribeResourcePortResponseBody) *DescribeResourcePortResponse {
	s.Body = v
	return s
}

func (s *DescribeResourcePortResponse) Validate() error {
	return dara.Validate(s)
}
