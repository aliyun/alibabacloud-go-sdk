// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspaceAlertNotification interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GrafanaWorkspaceAlertNotification
	GetId() *int64
	SetIsArms(v bool) *GrafanaWorkspaceAlertNotification
	GetIsArms() *bool
	SetIsDefault(v bool) *GrafanaWorkspaceAlertNotification
	GetIsDefault() *bool
	SetName(v string) *GrafanaWorkspaceAlertNotification
	GetName() *string
	SetSendReminder(v bool) *GrafanaWorkspaceAlertNotification
	GetSendReminder() *bool
	SetSettings(v string) *GrafanaWorkspaceAlertNotification
	GetSettings() *string
	SetType(v string) *GrafanaWorkspaceAlertNotification
	GetType() *string
	SetUid(v string) *GrafanaWorkspaceAlertNotification
	GetUid() *string
}

type GrafanaWorkspaceAlertNotification struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsArms *bool `json:"isArms,omitempty" xml:"isArms,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	SendReminder *bool `json:"sendReminder,omitempty" xml:"sendReminder,omitempty"`
	// example:
	//
	// {}
	Settings *string `json:"settings,omitempty" xml:"settings,omitempty"`
	// example:
	//
	// Email
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// Xfdf******
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s GrafanaWorkspaceAlertNotification) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceAlertNotification) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceAlertNotification) GetId() *int64 {
	return s.Id
}

func (s *GrafanaWorkspaceAlertNotification) GetIsArms() *bool {
	return s.IsArms
}

func (s *GrafanaWorkspaceAlertNotification) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GrafanaWorkspaceAlertNotification) GetName() *string {
	return s.Name
}

func (s *GrafanaWorkspaceAlertNotification) GetSendReminder() *bool {
	return s.SendReminder
}

func (s *GrafanaWorkspaceAlertNotification) GetSettings() *string {
	return s.Settings
}

func (s *GrafanaWorkspaceAlertNotification) GetType() *string {
	return s.Type
}

func (s *GrafanaWorkspaceAlertNotification) GetUid() *string {
	return s.Uid
}

func (s *GrafanaWorkspaceAlertNotification) SetId(v int64) *GrafanaWorkspaceAlertNotification {
	s.Id = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetIsArms(v bool) *GrafanaWorkspaceAlertNotification {
	s.IsArms = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetIsDefault(v bool) *GrafanaWorkspaceAlertNotification {
	s.IsDefault = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetName(v string) *GrafanaWorkspaceAlertNotification {
	s.Name = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetSendReminder(v bool) *GrafanaWorkspaceAlertNotification {
	s.SendReminder = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetSettings(v string) *GrafanaWorkspaceAlertNotification {
	s.Settings = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetType(v string) *GrafanaWorkspaceAlertNotification {
	s.Type = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) SetUid(v string) *GrafanaWorkspaceAlertNotification {
	s.Uid = &v
	return s
}

func (s *GrafanaWorkspaceAlertNotification) Validate() error {
	return dara.Validate(s)
}
