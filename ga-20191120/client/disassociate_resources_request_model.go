// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *DisassociateResourcesRequest
	GetAcceleratorId() *string
	SetAssociatedResourceId(v string) *DisassociateResourcesRequest
	GetAssociatedResourceId() *string
	SetAssociatedResourceRegionId(v string) *DisassociateResourcesRequest
	GetAssociatedResourceRegionId() *string
	SetAssociatedResourceType(v string) *DisassociateResourcesRequest
	GetAssociatedResourceType() *string
	SetDryRun(v bool) *DisassociateResourcesRequest
	GetDryRun() *bool
	SetRegionId(v string) *DisassociateResourcesRequest
	GetRegionId() *string
}

type DisassociateResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cd498437eb9a346c38be8482689800e91
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// example:
	//
	// cn-hangzhou
	AssociatedResourceRegionId *string `json:"AssociatedResourceRegionId,omitempty" xml:"AssociatedResourceRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// WAF
	AssociatedResourceType *string `json:"AssociatedResourceType,omitempty" xml:"AssociatedResourceType,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisassociateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateResourcesRequest) GoString() string {
	return s.String()
}

func (s *DisassociateResourcesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *DisassociateResourcesRequest) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *DisassociateResourcesRequest) GetAssociatedResourceRegionId() *string {
	return s.AssociatedResourceRegionId
}

func (s *DisassociateResourcesRequest) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *DisassociateResourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisassociateResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisassociateResourcesRequest) SetAcceleratorId(v string) *DisassociateResourcesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *DisassociateResourcesRequest) SetAssociatedResourceId(v string) *DisassociateResourcesRequest {
	s.AssociatedResourceId = &v
	return s
}

func (s *DisassociateResourcesRequest) SetAssociatedResourceRegionId(v string) *DisassociateResourcesRequest {
	s.AssociatedResourceRegionId = &v
	return s
}

func (s *DisassociateResourcesRequest) SetAssociatedResourceType(v string) *DisassociateResourcesRequest {
	s.AssociatedResourceType = &v
	return s
}

func (s *DisassociateResourcesRequest) SetDryRun(v bool) *DisassociateResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *DisassociateResourcesRequest) SetRegionId(v string) *DisassociateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DisassociateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
