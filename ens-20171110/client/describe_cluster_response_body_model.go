// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeClusterResponseBodyClusters) *DescribeClusterResponseBody
	GetClusters() []*DescribeClusterResponseBodyClusters
	SetRequestId(v string) *DescribeClusterResponseBody
	GetRequestId() *string
}

type DescribeClusterResponseBody struct {
	// An array that consists of the information about clusters.
	Clusters []*DescribeClusterResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) GetClusters() []*DescribeClusterResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterResponseBody) SetClusters(v []*DescribeClusterResponseBodyClusters) *DescribeClusterResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResponseBodyClusters struct {
	// The cluster ID.
	//
	// example:
	//
	// c8f0377146d104687ac562eef9403****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The version of the cluster.
	//
	// example:
	//
	// 1.18.8
	CurrentVersion *string `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// vc-a622bb**
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next version of the cluster.
	//
	// example:
	//
	// 1.20.8
	NextVersion *string `json:"NextVersion,omitempty" xml:"NextVersion,omitempty"`
	// The health status of the instance.
	//
	// Valid values:
	//
	// 	- healthy
	//
	// 	- unhealthy
	//
	// example:
	//
	// healthy
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterResponseBodyClusters) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeClusterResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeClusterResponseBodyClusters) GetNextVersion() *string {
	return s.NextVersion
}

func (s *DescribeClusterResponseBodyClusters) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterResponseBodyClusters) SetClusterId(v string) *DescribeClusterResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusters) SetCurrentVersion(v string) *DescribeClusterResponseBodyClusters {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyClusters) SetName(v string) *DescribeClusterResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusters) SetNextVersion(v string) *DescribeClusterResponseBodyClusters {
	s.NextVersion = &v
	return s
}

func (s *DescribeClusterResponseBodyClusters) SetStatus(v string) *DescribeClusterResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
