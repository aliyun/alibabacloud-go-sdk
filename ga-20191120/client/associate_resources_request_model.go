// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *AssociateResourcesRequest
	GetAcceleratorId() *string
	SetAssociatedMode(v string) *AssociateResourcesRequest
	GetAssociatedMode() *string
	SetAssociatedResourceId(v string) *AssociateResourcesRequest
	GetAssociatedResourceId() *string
	SetAssociatedResourceRegionId(v string) *AssociateResourcesRequest
	GetAssociatedResourceRegionId() *string
	SetAssociatedResourceType(v string) *AssociateResourcesRequest
	GetAssociatedResourceType() *string
	SetDryRun(v bool) *AssociateResourcesRequest
	GetDryRun() *bool
	SetRegionId(v string) *AssociateResourcesRequest
	GetRegionId() *string
}

type AssociateResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// example:
	//
	// Associated
	AssociatedMode *string `json:"AssociatedMode,omitempty" xml:"AssociatedMode,omitempty"`
	// example:
	//
	// waf_xx
	AssociatedResourceId *string `json:"AssociatedResourceId,omitempty" xml:"AssociatedResourceId,omitempty"`
	// This parameter is required.
	//
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

func (s AssociateResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateResourcesRequest) GoString() string {
	return s.String()
}

func (s *AssociateResourcesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *AssociateResourcesRequest) GetAssociatedMode() *string {
	return s.AssociatedMode
}

func (s *AssociateResourcesRequest) GetAssociatedResourceId() *string {
	return s.AssociatedResourceId
}

func (s *AssociateResourcesRequest) GetAssociatedResourceRegionId() *string {
	return s.AssociatedResourceRegionId
}

func (s *AssociateResourcesRequest) GetAssociatedResourceType() *string {
	return s.AssociatedResourceType
}

func (s *AssociateResourcesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateResourcesRequest) SetAcceleratorId(v string) *AssociateResourcesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AssociateResourcesRequest) SetAssociatedMode(v string) *AssociateResourcesRequest {
	s.AssociatedMode = &v
	return s
}

func (s *AssociateResourcesRequest) SetAssociatedResourceId(v string) *AssociateResourcesRequest {
	s.AssociatedResourceId = &v
	return s
}

func (s *AssociateResourcesRequest) SetAssociatedResourceRegionId(v string) *AssociateResourcesRequest {
	s.AssociatedResourceRegionId = &v
	return s
}

func (s *AssociateResourcesRequest) SetAssociatedResourceType(v string) *AssociateResourcesRequest {
	s.AssociatedResourceType = &v
	return s
}

func (s *AssociateResourcesRequest) SetDryRun(v bool) *AssociateResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateResourcesRequest) SetRegionId(v string) *AssociateResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateResourcesRequest) Validate() error {
	return dara.Validate(s)
}
