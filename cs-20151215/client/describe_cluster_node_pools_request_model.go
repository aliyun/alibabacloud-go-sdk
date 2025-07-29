// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNodePoolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodepoolName(v string) *DescribeClusterNodePoolsRequest
	GetNodepoolName() *string
}

type DescribeClusterNodePoolsRequest struct {
	// Node pool name.
	//
	// example:
	//
	// nodepool-test
	NodepoolName *string `json:"NodepoolName,omitempty" xml:"NodepoolName,omitempty"`
}

func (s DescribeClusterNodePoolsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNodePoolsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodePoolsRequest) GetNodepoolName() *string {
	return s.NodepoolName
}

func (s *DescribeClusterNodePoolsRequest) SetNodepoolName(v string) *DescribeClusterNodePoolsRequest {
	s.NodepoolName = &v
	return s
}

func (s *DescribeClusterNodePoolsRequest) Validate() error {
	return dara.Validate(s)
}
