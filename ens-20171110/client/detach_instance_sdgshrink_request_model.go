// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DetachInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
	SetSDGId(v string) *DetachInstanceSDGShrinkRequest
	GetSDGId() *string
}

type DetachInstanceSDGShrinkRequest struct {
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
	// sdg-xxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s DetachInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DetachInstanceSDGShrinkRequest) GetSDGId() *string {
	return s.SDGId
}

func (s *DetachInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *DetachInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DetachInstanceSDGShrinkRequest) SetSDGId(v string) *DetachInstanceSDGShrinkRequest {
	s.SDGId = &v
	return s
}

func (s *DetachInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
