// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupCapabilityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceType(v string) *ListResourceGroupCapabilityRequest
	GetResourceType() *string
	SetService(v string) *ListResourceGroupCapabilityRequest
	GetService() *string
	SetSupportResourceGroupEvent(v bool) *ListResourceGroupCapabilityRequest
	GetSupportResourceGroupEvent() *bool
}

type ListResourceGroupCapabilityRequest struct {
	// The resource type.
	//
	// You can obtain the resource type from the **Resource Type*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The cloud service ID.
	//
	// You can obtain the cloud service ID from the **Cloud Service Code*	- column in [Cloud services that support resource groups](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// Specifies whether resource group events are supported.
	//
	// example:
	//
	// true
	SupportResourceGroupEvent *bool `json:"SupportResourceGroupEvent,omitempty" xml:"SupportResourceGroupEvent,omitempty"`
}

func (s ListResourceGroupCapabilityRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupCapabilityRequest) GoString() string {
	return s.String()
}

func (s *ListResourceGroupCapabilityRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceGroupCapabilityRequest) GetService() *string {
	return s.Service
}

func (s *ListResourceGroupCapabilityRequest) GetSupportResourceGroupEvent() *bool {
	return s.SupportResourceGroupEvent
}

func (s *ListResourceGroupCapabilityRequest) SetResourceType(v string) *ListResourceGroupCapabilityRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceGroupCapabilityRequest) SetService(v string) *ListResourceGroupCapabilityRequest {
	s.Service = &v
	return s
}

func (s *ListResourceGroupCapabilityRequest) SetSupportResourceGroupEvent(v bool) *ListResourceGroupCapabilityRequest {
	s.SupportResourceGroupEvent = &v
	return s
}

func (s *ListResourceGroupCapabilityRequest) Validate() error {
	return dara.Validate(s)
}
