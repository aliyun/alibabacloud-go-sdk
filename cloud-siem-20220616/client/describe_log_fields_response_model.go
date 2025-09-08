// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogFieldsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogFieldsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogFieldsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogFieldsResponseBody) *DescribeLogFieldsResponse
	GetBody() *DescribeLogFieldsResponseBody
}

type DescribeLogFieldsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogFieldsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogFieldsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogFieldsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogFieldsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogFieldsResponse) GetBody() *DescribeLogFieldsResponseBody {
	return s.Body
}

func (s *DescribeLogFieldsResponse) SetHeaders(v map[string]*string) *DescribeLogFieldsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogFieldsResponse) SetStatusCode(v int32) *DescribeLogFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogFieldsResponse) SetBody(v *DescribeLogFieldsResponseBody) *DescribeLogFieldsResponse {
	s.Body = v
	return s
}

func (s *DescribeLogFieldsResponse) Validate() error {
	return dara.Validate(s)
}
