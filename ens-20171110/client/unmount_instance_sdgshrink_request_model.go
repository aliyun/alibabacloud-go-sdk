// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnmountInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *UnmountInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGId(v string) *UnmountInstanceSDGShrinkRequest
	GetSDGId() *string
}

type UnmountInstanceSDGShrinkRequest struct {
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

func (s UnmountInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UnmountInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *UnmountInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *UnmountInstanceSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *UnmountInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *UnmountInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *UnmountInstanceSDGShrinkRequest) SetSDGId(v string) *UnmountInstanceSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *UnmountInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
