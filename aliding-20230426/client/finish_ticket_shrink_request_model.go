// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyShrink(v string) *FinishTicketShrinkRequest
	GetNotifyShrink() *string
	SetOpenTeamId(v string) *FinishTicketShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *FinishTicketShrinkRequest
	GetOpenTicketId() *string
	SetTenantContextShrink(v string) *FinishTicketShrinkRequest
	GetTenantContextShrink() *string
	SetTicketMemoShrink(v string) *FinishTicketShrinkRequest
	GetTicketMemoShrink() *string
}

type FinishTicketShrinkRequest struct {
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
	OpenTicketId        *string `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	TicketMemoShrink    *string `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty"`
}

func (s FinishTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *FinishTicketShrinkRequest) GetNotifyShrink() *string {
	return s.NotifyShrink
}

func (s *FinishTicketShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *FinishTicketShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *FinishTicketShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *FinishTicketShrinkRequest) GetTicketMemoShrink() *string {
	return s.TicketMemoShrink
}

func (s *FinishTicketShrinkRequest) SetNotifyShrink(v string) *FinishTicketShrinkRequest {
	s.NotifyShrink = &v
	return s
}

func (s *FinishTicketShrinkRequest) SetOpenTeamId(v string) *FinishTicketShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *FinishTicketShrinkRequest) SetOpenTicketId(v string) *FinishTicketShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *FinishTicketShrinkRequest) SetTenantContextShrink(v string) *FinishTicketShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *FinishTicketShrinkRequest) SetTicketMemoShrink(v string) *FinishTicketShrinkRequest {
	s.TicketMemoShrink = &v
	return s
}

func (s *FinishTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
