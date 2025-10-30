// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIdpSyncConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSyncEnabled(v bool) *IdpSyncConfig
	GetAutoSyncEnabled() *bool
	SetIdpDepartmentInfos(v []*IdpSyncConfigIdpDepartmentInfos) *IdpSyncConfig
	GetIdpDepartmentInfos() []*IdpSyncConfigIdpDepartmentInfos
	SetScheduleSyncIntervalSecond(v int64) *IdpSyncConfig
	GetScheduleSyncIntervalSecond() *int64
	SetUserSyncEnabled(v bool) *IdpSyncConfig
	GetUserSyncEnabled() *bool
}

type IdpSyncConfig struct {
	AutoSyncEnabled            *bool                              `json:"AutoSyncEnabled,omitempty" xml:"AutoSyncEnabled,omitempty"`
	IdpDepartmentInfos         []*IdpSyncConfigIdpDepartmentInfos `json:"IdpDepartmentInfos,omitempty" xml:"IdpDepartmentInfos,omitempty" type:"Repeated"`
	ScheduleSyncIntervalSecond *int64                             `json:"ScheduleSyncIntervalSecond,omitempty" xml:"ScheduleSyncIntervalSecond,omitempty"`
	UserSyncEnabled            *bool                              `json:"UserSyncEnabled,omitempty" xml:"UserSyncEnabled,omitempty"`
}

func (s IdpSyncConfig) String() string {
	return dara.Prettify(s)
}

func (s IdpSyncConfig) GoString() string {
	return s.String()
}

func (s *IdpSyncConfig) GetAutoSyncEnabled() *bool {
	return s.AutoSyncEnabled
}

func (s *IdpSyncConfig) GetIdpDepartmentInfos() []*IdpSyncConfigIdpDepartmentInfos {
	return s.IdpDepartmentInfos
}

func (s *IdpSyncConfig) GetScheduleSyncIntervalSecond() *int64 {
	return s.ScheduleSyncIntervalSecond
}

func (s *IdpSyncConfig) GetUserSyncEnabled() *bool {
	return s.UserSyncEnabled
}

func (s *IdpSyncConfig) SetAutoSyncEnabled(v bool) *IdpSyncConfig {
	s.AutoSyncEnabled = &v
	return s
}

func (s *IdpSyncConfig) SetIdpDepartmentInfos(v []*IdpSyncConfigIdpDepartmentInfos) *IdpSyncConfig {
	s.IdpDepartmentInfos = v
	return s
}

func (s *IdpSyncConfig) SetScheduleSyncIntervalSecond(v int64) *IdpSyncConfig {
	s.ScheduleSyncIntervalSecond = &v
	return s
}

func (s *IdpSyncConfig) SetUserSyncEnabled(v bool) *IdpSyncConfig {
	s.UserSyncEnabled = &v
	return s
}

func (s *IdpSyncConfig) Validate() error {
	if s.IdpDepartmentInfos != nil {
		for _, item := range s.IdpDepartmentInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IdpSyncConfigIdpDepartmentInfos struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s IdpSyncConfigIdpDepartmentInfos) String() string {
	return dara.Prettify(s)
}

func (s IdpSyncConfigIdpDepartmentInfos) GoString() string {
	return s.String()
}

func (s *IdpSyncConfigIdpDepartmentInfos) GetId() *string {
	return s.Id
}

func (s *IdpSyncConfigIdpDepartmentInfos) GetName() *string {
	return s.Name
}

func (s *IdpSyncConfigIdpDepartmentInfos) SetId(v string) *IdpSyncConfigIdpDepartmentInfos {
	s.Id = &v
	return s
}

func (s *IdpSyncConfigIdpDepartmentInfos) SetName(v string) *IdpSyncConfigIdpDepartmentInfos {
	s.Name = &v
	return s
}

func (s *IdpSyncConfigIdpDepartmentInfos) Validate() error {
	return dara.Validate(s)
}
