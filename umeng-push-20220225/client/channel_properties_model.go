// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChannelProperties interface {
	dara.Model
	String() string
	GoString() string
	SetChannelActivity(v string) *ChannelProperties
	GetChannelActivity() *string
	SetChannelFcm(v string) *ChannelProperties
	GetChannelFcm() *string
	SetHarmonyChannelCategory(v string) *ChannelProperties
	GetHarmonyChannelCategory() *string
	SetHuaweiChannelCategory(v string) *ChannelProperties
	GetHuaweiChannelCategory() *string
	SetHuaweiChannelImportance(v string) *ChannelProperties
	GetHuaweiChannelImportance() *string
	SetHuaweiMessageUrgency(v string) *ChannelProperties
	GetHuaweiMessageUrgency() *string
	SetMainActivity(v string) *ChannelProperties
	GetMainActivity() *string
	SetOppoCategory(v string) *ChannelProperties
	GetOppoCategory() *string
	SetOppoChannelId(v string) *ChannelProperties
	GetOppoChannelId() *string
	SetOppoNotifyLevel(v string) *ChannelProperties
	GetOppoNotifyLevel() *string
	SetUseHuaweiMessage(v string) *ChannelProperties
	GetUseHuaweiMessage() *string
	SetUseHuaweiPlainMessage(v string) *ChannelProperties
	GetUseHuaweiPlainMessage() *string
	SetVivoAddBadge(v string) *ChannelProperties
	GetVivoAddBadge() *string
	SetVivoCategory(v string) *ChannelProperties
	GetVivoCategory() *string
	SetVivoPushMode(v string) *ChannelProperties
	GetVivoPushMode() *string
	SetXiaomiChannelId(v string) *ChannelProperties
	GetXiaomiChannelId() *string
}

type ChannelProperties struct {
	ChannelActivity         *string `json:"channelActivity,omitempty" xml:"channelActivity,omitempty"`
	ChannelFcm              *string `json:"channelFcm,omitempty" xml:"channelFcm,omitempty"`
	HarmonyChannelCategory  *string `json:"harmonyChannelCategory,omitempty" xml:"harmonyChannelCategory,omitempty"`
	HuaweiChannelCategory   *string `json:"huaweiChannelCategory,omitempty" xml:"huaweiChannelCategory,omitempty"`
	HuaweiChannelImportance *string `json:"huaweiChannelImportance,omitempty" xml:"huaweiChannelImportance,omitempty"`
	// example:
	//
	// 取值为"NORMAL"和"HIGH",默认为”NORMAL”
	HuaweiMessageUrgency *string `json:"huaweiMessageUrgency,omitempty" xml:"huaweiMessageUrgency,omitempty"`
	MainActivity         *string `json:"mainActivity,omitempty" xml:"mainActivity,omitempty"`
	OppoCategory         *string `json:"oppoCategory,omitempty" xml:"oppoCategory,omitempty"`
	OppoChannelId        *string `json:"oppoChannelId,omitempty" xml:"oppoChannelId,omitempty"`
	OppoNotifyLevel      *string `json:"oppoNotifyLevel,omitempty" xml:"oppoNotifyLevel,omitempty"`
	// example:
	//
	// "true" ,默认为"false"，可不填
	UseHuaweiMessage *string `json:"useHuaweiMessage,omitempty" xml:"useHuaweiMessage,omitempty"`
	// example:
	//
	// true
	UseHuaweiPlainMessage *string `json:"useHuaweiPlainMessage,omitempty" xml:"useHuaweiPlainMessage,omitempty"`
	// example:
	//
	// "true",默认"false"
	VivoAddBadge    *string `json:"vivoAddBadge,omitempty" xml:"vivoAddBadge,omitempty"`
	VivoCategory    *string `json:"vivoCategory,omitempty" xml:"vivoCategory,omitempty"`
	VivoPushMode    *string `json:"vivoPushMode,omitempty" xml:"vivoPushMode,omitempty"`
	XiaomiChannelId *string `json:"xiaomiChannelId,omitempty" xml:"xiaomiChannelId,omitempty"`
}

func (s ChannelProperties) String() string {
	return dara.Prettify(s)
}

func (s ChannelProperties) GoString() string {
	return s.String()
}

func (s *ChannelProperties) GetChannelActivity() *string {
	return s.ChannelActivity
}

func (s *ChannelProperties) GetChannelFcm() *string {
	return s.ChannelFcm
}

func (s *ChannelProperties) GetHarmonyChannelCategory() *string {
	return s.HarmonyChannelCategory
}

func (s *ChannelProperties) GetHuaweiChannelCategory() *string {
	return s.HuaweiChannelCategory
}

func (s *ChannelProperties) GetHuaweiChannelImportance() *string {
	return s.HuaweiChannelImportance
}

func (s *ChannelProperties) GetHuaweiMessageUrgency() *string {
	return s.HuaweiMessageUrgency
}

func (s *ChannelProperties) GetMainActivity() *string {
	return s.MainActivity
}

func (s *ChannelProperties) GetOppoCategory() *string {
	return s.OppoCategory
}

func (s *ChannelProperties) GetOppoChannelId() *string {
	return s.OppoChannelId
}

func (s *ChannelProperties) GetOppoNotifyLevel() *string {
	return s.OppoNotifyLevel
}

func (s *ChannelProperties) GetUseHuaweiMessage() *string {
	return s.UseHuaweiMessage
}

func (s *ChannelProperties) GetUseHuaweiPlainMessage() *string {
	return s.UseHuaweiPlainMessage
}

func (s *ChannelProperties) GetVivoAddBadge() *string {
	return s.VivoAddBadge
}

func (s *ChannelProperties) GetVivoCategory() *string {
	return s.VivoCategory
}

func (s *ChannelProperties) GetVivoPushMode() *string {
	return s.VivoPushMode
}

func (s *ChannelProperties) GetXiaomiChannelId() *string {
	return s.XiaomiChannelId
}

func (s *ChannelProperties) SetChannelActivity(v string) *ChannelProperties {
	s.ChannelActivity = &v
	return s
}

func (s *ChannelProperties) SetChannelFcm(v string) *ChannelProperties {
	s.ChannelFcm = &v
	return s
}

func (s *ChannelProperties) SetHarmonyChannelCategory(v string) *ChannelProperties {
	s.HarmonyChannelCategory = &v
	return s
}

func (s *ChannelProperties) SetHuaweiChannelCategory(v string) *ChannelProperties {
	s.HuaweiChannelCategory = &v
	return s
}

func (s *ChannelProperties) SetHuaweiChannelImportance(v string) *ChannelProperties {
	s.HuaweiChannelImportance = &v
	return s
}

func (s *ChannelProperties) SetHuaweiMessageUrgency(v string) *ChannelProperties {
	s.HuaweiMessageUrgency = &v
	return s
}

func (s *ChannelProperties) SetMainActivity(v string) *ChannelProperties {
	s.MainActivity = &v
	return s
}

func (s *ChannelProperties) SetOppoCategory(v string) *ChannelProperties {
	s.OppoCategory = &v
	return s
}

func (s *ChannelProperties) SetOppoChannelId(v string) *ChannelProperties {
	s.OppoChannelId = &v
	return s
}

func (s *ChannelProperties) SetOppoNotifyLevel(v string) *ChannelProperties {
	s.OppoNotifyLevel = &v
	return s
}

func (s *ChannelProperties) SetUseHuaweiMessage(v string) *ChannelProperties {
	s.UseHuaweiMessage = &v
	return s
}

func (s *ChannelProperties) SetUseHuaweiPlainMessage(v string) *ChannelProperties {
	s.UseHuaweiPlainMessage = &v
	return s
}

func (s *ChannelProperties) SetVivoAddBadge(v string) *ChannelProperties {
	s.VivoAddBadge = &v
	return s
}

func (s *ChannelProperties) SetVivoCategory(v string) *ChannelProperties {
	s.VivoCategory = &v
	return s
}

func (s *ChannelProperties) SetVivoPushMode(v string) *ChannelProperties {
	s.VivoPushMode = &v
	return s
}

func (s *ChannelProperties) SetXiaomiChannelId(v string) *ChannelProperties {
	s.XiaomiChannelId = &v
	return s
}

func (s *ChannelProperties) Validate() error {
	return dara.Validate(s)
}
