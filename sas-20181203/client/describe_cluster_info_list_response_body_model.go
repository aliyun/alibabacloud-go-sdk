// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v []*DescribeClusterInfoListResponseBodyClusterList) *DescribeClusterInfoListResponseBody
	GetClusterList() []*DescribeClusterInfoListResponseBodyClusterList
	SetRequestId(v string) *DescribeClusterInfoListResponseBody
	GetRequestId() *string
}

type DescribeClusterInfoListResponseBody struct {
	// An array that consists of the information about clusters.
	ClusterList []*DescribeClusterInfoListResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterInfoListResponseBody) GetClusterList() []*DescribeClusterInfoListResponseBodyClusterList {
	return s.ClusterList
}

func (s *DescribeClusterInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterInfoListResponseBody) SetClusterList(v []*DescribeClusterInfoListResponseBodyClusterList) *DescribeClusterInfoListResponseBody {
	s.ClusterList = v
	return s
}

func (s *DescribeClusterInfoListResponseBody) SetRequestId(v string) *DescribeClusterInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterInfoListResponseBody) Validate() error {
	if s.ClusterList != nil {
		for _, item := range s.ClusterList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterInfoListResponseBodyClusterList struct {
	// The ID of the container cluster.
	//
	// example:
	//
	// cfeb7a9f99ce740e98c5595d0fe37****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the container cluster.
	//
	// example:
	//
	// test111
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **ManagedKubernetes**: managed Kubernetes cluster.
	//
	// 	- **NotManagedKubernetes**: non-managed Kubernetes cluster.
	//
	// 	- **PrivateKubernetes**: private cluster.
	//
	// 	- **kubernetes**: dedicated Kubernetes cluster.
	//
	// 	- **ask**: dedicated serverless Kubernetes (ASK) cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The region in which the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **unavailable**: The cluster is unavailable.
	//
	// 	- **Available**: The cluster is available.
	//
	// 	- **Creating**: The cluster is being created.
	//
	// 	- **CreateFailed**: The cluster failed to be created.
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether container network topology was enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	TargetResult *bool `json:"TargetResult,omitempty" xml:"TargetResult,omitempty"`
}

func (s DescribeClusterInfoListResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterInfoListResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetState() *string {
	return s.State
}

func (s *DescribeClusterInfoListResponseBodyClusterList) GetTargetResult() *bool {
	return s.TargetResult
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetClusterId(v string) *DescribeClusterInfoListResponseBodyClusterList {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetClusterName(v string) *DescribeClusterInfoListResponseBodyClusterList {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetClusterType(v string) *DescribeClusterInfoListResponseBodyClusterList {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetRegionId(v string) *DescribeClusterInfoListResponseBodyClusterList {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetState(v string) *DescribeClusterInfoListResponseBodyClusterList {
	s.State = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) SetTargetResult(v bool) *DescribeClusterInfoListResponseBodyClusterList {
	s.TargetResult = &v
	return s
}

func (s *DescribeClusterInfoListResponseBodyClusterList) Validate() error {
	return dara.Validate(s)
}
