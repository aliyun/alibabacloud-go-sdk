// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClassListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClassListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClassListResponseBody) *DescribeClassListResponse
	GetBody() *DescribeClassListResponseBody
}

type DescribeClassListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClassListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClassListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassListResponse) GoString() string {
	return s.String()
}

func (s *DescribeClassListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClassListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClassListResponse) GetBody() *DescribeClassListResponseBody {
	return s.Body
}

func (s *DescribeClassListResponse) SetHeaders(v map[string]*string) *DescribeClassListResponse {
	s.Headers = v
	return s
}

func (s *DescribeClassListResponse) SetStatusCode(v int32) *DescribeClassListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClassListResponse) SetBody(v *DescribeClassListResponseBody) *DescribeClassListResponse {
	s.Body = v
	return s
}

func (s *DescribeClassListResponse) Validate() error {
	return dara.Validate(s)
}
