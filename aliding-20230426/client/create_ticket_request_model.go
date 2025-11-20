// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFields(v string) *CreateTicketRequest
	GetCustomFields() *string
	SetNotify(v *CreateTicketRequestNotify) *CreateTicketRequest
	GetNotify() *CreateTicketRequestNotify
	SetOpenTeamId(v string) *CreateTicketRequest
	GetOpenTeamId() *string
	SetOpenTemplateBizId(v string) *CreateTicketRequest
	GetOpenTemplateBizId() *string
	SetProcessorUserIds(v []*string) *CreateTicketRequest
	GetProcessorUserIds() []*string
	SetScene(v string) *CreateTicketRequest
	GetScene() *string
	SetSceneContext(v *CreateTicketRequestSceneContext) *CreateTicketRequest
	GetSceneContext() *CreateTicketRequestSceneContext
	SetTenantContext(v *CreateTicketRequestTenantContext) *CreateTicketRequest
	GetTenantContext() *CreateTicketRequestTenantContext
	SetTitle(v string) *CreateTicketRequest
	GetTitle() *string
}

type CreateTicketRequest struct {
	// example:
	//
	// []
	CustomFields *string                    `json:"CustomFields,omitempty" xml:"CustomFields,omitempty"`
	Notify       *CreateTicketRequestNotify `json:"Notify,omitempty" xml:"Notify,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// G3IOe205RLciE
	OpenTeamId *string `json:"OpenTeamId,omitempty" xml:"OpenTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	OpenTemplateBizId *string `json:"OpenTemplateBizId,omitempty" xml:"OpenTemplateBizId,omitempty"`
	// This parameter is required.
	ProcessorUserIds []*string `json:"ProcessorUserIds,omitempty" xml:"ProcessorUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// SG
	Scene         *string                           `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SceneContext  *CreateTicketRequestSceneContext  `json:"SceneContext,omitempty" xml:"SceneContext,omitempty" type:"Struct"`
	TenantContext *CreateTicketRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Killer Ball Wo Mouichido
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketRequest) GetCustomFields() *string {
	return s.CustomFields
}

func (s *CreateTicketRequest) GetNotify() *CreateTicketRequestNotify {
	return s.Notify
}

func (s *CreateTicketRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *CreateTicketRequest) GetOpenTemplateBizId() *string {
	return s.OpenTemplateBizId
}

func (s *CreateTicketRequest) GetProcessorUserIds() []*string {
	return s.ProcessorUserIds
}

func (s *CreateTicketRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateTicketRequest) GetSceneContext() *CreateTicketRequestSceneContext {
	return s.SceneContext
}

func (s *CreateTicketRequest) GetTenantContext() *CreateTicketRequestTenantContext {
	return s.TenantContext
}

func (s *CreateTicketRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketRequest) SetCustomFields(v string) *CreateTicketRequest {
	s.CustomFields = &v
	return s
}

func (s *CreateTicketRequest) SetNotify(v *CreateTicketRequestNotify) *CreateTicketRequest {
	s.Notify = v
	return s
}

func (s *CreateTicketRequest) SetOpenTeamId(v string) *CreateTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateTicketRequest) SetOpenTemplateBizId(v string) *CreateTicketRequest {
	s.OpenTemplateBizId = &v
	return s
}

func (s *CreateTicketRequest) SetProcessorUserIds(v []*string) *CreateTicketRequest {
	s.ProcessorUserIds = v
	return s
}

func (s *CreateTicketRequest) SetScene(v string) *CreateTicketRequest {
	s.Scene = &v
	return s
}

func (s *CreateTicketRequest) SetSceneContext(v *CreateTicketRequestSceneContext) *CreateTicketRequest {
	s.SceneContext = v
	return s
}

func (s *CreateTicketRequest) SetTenantContext(v *CreateTicketRequestTenantContext) *CreateTicketRequest {
	s.TenantContext = v
	return s
}

func (s *CreateTicketRequest) SetTitle(v string) *CreateTicketRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketRequest) Validate() error {
	if s.Notify != nil {
		if err := s.Notify.Validate(); err != nil {
			return err
		}
	}
	if s.SceneContext != nil {
		if err := s.SceneContext.Validate(); err != nil {
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

type CreateTicketRequestNotify struct {
	GroupNoticeReceiverUserIds []*string `json:"GroupNoticeReceiverUserIds,omitempty" xml:"GroupNoticeReceiverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NoticeAllGroupMember      *bool     `json:"NoticeAllGroupMember,omitempty" xml:"NoticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUserIds []*string `json:"WorkNoticeReceiverUserIds,omitempty" xml:"WorkNoticeReceiverUserIds,omitempty" type:"Repeated"`
}

func (s CreateTicketRequestNotify) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestNotify) GetGroupNoticeReceiverUserIds() []*string {
	return s.GroupNoticeReceiverUserIds
}

func (s *CreateTicketRequestNotify) GetNoticeAllGroupMember() *bool {
	return s.NoticeAllGroupMember
}

func (s *CreateTicketRequestNotify) GetWorkNoticeReceiverUserIds() []*string {
	return s.WorkNoticeReceiverUserIds
}

func (s *CreateTicketRequestNotify) SetGroupNoticeReceiverUserIds(v []*string) *CreateTicketRequestNotify {
	s.GroupNoticeReceiverUserIds = v
	return s
}

func (s *CreateTicketRequestNotify) SetNoticeAllGroupMember(v bool) *CreateTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *CreateTicketRequestNotify) SetWorkNoticeReceiverUserIds(v []*string) *CreateTicketRequestNotify {
	s.WorkNoticeReceiverUserIds = v
	return s
}

func (s *CreateTicketRequestNotify) Validate() error {
	return dara.Validate(s)
}

type CreateTicketRequestSceneContext struct {
	GroupMsgs []*CreateTicketRequestSceneContextGroupMsgs `json:"GroupMsgs,omitempty" xml:"GroupMsgs,omitempty" type:"Repeated"`
	// example:
	//
	// cidDKVAOW8yVWPEN+WZfwSSAQ==
	OpenConversationId *string   `json:"OpenConversationId,omitempty" xml:"OpenConversationId,omitempty"`
	RelevantorUserIds  []*string `json:"RelevantorUserIds,omitempty" xml:"RelevantorUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// 42674892
	TopicId *string `json:"TopicId,omitempty" xml:"TopicId,omitempty"`
}

func (s CreateTicketRequestSceneContext) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequestSceneContext) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContext) GetGroupMsgs() []*CreateTicketRequestSceneContextGroupMsgs {
	return s.GroupMsgs
}

func (s *CreateTicketRequestSceneContext) GetOpenConversationId() *string {
	return s.OpenConversationId
}

func (s *CreateTicketRequestSceneContext) GetRelevantorUserIds() []*string {
	return s.RelevantorUserIds
}

func (s *CreateTicketRequestSceneContext) GetTopicId() *string {
	return s.TopicId
}

func (s *CreateTicketRequestSceneContext) SetGroupMsgs(v []*CreateTicketRequestSceneContextGroupMsgs) *CreateTicketRequestSceneContext {
	s.GroupMsgs = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetOpenConversationId(v string) *CreateTicketRequestSceneContext {
	s.OpenConversationId = &v
	return s
}

func (s *CreateTicketRequestSceneContext) SetRelevantorUserIds(v []*string) *CreateTicketRequestSceneContext {
	s.RelevantorUserIds = v
	return s
}

func (s *CreateTicketRequestSceneContext) SetTopicId(v string) *CreateTicketRequestSceneContext {
	s.TopicId = &v
	return s
}

func (s *CreateTicketRequestSceneContext) Validate() error {
	if s.GroupMsgs != nil {
		for _, item := range s.GroupMsgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTicketRequestSceneContextGroupMsgs struct {
	// example:
	//
	// true
	Anchor *bool `json:"Anchor,omitempty" xml:"Anchor,omitempty"`
	// example:
	//
	// 1234567
	OpenMsgId *string `json:"OpenMsgId,omitempty" xml:"OpenMsgId,omitempty"`
}

func (s CreateTicketRequestSceneContextGroupMsgs) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequestSceneContextGroupMsgs) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestSceneContextGroupMsgs) GetAnchor() *bool {
	return s.Anchor
}

func (s *CreateTicketRequestSceneContextGroupMsgs) GetOpenMsgId() *string {
	return s.OpenMsgId
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetAnchor(v bool) *CreateTicketRequestSceneContextGroupMsgs {
	s.Anchor = &v
	return s
}

func (s *CreateTicketRequestSceneContextGroupMsgs) SetOpenMsgId(v string) *CreateTicketRequestSceneContextGroupMsgs {
	s.OpenMsgId = &v
	return s
}

func (s *CreateTicketRequestSceneContextGroupMsgs) Validate() error {
	return dara.Validate(s)
}

type CreateTicketRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateTicketRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateTicketRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateTicketRequestTenantContext) SetTenantId(v string) *CreateTicketRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateTicketRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
