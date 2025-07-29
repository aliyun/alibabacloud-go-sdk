// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonUpgradeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAddonUpgradeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAddonUpgradeStatusResponse
	GetStatusCode() *int32
	SetBody(v map[string]interface{}) *DescribeClusterAddonUpgradeStatusResponse
	GetBody() map[string]interface{}
}

type DescribeClusterAddonUpgradeStatusResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAddonUpgradeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonUpgradeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAddonUpgradeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAddonUpgradeStatusResponse) GetBody() map[string]interface{} {
	return s.Body
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetStatusCode(v int32) *DescribeClusterAddonUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponse) SetBody(v map[string]interface{}) *DescribeClusterAddonUpgradeStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAddonUpgradeStatusResponse) Validate() error {
	return dara.Validate(s)
}
