// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantHonorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GrantHonorRequestTenantContext) *GrantHonorRequest
	GetTenantContext() *GrantHonorRequestTenantContext
	SetExpirationTime(v int64) *GrantHonorRequest
	GetExpirationTime() *int64
	SetGrantReason(v string) *GrantHonorRequest
	GetGrantReason() *string
	SetGranterName(v string) *GrantHonorRequest
	GetGranterName() *string
	SetHonorId(v string) *GrantHonorRequest
	GetHonorId() *string
	SetNoticeAnnouncer(v bool) *GrantHonorRequest
	GetNoticeAnnouncer() *bool
	SetNoticeSingle(v bool) *GrantHonorRequest
	GetNoticeSingle() *bool
	SetOpenConversationIds(v []*string) *GrantHonorRequest
	GetOpenConversationIds() []*string
	SetOrgId(v int64) *GrantHonorRequest
	GetOrgId() *int64
	SetReceiverUserIds(v []*string) *GrantHonorRequest
	GetReceiverUserIds() []*string
	SetSenderUserId(v string) *GrantHonorRequest
	GetSenderUserId() *string
}

type GrantHonorRequest struct {
	TenantContext *GrantHonorRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// null
	ExpirationTime *int64 `json:"expirationTime,omitempty" xml:"expirationTime,omitempty"`
	// This parameter is required.
	GrantReason *string `json:"grantReason,omitempty" xml:"grantReason,omitempty"`
	// This parameter is required.
	GranterName *string `json:"granterName,omitempty" xml:"granterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 21659398
	HonorId *string `json:"honorId,omitempty" xml:"honorId,omitempty"`
	// example:
	//
	// false
	NoticeAnnouncer *bool `json:"noticeAnnouncer,omitempty" xml:"noticeAnnouncer,omitempty"`
	// example:
	//
	// false
	NoticeSingle        *bool     `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	OpenConversationIds []*string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 345391052
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	ReceiverUserIds []*string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s GrantHonorRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorRequest) GoString() string {
	return s.String()
}

func (s *GrantHonorRequest) GetTenantContext() *GrantHonorRequestTenantContext {
	return s.TenantContext
}

func (s *GrantHonorRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GrantHonorRequest) GetGrantReason() *string {
	return s.GrantReason
}

func (s *GrantHonorRequest) GetGranterName() *string {
	return s.GranterName
}

func (s *GrantHonorRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *GrantHonorRequest) GetNoticeAnnouncer() *bool {
	return s.NoticeAnnouncer
}

func (s *GrantHonorRequest) GetNoticeSingle() *bool {
	return s.NoticeSingle
}

func (s *GrantHonorRequest) GetOpenConversationIds() []*string {
	return s.OpenConversationIds
}

func (s *GrantHonorRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GrantHonorRequest) GetReceiverUserIds() []*string {
	return s.ReceiverUserIds
}

func (s *GrantHonorRequest) GetSenderUserId() *string {
	return s.SenderUserId
}

func (s *GrantHonorRequest) SetTenantContext(v *GrantHonorRequestTenantContext) *GrantHonorRequest {
	s.TenantContext = v
	return s
}

func (s *GrantHonorRequest) SetExpirationTime(v int64) *GrantHonorRequest {
	s.ExpirationTime = &v
	return s
}

func (s *GrantHonorRequest) SetGrantReason(v string) *GrantHonorRequest {
	s.GrantReason = &v
	return s
}

func (s *GrantHonorRequest) SetGranterName(v string) *GrantHonorRequest {
	s.GranterName = &v
	return s
}

func (s *GrantHonorRequest) SetHonorId(v string) *GrantHonorRequest {
	s.HonorId = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeAnnouncer(v bool) *GrantHonorRequest {
	s.NoticeAnnouncer = &v
	return s
}

func (s *GrantHonorRequest) SetNoticeSingle(v bool) *GrantHonorRequest {
	s.NoticeSingle = &v
	return s
}

func (s *GrantHonorRequest) SetOpenConversationIds(v []*string) *GrantHonorRequest {
	s.OpenConversationIds = v
	return s
}

func (s *GrantHonorRequest) SetOrgId(v int64) *GrantHonorRequest {
	s.OrgId = &v
	return s
}

func (s *GrantHonorRequest) SetReceiverUserIds(v []*string) *GrantHonorRequest {
	s.ReceiverUserIds = v
	return s
}

func (s *GrantHonorRequest) SetSenderUserId(v string) *GrantHonorRequest {
	s.SenderUserId = &v
	return s
}

func (s *GrantHonorRequest) Validate() error {
	return dara.Validate(s)
}

type GrantHonorRequestTenantContext struct {
	// example:
	//
	// null
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GrantHonorRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GrantHonorRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GrantHonorRequestTenantContext) SetTenantId(v string) *GrantHonorRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GrantHonorRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
