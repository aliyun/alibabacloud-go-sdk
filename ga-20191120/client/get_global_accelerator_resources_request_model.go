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
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// example:
	//
	// waf_v2_public_cn-x0r****gr1i
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	AssociatedResourceRegionId *string `json:"AssociatedResourceRegionId,omitempty" xml:"AssociatedResourceRegionId,omitempty"`
	// example:
	//
	// WAF
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
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
