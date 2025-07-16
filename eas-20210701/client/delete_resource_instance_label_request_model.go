// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceInstanceLabelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllInstances(v bool) *DeleteResourceInstanceLabelRequest
	GetAllInstances() *bool
	SetInstanceIds(v []*string) *DeleteResourceInstanceLabelRequest
	GetInstanceIds() []*string
	SetKeys(v []*string) *DeleteResourceInstanceLabelRequest
	GetKeys() []*string
}

type DeleteResourceInstanceLabelRequest struct {
	// Specifies whether the delete operation takes effect on all instances in the resource group. If you set this parameter to true, the InstanceIds parameter does not take effect.
	AllInstances *bool `json:"AllInstances,omitempty" xml:"AllInstances,omitempty"`
	// The instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The keys of the tags that you want to delete.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s DeleteResourceInstanceLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceInstanceLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstanceLabelRequest) GetAllInstances() *bool {
	return s.AllInstances
}

func (s *DeleteResourceInstanceLabelRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DeleteResourceInstanceLabelRequest) GetKeys() []*string {
	return s.Keys
}

func (s *DeleteResourceInstanceLabelRequest) SetAllInstances(v bool) *DeleteResourceInstanceLabelRequest {
	s.AllInstances = &v
	return s
}

func (s *DeleteResourceInstanceLabelRequest) SetInstanceIds(v []*string) *DeleteResourceInstanceLabelRequest {
	s.InstanceIds = v
	return s
}

func (s *DeleteResourceInstanceLabelRequest) SetKeys(v []*string) *DeleteResourceInstanceLabelRequest {
	s.Keys = v
	return s
}

func (s *DeleteResourceInstanceLabelRequest) Validate() error {
	return dara.Validate(s)
}
