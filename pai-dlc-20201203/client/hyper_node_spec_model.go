// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHyperNodeSpec interface {
	dara.Model
	String() string
	GoString() string
	SetHyperNodeName(v string) *HyperNodeSpec
	GetHyperNodeName() *string
	SetNodeNames(v string) *HyperNodeSpec
	GetNodeNames() *string
}

type HyperNodeSpec struct {
	HyperNodeName *string `json:"HyperNodeName,omitempty" xml:"HyperNodeName,omitempty"`
	NodeNames     *string `json:"NodeNames,omitempty" xml:"NodeNames,omitempty"`
}

func (s HyperNodeSpec) String() string {
	return dara.Prettify(s)
}

func (s HyperNodeSpec) GoString() string {
	return s.String()
}

func (s *HyperNodeSpec) GetHyperNodeName() *string {
	return s.HyperNodeName
}

func (s *HyperNodeSpec) GetNodeNames() *string {
	return s.NodeNames
}

func (s *HyperNodeSpec) SetHyperNodeName(v string) *HyperNodeSpec {
	s.HyperNodeName = &v
	return s
}

func (s *HyperNodeSpec) SetNodeNames(v string) *HyperNodeSpec {
	s.NodeNames = &v
	return s
}

func (s *HyperNodeSpec) Validate() error {
	return dara.Validate(s)
}
