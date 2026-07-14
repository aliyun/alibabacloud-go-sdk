// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionConfig interface {
	dara.Model
	String() string
	GoString() string
	SetFilterSetting(v *FilterSetting) *SubscriptionConfig
	GetFilterSetting() *FilterSetting
	SetSubscribeLegacyEvent(v bool) *SubscriptionConfig
	GetSubscribeLegacyEvent() *bool
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionConfig
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionConfig struct {
	// The event content filtering conditions.
	//
	// example:
	//
	// {}
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// Specifies whether to subscribe to legacy product events (events with an empty workspace from CMS 1.0, ARMS, or SLS). Valid values:
	//
	// - true: Subscribe.
	//
	// - false/null: Do not subscribe.
	//
	// example:
	//
	// false
	SubscribeLegacyEvent *bool `json:"subscribeLegacyEvent,omitempty" xml:"subscribeLegacyEvent,omitempty"`
	// The cross-workspace event routing (global subscription) settings.
	//
	// example:
	//
	// {}
	WorkspaceFilterSetting *WorkspaceFilterSetting `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s SubscriptionConfig) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionConfig) GoString() string {
	return s.String()
}

func (s *SubscriptionConfig) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionConfig) GetSubscribeLegacyEvent() *bool {
	return s.SubscribeLegacyEvent
}

func (s *SubscriptionConfig) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionConfig) SetFilterSetting(v *FilterSetting) *SubscriptionConfig {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionConfig) SetSubscribeLegacyEvent(v bool) *SubscriptionConfig {
	s.SubscribeLegacyEvent = &v
	return s
}

func (s *SubscriptionConfig) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionConfig {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionConfig) Validate() error {
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	if s.WorkspaceFilterSetting != nil {
		if err := s.WorkspaceFilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
