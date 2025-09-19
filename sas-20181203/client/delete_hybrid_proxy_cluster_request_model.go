// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DeleteHybridProxyClusterRequest
	GetClusterName() *string
}

type DeleteHybridProxyClusterRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
}

func (s DeleteHybridProxyClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DeleteHybridProxyClusterRequest) SetClusterName(v string) *DeleteHybridProxyClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteHybridProxyClusterRequest) Validate() error {
	return dara.Validate(s)
}
