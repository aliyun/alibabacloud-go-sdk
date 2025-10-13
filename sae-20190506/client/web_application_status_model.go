// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebApplicationStatus interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceCount(v int64) *WebApplicationStatus
	GetInstanceCount() *int64
	SetWebScalingConfig(v *WebScalingConfig) *WebApplicationStatus
	GetWebScalingConfig() *WebScalingConfig
}

type WebApplicationStatus struct {
	InstanceCount    *int64            `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	WebScalingConfig *WebScalingConfig `json:"WebScalingConfig,omitempty" xml:"WebScalingConfig,omitempty"`
}

func (s WebApplicationStatus) String() string {
	return dara.Prettify(s)
}

func (s WebApplicationStatus) GoString() string {
	return s.String()
}

func (s *WebApplicationStatus) GetInstanceCount() *int64 {
	return s.InstanceCount
}

func (s *WebApplicationStatus) GetWebScalingConfig() *WebScalingConfig {
	return s.WebScalingConfig
}

func (s *WebApplicationStatus) SetInstanceCount(v int64) *WebApplicationStatus {
	s.InstanceCount = &v
	return s
}

func (s *WebApplicationStatus) SetWebScalingConfig(v *WebScalingConfig) *WebApplicationStatus {
	s.WebScalingConfig = v
	return s
}

func (s *WebApplicationStatus) Validate() error {
	if s.WebScalingConfig != nil {
		if err := s.WebScalingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
