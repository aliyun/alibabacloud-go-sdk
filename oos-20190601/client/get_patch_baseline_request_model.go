// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatchBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetPatchBaselineRequest
	GetName() *string
	SetRegionId(v string) *GetPatchBaselineRequest
	GetRegionId() *string
}

type GetPatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the patch baseline whose details you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPatchBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineRequest) GetName() *string {
	return s.Name
}

func (s *GetPatchBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetPatchBaselineRequest) SetName(v string) *GetPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *GetPatchBaselineRequest) SetRegionId(v string) *GetPatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *GetPatchBaselineRequest) Validate() error {
	return dara.Validate(s)
}
