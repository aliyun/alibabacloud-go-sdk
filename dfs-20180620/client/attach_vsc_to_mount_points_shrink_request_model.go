// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachVscToMountPointsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachInfosShrink(v string) *AttachVscToMountPointsShrinkRequest
	GetAttachInfosShrink() *string
	SetInputRegionId(v string) *AttachVscToMountPointsShrinkRequest
	GetInputRegionId() *string
	SetUseAssumeRoleChkServerPerm(v bool) *AttachVscToMountPointsShrinkRequest
	GetUseAssumeRoleChkServerPerm() *bool
}

type AttachVscToMountPointsShrinkRequest struct {
	AttachInfosShrink *string `json:"AttachInfos,omitempty" xml:"AttachInfos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	InputRegionId *string `json:"InputRegionId,omitempty" xml:"InputRegionId,omitempty"`
	// example:
	//
	// false
	UseAssumeRoleChkServerPerm *bool `json:"UseAssumeRoleChkServerPerm,omitempty" xml:"UseAssumeRoleChkServerPerm,omitempty"`
}

func (s AttachVscToMountPointsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachVscToMountPointsShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachVscToMountPointsShrinkRequest) GetAttachInfosShrink() *string {
	return s.AttachInfosShrink
}

func (s *AttachVscToMountPointsShrinkRequest) GetInputRegionId() *string {
	return s.InputRegionId
}

func (s *AttachVscToMountPointsShrinkRequest) GetUseAssumeRoleChkServerPerm() *bool {
	return s.UseAssumeRoleChkServerPerm
}

func (s *AttachVscToMountPointsShrinkRequest) SetAttachInfosShrink(v string) *AttachVscToMountPointsShrinkRequest {
	s.AttachInfosShrink = &v
	return s
}

func (s *AttachVscToMountPointsShrinkRequest) SetInputRegionId(v string) *AttachVscToMountPointsShrinkRequest {
	s.InputRegionId = &v
	return s
}

func (s *AttachVscToMountPointsShrinkRequest) SetUseAssumeRoleChkServerPerm(v bool) *AttachVscToMountPointsShrinkRequest {
	s.UseAssumeRoleChkServerPerm = &v
	return s
}

func (s *AttachVscToMountPointsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
