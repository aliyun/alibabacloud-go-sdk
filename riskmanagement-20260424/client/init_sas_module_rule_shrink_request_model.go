// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitSasModuleRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBind(v int32) *InitSasModuleRuleShrinkRequest
	GetAutoBind() *int32
	SetInstancesShrink(v string) *InitSasModuleRuleShrinkRequest
	GetInstancesShrink() *string
	SetIsTrial(v bool) *InitSasModuleRuleShrinkRequest
	GetIsTrial() *bool
	SetRegionId(v string) *InitSasModuleRuleShrinkRequest
	GetRegionId() *string
}

type InitSasModuleRuleShrinkRequest struct {
	AutoBind        *int32  `json:"AutoBind,omitempty" xml:"AutoBind,omitempty"`
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	IsTrial         *bool   `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitSasModuleRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleShrinkRequest) GetAutoBind() *int32 {
	return s.AutoBind
}

func (s *InitSasModuleRuleShrinkRequest) GetInstancesShrink() *string {
	return s.InstancesShrink
}

func (s *InitSasModuleRuleShrinkRequest) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *InitSasModuleRuleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InitSasModuleRuleShrinkRequest) SetAutoBind(v int32) *InitSasModuleRuleShrinkRequest {
	s.AutoBind = &v
	return s
}

func (s *InitSasModuleRuleShrinkRequest) SetInstancesShrink(v string) *InitSasModuleRuleShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *InitSasModuleRuleShrinkRequest) SetIsTrial(v bool) *InitSasModuleRuleShrinkRequest {
	s.IsTrial = &v
	return s
}

func (s *InitSasModuleRuleShrinkRequest) SetRegionId(v string) *InitSasModuleRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *InitSasModuleRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
