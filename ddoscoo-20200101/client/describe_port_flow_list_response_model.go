// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortFlowListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePortFlowListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePortFlowListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePortFlowListResponseBody) *DescribePortFlowListResponse
	GetBody() *DescribePortFlowListResponseBody
}

type DescribePortFlowListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePortFlowListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePortFlowListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePortFlowListResponse) GoString() string {
	return s.String()
}

func (s *DescribePortFlowListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePortFlowListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePortFlowListResponse) GetBody() *DescribePortFlowListResponseBody {
	return s.Body
}

func (s *DescribePortFlowListResponse) SetHeaders(v map[string]*string) *DescribePortFlowListResponse {
	s.Headers = v
	return s
}

func (s *DescribePortFlowListResponse) SetStatusCode(v int32) *DescribePortFlowListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePortFlowListResponse) SetBody(v *DescribePortFlowListResponseBody) *DescribePortFlowListResponse {
	s.Body = v
	return s
}

func (s *DescribePortFlowListResponse) Validate() error {
	return dara.Validate(s)
}
