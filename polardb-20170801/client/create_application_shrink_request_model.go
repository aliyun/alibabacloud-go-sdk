// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationType(v string) *CreateApplicationShrinkRequest
	GetApplicationType() *string
	SetArchitecture(v string) *CreateApplicationShrinkRequest
	GetArchitecture() *string
	SetAutoRenew(v bool) *CreateApplicationShrinkRequest
	GetAutoRenew() *bool
	SetComponentsShrink(v string) *CreateApplicationShrinkRequest
	GetComponentsShrink() *string
	SetDBClusterId(v string) *CreateApplicationShrinkRequest
	GetDBClusterId() *string
	SetDescription(v string) *CreateApplicationShrinkRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateApplicationShrinkRequest
	GetDryRun() *bool
	SetEndpointsShrink(v string) *CreateApplicationShrinkRequest
	GetEndpointsShrink() *string
	SetPayType(v string) *CreateApplicationShrinkRequest
	GetPayType() *string
	SetPeriod(v string) *CreateApplicationShrinkRequest
	GetPeriod() *string
	SetPolarFSInstanceId(v string) *CreateApplicationShrinkRequest
	GetPolarFSInstanceId() *string
	SetRegionId(v string) *CreateApplicationShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateApplicationShrinkRequest
	GetResourceGroupId() *string
	SetUsedTime(v string) *CreateApplicationShrinkRequest
	GetUsedTime() *string
	SetVSwitchId(v string) *CreateApplicationShrinkRequest
	GetVSwitchId() *string
	SetZoneId(v string) *CreateApplicationShrinkRequest
	GetZoneId() *string
}

type CreateApplicationShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// supabase
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// x86
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// example:
	//
	// true
	AutoRenew        *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	DryRun          *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	EndpointsShrink *string `json:"Endpoints,omitempty" xml:"Endpoints,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// Year
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// pcs-********************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-********************
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1
	UsedTime *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// example:
	//
	// vsw-*********************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-beijing-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateApplicationShrinkRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *CreateApplicationShrinkRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateApplicationShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateApplicationShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApplicationShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateApplicationShrinkRequest) GetEndpointsShrink() *string {
	return s.EndpointsShrink
}

func (s *CreateApplicationShrinkRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateApplicationShrinkRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateApplicationShrinkRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *CreateApplicationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateApplicationShrinkRequest) GetUsedTime() *string {
	return s.UsedTime
}

func (s *CreateApplicationShrinkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateApplicationShrinkRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateApplicationShrinkRequest) SetApplicationType(v string) *CreateApplicationShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetArchitecture(v string) *CreateApplicationShrinkRequest {
	s.Architecture = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetAutoRenew(v bool) *CreateApplicationShrinkRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetComponentsShrink(v string) *CreateApplicationShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDBClusterId(v string) *CreateApplicationShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDescription(v string) *CreateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDryRun(v bool) *CreateApplicationShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetEndpointsShrink(v string) *CreateApplicationShrinkRequest {
	s.EndpointsShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPayType(v string) *CreateApplicationShrinkRequest {
	s.PayType = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPeriod(v string) *CreateApplicationShrinkRequest {
	s.Period = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetPolarFSInstanceId(v string) *CreateApplicationShrinkRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRegionId(v string) *CreateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetUsedTime(v string) *CreateApplicationShrinkRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVSwitchId(v string) *CreateApplicationShrinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetZoneId(v string) *CreateApplicationShrinkRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
