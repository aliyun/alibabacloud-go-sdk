// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourcesDeleteProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *UpdateResourcesDeleteProtectionResponseBody
	GetNamespace() *string
	SetProtection(v string) *UpdateResourcesDeleteProtectionResponseBody
	GetProtection() *string
	SetRequestId(v string) *UpdateResourcesDeleteProtectionResponseBody
	GetRequestId() *string
	SetResourceType(v string) *UpdateResourcesDeleteProtectionResponseBody
	GetResourceType() *string
	SetResources(v []*string) *UpdateResourcesDeleteProtectionResponseBody
	GetResources() []*string
}

type UpdateResourcesDeleteProtectionResponseBody struct {
	// The namespace of the resource.
	//
	// example:
	//
	// default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The deletion protection status of the resource.
	//
	// example:
	//
	// enable
	Protection *string `json:"protection,omitempty" xml:"protection,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0527ac9a-c899-4341-a21a-xxxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// namespaces
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	// The list of resources for which the deletion protection status is updated.
	Resources []*string `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s UpdateResourcesDeleteProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourcesDeleteProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourcesDeleteProtectionResponseBody) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateResourcesDeleteProtectionResponseBody) GetProtection() *string {
	return s.Protection
}

func (s *UpdateResourcesDeleteProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResourcesDeleteProtectionResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateResourcesDeleteProtectionResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *UpdateResourcesDeleteProtectionResponseBody) SetNamespace(v string) *UpdateResourcesDeleteProtectionResponseBody {
	s.Namespace = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponseBody) SetProtection(v string) *UpdateResourcesDeleteProtectionResponseBody {
	s.Protection = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponseBody) SetRequestId(v string) *UpdateResourcesDeleteProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponseBody) SetResourceType(v string) *UpdateResourcesDeleteProtectionResponseBody {
	s.ResourceType = &v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponseBody) SetResources(v []*string) *UpdateResourcesDeleteProtectionResponseBody {
	s.Resources = v
	return s
}

func (s *UpdateResourcesDeleteProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
