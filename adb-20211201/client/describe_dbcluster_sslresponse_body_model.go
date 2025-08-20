// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *DescribeDBClusterSSLResponseBody
	GetConnectionString() *string
	SetExpireTime(v string) *DescribeDBClusterSSLResponseBody
	GetExpireTime() *string
	SetRequestId(v string) *DescribeDBClusterSSLResponseBody
	GetRequestId() *string
	SetSSLEnabled(v bool) *DescribeDBClusterSSLResponseBody
	GetSSLEnabled() *bool
}

type DescribeDBClusterSSLResponseBody struct {
	// example:
	//
	// amv-*********.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 2022-10-11T08:16:43Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 348303D8-6F42-5141-9B00-A6EECA1E37B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SSLEnabled *bool `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
}

func (s DescribeDBClusterSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterSSLResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDBClusterSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterSSLResponseBody) GetSSLEnabled() *bool {
	return s.SSLEnabled
}

func (s *DescribeDBClusterSSLResponseBody) SetConnectionString(v string) *DescribeDBClusterSSLResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetExpireTime(v string) *DescribeDBClusterSSLResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetRequestId(v string) *DescribeDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetSSLEnabled(v bool) *DescribeDBClusterSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
