// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnBindHybridProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *UnBindHybridProxyRequest
	GetClusterName() *string
	SetYundunUuids(v []*string) *UnBindHybridProxyRequest
	GetYundunUuids() []*string
}

type UnBindHybridProxyRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// proxy-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The UUIDs of the servers that you want to remove from the proxy cluster.
	//
	// This parameter is required.
	YundunUuids []*string `json:"YundunUuids,omitempty" xml:"YundunUuids,omitempty" type:"Repeated"`
}

func (s UnBindHybridProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s UnBindHybridProxyRequest) GoString() string {
	return s.String()
}

func (s *UnBindHybridProxyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UnBindHybridProxyRequest) GetYundunUuids() []*string {
	return s.YundunUuids
}

func (s *UnBindHybridProxyRequest) SetClusterName(v string) *UnBindHybridProxyRequest {
	s.ClusterName = &v
	return s
}

func (s *UnBindHybridProxyRequest) SetYundunUuids(v []*string) *UnBindHybridProxyRequest {
	s.YundunUuids = v
	return s
}

func (s *UnBindHybridProxyRequest) Validate() error {
	return dara.Validate(s)
}
