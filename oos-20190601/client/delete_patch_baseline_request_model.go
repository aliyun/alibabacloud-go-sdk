// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePatchBaselineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeletePatchBaselineRequest
	GetName() *string
	SetRegionId(v string) *DeletePatchBaselineRequest
	GetRegionId() *string
}

type DeletePatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePatchBaselineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineRequest) GetName() *string {
	return s.Name
}

func (s *DeletePatchBaselineRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePatchBaselineRequest) SetName(v string) *DeletePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *DeletePatchBaselineRequest) SetRegionId(v string) *DeletePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePatchBaselineRequest) Validate() error {
	return dara.Validate(s)
}
