// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceInstanceLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllInstances(v bool) *UpdateResourceInstanceLabelRequest
	GetAllInstances() *bool
	SetInstanceIds(v []*string) *UpdateResourceInstanceLabelRequest
	GetInstanceIds() []*string
	SetLabels(v map[string]*string) *UpdateResourceInstanceLabelRequest
	GetLabels() map[string]*string
}

type UpdateResourceInstanceLabelRequest struct {
	// Specifies whether the modification takes effect on all instances in the resource group. If you set this parameter to true, the InstanceIds parameter does not take effect.
	//
	// example:
	//
	// false
	AllInstances *bool `json:"AllInstances,omitempty" xml:"AllInstances,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The custom tag.
	Labels map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateResourceInstanceLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceInstanceLabelRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceInstanceLabelRequest) GetAllInstances() *bool {
	return s.AllInstances
}

func (s *UpdateResourceInstanceLabelRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UpdateResourceInstanceLabelRequest) GetLabels() map[string]*string {
	return s.Labels
}

func (s *UpdateResourceInstanceLabelRequest) SetAllInstances(v bool) *UpdateResourceInstanceLabelRequest {
	s.AllInstances = &v
	return s
}

func (s *UpdateResourceInstanceLabelRequest) SetInstanceIds(v []*string) *UpdateResourceInstanceLabelRequest {
	s.InstanceIds = v
	return s
}

func (s *UpdateResourceInstanceLabelRequest) SetLabels(v map[string]*string) *UpdateResourceInstanceLabelRequest {
	s.Labels = v
	return s
}

func (s *UpdateResourceInstanceLabelRequest) Validate() error {
	return dara.Validate(s)
}
