// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceSpec interface {
	dara.Model
	String() string
	GoString() string
	SetNodeSpecs(v []*NodeSpec) *ResourceSpec
	GetNodeSpecs() []*NodeSpec
}

type ResourceSpec struct {
	NodeSpecs []*NodeSpec `json:"NodeSpecs,omitempty" xml:"NodeSpecs,omitempty" type:"Repeated"`
}

func (s ResourceSpec) String() string {
	return dara.Prettify(s)
}

func (s ResourceSpec) GoString() string {
	return s.String()
}

func (s *ResourceSpec) GetNodeSpecs() []*NodeSpec {
	return s.NodeSpecs
}

func (s *ResourceSpec) SetNodeSpecs(v []*NodeSpec) *ResourceSpec {
	s.NodeSpecs = v
	return s
}

func (s *ResourceSpec) Validate() error {
	return dara.Validate(s)
}
