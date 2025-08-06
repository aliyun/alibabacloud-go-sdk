// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserTotalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayUserTotalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayUserTotalResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayUserTotalResponseBody) *DescribePlayUserTotalResponse
	GetBody() *DescribePlayUserTotalResponseBody
}

type DescribePlayUserTotalResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayUserTotalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayUserTotalResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserTotalResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayUserTotalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayUserTotalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayUserTotalResponse) GetBody() *DescribePlayUserTotalResponseBody {
	return s.Body
}

func (s *DescribePlayUserTotalResponse) SetHeaders(v map[string]*string) *DescribePlayUserTotalResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayUserTotalResponse) SetStatusCode(v int32) *DescribePlayUserTotalResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayUserTotalResponse) SetBody(v *DescribePlayUserTotalResponseBody) *DescribePlayUserTotalResponse {
	s.Body = v
	return s
}

func (s *DescribePlayUserTotalResponse) Validate() error {
	return dara.Validate(s)
}
