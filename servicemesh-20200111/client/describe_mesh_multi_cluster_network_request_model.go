// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMeshMultiClusterNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceMeshId(v string) *DescribeMeshMultiClusterNetworkRequest
	GetServiceMeshId() *string
}

type DescribeMeshMultiClusterNetworkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccb37ff104caf419fbf48fb38e6f3****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DescribeMeshMultiClusterNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMeshMultiClusterNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMeshMultiClusterNetworkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeMeshMultiClusterNetworkRequest) SetServiceMeshId(v string) *DescribeMeshMultiClusterNetworkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeMeshMultiClusterNetworkRequest) Validate() error {
	return dara.Validate(s)
}
