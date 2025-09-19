// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulListPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulListPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulListPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulListPageResponseBody) *DescribeVulListPageResponse
	GetBody() *DescribeVulListPageResponseBody
}

type DescribeVulListPageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulListPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulListPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulListPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulListPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulListPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulListPageResponse) GetBody() *DescribeVulListPageResponseBody {
	return s.Body
}

func (s *DescribeVulListPageResponse) SetHeaders(v map[string]*string) *DescribeVulListPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulListPageResponse) SetStatusCode(v int32) *DescribeVulListPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulListPageResponse) SetBody(v *DescribeVulListPageResponseBody) *DescribeVulListPageResponse {
	s.Body = v
	return s
}

func (s *DescribeVulListPageResponse) Validate() error {
	return dara.Validate(s)
}
