// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScenegroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddFriendForbidden(v int64) *CreateScenegroupRequest
	GetAddFriendForbidden() *int64
	SetAllMembersCanCreateCalendar(v int64) *CreateScenegroupRequest
	GetAllMembersCanCreateCalendar() *int64
	SetAllMembersCanCreateMcsConf(v int64) *CreateScenegroupRequest
	GetAllMembersCanCreateMcsConf() *int64
	SetChatBannedType(v int64) *CreateScenegroupRequest
	GetChatBannedType() *int64
	SetGroupEmailDisabled(v int64) *CreateScenegroupRequest
	GetGroupEmailDisabled() *int64
	SetGroupLiveSwitch(v int64) *CreateScenegroupRequest
	GetGroupLiveSwitch() *int64
	SetIcon(v string) *CreateScenegroupRequest
	GetIcon() *string
	SetManagementType(v int64) *CreateScenegroupRequest
	GetManagementType() *int64
	SetMembersToAdminChat(v int64) *CreateScenegroupRequest
	GetMembersToAdminChat() *int64
	SetMentionAllAuthority(v int64) *CreateScenegroupRequest
	GetMentionAllAuthority() *int64
	SetOnlyAdminCanDing(v int64) *CreateScenegroupRequest
	GetOnlyAdminCanDing() *int64
	SetOnlyAdminCanSetMsgTop(v int64) *CreateScenegroupRequest
	GetOnlyAdminCanSetMsgTop() *int64
	SetSearchable(v int64) *CreateScenegroupRequest
	GetSearchable() *int64
	SetShowHistoryType(v int64) *CreateScenegroupRequest
	GetShowHistoryType() *int64
	SetSubadminIds(v string) *CreateScenegroupRequest
	GetSubadminIds() *string
	SetTemplateId(v string) *CreateScenegroupRequest
	GetTemplateId() *string
	SetTitle(v string) *CreateScenegroupRequest
	GetTitle() *string
	SetUserIds(v string) *CreateScenegroupRequest
	GetUserIds() *string
	SetUuid(v string) *CreateScenegroupRequest
	GetUuid() *string
	SetValidationType(v int64) *CreateScenegroupRequest
	GetValidationType() *int64
}

type CreateScenegroupRequest struct {
	// example:
	//
	// 0
	AddFriendForbidden *int64 `json:"AddFriendForbidden,omitempty" xml:"AddFriendForbidden,omitempty"`
	// example:
	//
	// 0
	AllMembersCanCreateCalendar *int64 `json:"AllMembersCanCreateCalendar,omitempty" xml:"AllMembersCanCreateCalendar,omitempty"`
	// example:
	//
	// 0
	AllMembersCanCreateMcsConf *int64 `json:"AllMembersCanCreateMcsConf,omitempty" xml:"AllMembersCanCreateMcsConf,omitempty"`
	// example:
	//
	// 0
	ChatBannedType *int64 `json:"ChatBannedType,omitempty" xml:"ChatBannedType,omitempty"`
	// example:
	//
	// 0
	GroupEmailDisabled *int64 `json:"GroupEmailDisabled,omitempty" xml:"GroupEmailDisabled,omitempty"`
	// example:
	//
	// 1
	GroupLiveSwitch *int64 `json:"GroupLiveSwitch,omitempty" xml:"GroupLiveSwitch,omitempty"`
	// example:
	//
	// @lADOADma*****QKA
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 0
	ManagementType *int64 `json:"ManagementType,omitempty" xml:"ManagementType,omitempty"`
	// example:
	//
	// 0
	MembersToAdminChat *int64 `json:"MembersToAdminChat,omitempty" xml:"MembersToAdminChat,omitempty"`
	// example:
	//
	// 0
	MentionAllAuthority *int64 `json:"MentionAllAuthority,omitempty" xml:"MentionAllAuthority,omitempty"`
	// example:
	//
	// 0
	OnlyAdminCanDing *int64 `json:"OnlyAdminCanDing,omitempty" xml:"OnlyAdminCanDing,omitempty"`
	// example:
	//
	// 0
	OnlyAdminCanSetMsgTop *int64 `json:"OnlyAdminCanSetMsgTop,omitempty" xml:"OnlyAdminCanSetMsgTop,omitempty"`
	// example:
	//
	// 0
	Searchable *int64 `json:"Searchable,omitempty" xml:"Searchable,omitempty"`
	// example:
	//
	// 0
	ShowHistoryType *int64 `json:"ShowHistoryType,omitempty" xml:"ShowHistoryType,omitempty"`
	// example:
	//
	// 072*****,013*****
	SubadminIds *string `json:"SubadminIds,omitempty" xml:"SubadminIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c354***-***-***-b4ea-6f1ab***65
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 测试群
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 072*****,013*****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// example:
	//
	// axcf*-*****-*****-23da*
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// example:
	//
	// 0
	ValidationType *int64 `json:"ValidationType,omitempty" xml:"ValidationType,omitempty"`
}

func (s CreateScenegroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScenegroupRequest) GoString() string {
	return s.String()
}

func (s *CreateScenegroupRequest) GetAddFriendForbidden() *int64 {
	return s.AddFriendForbidden
}

func (s *CreateScenegroupRequest) GetAllMembersCanCreateCalendar() *int64 {
	return s.AllMembersCanCreateCalendar
}

func (s *CreateScenegroupRequest) GetAllMembersCanCreateMcsConf() *int64 {
	return s.AllMembersCanCreateMcsConf
}

func (s *CreateScenegroupRequest) GetChatBannedType() *int64 {
	return s.ChatBannedType
}

func (s *CreateScenegroupRequest) GetGroupEmailDisabled() *int64 {
	return s.GroupEmailDisabled
}

func (s *CreateScenegroupRequest) GetGroupLiveSwitch() *int64 {
	return s.GroupLiveSwitch
}

func (s *CreateScenegroupRequest) GetIcon() *string {
	return s.Icon
}

func (s *CreateScenegroupRequest) GetManagementType() *int64 {
	return s.ManagementType
}

func (s *CreateScenegroupRequest) GetMembersToAdminChat() *int64 {
	return s.MembersToAdminChat
}

func (s *CreateScenegroupRequest) GetMentionAllAuthority() *int64 {
	return s.MentionAllAuthority
}

func (s *CreateScenegroupRequest) GetOnlyAdminCanDing() *int64 {
	return s.OnlyAdminCanDing
}

func (s *CreateScenegroupRequest) GetOnlyAdminCanSetMsgTop() *int64 {
	return s.OnlyAdminCanSetMsgTop
}

func (s *CreateScenegroupRequest) GetSearchable() *int64 {
	return s.Searchable
}

func (s *CreateScenegroupRequest) GetShowHistoryType() *int64 {
	return s.ShowHistoryType
}

func (s *CreateScenegroupRequest) GetSubadminIds() *string {
	return s.SubadminIds
}

func (s *CreateScenegroupRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateScenegroupRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateScenegroupRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *CreateScenegroupRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateScenegroupRequest) GetValidationType() *int64 {
	return s.ValidationType
}

func (s *CreateScenegroupRequest) SetAddFriendForbidden(v int64) *CreateScenegroupRequest {
	s.AddFriendForbidden = &v
	return s
}

func (s *CreateScenegroupRequest) SetAllMembersCanCreateCalendar(v int64) *CreateScenegroupRequest {
	s.AllMembersCanCreateCalendar = &v
	return s
}

func (s *CreateScenegroupRequest) SetAllMembersCanCreateMcsConf(v int64) *CreateScenegroupRequest {
	s.AllMembersCanCreateMcsConf = &v
	return s
}

func (s *CreateScenegroupRequest) SetChatBannedType(v int64) *CreateScenegroupRequest {
	s.ChatBannedType = &v
	return s
}

func (s *CreateScenegroupRequest) SetGroupEmailDisabled(v int64) *CreateScenegroupRequest {
	s.GroupEmailDisabled = &v
	return s
}

func (s *CreateScenegroupRequest) SetGroupLiveSwitch(v int64) *CreateScenegroupRequest {
	s.GroupLiveSwitch = &v
	return s
}

func (s *CreateScenegroupRequest) SetIcon(v string) *CreateScenegroupRequest {
	s.Icon = &v
	return s
}

func (s *CreateScenegroupRequest) SetManagementType(v int64) *CreateScenegroupRequest {
	s.ManagementType = &v
	return s
}

func (s *CreateScenegroupRequest) SetMembersToAdminChat(v int64) *CreateScenegroupRequest {
	s.MembersToAdminChat = &v
	return s
}

func (s *CreateScenegroupRequest) SetMentionAllAuthority(v int64) *CreateScenegroupRequest {
	s.MentionAllAuthority = &v
	return s
}

func (s *CreateScenegroupRequest) SetOnlyAdminCanDing(v int64) *CreateScenegroupRequest {
	s.OnlyAdminCanDing = &v
	return s
}

func (s *CreateScenegroupRequest) SetOnlyAdminCanSetMsgTop(v int64) *CreateScenegroupRequest {
	s.OnlyAdminCanSetMsgTop = &v
	return s
}

func (s *CreateScenegroupRequest) SetSearchable(v int64) *CreateScenegroupRequest {
	s.Searchable = &v
	return s
}

func (s *CreateScenegroupRequest) SetShowHistoryType(v int64) *CreateScenegroupRequest {
	s.ShowHistoryType = &v
	return s
}

func (s *CreateScenegroupRequest) SetSubadminIds(v string) *CreateScenegroupRequest {
	s.SubadminIds = &v
	return s
}

func (s *CreateScenegroupRequest) SetTemplateId(v string) *CreateScenegroupRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateScenegroupRequest) SetTitle(v string) *CreateScenegroupRequest {
	s.Title = &v
	return s
}

func (s *CreateScenegroupRequest) SetUserIds(v string) *CreateScenegroupRequest {
	s.UserIds = &v
	return s
}

func (s *CreateScenegroupRequest) SetUuid(v string) *CreateScenegroupRequest {
	s.Uuid = &v
	return s
}

func (s *CreateScenegroupRequest) SetValidationType(v int64) *CreateScenegroupRequest {
	s.ValidationType = &v
	return s
}

func (s *CreateScenegroupRequest) Validate() error {
	return dara.Validate(s)
}
