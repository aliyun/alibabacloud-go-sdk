// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaPluginStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterIds(v []*string) *GetOpaPluginStatusRequest
	GetClusterIds() []*string
}

type GetOpaPluginStatusRequest struct {
	// The list of cluster IDs.
	//
	// > You can query this value by calling the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation.
	//
	// This parameter is required.
	ClusterIds []*string `json:"ClusterIds,omitempty" xml:"ClusterIds,omitempty" type:"Repeated"`
}

func (s GetOpaPluginStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpaPluginStatusRequest) GoString() string {
	return s.String()
}

func (s *GetOpaPluginStatusRequest) GetClusterIds() []*string {
	return s.ClusterIds
}

func (s *GetOpaPluginStatusRequest) SetClusterIds(v []*string) *GetOpaPluginStatusRequest {
	s.ClusterIds = v
	return s
}

func (s *GetOpaPluginStatusRequest) Validate() error {
	return dara.Validate(s)
}
