// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVulTargetConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVulTargetConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVulTargetConfigResponseBody) *DescribeVulTargetConfigResponse
	GetBody() *DescribeVulTargetConfigResponseBody
}

type DescribeVulTargetConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVulTargetConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVulTargetConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVulTargetConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVulTargetConfigResponse) GetBody() *DescribeVulTargetConfigResponseBody {
	return s.Body
}

func (s *DescribeVulTargetConfigResponse) SetHeaders(v map[string]*string) *DescribeVulTargetConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVulTargetConfigResponse) SetStatusCode(v int32) *DescribeVulTargetConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVulTargetConfigResponse) SetBody(v *DescribeVulTargetConfigResponseBody) *DescribeVulTargetConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVulTargetConfigResponse) Validate() error {
	return dara.Validate(s)
}
