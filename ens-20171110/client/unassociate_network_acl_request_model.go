// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclId(v string) *UnassociateNetworkAclRequest
	GetNetworkAclId() *string
	SetResource(v []*UnassociateNetworkAclRequestResource) *UnassociateNetworkAclRequest
	GetResource() []*UnassociateNetworkAclRequestResource
}

type UnassociateNetworkAclRequest struct {
	// The ID of the network ACL that you want to disassociate from a resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0sp****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// Resources that you want to disassociate. Valid values of **N**: 0 to 29. A maximum of 30 resources can be unbound.
	//
	// This parameter is required.
	Resource []*UnassociateNetworkAclRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s UnassociateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *UnassociateNetworkAclRequest) GetResource() []*UnassociateNetworkAclRequestResource {
	return s.Resource
}

func (s *UnassociateNetworkAclRequest) SetNetworkAclId(v string) *UnassociateNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *UnassociateNetworkAclRequest) SetResource(v []*UnassociateNetworkAclRequestResource) *UnassociateNetworkAclRequest {
	s.Resource = v
	return s
}

func (s *UnassociateNetworkAclRequest) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UnassociateNetworkAclRequestResource struct {
	// The ID of the resource that you want to disassociate.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-5***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource that you want to disassociate. Valid values:
	//
	// 	- Network
	//
	// This parameter is required.
	//
	// example:
	//
	// Network
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UnassociateNetworkAclRequestResource) String() string {
	return dara.Prettify(s)
}

func (s UnassociateNetworkAclRequestResource) GoString() string {
	return s.String()
}

func (s *UnassociateNetworkAclRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *UnassociateNetworkAclRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnassociateNetworkAclRequestResource) SetResourceId(v string) *UnassociateNetworkAclRequestResource {
	s.ResourceId = &v
	return s
}

func (s *UnassociateNetworkAclRequestResource) SetResourceType(v string) *UnassociateNetworkAclRequestResource {
	s.ResourceType = &v
	return s
}

func (s *UnassociateNetworkAclRequestResource) Validate() error {
	return dara.Validate(s)
}
