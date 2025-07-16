// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantHonorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GrantHonorShrinkRequest
	GetTenantContextShrink() *string
	SetExpirationTime(v int64) *GrantHonorShrinkRequest
	GetExpirationTime() *int64
	SetGrantReason(v string) *GrantHonorShrinkRequest
	GetGrantReason() *string
	SetGranterName(v string) *GrantHonorShrinkRequest
	GetGranterName() *string
	SetHonorId(v string) *GrantHonorShrinkRequest
	GetHonorId() *string
	SetNoticeAnnouncer(v bool) *GrantHonorShrinkRequest
	GetNoticeAnnouncer() *bool
	SetNoticeSingle(v bool) *GrantHonorShrinkRequest
	GetNoticeSingle() *bool
	SetOpenConversationIdsShrink(v string) *GrantHonorShrinkRequest
	GetOpenConversationIdsShrink() *string
	SetOrgId(v int64) *GrantHonorShrinkRequest
	GetOrgId() *int64
	SetReceiverUserIdsShrink(v string) *GrantHonorShrinkRequest
	GetReceiverUserIdsShrink() *string
	SetSenderUserId(v string) *GrantHonorShrinkRequest
	GetSenderUserId() *string
}

type GrantHonorShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
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
	NoticeSingle              *bool   `json:"noticeSingle,omitempty" xml:"noticeSingle,omitempty"`
	OpenConversationIdsShrink *string `json:"openConversationIds,omitempty" xml:"openConversationIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 345391052
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	ReceiverUserIdsShrink *string `json:"receiverUserIds,omitempty" xml:"receiverUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 363784
	SenderUserId *string `json:"senderUserId,omitempty" xml:"senderUserId,omitempty"`
}

func (s GrantHonorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorShrinkRequest) GoString() string {
	return s.String()
}

func (s *GrantHonorShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GrantHonorShrinkRequest) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GrantHonorShrinkRequest) GetGrantReason() *string {
	return s.GrantReason
}

func (s *GrantHonorShrinkRequest) GetGranterName() *string {
	return s.GranterName
}

func (s *GrantHonorShrinkRequest) GetHonorId() *string {
	return s.HonorId
}

func (s *GrantHonorShrinkRequest) GetNoticeAnnouncer() *bool {
	return s.NoticeAnnouncer
}

func (s *GrantHonorShrinkRequest) GetNoticeSingle() *bool {
	return s.NoticeSingle
}

func (s *GrantHonorShrinkRequest) GetOpenConversationIdsShrink() *string {
	return s.OpenConversationIdsShrink
}

func (s *GrantHonorShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *GrantHonorShrinkRequest) GetReceiverUserIdsShrink() *string {
	return s.ReceiverUserIdsShrink
}

func (s *GrantHonorShrinkRequest) GetSenderUserId() *string {
	return s.SenderUserId
}

func (s *GrantHonorShrinkRequest) SetTenantContextShrink(v string) *GrantHonorShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetExpirationTime(v int64) *GrantHonorShrinkRequest {
	s.ExpirationTime = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetGrantReason(v string) *GrantHonorShrinkRequest {
	s.GrantReason = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetGranterName(v string) *GrantHonorShrinkRequest {
	s.GranterName = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetHonorId(v string) *GrantHonorShrinkRequest {
	s.HonorId = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetNoticeAnnouncer(v bool) *GrantHonorShrinkRequest {
	s.NoticeAnnouncer = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetNoticeSingle(v bool) *GrantHonorShrinkRequest {
	s.NoticeSingle = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetOpenConversationIdsShrink(v string) *GrantHonorShrinkRequest {
	s.OpenConversationIdsShrink = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetOrgId(v int64) *GrantHonorShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetReceiverUserIdsShrink(v string) *GrantHonorShrinkRequest {
	s.ReceiverUserIdsShrink = &v
	return s
}

func (s *GrantHonorShrinkRequest) SetSenderUserId(v string) *GrantHonorShrinkRequest {
	s.SenderUserId = &v
	return s
}

func (s *GrantHonorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
