// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDefaultPatchBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RegisterDefaultPatchBaselineRequest
	GetName() *string
	SetRegionId(v string) *RegisterDefaultPatchBaselineRequest
	GetRegionId() *string
}

type RegisterDefaultPatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterDefaultPatchBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s RegisterDefaultPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineRequest) GetName() *string {
	return s.Name
}

func (s *RegisterDefaultPatchBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RegisterDefaultPatchBaselineRequest) SetName(v string) *RegisterDefaultPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *RegisterDefaultPatchBaselineRequest) SetRegionId(v string) *RegisterDefaultPatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineRequest) Validate() error {
	return dara.Validate(s)
}
