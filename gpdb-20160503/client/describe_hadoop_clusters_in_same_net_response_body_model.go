// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopClustersInSameNetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*string) *DescribeHadoopClustersInSameNetResponseBody
	GetClusters() []*string
	SetRequestId(v string) *DescribeHadoopClustersInSameNetResponseBody
	GetRequestId() *string
}

type DescribeHadoopClustersInSameNetResponseBody struct {
	Clusters []*string `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHadoopClustersInSameNetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopClustersInSameNetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHadoopClustersInSameNetResponseBody) GetClusters() []*string {
	return s.Clusters
}

func (s *DescribeHadoopClustersInSameNetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHadoopClustersInSameNetResponseBody) SetClusters(v []*string) *DescribeHadoopClustersInSameNetResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeHadoopClustersInSameNetResponseBody) SetRequestId(v string) *DescribeHadoopClustersInSameNetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHadoopClustersInSameNetResponseBody) Validate() error {
	return dara.Validate(s)
}
