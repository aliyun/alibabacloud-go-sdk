// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayEventListResponseBody) *DescribePlayEventListResponse
	GetBody() *DescribePlayEventListResponseBody
}

type DescribePlayEventListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayEventListResponse) GetBody() *DescribePlayEventListResponseBody {
	return s.Body
}

func (s *DescribePlayEventListResponse) SetHeaders(v map[string]*string) *DescribePlayEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayEventListResponse) SetStatusCode(v int32) *DescribePlayEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayEventListResponse) SetBody(v *DescribePlayEventListResponseBody) *DescribePlayEventListResponse {
	s.Body = v
	return s
}

func (s *DescribePlayEventListResponse) Validate() error {
	return dara.Validate(s)
}
