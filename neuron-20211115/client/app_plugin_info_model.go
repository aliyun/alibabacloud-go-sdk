// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppPluginInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *AppPluginInfo
	GetAppId() *int64
	SetAppVersion(v string) *AppPluginInfo
	GetAppVersion() *string
	SetAppVersionId(v int64) *AppPluginInfo
	GetAppVersionId() *int64
	SetConfig(v *MobiPluginConfig) *AppPluginInfo
	GetConfig() *MobiPluginConfig
	SetGmtCreate(v string) *AppPluginInfo
	GetGmtCreate() *string
	SetGmtModified(v string) *AppPluginInfo
	GetGmtModified() *string
	SetId(v int64) *AppPluginInfo
	GetId() *int64
	SetPluginKey(v string) *AppPluginInfo
	GetPluginKey() *string
	SetPluginType(v string) *AppPluginInfo
	GetPluginType() *string
	SetType(v string) *AppPluginInfo
	GetType() *string
}

type AppPluginInfo struct {
	AppId        *int64            `json:"appId,omitempty" xml:"appId,omitempty"`
	AppVersion   *string           `json:"appVersion,omitempty" xml:"appVersion,omitempty"`
	AppVersionId *int64            `json:"appVersionId,omitempty" xml:"appVersionId,omitempty"`
	Config       *MobiPluginConfig `json:"config,omitempty" xml:"config,omitempty"`
	GmtCreate    *string           `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified  *string           `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id           *int64            `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// MULTIPLEMALL_CONSOLE_LAYOUT_APPSTORE
	PluginKey *string `json:"pluginKey,omitempty" xml:"pluginKey,omitempty"`
	// example:
	//
	// PAGE
	PluginType *string `json:"pluginType,omitempty" xml:"pluginType,omitempty"`
	// example:
	//
	// MOBI
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AppPluginInfo) String() string {
	return dara.Prettify(s)
}

func (s AppPluginInfo) GoString() string {
	return s.String()
}

func (s *AppPluginInfo) GetAppId() *int64 {
	return s.AppId
}

func (s *AppPluginInfo) GetAppVersion() *string {
	return s.AppVersion
}

func (s *AppPluginInfo) GetAppVersionId() *int64 {
	return s.AppVersionId
}

func (s *AppPluginInfo) GetConfig() *MobiPluginConfig {
	return s.Config
}

func (s *AppPluginInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AppPluginInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppPluginInfo) GetId() *int64 {
	return s.Id
}

func (s *AppPluginInfo) GetPluginKey() *string {
	return s.PluginKey
}

func (s *AppPluginInfo) GetPluginType() *string {
	return s.PluginType
}

func (s *AppPluginInfo) GetType() *string {
	return s.Type
}

func (s *AppPluginInfo) SetAppId(v int64) *AppPluginInfo {
	s.AppId = &v
	return s
}

func (s *AppPluginInfo) SetAppVersion(v string) *AppPluginInfo {
	s.AppVersion = &v
	return s
}

func (s *AppPluginInfo) SetAppVersionId(v int64) *AppPluginInfo {
	s.AppVersionId = &v
	return s
}

func (s *AppPluginInfo) SetConfig(v *MobiPluginConfig) *AppPluginInfo {
	s.Config = v
	return s
}

func (s *AppPluginInfo) SetGmtCreate(v string) *AppPluginInfo {
	s.GmtCreate = &v
	return s
}

func (s *AppPluginInfo) SetGmtModified(v string) *AppPluginInfo {
	s.GmtModified = &v
	return s
}

func (s *AppPluginInfo) SetId(v int64) *AppPluginInfo {
	s.Id = &v
	return s
}

func (s *AppPluginInfo) SetPluginKey(v string) *AppPluginInfo {
	s.PluginKey = &v
	return s
}

func (s *AppPluginInfo) SetPluginType(v string) *AppPluginInfo {
	s.PluginType = &v
	return s
}

func (s *AppPluginInfo) SetType(v string) *AppPluginInfo {
	s.Type = &v
	return s
}

func (s *AppPluginInfo) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}
