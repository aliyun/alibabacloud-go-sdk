// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *ListTicketOperateRecordShrinkRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *ListTicketOperateRecordShrinkRequest
	GetOpenTicketId() *string
	SetTenantContextShrink(v string) *ListTicketOperateRecordShrinkRequest
	GetTenantContextShrink() *string
}

type ListTicketOperateRecordShrinkRequest struct {
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

func (s ListTicketOperateRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordShrinkRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *ListTicketOperateRecordShrinkRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *ListTicketOperateRecordShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *ListTicketOperateRecordShrinkRequest) SetOpenTeamId(v string) *ListTicketOperateRecordShrinkRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ListTicketOperateRecordShrinkRequest) SetOpenTicketId(v string) *ListTicketOperateRecordShrinkRequest {
	s.OpenTicketId = &v
	return s
}

func (s *ListTicketOperateRecordShrinkRequest) SetTenantContextShrink(v string) *ListTicketOperateRecordShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *ListTicketOperateRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
