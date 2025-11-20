// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *GetTicketRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *GetTicketRequest
	GetOpenTicketId() *string
	SetTenantContext(v *GetTicketRequestTenantContext) *GetTicketRequest
	GetTenantContext() *GetTicketRequestTenantContext
}

type GetTicketRequest struct {
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
	OpenTicketId  *string                        `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	TenantContext *GetTicketRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s GetTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTicketRequest) GoString() string {
	return s.String()
}

func (s *GetTicketRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *GetTicketRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *GetTicketRequest) GetTenantContext() *GetTicketRequestTenantContext {
	return s.TenantContext
}

func (s *GetTicketRequest) SetOpenTeamId(v string) *GetTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *GetTicketRequest) SetOpenTicketId(v string) *GetTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *GetTicketRequest) SetTenantContext(v *GetTicketRequestTenantContext) *GetTicketRequest {
	s.TenantContext = v
	return s
}

func (s *GetTicketRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTicketRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetTicketRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetTicketRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetTicketRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetTicketRequestTenantContext) SetTenantId(v string) *GetTicketRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetTicketRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
