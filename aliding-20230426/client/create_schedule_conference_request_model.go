// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *CreateScheduleConferenceRequest
	GetEndTime() *int64
	SetScheduleConfSettingModel(v *CreateScheduleConferenceRequestScheduleConfSettingModel) *CreateScheduleConferenceRequest
	GetScheduleConfSettingModel() *CreateScheduleConferenceRequestScheduleConfSettingModel
	SetStartTime(v int64) *CreateScheduleConferenceRequest
	GetStartTime() *int64
	SetTenantContext(v *CreateScheduleConferenceRequestTenantContext) *CreateScheduleConferenceRequest
	GetTenantContext() *CreateScheduleConferenceRequestTenantContext
	SetTitle(v string) *CreateScheduleConferenceRequest
	GetTitle() *string
}

type CreateScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000L
	EndTime                  *int64                                                   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ScheduleConfSettingModel *CreateScheduleConferenceRequestScheduleConfSettingModel `json:"ScheduleConfSettingModel,omitempty" xml:"ScheduleConfSettingModel,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1687924800000L
	StartTime     *int64                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *CreateScheduleConferenceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 预约会议标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateScheduleConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateScheduleConferenceRequest) GetScheduleConfSettingModel() *CreateScheduleConferenceRequestScheduleConfSettingModel {
	return s.ScheduleConfSettingModel
}

func (s *CreateScheduleConferenceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateScheduleConferenceRequest) GetTenantContext() *CreateScheduleConferenceRequestTenantContext {
	return s.TenantContext
}

func (s *CreateScheduleConferenceRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateScheduleConferenceRequest) SetEndTime(v int64) *CreateScheduleConferenceRequest {
	s.EndTime = &v
	return s
}

func (s *CreateScheduleConferenceRequest) SetScheduleConfSettingModel(v *CreateScheduleConferenceRequestScheduleConfSettingModel) *CreateScheduleConferenceRequest {
	s.ScheduleConfSettingModel = v
	return s
}

func (s *CreateScheduleConferenceRequest) SetStartTime(v int64) *CreateScheduleConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *CreateScheduleConferenceRequest) SetTenantContext(v *CreateScheduleConferenceRequestTenantContext) *CreateScheduleConferenceRequest {
	s.TenantContext = v
	return s
}

func (s *CreateScheduleConferenceRequest) SetTitle(v string) *CreateScheduleConferenceRequest {
	s.Title = &v
	return s
}

func (s *CreateScheduleConferenceRequest) Validate() error {
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

type CreateScheduleConferenceRequestScheduleConfSettingModel struct {
	CohostUserIds []*string `json:"CohostUserIds,omitempty" xml:"CohostUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// dingc02f685fa06381c44ac5d6980864d335
	ConfAllowedCorpId *string `json:"ConfAllowedCorpId,omitempty" xml:"ConfAllowedCorpId,omitempty"`
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	HostUserId *string `json:"HostUserId,omitempty" xml:"HostUserId,omitempty"`
	// example:
	//
	// 0
	LockRoom                    *int32                                                                              `json:"LockRoom,omitempty" xml:"LockRoom,omitempty"`
	MoziConfOpenRecordSetting   *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting   `json:"MoziConfOpenRecordSetting,omitempty" xml:"MoziConfOpenRecordSetting,omitempty" type:"Struct"`
	MoziConfVirtualExtraSetting *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting `json:"MoziConfVirtualExtraSetting,omitempty" xml:"MoziConfVirtualExtraSetting,omitempty" type:"Struct"`
	// example:
	//
	// 1
	MuteOnJoin *int32 `json:"MuteOnJoin,omitempty" xml:"MuteOnJoin,omitempty"`
	// example:
	//
	// 0
	ScreenShareForbidden *int32 `json:"ScreenShareForbidden,omitempty" xml:"ScreenShareForbidden,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModel) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModel) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetCohostUserIds() []*string {
	return s.CohostUserIds
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetConfAllowedCorpId() *string {
	return s.ConfAllowedCorpId
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetHostUserId() *string {
	return s.HostUserId
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetLockRoom() *int32 {
	return s.LockRoom
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetMoziConfOpenRecordSetting() *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	return s.MoziConfOpenRecordSetting
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetMoziConfVirtualExtraSetting() *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	return s.MoziConfVirtualExtraSetting
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetMuteOnJoin() *int32 {
	return s.MuteOnJoin
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) GetScreenShareForbidden() *int32 {
	return s.ScreenShareForbidden
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetCohostUserIds(v []*string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.CohostUserIds = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetConfAllowedCorpId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.ConfAllowedCorpId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetHostUserId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.HostUserId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetLockRoom(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.LockRoom = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMoziConfOpenRecordSetting(v *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MoziConfOpenRecordSetting = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMoziConfVirtualExtraSetting(v *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MoziConfVirtualExtraSetting = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetMuteOnJoin(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.MuteOnJoin = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) SetScreenShareForbidden(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModel {
	s.ScreenShareForbidden = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModel) Validate() error {
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

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting struct {
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

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetIsFollowHost() *bool {
	return s.IsFollowHost
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetMode() *string {
	return s.Mode
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetRecordAutoStart() *int32 {
	return s.RecordAutoStart
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) GetRecordAutoStartType() *int32 {
	return s.RecordAutoStartType
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetIsFollowHost(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.IsFollowHost = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetMode(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.Mode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStart(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStart = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) SetRecordAutoStartType(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting {
	s.RecordAutoStartType = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfOpenRecordSetting) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting struct {
	// example:
	//
	// 2iPOLbpUNMLzB5LuwggiiqiPwiEiE
	CloudRecordOwnerUserId *string `json:"CloudRecordOwnerUserId,omitempty" xml:"CloudRecordOwnerUserId,omitempty"`
	// example:
	//
	// 0
	EnableChat             *int32 `json:"EnableChat,omitempty" xml:"EnableChat,omitempty"`
	EnableWebAnonymousJoin *bool  `json:"EnableWebAnonymousJoin,omitempty" xml:"EnableWebAnonymousJoin,omitempty"`
	// example:
	//
	// 0
	JoinBeforeHost *int32 `json:"JoinBeforeHost,omitempty" xml:"JoinBeforeHost,omitempty"`
	// example:
	//
	// 0
	LockMediaStatusMicMute *int32 `json:"LockMediaStatusMicMute,omitempty" xml:"LockMediaStatusMicMute,omitempty"`
	// example:
	//
	// 0
	LockNick                     *int32                                                                                                            `json:"LockNick,omitempty" xml:"LockNick,omitempty"`
	MinutesOwnerUserId           *string                                                                                                           `json:"MinutesOwnerUserId,omitempty" xml:"MinutesOwnerUserId,omitempty"`
	MoziConfExtensionAppSettings []*CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings `json:"MoziConfExtensionAppSettings,omitempty" xml:"MoziConfExtensionAppSettings,omitempty" type:"Repeated"`
	PushAllMeetingRecords        *bool                                                                                                             `json:"PushAllMeetingRecords,omitempty" xml:"PushAllMeetingRecords,omitempty"`
	PushCloudRecordCard          *bool                                                                                                             `json:"PushCloudRecordCard,omitempty" xml:"PushCloudRecordCard,omitempty"`
	PushMinutesCard              *bool                                                                                                             `json:"PushMinutesCard,omitempty" xml:"PushMinutesCard,omitempty"`
	// example:
	//
	// 1
	WaitingRoom *int32 `json:"WaitingRoom,omitempty" xml:"WaitingRoom,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetCloudRecordOwnerUserId() *string {
	return s.CloudRecordOwnerUserId
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetEnableChat() *int32 {
	return s.EnableChat
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetEnableWebAnonymousJoin() *bool {
	return s.EnableWebAnonymousJoin
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetJoinBeforeHost() *int32 {
	return s.JoinBeforeHost
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetLockMediaStatusMicMute() *int32 {
	return s.LockMediaStatusMicMute
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetLockNick() *int32 {
	return s.LockNick
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetMinutesOwnerUserId() *string {
	return s.MinutesOwnerUserId
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetMoziConfExtensionAppSettings() []*CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	return s.MoziConfExtensionAppSettings
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushAllMeetingRecords() *bool {
	return s.PushAllMeetingRecords
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushCloudRecordCard() *bool {
	return s.PushCloudRecordCard
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetPushMinutesCard() *bool {
	return s.PushMinutesCard
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) GetWaitingRoom() *int32 {
	return s.WaitingRoom
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetCloudRecordOwnerUserId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.CloudRecordOwnerUserId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableChat(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableChat = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetEnableWebAnonymousJoin(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.EnableWebAnonymousJoin = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetJoinBeforeHost(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.JoinBeforeHost = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockMediaStatusMicMute(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockMediaStatusMicMute = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetLockNick(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.LockNick = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMinutesOwnerUserId(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MinutesOwnerUserId = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetMoziConfExtensionAppSettings(v []*CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.MoziConfExtensionAppSettings = v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushAllMeetingRecords(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushAllMeetingRecords = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushCloudRecordCard(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushCloudRecordCard = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetPushMinutesCard(v bool) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.PushMinutesCard = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) SetWaitingRoom(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting {
	s.WaitingRoom = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSetting) Validate() error {
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

type CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings struct {
	// example:
	//
	// 0
	AutoOpenMode *int32 `json:"AutoOpenMode,omitempty" xml:"AutoOpenMode,omitempty"`
	// example:
	//
	// xxx
	CoolAppCode *string `json:"CoolAppCode,omitempty" xml:"CoolAppCode,omitempty"`
	// example:
	//
	// xxx
	ExtensionAppBizData *string `json:"ExtensionAppBizData,omitempty" xml:"ExtensionAppBizData,omitempty"`
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetAutoOpenMode() *int32 {
	return s.AutoOpenMode
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetCoolAppCode() *string {
	return s.CoolAppCode
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) GetExtensionAppBizData() *string {
	return s.ExtensionAppBizData
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetAutoOpenMode(v int32) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.AutoOpenMode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetCoolAppCode(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.CoolAppCode = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) SetExtensionAppBizData(v string) *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings {
	s.ExtensionAppBizData = &v
	return s
}

func (s *CreateScheduleConferenceRequestScheduleConfSettingModelMoziConfVirtualExtraSettingMoziConfExtensionAppSettings) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleConferenceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateScheduleConferenceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateScheduleConferenceRequestTenantContext) SetTenantId(v string) *CreateScheduleConferenceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateScheduleConferenceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
