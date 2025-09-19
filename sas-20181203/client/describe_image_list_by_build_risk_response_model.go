// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageListByBuildRiskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageListByBuildRiskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageListByBuildRiskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageListByBuildRiskResponseBody) *DescribeImageListByBuildRiskResponse
	GetBody() *DescribeImageListByBuildRiskResponseBody
}

type DescribeImageListByBuildRiskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageListByBuildRiskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageListByBuildRiskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageListByBuildRiskResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageListByBuildRiskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageListByBuildRiskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageListByBuildRiskResponse) GetBody() *DescribeImageListByBuildRiskResponseBody {
	return s.Body
}

func (s *DescribeImageListByBuildRiskResponse) SetHeaders(v map[string]*string) *DescribeImageListByBuildRiskResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageListByBuildRiskResponse) SetStatusCode(v int32) *DescribeImageListByBuildRiskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageListByBuildRiskResponse) SetBody(v *DescribeImageListByBuildRiskResponseBody) *DescribeImageListByBuildRiskResponse {
	s.Body = v
	return s
}

func (s *DescribeImageListByBuildRiskResponse) Validate() error {
	return dara.Validate(s)
}
