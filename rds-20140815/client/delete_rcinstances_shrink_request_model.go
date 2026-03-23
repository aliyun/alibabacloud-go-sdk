// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteRCInstancesShrinkRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteRCInstancesShrinkRequest
	GetForce() *bool
	SetInstanceIdShrink(v string) *DeleteRCInstancesShrinkRequest
	GetInstanceIdShrink() *string
	SetRegionId(v string) *DeleteRCInstancesShrinkRequest
	GetRegionId() *string
	SetTerminateSubscription(v bool) *DeleteRCInstancesShrinkRequest
	GetTerminateSubscription() *bool
}

type DeleteRCInstancesShrinkRequest struct {
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Force  *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	InstanceIdShrink      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TerminateSubscription *bool   `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s DeleteRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteRCInstancesShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRCInstancesShrinkRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteRCInstancesShrinkRequest) GetInstanceIdShrink() *string {
	return s.InstanceIdShrink
}

func (s *DeleteRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRCInstancesShrinkRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteRCInstancesShrinkRequest) SetDryRun(v bool) *DeleteRCInstancesShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRCInstancesShrinkRequest) SetForce(v bool) *DeleteRCInstancesShrinkRequest {
	s.Force = &v
	return s
}

func (s *DeleteRCInstancesShrinkRequest) SetInstanceIdShrink(v string) *DeleteRCInstancesShrinkRequest {
	s.InstanceIdShrink = &v
	return s
}

func (s *DeleteRCInstancesShrinkRequest) SetRegionId(v string) *DeleteRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRCInstancesShrinkRequest) SetTerminateSubscription(v bool) *DeleteRCInstancesShrinkRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
