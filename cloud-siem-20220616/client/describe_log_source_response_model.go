// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogSourceResponseBody) *DescribeLogSourceResponse
	GetBody() *DescribeLogSourceResponseBody
}

type DescribeLogSourceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogSourceResponse) GetBody() *DescribeLogSourceResponseBody {
	return s.Body
}

func (s *DescribeLogSourceResponse) SetHeaders(v map[string]*string) *DescribeLogSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogSourceResponse) SetStatusCode(v int32) *DescribeLogSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogSourceResponse) SetBody(v *DescribeLogSourceResponseBody) *DescribeLogSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeLogSourceResponse) Validate() error {
	return dara.Validate(s)
}
