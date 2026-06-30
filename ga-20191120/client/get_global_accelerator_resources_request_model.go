// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalAcceleratorResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetGlobalAcceleratorResourcesRequest
	GetAcceleratorId() *string
	SetAssociatedResourceId(v string) *GetGlobalAcceleratorResourcesRequest
	GetAssociatedResourceId() *string
	SetAssociatedResourceRegionId(v string) *GetGlobalAcceleratorResourcesRequest
	GetAssociatedResourceRegionId() *string
	SetAssociatedResourceType(v string) *GetGlobalAcceleratorResourcesRequest
	GetAssociatedResourceType() *string
	SetRegionId(v string) *GetGlobalAcceleratorResourcesRequest
	GetRegionId() *string
}

type GetGlobalAcceleratorResourcesRequest struct {
	// The instance ID of the Alibaba Cloud Global Accelerator (GA).
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the linked peripheral resource.
	//
	// example:
	//
	// waf_v2_public_cn-x0r****gr1i
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// The Region ID where the linked instance is located.
	//
	// example:
	//
	// cn-hangzhou
	AssociatedResourceRegionId *string `json:"AssociatedResourceRegionId,omitempty" xml:"AssociatedResourceRegionId,omitempty"`
	// The resource type of the linked peripheral resource.
	//
	// example:
	//
	// WAF
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// The Region ID where the Alibaba Cloud Global Accelerator (GA) instance is located. The only valid value is cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGlobalAcceleratorResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalAcceleratorResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalAcceleratorResourcesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetGlobalAcceleratorResourcesRequest) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *GetGlobalAcceleratorResourcesRequest) GetAssociatedResourceRegionId() *string {
	return s.AssociatedResourceRegionId
}

func (s *GetGlobalAcceleratorResourcesRequest) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *GetGlobalAcceleratorResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetGlobalAcceleratorResourcesRequest) SetAcceleratorId(v string) *GetGlobalAcceleratorResourcesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesRequest) SetAssociatedResourceId(v string) *GetGlobalAcceleratorResourcesRequest {
	s.AssociatedResourceId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesRequest) SetAssociatedResourceRegionId(v string) *GetGlobalAcceleratorResourcesRequest {
	s.AssociatedResourceRegionId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesRequest) SetAssociatedResourceType(v string) *GetGlobalAcceleratorResourcesRequest {
	s.AssociatedResourceType = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesRequest) SetRegionId(v string) *GetGlobalAcceleratorResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *GetGlobalAcceleratorResourcesRequest) Validate() error {
	return dara.Validate(s)
}
