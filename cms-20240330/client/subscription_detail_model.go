// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionDetail interface {
	dara.Model
	String() string
	GoString() string
	SetFilterSetting(v *FilterSetting) *SubscriptionDetail
	GetFilterSetting() *FilterSetting
	SetSubscribeLegacyEvent(v bool) *SubscriptionDetail
	GetSubscribeLegacyEvent() *bool
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionDetail
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionDetail struct {
	// The filter conditions for event content.
	//
	// example:
	//
	// {}
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// Specifies whether to subscribe to legacy product events (events with an empty workspace from CMS 1.0, ARMS, or SLS).
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

func (s SubscriptionDetail) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionDetail) GoString() string {
	return s.String()
}

func (s *SubscriptionDetail) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionDetail) GetSubscribeLegacyEvent() *bool {
	return s.SubscribeLegacyEvent
}

func (s *SubscriptionDetail) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionDetail) SetFilterSetting(v *FilterSetting) *SubscriptionDetail {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionDetail) SetSubscribeLegacyEvent(v bool) *SubscriptionDetail {
	s.SubscribeLegacyEvent = &v
	return s
}

func (s *SubscriptionDetail) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionDetail {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionDetail) Validate() error {
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
