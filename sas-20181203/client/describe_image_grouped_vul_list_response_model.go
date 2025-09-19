// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageGroupedVulListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageGroupedVulListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageGroupedVulListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageGroupedVulListResponseBody) *DescribeImageGroupedVulListResponse
	GetBody() *DescribeImageGroupedVulListResponseBody
}

type DescribeImageGroupedVulListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageGroupedVulListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageGroupedVulListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageGroupedVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageGroupedVulListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageGroupedVulListResponse) GetBody() *DescribeImageGroupedVulListResponseBody {
	return s.Body
}

func (s *DescribeImageGroupedVulListResponse) SetHeaders(v map[string]*string) *DescribeImageGroupedVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageGroupedVulListResponse) SetStatusCode(v int32) *DescribeImageGroupedVulListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageGroupedVulListResponse) SetBody(v *DescribeImageGroupedVulListResponseBody) *DescribeImageGroupedVulListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageGroupedVulListResponse) Validate() error {
	return dara.Validate(s)
}
