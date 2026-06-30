// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalAcceleratorResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedResources(v []*GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) *GetGlobalAcceleratorResourcesResponseBody
	GetAssociatedResources() []*GetGlobalAcceleratorResourcesResponseBodyAssociatedResources
	SetRequestId(v string) *GetGlobalAcceleratorResourcesResponseBody
	GetRequestId() *string
}

type GetGlobalAcceleratorResourcesResponseBody struct {
	// List of linked instances.
	AssociatedResources []*GetGlobalAcceleratorResourcesResponseBodyAssociatedResources `json:"AssociatedResources,omitempty" xml:"AssociatedResources,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGlobalAcceleratorResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalAcceleratorResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetGlobalAcceleratorResourcesResponseBody) GetAssociatedResources() []*GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	return s.AssociatedResources
}

func (s *GetGlobalAcceleratorResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGlobalAcceleratorResourcesResponseBody) SetAssociatedResources(v []*GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) *GetGlobalAcceleratorResourcesResponseBody {
	s.AssociatedResources = v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBody) SetRequestId(v string) *GetGlobalAcceleratorResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBody) Validate() error {
	if s.AssociatedResources != nil {
		for _, item := range s.AssociatedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGlobalAcceleratorResourcesResponseBodyAssociatedResources struct {
	// Instance ID of Alibaba Cloud Global Accelerator (GA).
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// Association pattern:
	//
	// - **Managed**: Managed mode, where GA restricts user operations based on management policies.
	//
	// - **Associated*	- (default): Loose coupling association, where GA does not restrict user operations; WAF uses loose coupling.
	//
	// example:
	//
	// Associated
	AssociatedMode *string `json:"AssociatedMode,omitempty" xml:"AssociatedMode,omitempty"`
	// ID of the linked instance.
	//
	// example:
	//
	// waf_v2_public_cn-x0r****gr1i
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// Region ID of the linked instance.
	//
	// example:
	//
	// cn-hangzhou
	AssociatedResourceRegionId *string `json:"AssociatedResourceRegionId,omitempty" xml:"AssociatedResourceRegionId,omitempty"`
	// Resource type of the linked instance.
	//
	// example:
	//
	// WAF
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// Status of the association between the Alibaba Cloud Global Accelerator (GA) instance and other cloud product resources.
	//
	// - **active**: The resource is active.
	//
	// - **binding**: The resource is being attached.
	//
	// - **unBinding**: The resource is being detached.
	//
	// - **wafRelease**: The WAF instance has been released.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GoString() string {
	return s.String()
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetAssociatedMode() *string {
	return s.AssociatedMode
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetAssociatedResourceRegionId() *string {
	return s.AssociatedResourceRegionId
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) GetState() *string {
	return s.State
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetAcceleratorId(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.AcceleratorId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetAssociatedMode(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.AssociatedMode = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetAssociatedResourceId(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.AssociatedResourceId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetAssociatedResourceRegionId(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.AssociatedResourceRegionId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetAssociatedResourceType(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.AssociatedResourceType = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) SetState(v string) *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources {
	s.State = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesResponseBodyAssociatedResources) Validate() error {
	return dara.Validate(s)
}
