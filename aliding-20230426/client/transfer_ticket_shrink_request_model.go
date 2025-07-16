// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyShrink(v string) *TransferTicketShrinkRequest
	GetNotifyShrink() *string
	SetOpenTeamId(v string) *TransferTicketShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *TransferTicketShrinkRequest
	GetOpenTicketId() *string
	SetProcessorUserIdsShrink(v string) *TransferTicketShrinkRequest
	GetProcessorUserIdsShrink() *string
	SetTenantContextShrink(v string) *TransferTicketShrinkRequest
	GetTenantContextShrink() *string
	SetTicketMemoShrink(v string) *TransferTicketShrinkRequest
	GetTicketMemoShrink() *string
}

type TransferTicketShrinkRequest struct {
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

func (s TransferTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *TransferTicketShrinkRequest) GetNotifyShrink() *string {
	return s.NotifyShrink
}

func (s *TransferTicketShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *TransferTicketShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *TransferTicketShrinkRequest) GetProcessorUserIdsShrink() *string {
	return s.ProcessorUserIdsShrink
}

func (s *TransferTicketShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *TransferTicketShrinkRequest) GetTicketMemoShrink() *string {
	return s.TicketMemoShrink
}

func (s *TransferTicketShrinkRequest) SetNotifyShrink(v string) *TransferTicketShrinkRequest {
	s.NotifyShrink = &v
	return s
}

func (s *TransferTicketShrinkRequest) SetOpenTeamId(v string) *TransferTicketShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TransferTicketShrinkRequest) SetOpenTicketId(v string) *TransferTicketShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *TransferTicketShrinkRequest) SetProcessorUserIdsShrink(v string) *TransferTicketShrinkRequest {
	s.ProcessorUserIdsShrink = &v
	return s
}

func (s *TransferTicketShrinkRequest) SetTenantContextShrink(v string) *TransferTicketShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *TransferTicketShrinkRequest) SetTicketMemoShrink(v string) *TransferTicketShrinkRequest {
	s.TicketMemoShrink = &v
	return s
}

func (s *TransferTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
