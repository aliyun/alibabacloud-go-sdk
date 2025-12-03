// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceClusterListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceClusters(v *DescribeInstanceClusterListResponseBodyInstanceClusters) *DescribeInstanceClusterListResponseBody
	GetInstanceClusters() *DescribeInstanceClusterListResponseBodyInstanceClusters
	SetPageNumber(v int32) *DescribeInstanceClusterListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceClusterListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceClusterListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceClusterListResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceClusterListResponseBody struct {
	// The instance cluster list.
	InstanceClusters *DescribeInstanceClusterListResponseBodyInstanceClusters `json:"InstanceClusters,omitempty" xml:"InstanceClusters,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ015
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceClusterListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterListResponseBody) GetInstanceClusters() *DescribeInstanceClusterListResponseBodyInstanceClusters {
	return s.InstanceClusters
}

func (s *DescribeInstanceClusterListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceClusterListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceClusterListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceClusterListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceClusterListResponseBody) SetInstanceClusters(v *DescribeInstanceClusterListResponseBodyInstanceClusters) *DescribeInstanceClusterListResponseBody {
	s.InstanceClusters = v
	return s
}

func (s *DescribeInstanceClusterListResponseBody) SetPageNumber(v int32) *DescribeInstanceClusterListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBody) SetPageSize(v int32) *DescribeInstanceClusterListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBody) SetRequestId(v string) *DescribeInstanceClusterListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBody) SetTotalCount(v int32) *DescribeInstanceClusterListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBody) Validate() error {
	if s.InstanceClusters != nil {
		if err := s.InstanceClusters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceClusterListResponseBodyInstanceClusters struct {
	InstanceCluster []*DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster `json:"InstanceCluster,omitempty" xml:"InstanceCluster,omitempty" type:"Repeated"`
}

func (s DescribeInstanceClusterListResponseBodyInstanceClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterListResponseBodyInstanceClusters) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClusters) GetInstanceCluster() []*DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	return s.InstanceCluster
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClusters) SetInstanceCluster(v []*DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) *DescribeInstanceClusterListResponseBodyInstanceClusters {
	s.InstanceCluster = v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClusters) Validate() error {
	if s.InstanceCluster != nil {
		for _, item := range s.InstanceCluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster struct {
	// The time when the cluster was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-08-14T17:46:59+08:00
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Test TrafficControl
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// apigateway-cluster-hz-xxxxxxxxxxxx
	InstanceClusterId *string `json:"InstanceClusterId,omitempty" xml:"InstanceClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test-cluster
	InstanceClusterName *string `json:"InstanceClusterName,omitempty" xml:"InstanceClusterName,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// RUNNING
	InstanceClusterStatus *string `json:"InstanceClusterStatus,omitempty" xml:"InstanceClusterStatus,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// normal
	InstanceClusterType *string `json:"InstanceClusterType,omitempty" xml:"InstanceClusterType,omitempty"`
	// The time when the cluster was last modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-01-12T10:11:08+08:00
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region ID of the cluster.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GoString() string {
	return s.String()
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetInstanceClusterId() *string {
	return s.InstanceClusterId
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetInstanceClusterName() *string {
	return s.InstanceClusterName
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetInstanceClusterStatus() *string {
	return s.InstanceClusterStatus
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetInstanceClusterType() *string {
	return s.InstanceClusterType
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetCreatedTime(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.CreatedTime = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetDescription(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.Description = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetInstanceClusterId(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.InstanceClusterId = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetInstanceClusterName(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.InstanceClusterName = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetInstanceClusterStatus(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.InstanceClusterStatus = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetInstanceClusterType(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.InstanceClusterType = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetModifiedTime(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) SetRegionId(v string) *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceClusterListResponseBodyInstanceClustersInstanceCluster) Validate() error {
	return dara.Validate(s)
}
