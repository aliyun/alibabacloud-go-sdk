// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterSpec interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *ClusterSpec
	GetClusterType() *string
	SetDataSources(v []*DataSource) *ClusterSpec
	GetDataSources() []*DataSource
	SetImage(v string) *ClusterSpec
	GetImage() *string
}

type ClusterSpec struct {
	// The cluster type.
	//
	// example:
	//
	// RayCluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The datasource config.
	DataSources []*DataSource `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The container image address.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai-training/kmeans:v1.0.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
}

func (s ClusterSpec) String() string {
	return dara.Prettify(s)
}

func (s ClusterSpec) GoString() string {
	return s.String()
}

func (s *ClusterSpec) GetClusterType() *string {
	return s.ClusterType
}

func (s *ClusterSpec) GetDataSources() []*DataSource {
	return s.DataSources
}

func (s *ClusterSpec) GetImage() *string {
	return s.Image
}

func (s *ClusterSpec) SetClusterType(v string) *ClusterSpec {
	s.ClusterType = &v
	return s
}

func (s *ClusterSpec) SetDataSources(v []*DataSource) *ClusterSpec {
	s.DataSources = v
	return s
}

func (s *ClusterSpec) SetImage(v string) *ClusterSpec {
	s.Image = &v
	return s
}

func (s *ClusterSpec) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
