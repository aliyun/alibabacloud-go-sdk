// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineItemListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBaselineItemListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBaselineItemListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBaselineItemListResponseBody) *DescribeImageBaselineItemListResponse
	GetBody() *DescribeImageBaselineItemListResponseBody
}

type DescribeImageBaselineItemListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBaselineItemListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBaselineItemListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineItemListResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineItemListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBaselineItemListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBaselineItemListResponse) GetBody() *DescribeImageBaselineItemListResponseBody {
	return s.Body
}

func (s *DescribeImageBaselineItemListResponse) SetHeaders(v map[string]*string) *DescribeImageBaselineItemListResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBaselineItemListResponse) SetStatusCode(v int32) *DescribeImageBaselineItemListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBaselineItemListResponse) SetBody(v *DescribeImageBaselineItemListResponseBody) *DescribeImageBaselineItemListResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBaselineItemListResponse) Validate() error {
	return dara.Validate(s)
}
