// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteGrafanaResourceRequest
	GetClusterId() *string
	SetClusterName(v string) *DeleteGrafanaResourceRequest
	GetClusterName() *string
	SetRegionId(v string) *DeleteGrafanaResourceRequest
	GetRegionId() *string
}

type DeleteGrafanaResourceRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// clusterNameOfTest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGrafanaResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteGrafanaResourceRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DeleteGrafanaResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGrafanaResourceRequest) SetClusterId(v string) *DeleteGrafanaResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetClusterName(v string) *DeleteGrafanaResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) SetRegionId(v string) *DeleteGrafanaResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGrafanaResourceRequest) Validate() error {
	return dara.Validate(s)
}
