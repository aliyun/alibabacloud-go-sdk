// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *ListTicketOperateRecordRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *ListTicketOperateRecordRequest
	GetOpenTicketId() *string
	SetTenantContext(v *ListTicketOperateRecordRequestTenantContext) *ListTicketOperateRecordRequest
	GetTenantContext() *ListTicketOperateRecordRequestTenantContext
}

type ListTicketOperateRecordRequest struct {
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
	OpenTicketId  *string                                      `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	TenantContext *ListTicketOperateRecordRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s ListTicketOperateRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordRequest) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *ListTicketOperateRecordRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *ListTicketOperateRecordRequest) GetTenantContext() *ListTicketOperateRecordRequestTenantContext {
	return s.TenantContext
}

func (s *ListTicketOperateRecordRequest) SetOpenTeamId(v string) *ListTicketOperateRecordRequest {
	s.OpenTeamId = &v
	return s
}

func (s *ListTicketOperateRecordRequest) SetOpenTicketId(v string) *ListTicketOperateRecordRequest {
	s.OpenTicketId = &v
	return s
}

func (s *ListTicketOperateRecordRequest) SetTenantContext(v *ListTicketOperateRecordRequestTenantContext) *ListTicketOperateRecordRequest {
	s.TenantContext = v
	return s
}

func (s *ListTicketOperateRecordRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTicketOperateRecordRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListTicketOperateRecordRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordRequestTenantContext) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *ListTicketOperateRecordRequestTenantContext) SetTenantId(v string) *ListTicketOperateRecordRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *ListTicketOperateRecordRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
