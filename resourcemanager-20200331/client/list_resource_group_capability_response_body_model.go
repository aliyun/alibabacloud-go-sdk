// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupCapabilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v []*ListResourceGroupCapabilityResponseBodyCapabilities) *ListResourceGroupCapabilityResponseBody
	GetCapabilities() []*ListResourceGroupCapabilityResponseBodyCapabilities
	SetRequestId(v string) *ListResourceGroupCapabilityResponseBody
	GetRequestId() *string
}

type ListResourceGroupCapabilityResponseBody struct {
	// Indicates whether a specific resource type or cloud service supports resource group events.
	Capabilities []*ListResourceGroupCapabilityResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2CB870A2-DF0F-5338-8223-F64737FF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceGroupCapabilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupCapabilityResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupCapabilityResponseBody) GetCapabilities() []*ListResourceGroupCapabilityResponseBodyCapabilities {
	return s.Capabilities
}

func (s *ListResourceGroupCapabilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupCapabilityResponseBody) SetCapabilities(v []*ListResourceGroupCapabilityResponseBodyCapabilities) *ListResourceGroupCapabilityResponseBody {
	s.Capabilities = v
	return s
}

func (s *ListResourceGroupCapabilityResponseBody) SetRequestId(v string) *ListResourceGroupCapabilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupCapabilityResponseBody) Validate() error {
	if s.Capabilities != nil {
		for _, item := range s.Capabilities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupCapabilityResponseBodyCapabilities struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceCenterResourceType *string `json:"ResourceCenterResourceType,omitempty" xml:"ResourceCenterResourceType,omitempty"`
	// The resource type.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The service code.
	//
	// You can obtain the code from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// Indicates whether a specific resource type or cloud service supports resource group events.
	//
	// example:
	//
	// true
	SupportResourceGroupEvent *bool `json:"SupportResourceGroupEvent,omitempty" xml:"SupportResourceGroupEvent,omitempty"`
}

func (s ListResourceGroupCapabilityResponseBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupCapabilityResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetResourceCenterResourceType() *string {
	return s.ResourceCenterResourceType
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetService() *string {
	return s.Service
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetSupportResourceGroupEvent() *bool {
	return s.SupportResourceGroupEvent
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) SetResourceCenterResourceType(v string) *ListResourceGroupCapabilityResponseBodyCapabilities {
	s.ResourceCenterResourceType = &v
	return s
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) SetResourceType(v string) *ListResourceGroupCapabilityResponseBodyCapabilities {
	s.ResourceType = &v
	return s
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) SetService(v string) *ListResourceGroupCapabilityResponseBodyCapabilities {
	s.Service = &v
	return s
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) SetSupportResourceGroupEvent(v bool) *ListResourceGroupCapabilityResponseBodyCapabilities {
	s.SupportResourceGroupEvent = &v
	return s
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) Validate() error {
	return dara.Validate(s)
}
