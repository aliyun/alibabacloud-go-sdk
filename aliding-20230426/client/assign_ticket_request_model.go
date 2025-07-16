// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotify(v *AssignTicketRequestNotify) *AssignTicketRequest
	GetNotify() *AssignTicketRequestNotify
	SetOpenTeamId(v string) *AssignTicketRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *AssignTicketRequest
	GetOpenTicketId() *string
	SetProcessorUserIds(v []*string) *AssignTicketRequest
	GetProcessorUserIds() []*string
	SetTenantContext(v *AssignTicketRequestTenantContext) *AssignTicketRequest
	GetTenantContext() *AssignTicketRequestTenantContext
	SetTicketMemo(v *AssignTicketRequestTicketMemo) *AssignTicketRequest
	GetTicketMemo() *AssignTicketRequestTicketMemo
}

type AssignTicketRequest struct {
	Notify *AssignTicketRequestNotify `json:"Notify,omitempty" xml:"Notify,omitempty" type:"Struct"`
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
	OpenTicketId     *string                           `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	ProcessorUserIds []*string                         `json:"ProcessorUserIds,omitempty" xml:"ProcessorUserIds,omitempty" type:"Repeated"`
	TenantContext    *AssignTicketRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	TicketMemo       *AssignTicketRequestTicketMemo    `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty" type:"Struct"`
}

func (s AssignTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketRequest) GoString() string {
	return s.String()
}

func (s *AssignTicketRequest) GetNotify() *AssignTicketRequestNotify {
	return s.Notify
}

func (s *AssignTicketRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *AssignTicketRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *AssignTicketRequest) GetProcessorUserIds() []*string {
	return s.ProcessorUserIds
}

func (s *AssignTicketRequest) GetTenantContext() *AssignTicketRequestTenantContext {
	return s.TenantContext
}

func (s *AssignTicketRequest) GetTicketMemo() *AssignTicketRequestTicketMemo {
	return s.TicketMemo
}

func (s *AssignTicketRequest) SetNotify(v *AssignTicketRequestNotify) *AssignTicketRequest {
	s.Notify = v
	return s
}

func (s *AssignTicketRequest) SetOpenTeamId(v string) *AssignTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *AssignTicketRequest) SetOpenTicketId(v string) *AssignTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *AssignTicketRequest) SetProcessorUserIds(v []*string) *AssignTicketRequest {
	s.ProcessorUserIds = v
	return s
}

func (s *AssignTicketRequest) SetTenantContext(v *AssignTicketRequestTenantContext) *AssignTicketRequest {
	s.TenantContext = v
	return s
}

func (s *AssignTicketRequest) SetTicketMemo(v *AssignTicketRequestTicketMemo) *AssignTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *AssignTicketRequest) Validate() error {
	return dara.Validate(s)
}

type AssignTicketRequestNotify struct {
	GroupNoticeReceiverUserIds []*string `json:"GroupNoticeReceiverUserIds,omitempty" xml:"GroupNoticeReceiverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember      *bool     `json:"NoticeAllGroupMember,omitempty" xml:"NoticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUserIds []*string `json:"WorkNoticeReceiverUserIds,omitempty" xml:"WorkNoticeReceiverUserIds,omitempty" type:"Repeated"`
}

func (s AssignTicketRequestNotify) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestNotify) GetGroupNoticeReceiverUserIds() []*string {
	return s.GroupNoticeReceiverUserIds
}

func (s *AssignTicketRequestNotify) GetNoticeAllGroupMember() *bool {
	return s.NoticeAllGroupMember
}

func (s *AssignTicketRequestNotify) GetWorkNoticeReceiverUserIds() []*string {
	return s.WorkNoticeReceiverUserIds
}

func (s *AssignTicketRequestNotify) SetGroupNoticeReceiverUserIds(v []*string) *AssignTicketRequestNotify {
	s.GroupNoticeReceiverUserIds = v
	return s
}

func (s *AssignTicketRequestNotify) SetNoticeAllGroupMember(v bool) *AssignTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *AssignTicketRequestNotify) SetWorkNoticeReceiverUserIds(v []*string) *AssignTicketRequestNotify {
	s.WorkNoticeReceiverUserIds = v
	return s
}

func (s *AssignTicketRequestNotify) Validate() error {
	return dara.Validate(s)
}

type AssignTicketRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AssignTicketRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AssignTicketRequestTenantContext) SetTenantId(v string) *AssignTicketRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AssignTicketRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type AssignTicketRequestTicketMemo struct {
	Attachments []*AssignTicketRequestTicketMemoAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	Memo        *string                                     `json:"Memo,omitempty" xml:"Memo,omitempty"`
}

func (s AssignTicketRequestTicketMemo) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemo) GetAttachments() []*AssignTicketRequestTicketMemoAttachments {
	return s.Attachments
}

func (s *AssignTicketRequestTicketMemo) GetMemo() *string {
	return s.Memo
}

func (s *AssignTicketRequestTicketMemo) SetAttachments(v []*AssignTicketRequestTicketMemoAttachments) *AssignTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *AssignTicketRequestTicketMemo) SetMemo(v string) *AssignTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *AssignTicketRequestTicketMemo) Validate() error {
	return dara.Validate(s)
}

type AssignTicketRequestTicketMemoAttachments struct {
	// example:
	//
	// ticket/image/44708069/43003/e27aec4499.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// wahaha.txt
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s AssignTicketRequestTicketMemoAttachments) String() string {
	return dara.Prettify(s)
}

func (s AssignTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *AssignTicketRequestTicketMemoAttachments) GetFileName() *string {
	return s.FileName
}

func (s *AssignTicketRequestTicketMemoAttachments) GetKey() *string {
	return s.Key
}

func (s *AssignTicketRequestTicketMemoAttachments) SetFileName(v string) *AssignTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *AssignTicketRequestTicketMemoAttachments) SetKey(v string) *AssignTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

func (s *AssignTicketRequestTicketMemoAttachments) Validate() error {
	return dara.Validate(s)
}
