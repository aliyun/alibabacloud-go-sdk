// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFields(v string) *CreateTicketShrinkRequest
	GetCustomFields() *string
	SetNotifyShrink(v string) *CreateTicketShrinkRequest
	GetNotifyShrink() *string
	SetOpenTeamId(v string) *CreateTicketShrinkRequest
	GetOpenTeamId() *string
	SetOpenTemplateBizId(v string) *CreateTicketShrinkRequest
	GetOpenTemplateBizId() *string
	SetProcessorUserIdsShrink(v string) *CreateTicketShrinkRequest
	GetProcessorUserIdsShrink() *string
	SetScene(v string) *CreateTicketShrinkRequest
	GetScene() *string
	SetSceneContextShrink(v string) *CreateTicketShrinkRequest
	GetSceneContextShrink() *string
	SetTenantContextShrink(v string) *CreateTicketShrinkRequest
	GetTenantContextShrink() *string
	SetTitle(v string) *CreateTicketShrinkRequest
	GetTitle() *string
}

type CreateTicketShrinkRequest struct {
	// example:
	//
	// []
	CustomFields *string `json:"CustomFields,omitempty" xml:"CustomFields,omitempty"`
	NotifyShrink *string `json:"Notify,omitempty" xml:"Notify,omitempty"`
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
	ProcessorUserIdsShrink *string `json:"ProcessorUserIds,omitempty" xml:"ProcessorUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SG
	Scene               *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SceneContextShrink  *string `json:"SceneContext,omitempty" xml:"SceneContext,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Killer Ball Wo Mouichido
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTicketShrinkRequest) GetCustomFields() *string {
	return s.CustomFields
}

func (s *CreateTicketShrinkRequest) GetNotifyShrink() *string {
	return s.NotifyShrink
}

func (s *CreateTicketShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *CreateTicketShrinkRequest) GetOpenTemplateBizId() *string {
	return s.OpenTemplateBizId
}

func (s *CreateTicketShrinkRequest) GetProcessorUserIdsShrink() *string {
	return s.ProcessorUserIdsShrink
}

func (s *CreateTicketShrinkRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateTicketShrinkRequest) GetSceneContextShrink() *string {
	return s.SceneContextShrink
}

func (s *CreateTicketShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CreateTicketShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateTicketShrinkRequest) SetCustomFields(v string) *CreateTicketShrinkRequest {
	s.CustomFields = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetNotifyShrink(v string) *CreateTicketShrinkRequest {
	s.NotifyShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetOpenTeamId(v string) *CreateTicketShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetOpenTemplateBizId(v string) *CreateTicketShrinkRequest {
	s.OpenTemplateBizId = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetProcessorUserIdsShrink(v string) *CreateTicketShrinkRequest {
	s.ProcessorUserIdsShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetScene(v string) *CreateTicketShrinkRequest {
	s.Scene = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetSceneContextShrink(v string) *CreateTicketShrinkRequest {
	s.SceneContextShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetTenantContextShrink(v string) *CreateTicketShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CreateTicketShrinkRequest) SetTitle(v string) *CreateTicketShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
