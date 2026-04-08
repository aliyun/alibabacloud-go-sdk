// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceManagerResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceId() *string
	SetResourceManagerResourceGroupId(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *ChangeResourceManagerResourceGroupRequest
	GetResourceType() *string
}

type ChangeResourceManagerResourceGroupRequest struct {
	// Indicates whether the request was successful.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// ChangeResourceManagerResourceGroup
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The HTTP status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// project
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceManagerResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceManagerResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ChangeResourceManagerResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceId(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceManagerResourceGroupId(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) SetResourceType(v string) *ChangeResourceManagerResourceGroupRequest {
	s.ResourceType = &v
	return s
}

func (s *ChangeResourceManagerResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
