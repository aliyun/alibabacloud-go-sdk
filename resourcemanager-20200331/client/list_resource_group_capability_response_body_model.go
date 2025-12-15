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
	Capabilities []*ListResourceGroupCapabilityResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
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
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
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

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetService() *string {
	return s.Service
}

func (s *ListResourceGroupCapabilityResponseBodyCapabilities) GetSupportResourceGroupEvent() *bool {
	return s.SupportResourceGroupEvent
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
