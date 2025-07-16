// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDbListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDbListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDbListResponseBody) *DescribeDbListResponse
	GetBody() *DescribeDbListResponseBody
}

type DescribeDbListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDbListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDbListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDbListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDbListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDbListResponse) GetBody() *DescribeDbListResponseBody {
	return s.Body
}

func (s *DescribeDbListResponse) SetHeaders(v map[string]*string) *DescribeDbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDbListResponse) SetStatusCode(v int32) *DescribeDbListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDbListResponse) SetBody(v *DescribeDbListResponseBody) *DescribeDbListResponse {
	s.Body = v
	return s
}

func (s *DescribeDbListResponse) Validate() error {
	return dara.Validate(s)
}
