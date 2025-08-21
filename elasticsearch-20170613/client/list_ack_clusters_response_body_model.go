// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAckClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAckClustersResponseBody
	GetRequestId() *string
	SetResult(v []*ListAckClustersResponseBodyResult) *ListAckClustersResponseBody
	GetResult() []*ListAckClustersResponseBodyResult
}

type ListAckClustersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F93EAA49-284F-4FCE-9E67-FA23FB4BB512
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListAckClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListAckClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAckClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAckClustersResponseBody) GetResult() []*ListAckClustersResponseBodyResult {
	return s.Result
}

func (s *ListAckClustersResponseBody) SetRequestId(v string) *ListAckClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAckClustersResponseBody) SetResult(v []*ListAckClustersResponseBodyResult) *ListAckClustersResponseBody {
	s.Result = v
	return s
}

func (s *ListAckClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAckClustersResponseBodyResult struct {
	// The ID of cluster.
	//
	// example:
	//
	// c5ea2c2d9a3cf499481292f60425d****
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The type of the cluster. The value is fixed as ManagedKubernetes.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the VPC to which the cluster belongs.
	//
	// example:
	//
	// vpc-bp12nu14urf0upaf4****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListAckClustersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAckClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAckClustersResponseBodyResult) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAckClustersResponseBodyResult) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListAckClustersResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListAckClustersResponseBodyResult) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAckClustersResponseBodyResult) SetClusterId(v string) *ListAckClustersResponseBodyResult {
	s.ClusterId = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetClusterType(v string) *ListAckClustersResponseBodyResult {
	s.ClusterType = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetName(v string) *ListAckClustersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) SetVpcId(v string) *ListAckClustersResponseBodyResult {
	s.VpcId = &v
	return s
}

func (s *ListAckClustersResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
