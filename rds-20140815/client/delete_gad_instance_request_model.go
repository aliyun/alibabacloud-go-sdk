// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGadInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGadInstanceName(v string) *DeleteGadInstanceRequest
	GetGadInstanceName() *string
	SetRegionId(v string) *DeleteGadInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteGadInstanceRequest
	GetResourceGroupId() *string
}

type DeleteGadInstanceRequest struct {
	// The ID of the global active database cluster. You can call the GadInstanceName operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gad-rm-bp1npi2j8********
	GadInstanceName *string `json:"GadInstanceName,omitempty" xml:"GadInstanceName,omitempty"`
	// The region ID of the central node of the global active database cluster. The central node refers to the primary node. You can call the DescribeGadInstances operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteGadInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGadInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGadInstanceRequest) GetGadInstanceName() *string {
	return s.GadInstanceName
}

func (s *DeleteGadInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGadInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteGadInstanceRequest) SetGadInstanceName(v string) *DeleteGadInstanceRequest {
	s.GadInstanceName = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetRegionId(v string) *DeleteGadInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGadInstanceRequest) SetResourceGroupId(v string) *DeleteGadInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteGadInstanceRequest) Validate() error {
	return dara.Validate(s)
}
