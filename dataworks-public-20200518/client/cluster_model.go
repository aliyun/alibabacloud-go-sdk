// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCluster interface {
	dara.Model
	String() string
	GoString() string
	SetClusterBizId(v string) *Cluster
	GetClusterBizId() *string
	SetClusterId(v int64) *Cluster
	GetClusterId() *int64
}

type Cluster struct {
	// This parameter is required.
	//
	// example:
	//
	// c-d8a7523****
	ClusterBizId *string `json:"ClusterBizId,omitempty" xml:"ClusterBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ClusterId *int64 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s Cluster) String() string {
	return dara.Prettify(s)
}

func (s Cluster) GoString() string {
	return s.String()
}

func (s *Cluster) GetClusterBizId() *string {
	return s.ClusterBizId
}

func (s *Cluster) GetClusterId() *int64 {
	return s.ClusterId
}

func (s *Cluster) SetClusterBizId(v string) *Cluster {
	s.ClusterBizId = &v
	return s
}

func (s *Cluster) SetClusterId(v int64) *Cluster {
	s.ClusterId = &v
	return s
}

func (s *Cluster) Validate() error {
	return dara.Validate(s)
}
