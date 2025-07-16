// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *GetTicketShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *GetTicketShrinkRequest
	GetOpenTicketId() *string
	SetTenantContextShrink(v string) *GetTicketShrinkRequest
	GetTenantContextShrink() *string
}

type GetTicketShrinkRequest struct {
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
}

func (s GetTicketShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTicketShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *GetTicketShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *GetTicketShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetTicketShrinkRequest) SetOpenTeamId(v string) *GetTicketShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetTicketShrinkRequest) SetOpenTicketId(v string) *GetTicketShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *GetTicketShrinkRequest) SetTenantContextShrink(v string) *GetTicketShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetTicketShrinkRequest) Validate() error {
	return dara.Validate(s)
}
