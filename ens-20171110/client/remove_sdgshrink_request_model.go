// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *RemoveSDGShrinkRequest
	GetInstanceIdsShrink() *string
}

type RemoveSDGShrinkRequest struct {
	// IDs of Android in Container (AIC) instances.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s RemoveSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RemoveSDGShrinkRequest) SetInstanceIdsShrink(v string) *RemoveSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RemoveSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
