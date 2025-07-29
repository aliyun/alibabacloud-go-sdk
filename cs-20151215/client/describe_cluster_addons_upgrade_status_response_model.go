// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonsUpgradeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAddonsUpgradeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAddonsUpgradeStatusResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DescribeClusterAddonsUpgradeStatusResponse
	GetBody() map[string]interface{}
}

type DescribeClusterAddonsUpgradeStatusResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAddonsUpgradeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonsUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonsUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) SetStatusCode(v int32) *DescribeClusterAddonsUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) SetBody(v map[string]interface{}) *DescribeClusterAddonsUpgradeStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAddonsUpgradeStatusResponse) Validate() error {
	return dara.Validate(s)
}
