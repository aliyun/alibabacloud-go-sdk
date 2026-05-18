// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachVscFromMountPointsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDetachInfosShrink(v string) *DetachVscFromMountPointsShrinkRequest
	GetDetachInfosShrink() *string
	SetInputRegionId(v string) *DetachVscFromMountPointsShrinkRequest
	GetInputRegionId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *DetachVscFromMountPointsShrinkRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type DetachVscFromMountPointsShrinkRequest struct {
	DetachInfosShrink *string `json:"DetachInfos,omitempty" xml:"DetachInfos,omitempty"`
	// This parameter is required.
	InputRegionId              *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	UseAssumeRoleChkServerPerm *bool   `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s DetachVscFromMountPointsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachVscFromMountPointsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachVscFromMountPointsShrinkRequest) GetDetachInfosShrink() *string {
	return s.DetachInfosShrink
}

func (s *DetachVscFromMountPointsShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *DetachVscFromMountPointsShrinkRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *DetachVscFromMountPointsShrinkRequest) SetDetachInfosShrink(v string) *DetachVscFromMountPointsShrinkRequest {
	s.DetachInfosShrink = &v
	return s
}

func (s *DetachVscFromMountPointsShrinkRequest) SetInputRegionId(v string) *DetachVscFromMountPointsShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *DetachVscFromMountPointsShrinkRequest) SetUseAssumeRoleChkServerPerm(v bool) *DetachVscFromMountPointsShrinkRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *DetachVscFromMountPointsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
