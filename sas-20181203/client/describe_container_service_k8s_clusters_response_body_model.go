// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerServiceK8sClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetK8sClusters(v []*DescribeContainerServiceK8sClustersResponseBodyK8sClusters) *DescribeContainerServiceK8sClustersResponseBody
	GetK8sClusters() []*DescribeContainerServiceK8sClustersResponseBodyK8sClusters
	SetRequestId(v string) *DescribeContainerServiceK8sClustersResponseBody
	GetRequestId() *string
}

type DescribeContainerServiceK8sClustersResponseBody struct {
	// The information about the clusters.
	K8sClusters []*DescribeContainerServiceK8sClustersResponseBodyK8sClusters `json:"K8sClusters,omitempty" xml:"K8sClusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 20456DD5-5CBF-5015-9173-12CA4246B***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerServiceK8sClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClustersResponseBody) GetK8sClusters() []*DescribeContainerServiceK8sClustersResponseBodyK8sClusters {
	return s.K8sClusters
}

func (s *DescribeContainerServiceK8sClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContainerServiceK8sClustersResponseBody) SetK8sClusters(v []*DescribeContainerServiceK8sClustersResponseBodyK8sClusters) *DescribeContainerServiceK8sClustersResponseBody {
	s.K8sClusters = v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponseBody) SetRequestId(v string) *DescribeContainerServiceK8sClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContainerServiceK8sClustersResponseBodyK8sClusters struct {
	// The ID of the cluster.
	//
	// example:
	//
	// cdbbe7aa56cbf4b8f830f83718d26****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeContainerServiceK8sClustersResponseBodyK8sClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerServiceK8sClustersResponseBodyK8sClusters) GoString() string {
	return s.String()
}

func (s *DescribeContainerServiceK8sClustersResponseBodyK8sClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeContainerServiceK8sClustersResponseBodyK8sClusters) GetName() *string {
	return s.Name
}

func (s *DescribeContainerServiceK8sClustersResponseBodyK8sClusters) SetClusterId(v string) *DescribeContainerServiceK8sClustersResponseBodyK8sClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponseBodyK8sClusters) SetName(v string) *DescribeContainerServiceK8sClustersResponseBodyK8sClusters {
	s.Name = &v
	return s
}

func (s *DescribeContainerServiceK8sClustersResponseBodyK8sClusters) Validate() error {
	return dara.Validate(s)
}
