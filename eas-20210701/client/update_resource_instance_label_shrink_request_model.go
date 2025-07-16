// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceLabelShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllInstances(v bool) *UpdateResourceInstanceLabelShrinkRequest
	GetAllInstances() *bool
	SetInstanceIdsShrink(v string) *UpdateResourceInstanceLabelShrinkRequest
	GetInstanceIdsShrink() *string
	SetLabels(v map[string]*string) *UpdateResourceInstanceLabelShrinkRequest
	GetLabels() map[string]*string
}

type UpdateResourceInstanceLabelShrinkRequest struct {
	// Specifies whether the modification takes effect on all instances in the resource group. If you set this parameter to true, the InstanceIds parameter does not take effect.
	//
	// example:
	//
	// false
	AllInstances *bool `json:"AllInstances,omitempty" xml:"AllInstances,omitempty"`
	// The instance IDs.
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The custom tag.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateResourceInstanceLabelShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceLabelShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceLabelShrinkRequest) GetAllInstances() *bool {
	return s.AllInstances
}

func (s *UpdateResourceInstanceLabelShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *UpdateResourceInstanceLabelShrinkRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *UpdateResourceInstanceLabelShrinkRequest) SetAllInstances(v bool) *UpdateResourceInstanceLabelShrinkRequest {
	s.AllInstances = &v
	return s
}

func (s *UpdateResourceInstanceLabelShrinkRequest) SetInstanceIdsShrink(v string) *UpdateResourceInstanceLabelShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *UpdateResourceInstanceLabelShrinkRequest) SetLabels(v map[string]*string) *UpdateResourceInstanceLabelShrinkRequest {
	s.Labels = v
	return s
}

func (s *UpdateResourceInstanceLabelShrinkRequest) Validate() error {
	return dara.Validate(s)
}
