// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ChangeResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
	SetService(v string) *ChangeResourceGroupRequest
	GetService() *string
}

type ChangeResourceGroupRequest struct {
	// Target resource group ID.
	//
	// example:
	//
	// rg-aekzdrfx2xdnaja
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// gw-ct4i14um1hkn0tpqfae0
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource type
	//
	// example:
	//
	// gateway
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service name, fixed value apig
	//
	// example:
	//
	// apig
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) GetService() *string {
	return s.Service
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetService(v string) *ChangeResourceGroupRequest {
	s.Service = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
