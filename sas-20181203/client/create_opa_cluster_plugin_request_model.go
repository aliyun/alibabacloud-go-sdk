// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpaClusterPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *CreateOpaClusterPluginRequest
	GetClusterIds() []*string
}

type CreateOpaClusterPluginRequest struct {
	// The cluster IDs.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// This parameter is required.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
}

func (s CreateOpaClusterPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpaClusterPluginRequest) GoString() string {
	return s.String()
}

func (s *CreateOpaClusterPluginRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *CreateOpaClusterPluginRequest) SetClusterIds(v []*string) *CreateOpaClusterPluginRequest {
	s.ClusterIds = v
	return s
}

func (s *CreateOpaClusterPluginRequest) Validate() error {
	return dara.Validate(s)
}
