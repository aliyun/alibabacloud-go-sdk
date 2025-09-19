// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindHybridProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *BindHybridProxyRequest
	GetClusterName() *string
	SetYundunUuids(v []*string) *BindHybridProxyRequest
	GetYundunUuids() []*string
}

type BindHybridProxyRequest struct {
	// The name of the proxy cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// sas-proxy
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The UUIDs of servers that you want to add to Security Center over the proxy server.
	//
	// This parameter is required.
	YundunUuids []*string `json:"YundunUuids,omitempty" xml:"YundunUuids,omitempty" type:"Repeated"`
}

func (s BindHybridProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s BindHybridProxyRequest) GoString() string {
	return s.String()
}

func (s *BindHybridProxyRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *BindHybridProxyRequest) GetYundunUuids() []*string {
	return s.YundunUuids
}

func (s *BindHybridProxyRequest) SetClusterName(v string) *BindHybridProxyRequest {
	s.ClusterName = &v
	return s
}

func (s *BindHybridProxyRequest) SetYundunUuids(v []*string) *BindHybridProxyRequest {
	s.YundunUuids = v
	return s
}

func (s *BindHybridProxyRequest) Validate() error {
	return dara.Validate(s)
}
