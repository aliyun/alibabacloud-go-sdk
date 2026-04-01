// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeRCClustersResponseBodyClusters) *DescribeRCClustersResponseBody
	GetClusters() []*DescribeRCClustersResponseBodyClusters
	SetRequestId(v string) *DescribeRCClustersResponseBody
	GetRequestId() *string
}

type DescribeRCClustersResponseBody struct {
	// The information about the clusters.
	Clusters []*DescribeRCClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2875D608-A228-53D7-B8C9-35F13EDCF36D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRCClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCClustersResponseBody) GetClusters() []*DescribeRCClustersResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeRCClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCClustersResponseBody) SetClusters(v []*DescribeRCClustersResponseBodyClusters) *DescribeRCClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeRCClustersResponseBody) SetRequestId(v string) *DescribeRCClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCClustersResponseBodyClusters struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// test01
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2024-10-30T02:16:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Profile    *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// The cluster status. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Running**
	//
	// 	- **Starting**
	//
	// 	- **Stopping**
	//
	// 	- **Stopped**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId  *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeRCClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeRCClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeRCClustersResponseBodyClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeRCClustersResponseBodyClusters) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRCClustersResponseBodyClusters) GetProfile() *string {
	return s.Profile
}

func (s *DescribeRCClustersResponseBodyClusters) GetStatus() *string {
	return s.Status
}

func (s *DescribeRCClustersResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRCClustersResponseBodyClusters) SetClusterId(v string) *DescribeRCClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) SetClusterName(v string) *DescribeRCClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) SetCreateTime(v string) *DescribeRCClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) SetProfile(v string) *DescribeRCClustersResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) SetStatus(v string) *DescribeRCClustersResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) SetVpcId(v string) *DescribeRCClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeRCClustersResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
