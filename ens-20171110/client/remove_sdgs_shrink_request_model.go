// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *RemoveSDGsShrinkRequest
	GetInstanceIdsShrink() *string
	SetSdgIdsShrink(v string) *RemoveSDGsShrinkRequest
	GetSdgIdsShrink() *string
}

type RemoveSDGsShrinkRequest struct {
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	SdgIdsShrink      *string `json:"SdgIds,omitempty" xml:"SdgIds,omitempty"`
}

func (s RemoveSDGsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveSDGsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RemoveSDGsShrinkRequest) GetSdgIdsShrink() *string {
	return s.SdgIdsShrink
}

func (s *RemoveSDGsShrinkRequest) SetInstanceIdsShrink(v string) *RemoveSDGsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RemoveSDGsShrinkRequest) SetSdgIdsShrink(v string) *RemoveSDGsShrinkRequest {
	s.SdgIdsShrink = &v
	return s
}

func (s *RemoveSDGsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
