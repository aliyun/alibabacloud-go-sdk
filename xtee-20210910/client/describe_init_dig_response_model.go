// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInitDigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInitDigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInitDigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInitDigResponseBody) *DescribeInitDigResponse
	GetBody() *DescribeInitDigResponseBody
}

type DescribeInitDigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInitDigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInitDigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInitDigResponse) GoString() string {
	return s.String()
}

func (s *DescribeInitDigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInitDigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInitDigResponse) GetBody() *DescribeInitDigResponseBody {
	return s.Body
}

func (s *DescribeInitDigResponse) SetHeaders(v map[string]*string) *DescribeInitDigResponse {
	s.Headers = v
	return s
}

func (s *DescribeInitDigResponse) SetStatusCode(v int32) *DescribeInitDigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInitDigResponse) SetBody(v *DescribeInitDigResponseBody) *DescribeInitDigResponse {
	s.Body = v
	return s
}

func (s *DescribeInitDigResponse) Validate() error {
	return dara.Validate(s)
}
