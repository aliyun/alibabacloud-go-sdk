// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutingConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetVersionWeights(v []*VersionWeight) *RoutingConfiguration
	GetVersionWeights() []*VersionWeight
}

type RoutingConfiguration struct {
	// 不同版本的流量权重配置
	VersionWeights []*VersionWeight `json:"versionWeights,omitempty" xml:"versionWeights,omitempty" type:"Repeated"`
}

func (s RoutingConfiguration) String() string {
	return dara.Prettify(s)
}

func (s RoutingConfiguration) GoString() string {
	return s.String()
}

func (s *RoutingConfiguration) GetVersionWeights() []*VersionWeight {
	return s.VersionWeights
}

func (s *RoutingConfiguration) SetVersionWeights(v []*VersionWeight) *RoutingConfiguration {
	s.VersionWeights = v
	return s
}

func (s *RoutingConfiguration) Validate() error {
	return dara.Validate(s)
}
