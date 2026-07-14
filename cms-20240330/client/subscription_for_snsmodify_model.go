// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionForSNSModify interface {
	dara.Model
	String() string
	GoString() string
	SetFilterSetting(v *FilterSetting) *SubscriptionForSNSModify
	GetFilterSetting() *FilterSetting
	SetSubscribeLegacyEvent(v bool) *SubscriptionForSNSModify
	GetSubscribeLegacyEvent() *bool
	SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForSNSModify
	GetWorkspaceFilterSetting() *WorkspaceFilterSetting
}

type SubscriptionForSNSModify struct {
	FilterSetting *FilterSetting `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	// Specifies whether to subscribe to legacy product events (CMS 1.0 / ARMS / SLS events where workspace=null). Valid values:
	//
	// - true: Subscribe.
	//
	// - false or null: Do not subscribe.
	SubscribeLegacyEvent   *bool                   `json:"subscribeLegacyEvent,omitempty" xml:"subscribeLegacyEvent,omitempty"`
	WorkspaceFilterSetting *WorkspaceFilterSetting `json:"workspaceFilterSetting,omitempty" xml:"workspaceFilterSetting,omitempty"`
}

func (s SubscriptionForSNSModify) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionForSNSModify) GoString() string {
	return s.String()
}

func (s *SubscriptionForSNSModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *SubscriptionForSNSModify) GetSubscribeLegacyEvent() *bool {
	return s.SubscribeLegacyEvent
}

func (s *SubscriptionForSNSModify) GetWorkspaceFilterSetting() *WorkspaceFilterSetting {
	return s.WorkspaceFilterSetting
}

func (s *SubscriptionForSNSModify) SetFilterSetting(v *FilterSetting) *SubscriptionForSNSModify {
	s.FilterSetting = v
	return s
}

func (s *SubscriptionForSNSModify) SetSubscribeLegacyEvent(v bool) *SubscriptionForSNSModify {
	s.SubscribeLegacyEvent = &v
	return s
}

func (s *SubscriptionForSNSModify) SetWorkspaceFilterSetting(v *WorkspaceFilterSetting) *SubscriptionForSNSModify {
	s.WorkspaceFilterSetting = v
	return s
}

func (s *SubscriptionForSNSModify) Validate() error {
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
