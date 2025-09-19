// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridProxyPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DescribeHybridProxyPolicyRequest
	GetClusterName() *string
}

type DescribeHybridProxyPolicyRequest struct {
	// The name of the proxy cluster. You can query the name of the proxy cluster in the Security Center console.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-idc
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s DescribeHybridProxyPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridProxyPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridProxyPolicyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeHybridProxyPolicyRequest) SetClusterName(v string) *DescribeHybridProxyPolicyRequest {
	s.ClusterName = &v
	return s
}

func (s *DescribeHybridProxyPolicyRequest) Validate() error {
	return dara.Validate(s)
}
