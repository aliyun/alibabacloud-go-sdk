// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesDeleteProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *UpdateResourcesDeleteProtectionRequest
	GetEnable() *bool
	SetNamespace(v string) *UpdateResourcesDeleteProtectionRequest
	GetNamespace() *string
	SetResourceType(v string) *UpdateResourcesDeleteProtectionRequest
	GetResourceType() *string
	SetResources(v []*string) *UpdateResourcesDeleteProtectionRequest
	GetResources() []*string
}

type UpdateResourcesDeleteProtectionRequest struct {
	// Specifies whether to enable deletion protection. Set the value to true to enable deletion protection and set the value to false to disable deletion protection.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The namespace to which the resource belongs.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The type of resource for which deletion protection is enabled or disabled. You can specify namespaces or Services.
	//
	// example:
	//
	// services
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The resources list.
	Resources []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s UpdateResourcesDeleteProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesDeleteProtectionRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourcesDeleteProtectionRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateResourcesDeleteProtectionRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateResourcesDeleteProtectionRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateResourcesDeleteProtectionRequest) GetResources() []*string {
	return s.Resources
}

func (s *UpdateResourcesDeleteProtectionRequest) SetEnable(v bool) *UpdateResourcesDeleteProtectionRequest {
	s.Enable = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionRequest) SetNamespace(v string) *UpdateResourcesDeleteProtectionRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionRequest) SetResourceType(v string) *UpdateResourcesDeleteProtectionRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionRequest) SetResources(v []*string) *UpdateResourcesDeleteProtectionRequest {
	s.Resources = v
	return s
}

func (s *UpdateResourcesDeleteProtectionRequest) Validate() error {
	return dara.Validate(s)
}
