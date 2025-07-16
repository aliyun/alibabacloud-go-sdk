// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotify(v *FinishTicketRequestNotify) *FinishTicketRequest
	GetNotify() *FinishTicketRequestNotify
	SetOpenTeamId(v string) *FinishTicketRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *FinishTicketRequest
	GetOpenTicketId() *string
	SetTenantContext(v *FinishTicketRequestTenantContext) *FinishTicketRequest
	GetTenantContext() *FinishTicketRequestTenantContext
	SetTicketMemo(v *FinishTicketRequestTicketMemo) *FinishTicketRequest
	GetTicketMemo() *FinishTicketRequestTicketMemo
}

type FinishTicketRequest struct {
	Notify *FinishTicketRequestNotify `json:"Notify,omitempty" xml:"Notify,omitempty" type:"Struct"`
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
	OpenTicketId  *string                           `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	TenantContext *FinishTicketRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	TicketMemo    *FinishTicketRequestTicketMemo    `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty" type:"Struct"`
}

func (s FinishTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketRequest) GoString() string {
	return s.String()
}

func (s *FinishTicketRequest) GetNotify() *FinishTicketRequestNotify {
	return s.Notify
}

func (s *FinishTicketRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *FinishTicketRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *FinishTicketRequest) GetTenantContext() *FinishTicketRequestTenantContext {
	return s.TenantContext
}

func (s *FinishTicketRequest) GetTicketMemo() *FinishTicketRequestTicketMemo {
	return s.TicketMemo
}

func (s *FinishTicketRequest) SetNotify(v *FinishTicketRequestNotify) *FinishTicketRequest {
	s.Notify = v
	return s
}

func (s *FinishTicketRequest) SetOpenTeamId(v string) *FinishTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *FinishTicketRequest) SetOpenTicketId(v string) *FinishTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *FinishTicketRequest) SetTenantContext(v *FinishTicketRequestTenantContext) *FinishTicketRequest {
	s.TenantContext = v
	return s
}

func (s *FinishTicketRequest) SetTicketMemo(v *FinishTicketRequestTicketMemo) *FinishTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *FinishTicketRequest) Validate() error {
	return dara.Validate(s)
}

type FinishTicketRequestNotify struct {
	GroupNoticeReceiverUserIds []*string `json:"GroupNoticeReceiverUserIds,omitempty" xml:"GroupNoticeReceiverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember      *bool     `json:"NoticeAllGroupMember,omitempty" xml:"NoticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUserIds []*string `json:"WorkNoticeReceiverUserIds,omitempty" xml:"WorkNoticeReceiverUserIds,omitempty" type:"Repeated"`
}

func (s FinishTicketRequestNotify) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestNotify) GetGroupNoticeReceiverUserIds() []*string {
	return s.GroupNoticeReceiverUserIds
}

func (s *FinishTicketRequestNotify) GetNoticeAllGroupMember() *bool {
	return s.NoticeAllGroupMember
}

func (s *FinishTicketRequestNotify) GetWorkNoticeReceiverUserIds() []*string {
	return s.WorkNoticeReceiverUserIds
}

func (s *FinishTicketRequestNotify) SetGroupNoticeReceiverUserIds(v []*string) *FinishTicketRequestNotify {
	s.GroupNoticeReceiverUserIds = v
	return s
}

func (s *FinishTicketRequestNotify) SetNoticeAllGroupMember(v bool) *FinishTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *FinishTicketRequestNotify) SetWorkNoticeReceiverUserIds(v []*string) *FinishTicketRequestNotify {
	s.WorkNoticeReceiverUserIds = v
	return s
}

func (s *FinishTicketRequestNotify) Validate() error {
	return dara.Validate(s)
}

type FinishTicketRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s FinishTicketRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketRequestTenantContext) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *FinishTicketRequestTenantContext) SetTenantId(v string) *FinishTicketRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *FinishTicketRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type FinishTicketRequestTicketMemo struct {
	Attachments []*FinishTicketRequestTicketMemoAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	Memo        *string                                     `json:"Memo,omitempty" xml:"Memo,omitempty"`
}

func (s FinishTicketRequestTicketMemo) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemo) GetAttachments() []*FinishTicketRequestTicketMemoAttachments {
	return s.Attachments
}

func (s *FinishTicketRequestTicketMemo) GetMemo() *string {
	return s.Memo
}

func (s *FinishTicketRequestTicketMemo) SetAttachments(v []*FinishTicketRequestTicketMemoAttachments) *FinishTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *FinishTicketRequestTicketMemo) SetMemo(v string) *FinishTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *FinishTicketRequestTicketMemo) Validate() error {
	return dara.Validate(s)
}

type FinishTicketRequestTicketMemoAttachments struct {
	// example:
	//
	// wahaha.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// ticket/image/44xxxx9/43003/e27xxxx1640499.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s FinishTicketRequestTicketMemoAttachments) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *FinishTicketRequestTicketMemoAttachments) GetFileName() *string {
	return s.FileName
}

func (s *FinishTicketRequestTicketMemoAttachments) GetKey() *string {
	return s.Key
}

func (s *FinishTicketRequestTicketMemoAttachments) SetFileName(v string) *FinishTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *FinishTicketRequestTicketMemoAttachments) SetKey(v string) *FinishTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

func (s *FinishTicketRequestTicketMemoAttachments) Validate() error {
	return dara.Validate(s)
}
