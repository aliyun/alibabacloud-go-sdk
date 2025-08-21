// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccosicateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkAclId(v string) *AccosicateNetworkAclRequest
	GetNetworkAclId() *string
	SetResource(v []*AccosicateNetworkAclRequestResource) *AccosicateNetworkAclRequest
	GetResource() []*AccosicateNetworkAclRequestResource
}

type AccosicateNetworkAclRequest struct {
	// The ID of the network ACL.
	//
	// This parameter is required.
	//
	// example:
	//
	// nacl-a2do9e413e0sp****
	NetworkAclId *string `json:"NetworkAclId,omitempty" xml:"NetworkAclId,omitempty"`
	// The type of resource with which you want to associate the network ACL.
	//
	// This parameter is required.
	Resource []*AccosicateNetworkAclRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s AccosicateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s AccosicateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *AccosicateNetworkAclRequest) GetNetworkAclId() *string {
	return s.NetworkAclId
}

func (s *AccosicateNetworkAclRequest) GetResource() []*AccosicateNetworkAclRequestResource {
	return s.Resource
}

func (s *AccosicateNetworkAclRequest) SetNetworkAclId(v string) *AccosicateNetworkAclRequest {
	s.NetworkAclId = &v
	return s
}

func (s *AccosicateNetworkAclRequest) SetResource(v []*AccosicateNetworkAclRequestResource) *AccosicateNetworkAclRequest {
	s.Resource = v
	return s
}

func (s *AccosicateNetworkAclRequest) Validate() error {
	return dara.Validate(s)
}

type AccosicateNetworkAclRequestResource struct {
	// The ID of the associated resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-5****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the associated resource. Set the value to **Network**.
	//
	// Valid values of **N**: 0 to 29. You can associate a network ACL with at most 30 resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// Network
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AccosicateNetworkAclRequestResource) String() string {
	return dara.Prettify(s)
}

func (s AccosicateNetworkAclRequestResource) GoString() string {
	return s.String()
}

func (s *AccosicateNetworkAclRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *AccosicateNetworkAclRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *AccosicateNetworkAclRequestResource) SetResourceId(v string) *AccosicateNetworkAclRequestResource {
	s.ResourceId = &v
	return s
}

func (s *AccosicateNetworkAclRequestResource) SetResourceType(v string) *AccosicateNetworkAclRequestResource {
	s.ResourceType = &v
	return s
}

func (s *AccosicateNetworkAclRequestResource) Validate() error {
	return dara.Validate(s)
}
