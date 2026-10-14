// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceSDGShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *RemoveInstanceSDGShrinkRequest
	GetInstanceIdsShrink() *string
}

type RemoveInstanceSDGShrinkRequest struct {
	// The list of instance IDs. The value is a JSON array. A maximum of 100 IDs can be specified.
	//
	// This parameter is required.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s RemoveInstanceSDGShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *RemoveInstanceSDGShrinkRequest) SetInstanceIdsShrink(v string) *RemoveInstanceSDGShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *RemoveInstanceSDGShrinkRequest) Validate() error {
	return dara.Validate(s)
}
