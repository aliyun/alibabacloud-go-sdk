// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeshMultiClusterNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMultiClusterNetworks(v map[string]*MultiClusterNetworksValue) *DescribeMeshMultiClusterNetworkResponseBody
	GetMultiClusterNetworks() map[string]*MultiClusterNetworksValue
	SetRequestId(v string) *DescribeMeshMultiClusterNetworkResponseBody
	GetRequestId() *string
}

type DescribeMeshMultiClusterNetworkResponseBody struct {
	MultiClusterNetworks map[string]*MultiClusterNetworksValue `json:"MultiClusterNetworks,omitempty" xml:"MultiClusterNetworks,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 71680038-8009-5073-B43E-C057E9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMeshMultiClusterNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeshMultiClusterNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMeshMultiClusterNetworkResponseBody) GetMultiClusterNetworks() map[string]*MultiClusterNetworksValue {
	return s.MultiClusterNetworks
}

func (s *DescribeMeshMultiClusterNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMeshMultiClusterNetworkResponseBody) SetMultiClusterNetworks(v map[string]*MultiClusterNetworksValue) *DescribeMeshMultiClusterNetworkResponseBody {
	s.MultiClusterNetworks = v
	return s
}

func (s *DescribeMeshMultiClusterNetworkResponseBody) SetRequestId(v string) *DescribeMeshMultiClusterNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMeshMultiClusterNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
