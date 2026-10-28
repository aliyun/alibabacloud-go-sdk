// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuotaCluster interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *QuotaCluster
	GetClusterType() *string
	SetDataSources(v []*DataSource) *QuotaCluster
	GetDataSources() []*DataSource
	SetEndpoints(v map[string]*string) *QuotaCluster
	GetEndpoints() map[string]*string
	SetImage(v string) *QuotaCluster
	GetImage() *string
	SetStatus(v string) *QuotaCluster
	GetStatus() *string
}

type QuotaCluster struct {
	// The cluster type.
	//
	// example:
	//
	// RayCluster
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The datasource config.
	DataSources []*DataSource `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The service access endpoints.
	Endpoints map[string]*string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// The container image address.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/pai-training/kmeans:v1.0.0
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuotaCluster) String() string {
	return dara.Prettify(s)
}

func (s QuotaCluster) GoString() string {
	return s.String()
}

func (s *QuotaCluster) GetClusterType() *string {
	return s.ClusterType
}

func (s *QuotaCluster) GetDataSources() []*DataSource {
	return s.DataSources
}

func (s *QuotaCluster) GetEndpoints() map[string]*string {
	return s.Endpoints
}

func (s *QuotaCluster) GetImage() *string {
	return s.Image
}

func (s *QuotaCluster) GetStatus() *string {
	return s.Status
}

func (s *QuotaCluster) SetClusterType(v string) *QuotaCluster {
	s.ClusterType = &v
	return s
}

func (s *QuotaCluster) SetDataSources(v []*DataSource) *QuotaCluster {
	s.DataSources = v
	return s
}

func (s *QuotaCluster) SetEndpoints(v map[string]*string) *QuotaCluster {
	s.Endpoints = v
	return s
}

func (s *QuotaCluster) SetImage(v string) *QuotaCluster {
	s.Image = &v
	return s
}

func (s *QuotaCluster) SetStatus(v string) *QuotaCluster {
	s.Status = &v
	return s
}

func (s *QuotaCluster) Validate() error {
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
