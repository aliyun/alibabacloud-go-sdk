// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagListResponseBody) *DescribeTagListResponse
	GetBody() *DescribeTagListResponseBody
}

type DescribeTagListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagListResponse) GetBody() *DescribeTagListResponseBody {
	return s.Body
}

func (s *DescribeTagListResponse) SetHeaders(v map[string]*string) *DescribeTagListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagListResponse) SetStatusCode(v int32) *DescribeTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagListResponse) SetBody(v *DescribeTagListResponseBody) *DescribeTagListResponse {
	s.Body = v
	return s
}

func (s *DescribeTagListResponse) Validate() error {
	return dara.Validate(s)
}
