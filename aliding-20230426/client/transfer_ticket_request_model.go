// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNotify(v *TransferTicketRequestNotify) *TransferTicketRequest
	GetNotify() *TransferTicketRequestNotify
	SetOpenTeamId(v string) *TransferTicketRequest
	GetOpenTeamId() *string
	SetOpenTicketId(v string) *TransferTicketRequest
	GetOpenTicketId() *string
	SetProcessorUserIds(v []*string) *TransferTicketRequest
	GetProcessorUserIds() []*string
	SetTenantContext(v *TransferTicketRequestTenantContext) *TransferTicketRequest
	GetTenantContext() *TransferTicketRequestTenantContext
	SetTicketMemo(v *TransferTicketRequestTicketMemo) *TransferTicketRequest
	GetTicketMemo() *TransferTicketRequestTicketMemo
}

type TransferTicketRequest struct {
	Notify *TransferTicketRequestNotify `json:"Notify,omitempty" xml:"Notify,omitempty" type:"Struct"`
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
	OpenTicketId     *string                             `json:"OpenTicketId,omitempty" xml:"OpenTicketId,omitempty"`
	ProcessorUserIds []*string                           `json:"ProcessorUserIds,omitempty" xml:"ProcessorUserIds,omitempty" type:"Repeated"`
	TenantContext    *TransferTicketRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	TicketMemo       *TransferTicketRequestTicketMemo    `json:"TicketMemo,omitempty" xml:"TicketMemo,omitempty" type:"Struct"`
}

func (s TransferTicketRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketRequest) GoString() string {
	return s.String()
}

func (s *TransferTicketRequest) GetNotify() *TransferTicketRequestNotify {
	return s.Notify
}

func (s *TransferTicketRequest) GetOpenTeamId() *string {
	return s.OpenTeamId
}

func (s *TransferTicketRequest) GetOpenTicketId() *string {
	return s.OpenTicketId
}

func (s *TransferTicketRequest) GetProcessorUserIds() []*string {
	return s.ProcessorUserIds
}

func (s *TransferTicketRequest) GetTenantContext() *TransferTicketRequestTenantContext {
	return s.TenantContext
}

func (s *TransferTicketRequest) GetTicketMemo() *TransferTicketRequestTicketMemo {
	return s.TicketMemo
}

func (s *TransferTicketRequest) SetNotify(v *TransferTicketRequestNotify) *TransferTicketRequest {
	s.Notify = v
	return s
}

func (s *TransferTicketRequest) SetOpenTeamId(v string) *TransferTicketRequest {
	s.OpenTeamId = &v
	return s
}

func (s *TransferTicketRequest) SetOpenTicketId(v string) *TransferTicketRequest {
	s.OpenTicketId = &v
	return s
}

func (s *TransferTicketRequest) SetProcessorUserIds(v []*string) *TransferTicketRequest {
	s.ProcessorUserIds = v
	return s
}

func (s *TransferTicketRequest) SetTenantContext(v *TransferTicketRequestTenantContext) *TransferTicketRequest {
	s.TenantContext = v
	return s
}

func (s *TransferTicketRequest) SetTicketMemo(v *TransferTicketRequestTicketMemo) *TransferTicketRequest {
	s.TicketMemo = v
	return s
}

func (s *TransferTicketRequest) Validate() error {
	if s.Notify != nil {
		if err := s.Notify.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	if s.TicketMemo != nil {
		if err := s.TicketMemo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TransferTicketRequestNotify struct {
	GroupNoticeReceiverUserIds []*string `json:"GroupNoticeReceiverUserIds,omitempty" xml:"GroupNoticeReceiverUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// true
	NoticeAllGroupMember      *bool     `json:"NoticeAllGroupMember,omitempty" xml:"NoticeAllGroupMember,omitempty"`
	WorkNoticeReceiverUserIds []*string `json:"WorkNoticeReceiverUserIds,omitempty" xml:"WorkNoticeReceiverUserIds,omitempty" type:"Repeated"`
}

func (s TransferTicketRequestNotify) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketRequestNotify) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestNotify) GetGroupNoticeReceiverUserIds() []*string {
	return s.GroupNoticeReceiverUserIds
}

func (s *TransferTicketRequestNotify) GetNoticeAllGroupMember() *bool {
	return s.NoticeAllGroupMember
}

func (s *TransferTicketRequestNotify) GetWorkNoticeReceiverUserIds() []*string {
	return s.WorkNoticeReceiverUserIds
}

func (s *TransferTicketRequestNotify) SetGroupNoticeReceiverUserIds(v []*string) *TransferTicketRequestNotify {
	s.GroupNoticeReceiverUserIds = v
	return s
}

func (s *TransferTicketRequestNotify) SetNoticeAllGroupMember(v bool) *TransferTicketRequestNotify {
	s.NoticeAllGroupMember = &v
	return s
}

func (s *TransferTicketRequestNotify) SetWorkNoticeReceiverUserIds(v []*string) *TransferTicketRequestNotify {
	s.WorkNoticeReceiverUserIds = v
	return s
}

func (s *TransferTicketRequestNotify) Validate() error {
	return dara.Validate(s)
}

type TransferTicketRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TransferTicketRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketRequestTenantContext) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *TransferTicketRequestTenantContext) SetTenantId(v string) *TransferTicketRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *TransferTicketRequestTenantContext) Validate() error {
	return dara.Validate(s)
}

type TransferTicketRequestTicketMemo struct {
	Attachments []*TransferTicketRequestTicketMemoAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Repeated"`
	Memo        *string                                       `json:"Memo,omitempty" xml:"Memo,omitempty"`
}

func (s TransferTicketRequestTicketMemo) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketRequestTicketMemo) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemo) GetAttachments() []*TransferTicketRequestTicketMemoAttachments {
	return s.Attachments
}

func (s *TransferTicketRequestTicketMemo) GetMemo() *string {
	return s.Memo
}

func (s *TransferTicketRequestTicketMemo) SetAttachments(v []*TransferTicketRequestTicketMemoAttachments) *TransferTicketRequestTicketMemo {
	s.Attachments = v
	return s
}

func (s *TransferTicketRequestTicketMemo) SetMemo(v string) *TransferTicketRequestTicketMemo {
	s.Memo = &v
	return s
}

func (s *TransferTicketRequestTicketMemo) Validate() error {
	if s.Attachments != nil {
		for _, item := range s.Attachments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TransferTicketRequestTicketMemoAttachments struct {
	// example:
	//
	// auto-test-1727143229007.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s TransferTicketRequestTicketMemoAttachments) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketRequestTicketMemoAttachments) GoString() string {
	return s.String()
}

func (s *TransferTicketRequestTicketMemoAttachments) GetFileName() *string {
	return s.FileName
}

func (s *TransferTicketRequestTicketMemoAttachments) GetKey() *string {
	return s.Key
}

func (s *TransferTicketRequestTicketMemoAttachments) SetFileName(v string) *TransferTicketRequestTicketMemoAttachments {
	s.FileName = &v
	return s
}

func (s *TransferTicketRequestTicketMemoAttachments) SetKey(v string) *TransferTicketRequestTicketMemoAttachments {
	s.Key = &v
	return s
}

func (s *TransferTicketRequestTicketMemoAttachments) Validate() error {
	return dara.Validate(s)
}
