// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmgVulItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEmgVulItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEmgVulItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEmgVulItemResponseBody) *DescribeEmgVulItemResponse
	GetBody() *DescribeEmgVulItemResponseBody
}

type DescribeEmgVulItemResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEmgVulItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEmgVulItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgVulItemResponse) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEmgVulItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEmgVulItemResponse) GetBody() *DescribeEmgVulItemResponseBody {
	return s.Body
}

func (s *DescribeEmgVulItemResponse) SetHeaders(v map[string]*string) *DescribeEmgVulItemResponse {
	s.Headers = v
	return s
}

func (s *DescribeEmgVulItemResponse) SetStatusCode(v int32) *DescribeEmgVulItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEmgVulItemResponse) SetBody(v *DescribeEmgVulItemResponseBody) *DescribeEmgVulItemResponse {
	s.Body = v
	return s
}

func (s *DescribeEmgVulItemResponse) Validate() error {
	return dara.Validate(s)
}
