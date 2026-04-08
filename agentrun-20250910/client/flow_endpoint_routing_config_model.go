// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowEndpointRoutingConfig interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *FlowEndpointRoutingConfig
	GetVersion() *string
	SetWeight(v int) *FlowEndpointRoutingConfig
	GetWeight() *int
}

type FlowEndpointRoutingConfig struct {
	// 目标工作流版本号
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// 该版本在流量分配中的权重，0-100
	//
	// example:
	//
	// 100
	Weight *int `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s FlowEndpointRoutingConfig) String() string {
	return dara.Prettify(s)
}

func (s FlowEndpointRoutingConfig) GoString() string {
	return s.String()
}

func (s *FlowEndpointRoutingConfig) GetVersion() *string {
	return s.Version
}

func (s *FlowEndpointRoutingConfig) GetWeight() *int {
	return s.Weight
}

func (s *FlowEndpointRoutingConfig) SetVersion(v string) *FlowEndpointRoutingConfig {
	s.Version = &v
	return s
}

func (s *FlowEndpointRoutingConfig) SetWeight(v int) *FlowEndpointRoutingConfig {
	s.Weight = &v
	return s
}

func (s *FlowEndpointRoutingConfig) Validate() error {
	return dara.Validate(s)
}
