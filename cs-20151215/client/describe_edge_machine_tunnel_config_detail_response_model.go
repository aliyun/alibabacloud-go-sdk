// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdgeMachineTunnelConfigDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEdgeMachineTunnelConfigDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEdgeMachineTunnelConfigDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEdgeMachineTunnelConfigDetailResponseBody) *DescribeEdgeMachineTunnelConfigDetailResponse
	GetBody() *DescribeEdgeMachineTunnelConfigDetailResponseBody
}

type DescribeEdgeMachineTunnelConfigDetailResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEdgeMachineTunnelConfigDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEdgeMachineTunnelConfigDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdgeMachineTunnelConfigDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) GetBody() *DescribeEdgeMachineTunnelConfigDetailResponseBody {
	return s.Body
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) SetHeaders(v map[string]*string) *DescribeEdgeMachineTunnelConfigDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) SetStatusCode(v int32) *DescribeEdgeMachineTunnelConfigDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) SetBody(v *DescribeEdgeMachineTunnelConfigDetailResponseBody) *DescribeEdgeMachineTunnelConfigDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeEdgeMachineTunnelConfigDetailResponse) Validate() error {
	return dara.Validate(s)
}
