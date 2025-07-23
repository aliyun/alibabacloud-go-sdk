// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ConfigClusterNameRequest
	GetClusterId() *string
	SetClusterName(v string) *ConfigClusterNameRequest
	GetClusterName() *string
}

type ConfigClusterNameRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster-BqxX63Bsgfaisdf****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s ConfigClusterNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ConfigClusterNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ConfigClusterNameRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ConfigClusterNameRequest) SetClusterId(v string) *ConfigClusterNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ConfigClusterNameRequest) SetClusterName(v string) *ConfigClusterNameRequest {
	s.ClusterName = &v
	return s
}

func (s *ConfigClusterNameRequest) Validate() error {
	return dara.Validate(s)
}
