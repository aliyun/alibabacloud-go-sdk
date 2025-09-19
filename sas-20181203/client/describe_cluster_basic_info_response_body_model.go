// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterInfo(v *DescribeClusterBasicInfoResponseBodyClusterInfo) *DescribeClusterBasicInfoResponseBody
	GetClusterInfo() *DescribeClusterBasicInfoResponseBodyClusterInfo
	SetRequestId(v string) *DescribeClusterBasicInfoResponseBody
	GetRequestId() *string
}

type DescribeClusterBasicInfoResponseBody struct {
	// The detailed information about the cluster.
	ClusterInfo *DescribeClusterBasicInfoResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterBasicInfoResponseBody) GetClusterInfo() *DescribeClusterBasicInfoResponseBodyClusterInfo {
	return s.ClusterInfo
}

func (s *DescribeClusterBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterBasicInfoResponseBody) SetClusterInfo(v *DescribeClusterBasicInfoResponseBodyClusterInfo) *DescribeClusterBasicInfoResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterBasicInfoResponseBody) SetRequestId(v string) *DescribeClusterBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterBasicInfoResponseBodyClusterInfo struct {
	// The ID of cluster.
	//
	// example:
	//
	// c870ec78ecbcb41d2a35c679823ef****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// testackpro
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- **ManagedKubernetes**: managed Kubernetes cluster
	//
	// 	- **NotManagedKubernetes**: non-managed Kubernetes cluster
	//
	// 	- **PrivateKubernetes**: private cluster
	//
	// 	- **kubernetes**: dedicated Kubernetes cluster
	//
	// 	- **ask**: dedicated ASK cluster
	//
	// example:
	//
	// kubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The timestamp when the cluster was created. Unit: milliseconds.
	//
	// example:
	//
	// 1662038134000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.22.10-aliyun.1
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The number of instances in the cluster.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the region in which the cluster is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **unavailable**
	//
	// 	- **Available**
	//
	// 	- **Creating**
	//
	// 	- **CreateFailed**
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the cluster is enabled. Valid values:
	//
	// 	- **true**: The cluster is enabled.
	//
	// 	- **false**: The cluster is disabled.
	//
	// example:
	//
	// true
	TargetResult *bool `json:"TargetResult,omitempty" xml:"TargetResult,omitempty"`
}

func (s DescribeClusterBasicInfoResponseBodyClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBasicInfoResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetState() *string {
	return s.State
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) GetTargetResult() *bool {
	return s.TargetResult
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetClusterId(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetClusterName(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetClusterType(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetCreateTime(v int64) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetCurrentVersion(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetInstanceCount(v int32) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.InstanceCount = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetState(v string) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.State = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) SetTargetResult(v bool) *DescribeClusterBasicInfoResponseBodyClusterInfo {
	s.TargetResult = &v
	return s
}

func (s *DescribeClusterBasicInfoResponseBodyClusterInfo) Validate() error {
	return dara.Validate(s)
}
