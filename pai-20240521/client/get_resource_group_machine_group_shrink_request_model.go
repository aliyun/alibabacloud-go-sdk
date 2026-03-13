// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupMachineGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagShrink(v string) *GetResourceGroupMachineGroupShrinkRequest
	GetTagShrink() *string
}

type GetResourceGroupMachineGroupShrinkRequest struct {
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetResourceGroupMachineGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *GetResourceGroupMachineGroupShrinkRequest) SetTagShrink(v string) *GetResourceGroupMachineGroupShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *GetResourceGroupMachineGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
