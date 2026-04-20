// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseAutoScalePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScale(v bool) *ModifySupabaseAutoScalePolicyRequest
	GetAutoScale() *bool
	SetProjectId(v string) *ModifySupabaseAutoScalePolicyRequest
	GetProjectId() *string
	SetRegionId(v string) *ModifySupabaseAutoScalePolicyRequest
	GetRegionId() *string
}

type ModifySupabaseAutoScalePolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoScale *bool `json:"AutoScale,omitempty" xml:"AutoScale,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sbp-tyarplz****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifySupabaseAutoScalePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseAutoScalePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySupabaseAutoScalePolicyRequest) GetAutoScale() *bool {
	return s.AutoScale
}

func (s *ModifySupabaseAutoScalePolicyRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ModifySupabaseAutoScalePolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySupabaseAutoScalePolicyRequest) SetAutoScale(v bool) *ModifySupabaseAutoScalePolicyRequest {
	s.AutoScale = &v
	return s
}

func (s *ModifySupabaseAutoScalePolicyRequest) SetProjectId(v string) *ModifySupabaseAutoScalePolicyRequest {
	s.ProjectId = &v
	return s
}

func (s *ModifySupabaseAutoScalePolicyRequest) SetRegionId(v string) *ModifySupabaseAutoScalePolicyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySupabaseAutoScalePolicyRequest) Validate() error {
	return dara.Validate(s)
}
