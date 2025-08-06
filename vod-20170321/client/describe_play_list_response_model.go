// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayListResponseBody) *DescribePlayListResponse
	GetBody() *DescribePlayListResponseBody
}

type DescribePlayListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayListResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayListResponse) GetBody() *DescribePlayListResponseBody {
	return s.Body
}

func (s *DescribePlayListResponse) SetHeaders(v map[string]*string) *DescribePlayListResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayListResponse) SetStatusCode(v int32) *DescribePlayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayListResponse) SetBody(v *DescribePlayListResponseBody) *DescribePlayListResponse {
	s.Body = v
	return s
}

func (s *DescribePlayListResponse) Validate() error {
	return dara.Validate(s)
}
