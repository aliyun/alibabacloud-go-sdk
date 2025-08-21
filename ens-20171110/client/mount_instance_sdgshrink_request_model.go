// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMountInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *MountInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGId(v string) *MountInstanceSDGShrinkRequest
	GetSDGId() *string
}

type MountInstanceSDGShrinkRequest struct {
	// The IDs of the instances.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The ID of the SDG.
	//
	// This parameter is required.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s MountInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MountInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *MountInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *MountInstanceSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *MountInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *MountInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *MountInstanceSDGShrinkRequest) SetSDGId(v string) *MountInstanceSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *MountInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
