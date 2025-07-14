// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevisionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetContainers(v []*Container) *RevisionConfig
	GetContainers() []*Container
	SetEnableArmsMetrics(v bool) *RevisionConfig
	GetEnableArmsMetrics() *bool
	SetWebNetworkConfig(v *WebNetworkConfig) *RevisionConfig
	GetWebNetworkConfig() *WebNetworkConfig
}

type RevisionConfig struct {
	// This parameter is required.
	Containers        []*Container      `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	EnableArmsMetrics *bool             `json:"EnableArmsMetrics,omitempty" xml:"EnableArmsMetrics,omitempty"`
	WebNetworkConfig  *WebNetworkConfig `json:"WebNetworkConfig,omitempty" xml:"WebNetworkConfig,omitempty"`
}

func (s RevisionConfig) String() string {
	return dara.Prettify(s)
}

func (s RevisionConfig) GoString() string {
	return s.String()
}

func (s *RevisionConfig) GetContainers() []*Container {
	return s.Containers
}

func (s *RevisionConfig) GetEnableArmsMetrics() *bool {
	return s.EnableArmsMetrics
}

func (s *RevisionConfig) GetWebNetworkConfig() *WebNetworkConfig {
	return s.WebNetworkConfig
}

func (s *RevisionConfig) SetContainers(v []*Container) *RevisionConfig {
	s.Containers = v
	return s
}

func (s *RevisionConfig) SetEnableArmsMetrics(v bool) *RevisionConfig {
	s.EnableArmsMetrics = &v
	return s
}

func (s *RevisionConfig) SetWebNetworkConfig(v *WebNetworkConfig) *RevisionConfig {
	s.WebNetworkConfig = v
	return s
}

func (s *RevisionConfig) Validate() error {
	return dara.Validate(s)
}
