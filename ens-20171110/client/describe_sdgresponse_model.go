// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSDGResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSDGResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSDGResponseBody) *DescribeSDGResponse
	GetBody() *DescribeSDGResponseBody
}

type DescribeSDGResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSDGResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSDGResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGResponse) GoString() string {
	return s.String()
}

func (s *DescribeSDGResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSDGResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSDGResponse) GetBody() *DescribeSDGResponseBody {
	return s.Body
}

func (s *DescribeSDGResponse) SetHeaders(v map[string]*string) *DescribeSDGResponse {
	s.Headers = v
	return s
}

func (s *DescribeSDGResponse) SetStatusCode(v int32) *DescribeSDGResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSDGResponse) SetBody(v *DescribeSDGResponseBody) *DescribeSDGResponse {
	s.Body = v
	return s
}

func (s *DescribeSDGResponse) Validate() error {
	return dara.Validate(s)
}
