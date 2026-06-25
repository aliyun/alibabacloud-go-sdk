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
	SetLabelKeysShrink(v string) *DeleteResourceInstanceLabelShrinkRequest
	GetLabelKeysShrink() *string
}

type DeleteResourceInstanceLabelShrinkRequest struct {
	// Specifies whether the operation takes effect on all instances in the resource group. If this parameter is set to true, the InstanceIds parameter does not take effect.
	AllInstances *bool `json:"AllInstances,omitempty" xml:"AllInstances,omitempty"`
	// The machine instance IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// Deprecated
	//
	// The tags to delete. (Deprecated. Use LabelKeys instead.).
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The tags to delete.
	LabelKeysShrink *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
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

func (s *DeleteResourceInstanceLabelShrinkRequest) GetLabelKeysShrink() *string {
	return s.LabelKeysShrink
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

func (s *DeleteResourceInstanceLabelShrinkRequest) SetLabelKeysShrink(v string) *DeleteResourceInstanceLabelShrinkRequest {
	s.LabelKeysShrink = &v
	return s
}

func (s *DeleteResourceInstanceLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
