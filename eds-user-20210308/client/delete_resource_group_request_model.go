// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DeleteResourceGroupRequest
	GetBusinessChannel() *string
	SetResourceGroupId(v string) *DeleteResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceGroupIds(v []*string) *DeleteResourceGroupRequest
	GetResourceGroupIds() []*string
}

type DeleteResourceGroupRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// The ID of the resource group to delete.
	//
	// > If the `ResourceGroupIds` parameter is not empty, both parameters take effect.
	//
	// example:
	//
	// rg-aj01tck67a0szp***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of IDs of the resource groups to delete.
	//
	// > If the `ResourceGroupIds` parameter is not empty, both parameters take effect.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DeleteResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteResourceGroupRequest) GetResourceGroupIds() []*string {
	return s.ResourceGroupIds
}

func (s *DeleteResourceGroupRequest) SetBusinessChannel(v string) *DeleteResourceGroupRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DeleteResourceGroupRequest) SetResourceGroupId(v string) *DeleteResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteResourceGroupRequest) SetResourceGroupIds(v []*string) *DeleteResourceGroupRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *DeleteResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
