// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDBClusterNodeResponseBody
	GetDBClusterId() *string
	SetNodes(v []*DescribeDBClusterNodeResponseBodyNodes) *DescribeDBClusterNodeResponseBody
	GetNodes() []*DescribeDBClusterNodeResponseBodyNodes
	SetRequestId(v string) *DescribeDBClusterNodeResponseBody
	GetRequestId() *string
}

type DescribeDBClusterNodeResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the node.
	Nodes []*DescribeDBClusterNodeResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 87E44B48-B306-4AD3-A63B-C8**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterNodeResponseBody) GetNodes() []*DescribeDBClusterNodeResponseBodyNodes {
	return s.Nodes
}

func (s *DescribeDBClusterNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterNodeResponseBody) SetDBClusterId(v string) *DescribeDBClusterNodeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetNodes(v []*DescribeDBClusterNodeResponseBodyNodes) *DescribeDBClusterNodeResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) SetRequestId(v string) *DescribeDBClusterNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNodeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterNodeResponseBodyNodes struct {
	// The name of the node.
	//
	// example:
	//
	// master-10*******
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterNodeResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNodeResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNodeResponseBodyNodes) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterNodeResponseBodyNodes) SetName(v string) *DescribeDBClusterNodeResponseBodyNodes {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterNodeResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}
