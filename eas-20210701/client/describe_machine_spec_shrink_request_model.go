// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineSpecShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypesShrink(v string) *DescribeMachineSpecShrinkRequest
	GetInstanceTypesShrink() *string
}

type DescribeMachineSpecShrinkRequest struct {
	// Deprecated
	//
	// This parameter is deprecated.
	InstanceTypesShrink *string `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty"`
}

func (s DescribeMachineSpecShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeMachineSpecShrinkRequest) GetInstanceTypesShrink() *string {
	return s.InstanceTypesShrink
}

func (s *DescribeMachineSpecShrinkRequest) SetInstanceTypesShrink(v string) *DescribeMachineSpecShrinkRequest {
	s.InstanceTypesShrink = &v
	return s
}

func (s *DescribeMachineSpecShrinkRequest) Validate() error {
	return dara.Validate(s)
}
