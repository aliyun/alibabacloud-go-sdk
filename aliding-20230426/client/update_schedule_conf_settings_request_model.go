// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleConfSettingModel(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) *UpdateScheduleConfSettingsRequest
	GetScheduleConfSettingModel() *UpdateScheduleConfSettingsRequestScheduleConfSettingModel
	SetScheduleConferenceId(v string) *UpdateScheduleConfSettingsRequest
	GetScheduleConferenceId() *string
	SetTenantContext(v *UpdateScheduleConfSettingsRequestTenantContext) *UpdateScheduleConfSettingsRequest
	GetTenantContext() *UpdateScheduleConfSettingsRequestTenantContext
}

type UpdateScheduleConfSettingsRequest struct {
	ScheduleConfSettingModel *UpdateScheduleConfSettingsRequestScheduleConfSettingModel `json:"ScheduleConfSettingModel,omitempty" xml:"ScheduleConfSettingModel,omitempty" type:"Struct"`
	// example:
	//
	// f6fb627e-a7e8-403e-b1f8-26e85450f4a9
	ScheduleConferenceId *string                                         `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContext        *UpdateScheduleConfSettingsRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateScheduleConfSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequest) GetScheduleConfSettingModel() *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	return s.ScheduleConfSettingModel
}

func (s *UpdateScheduleConfSettingsRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *UpdateScheduleConfSettingsRequest) GetTenantContext() *UpdateScheduleConfSettingsRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateScheduleConfSettingsRequest) SetScheduleConfSettingModel(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) *UpdateScheduleConfSettingsRequest {
	s.ScheduleConfSettingModel = v
	return s
}

func (s *UpdateScheduleConfSettingsRequest) SetScheduleConferenceId(v string) *UpdateScheduleConfSettingsRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequest) SetTenantContext(v *UpdateScheduleConfSettingsRequestTenantContext) *UpdateScheduleConfSettingsRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateScheduleConfSettingsRequest) Validate() error {
	if s.ScheduleConfSettingModel != nil {
		if err := s.ScheduleConfSettingModel.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModel struct {
	CohostUserIds []*string `json:"CohostUserIds,omitempty" xml:"CohostUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// ding********
	ConfAllowedCorpId *string `json:"ConfAllowedCorpId,omitempty" xml:"ConfAllowedCorpId,omitempty"`
	// example:
	//
	// 012345
	HostUserId *string `json:"HostUserId,omitempty" xml:"HostUserId,omitempty"`
	// example:
	//
	// 1
	LockRoom                    *int32                                                                                `json:"LockRoom,omitempty" xml:"LockRoom,omitempty"`
	MoziConfOpenRecordSetting   *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting   `json:"MoziConfOpenRecordSetting,omitempty" xml:"MoziConfOpenRecordSetting,omitempty" type:"Struct"`
	MoziConfVirtualExtraSetting *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting `json:"MoziConfVirtualExtraSetting,omitempty" xml:"MoziConfVirtualExtraSetting,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MuteOnJoin *int32 `json:"MuteOnJoin,omitempty" xml:"MuteOnJoin,omitempty"`
	// example:
	//
	// 1
	ScreenShareForbidden *int32 `json:"ScreenShareForbidden,omitempty" xml:"ScreenShareForbidden,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModel) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetCohostUserIds() []*string {
	return s.CohostUserIds
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetConfAllowedCorpId() *string {
	return s.ConfAllowedCorpId
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetHostUserId() *string {
	return s.HostUserId
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetLockRoom() *int32 {
	return s.LockRoom
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetMoziConfOpenRecordSetting() *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	return s.MoziConfOpenRecordSetting
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetMoziConfVirtualExtraSetting() *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	return s.MoziConfVirtualExtraSetting
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetMuteOnJoin() *int32 {
	return s.MuteOnJoin
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) GetScreenShareForbidden() *int32 {
	return s.ScreenShareForbidden
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetCohostUserIds(v []*string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.CohostUserIds = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetConfAllowedCorpId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.ConfAllowedCorpId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetHostUserId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.HostUserId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetLockRoom(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.LockRoom = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMoziConfOpenRecordSetting(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MoziConfOpenRecordSetting = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMoziConfVirtualExtraSetting(v *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MoziConfVirtualExtraSetting = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetMuteOnJoin(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.MuteOnJoin = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) SetScreenShareForbidden(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModel {
	s.ScreenShareForbidden = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModel) Validate() error {
	if s.MoziConfOpenRecordSetting != nil {
		if err := s.MoziConfOpenRecordSetting.Validate(); err != nil {
			return err
		}
	}
	if s.MoziConfVirtualExtraSetting != nil {
		if err := s.MoziConfVirtualExtraSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting struct {
	IsFollowHost *bool `json:"IsFollowHost,omitempty" xml:"IsFollowHost,omitempty"`
	// example:
	//
	// grid
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 0
	RecordAutoStart *int32 `json:"RecordAutoStart,omitempty" xml:"RecordAutoStart,omitempty"`
	// example:
	//
	// 0
	RecordAutoStartType *int32 `json:"RecordAutoStartType,omitempty" xml:"RecordAutoStartType,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetIsFollowHost() *bool {
	return s.IsFollowHost
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetMode() *string {
	return s.Mode
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetRecordAutoStart() *int32 {
	return s.RecordAutoStart
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetRecordAutoStartType() *int32 {
	return s.RecordAutoStartType
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetIsFollowHost(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.IsFollowHost = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetMode(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.Mode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStart(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStart = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStartType(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStartType = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfOpenRecordSetting) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting struct {
	CloudRecordOwnerUserId *string `json:"CloudRecordOwnerUserId,omitempty" xml:"CloudRecordOwnerUserId,omitempty"`
	// example:
	//
	// 1
	EnableChat             *int32 `json:"EnableChat,omitempty" xml:"EnableChat,omitempty"`
	EnableWebAnonymousJoin *bool  `json:"EnableWebAnonymousJoin,omitempty" xml:"EnableWebAnonymousJoin,omitempty"`
	// example:
	//
	// 1
	JoinBeforeHost *int32 `json:"JoinBeforeHost,omitempty" xml:"JoinBeforeHost,omitempty"`
	// example:
	//
	// 1
	LockMediaStatusMicMute *int32 `json:"LockMediaStatusMicMute,omitempty" xml:"LockMediaStatusMicMute,omitempty"`
	// example:
	//
	// 1
	LockNick                     *int32                                                                                                              `json:"LockNick,omitempty" xml:"LockNick,omitempty"`
	MinutesOwnerUserId           *string                                                                                                             `json:"MinutesOwnerUserId,omitempty" xml:"MinutesOwnerUserId,omitempty"`
	MoziConfExtensionAppSettings []*UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings `json:"MoziConfExtensionAppSettings,omitempty" xml:"MoziConfExtensionAppSettings,omitempty" type:"Repeated"`
	PushAllMeetingRecords        *bool                                                                                                               `json:"PushAllMeetingRecords,omitempty" xml:"PushAllMeetingRecords,omitempty"`
	PushCloudRecordCard          *bool                                                                                                               `json:"PushCloudRecordCard,omitempty" xml:"PushCloudRecordCard,omitempty"`
	PushMinutesCard              *bool                                                                                                               `json:"PushMinutesCard,omitempty" xml:"PushMinutesCard,omitempty"`
	// example:
	//
	// 1
	WaitingRoom *int32 `json:"WaitingRoom,omitempty" xml:"WaitingRoom,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetCloudRecordOwnerUserId() *string {
	return s.CloudRecordOwnerUserId
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetEnableChat() *int32 {
	return s.EnableChat
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetEnableWebAnonymousJoin() *bool {
	return s.EnableWebAnonymousJoin
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetJoinBeforeHost() *int32 {
	return s.JoinBeforeHost
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetLockMediaStatusMicMute() *int32 {
	return s.LockMediaStatusMicMute
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetLockNick() *int32 {
	return s.LockNick
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetMinutesOwnerUserId() *string {
	return s.MinutesOwnerUserId
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetMoziConfExtensionAppSettings() []*UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	return s.MoziConfExtensionAppSettings
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushAllMeetingRecords() *bool {
	return s.PushAllMeetingRecords
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushCloudRecordCard() *bool {
	return s.PushCloudRecordCard
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushMinutesCard() *bool {
	return s.PushMinutesCard
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetWaitingRoom() *int32 {
	return s.WaitingRoom
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetCloudRecordOwnerUserId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.CloudRecordOwnerUserId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableChat(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableChat = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableWebAnonymousJoin(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableWebAnonymousJoin = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetJoinBeforeHost(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.JoinBeforeHost = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockMediaStatusMicMute(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockMediaStatusMicMute = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockNick(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockNick = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMinutesOwnerUserId(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MinutesOwnerUserId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMoziConfExtensionAppSettings(v []*UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MoziConfExtensionAppSettings = v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushAllMeetingRecords(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushAllMeetingRecords = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushCloudRecordCard(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushCloudRecordCard = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushMinutesCard(v bool) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushMinutesCard = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetWaitingRoom(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.WaitingRoom = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) Validate() error {
	if s.MoziConfExtensionAppSettings != nil {
		for _, item := range s.MoziConfExtensionAppSettings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings struct {
	// example:
	//
	// 0
	AutoOpenMode *int32 `json:"AutoOpenMode,omitempty" xml:"AutoOpenMode,omitempty"`
	// example:
	//
	// xxxx
	CoolAppCode *string `json:"CoolAppCode,omitempty" xml:"CoolAppCode,omitempty"`
	// example:
	//
	// xxx
	ExtensionAppBizData *string `json:"ExtensionAppBizData,omitempty" xml:"ExtensionAppBizData,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetAutoOpenMode() *int32 {
	return s.AutoOpenMode
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetCoolAppCode() *string {
	return s.CoolAppCode
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetExtensionAppBizData() *string {
	return s.ExtensionAppBizData
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetAutoOpenMode(v int32) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetCoolAppCode(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.CoolAppCode = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetExtensionAppBizData(v string) *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) Validate() error {
	return dara.Validate(s)
}

type UpdateScheduleConfSettingsRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateScheduleConfSettingsRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateScheduleConfSettingsRequestTenantContext) SetTenantId(v string) *UpdateScheduleConfSettingsRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateScheduleConfSettingsRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
