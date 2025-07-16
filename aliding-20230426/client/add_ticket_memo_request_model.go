// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketMemoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenTeamId(v string) *AddTicketMemoRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *AddTicketMemoRequest
	GetOpenTicketId() *string
	SetTenantContext(v *AddTicketMemoRequestTenantContext) *AddTicketMemoRequest
	GetTenantContext() *AddTicketMemoRequestTenantContext
	SetTicketMemo(v *AddTicketMemoRequestTicketMemo) *AddTicketMemoRequest
	GetTicketMemo() *AddTicketMemoRequestTicketMemo
}

type AddTicketMemoRequest struct {
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
	OpenTicketId  *string                            `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	TenantContext *AddTicketMemoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	TicketMemo    *AddTicketMemoRequestTicketMemo    `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty" type:"Struct"`
}

func (s AddTicketMemoRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoRequest) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *AddTicketMemoRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *AddTicketMemoRequest) GetTenantContext() *AddTicketMemoRequestTenantContext {
	return s.TenantContext
}

func (s *AddTicketMemoRequest) GetTicketMemo() *AddTicketMemoRequestTicketMemo {
	return s.TicketMemo
}

func (s *AddTicketMemoRequest) SetOpenTeamId(v string) *AddTicketMemoRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AddTicketMemoRequest) SetOpenTicketId(v string) *AddTicketMemoRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AddTicketMemoRequest) SetTenantContext(v *AddTicketMemoRequestTenantContext) *AddTicketMemoRequest {
	s.TenantContext = v
	return s
}

func (s *AddTicketMemoRequest) SetTicketMemo(v *AddTicketMemoRequestTicketMemo) *AddTicketMemoRequest {
	s.TicketMemo = v
	return s
}

func (s *AddTicketMemoRequest) Validate() error {
	return dara.Validate(s)
}

type AddTicketMemoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddTicketMemoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddTicketMemoRequestTenantContext) SetTenantId(v string) *AddTicketMemoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddTicketMemoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type AddTicketMemoRequestTicketMemo struct {
	Attachments []*AddTicketMemoRequestTicketMemoAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	Memo        *string                                      `json:"Memo,omitempty" xml:"Memo,omitempty"`
}

func (s AddTicketMemoRequestTicketMemo) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemo) GetAttachments() []*AddTicketMemoRequestTicketMemoAttachments {
	return s.Attachments
}

func (s *AddTicketMemoRequestTicketMemo) GetMemo() *string {
	return s.Memo
}

func (s *AddTicketMemoRequestTicketMemo) SetAttachments(v []*AddTicketMemoRequestTicketMemoAttachments) *AddTicketMemoRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *AddTicketMemoRequestTicketMemo) SetMemo(v string) *AddTicketMemoRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemo) Validate() error {
	return dara.Validate(s)
}

type AddTicketMemoRequestTicketMemoAttachments struct {
	// example:
	//
	// "ticket/image/44xxxx9/43003/e27204b38xxxx1640499.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// wahaha.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s AddTicketMemoRequestTicketMemoAttachments) String() string {
	return dara.Prettify(s)
}

func (s AddTicketMemoRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AddTicketMemoRequestTicketMemoAttachments) GetFileName() *string {
	return s.FileName
}

func (s *AddTicketMemoRequestTicketMemoAttachments) GetKey() *string {
	return s.Key
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetFileName(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemoAttachments) SetKey(v string) *AddTicketMemoRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

func (s *AddTicketMemoRequestTicketMemoAttachments) Validate() error {
	return dara.Validate(s)
}
