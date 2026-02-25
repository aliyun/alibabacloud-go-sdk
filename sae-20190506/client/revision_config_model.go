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
	// The container configurations. You can deploy only one container for each application. The maximum length of this array is 1.
	//
	// This parameter is required.
	Containers []*Container `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// Specifies whether to enable Application Real-Time Monitoring Service (ARMS) monitoring. Valid values:
	//
	// 	- **`true`**: Enables the ARMS monitoring.
	//
	// 	- **`false`**: Disables the ARMS monitoring.
	//
	// example:
	//
	// true
	EnableArmsMetrics *bool `json:"EnableArmsMetrics,omitempty" xml:"EnableArmsMetrics,omitempty"`
	// The network configurations.
	//
	// >  This parameter is used to configure network settings for a version of the application.
	WebNetworkConfig *WebNetworkConfig `json:"WebNetworkConfig,omitempty" xml:"WebNetworkConfig,omitempty"`
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
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebNetworkConfig != nil {
		if err := s.WebNetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
