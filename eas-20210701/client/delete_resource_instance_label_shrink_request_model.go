// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstanceLabelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllInstances(v bool) *DeleteResourceInstanceLabelShrinkRequest
	GetAllInstances() *bool
	SetInstanceIdsShrink(v string) *DeleteResourceInstanceLabelShrinkRequest
	GetInstanceIdsShrink() *string
	SetKeysShrink(v string) *DeleteResourceInstanceLabelShrinkRequest
	GetKeysShrink() *string
}

type DeleteResourceInstanceLabelShrinkRequest struct {
	// Specifies whether the delete operation takes effect on all instances in the resource group. If you set this parameter to true, the InstanceIds parameter does not take effect.
	AllInstances *bool `json:"AllInstances,omitempty" xml:"AllInstances,omitempty"`
	// The instance IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The keys of the tags that you want to delete.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
}

func (s DeleteResourceInstanceLabelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstanceLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstanceLabelShrinkRequest) GetAllInstances() *bool {
	return s.AllInstances
}

func (s *DeleteResourceInstanceLabelShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DeleteResourceInstanceLabelShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *DeleteResourceInstanceLabelShrinkRequest) SetAllInstances(v bool) *DeleteResourceInstanceLabelShrinkRequest {
	s.AllInstances = &v
	return s
}

func (s *DeleteResourceInstanceLabelShrinkRequest) SetInstanceIdsShrink(v string) *DeleteResourceInstanceLabelShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DeleteResourceInstanceLabelShrinkRequest) SetKeysShrink(v string) *DeleteResourceInstanceLabelShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *DeleteResourceInstanceLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
