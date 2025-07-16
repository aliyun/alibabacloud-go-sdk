// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyShrink(v string) *AssignTicketShrinkRequest
	GetNotifyShrink() *string
	SetOpenTeamId(v string) *AssignTicketShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *AssignTicketShrinkRequest
	GetOpenTicketId() *string
	SetProcessorUserIdsShrink(v string) *AssignTicketShrinkRequest
	GetProcessorUserIdsShrink() *string
	SetTenantContextShrink(v string) *AssignTicketShrinkRequest
	GetTenantContextShrink() *string
	SetTicketMemoShrink(v string) *AssignTicketShrinkRequest
	GetTicketMemoShrink() *string
}

type AssignTicketShrinkRequest struct {
	NotifyShrink *string `json:"Notify,omitempty" xml:"Notify,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eKWh3xxxxiE
	OpenTeamId *string `json:"OpenTeamId,omitempty" xml:"OpenTeamId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Dq9hP8Sk2v6vQxxxxiE
	OpenTicketId           *string `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	ProcessorUserIdsShrink *string `json:"ProcessorUserIds,omitempty" xml:"ProcessorUserIds,omitempty"`
	TenantContextShrink    *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	TicketMemoShrink       *string `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty"`
}

func (s AssignTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketShrinkRequest) GetNotifyShrink() *string {
	return s.NotifyShrink
}

func (s *AssignTicketShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *AssignTicketShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *AssignTicketShrinkRequest) GetProcessorUserIdsShrink() *string {
	return s.ProcessorUserIdsShrink
}

func (s *AssignTicketShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AssignTicketShrinkRequest) GetTicketMemoShrink() *string {
	return s.TicketMemoShrink
}

func (s *AssignTicketShrinkRequest) SetNotifyShrink(v string) *AssignTicketShrinkRequest {
	s.NotifyShrink = &v
	return s
}

func (s *AssignTicketShrinkRequest) SetOpenTeamId(v string) *AssignTicketShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AssignTicketShrinkRequest) SetOpenTicketId(v string) *AssignTicketShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AssignTicketShrinkRequest) SetProcessorUserIdsShrink(v string) *AssignTicketShrinkRequest {
	s.ProcessorUserIdsShrink = &v
	return s
}

func (s *AssignTicketShrinkRequest) SetTenantContextShrink(v string) *AssignTicketShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AssignTicketShrinkRequest) SetTicketMemoShrink(v string) *AssignTicketShrinkRequest {
	s.TicketMemoShrink = &v
	return s
}

func (s *AssignTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
