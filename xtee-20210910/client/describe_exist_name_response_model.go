// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExistNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExistNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExistNameResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExistNameResponseBody) *DescribeExistNameResponse
	GetBody() *DescribeExistNameResponseBody
}

type DescribeExistNameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExistNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExistNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExistNameResponse) GoString() string {
	return s.String()
}

func (s *DescribeExistNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExistNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExistNameResponse) GetBody() *DescribeExistNameResponseBody {
	return s.Body
}

func (s *DescribeExistNameResponse) SetHeaders(v map[string]*string) *DescribeExistNameResponse {
	s.Headers = v
	return s
}

func (s *DescribeExistNameResponse) SetStatusCode(v int32) *DescribeExistNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExistNameResponse) SetBody(v *DescribeExistNameResponseBody) *DescribeExistNameResponse {
	s.Body = v
	return s
}

func (s *DescribeExistNameResponse) Validate() error {
	return dara.Validate(s)
}
