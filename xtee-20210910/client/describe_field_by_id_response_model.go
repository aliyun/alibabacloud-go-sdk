// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFieldByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFieldByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFieldByIdResponseBody) *DescribeFieldByIdResponse
	GetBody() *DescribeFieldByIdResponseBody
}

type DescribeFieldByIdResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFieldByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFieldByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeFieldByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFieldByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFieldByIdResponse) GetBody() *DescribeFieldByIdResponseBody {
	return s.Body
}

func (s *DescribeFieldByIdResponse) SetHeaders(v map[string]*string) *DescribeFieldByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeFieldByIdResponse) SetStatusCode(v int32) *DescribeFieldByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFieldByIdResponse) SetBody(v *DescribeFieldByIdResponseBody) *DescribeFieldByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeFieldByIdResponse) Validate() error {
	return dara.Validate(s)
}
