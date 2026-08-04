// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNodeDrainParameters interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *NodeDrainParameters
	GetForce() *bool
	SetPodFromSubProducts(v []*string) *NodeDrainParameters
	GetPodFromSubProducts() []*string
	SetPodNames(v []*string) *NodeDrainParameters
	GetPodNames() []*string
}

type NodeDrainParameters struct {
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Delete pods of the specified job types
	PodFromSubProducts []*string `json:"PodFromSubProducts,omitempty" xml:"PodFromSubProducts,omitempty" type:"Repeated"`
	// Delete specified pods
	PodNames []*string `json:"PodNames,omitempty" xml:"PodNames,omitempty" type:"Repeated"`
}

func (s NodeDrainParameters) String() string {
	return dara.Prettify(s)
}

func (s NodeDrainParameters) GoString() string {
	return s.String()
}

func (s *NodeDrainParameters) GetForce() *bool {
	return s.Force
}

func (s *NodeDrainParameters) GetPodFromSubProducts() []*string {
	return s.PodFromSubProducts
}

func (s *NodeDrainParameters) GetPodNames() []*string {
	return s.PodNames
}

func (s *NodeDrainParameters) SetForce(v bool) *NodeDrainParameters {
	s.Force = &v
	return s
}

func (s *NodeDrainParameters) SetPodFromSubProducts(v []*string) *NodeDrainParameters {
	s.PodFromSubProducts = v
	return s
}

func (s *NodeDrainParameters) SetPodNames(v []*string) *NodeDrainParameters {
	s.PodNames = v
	return s
}

func (s *NodeDrainParameters) Validate() error {
	return dara.Validate(s)
}
