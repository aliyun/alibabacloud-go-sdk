// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *AddTicketMemoShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *AddTicketMemoShrinkRequest
	GetOpenTicketId() *string
	SetTenantContextShrink(v string) *AddTicketMemoShrinkRequest
	GetTenantContextShrink() *string
	SetTicketMemoShrink(v string) *AddTicketMemoShrinkRequest
	GetTicketMemoShrink() *string
}

type AddTicketMemoShrinkRequest struct {
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

func (s AddTicketMemoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddTicketMemoShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *AddTicketMemoShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *AddTicketMemoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *AddTicketMemoShrinkRequest) GetTicketMemoShrink() *string {
	return s.TicketMemoShrink
}

func (s *AddTicketMemoShrinkRequest) SetOpenTeamId(v string) *AddTicketMemoShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddTicketMemoShrinkRequest) SetOpenTicketId(v string) *AddTicketMemoShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AddTicketMemoShrinkRequest) SetTenantContextShrink(v string) *AddTicketMemoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *AddTicketMemoShrinkRequest) SetTicketMemoShrink(v string) *AddTicketMemoShrinkRequest {
	s.TicketMemoShrink = &v
	return s
}

func (s *AddTicketMemoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
