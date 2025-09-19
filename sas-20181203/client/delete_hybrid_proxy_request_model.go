// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *DeleteHybridProxyRequest
	GetClusterName() *string
	SetUuid(v string) *DeleteHybridProxyRequest
	GetUuid() *string
}

type DeleteHybridProxyRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The UUID of the proxy node that you want to remove. The value starts with inet-proxy.
	//
	// This parameter is required.
	//
	// example:
	//
	// inet-proxy-d2d94e8b-bb25-4744-8004-1e08a53c****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DeleteHybridProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyRequest) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *DeleteHybridProxyRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DeleteHybridProxyRequest) SetClusterName(v string) *DeleteHybridProxyRequest {
	s.ClusterName = &v
	return s
}

func (s *DeleteHybridProxyRequest) SetUuid(v string) *DeleteHybridProxyRequest {
	s.Uuid = &v
	return s
}

func (s *DeleteHybridProxyRequest) Validate() error {
	return dara.Validate(s)
}
