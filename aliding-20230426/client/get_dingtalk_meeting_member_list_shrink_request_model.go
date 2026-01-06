// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkMeetingMemberListShrinkRequest
	GetTenantContextShrink() *string
	SetClusterName(v string) *GetDingtalkMeetingMemberListShrinkRequest
	GetClusterName() *string
	SetConferenceId(v string) *GetDingtalkMeetingMemberListShrinkRequest
	GetConferenceId() *string
	SetCurrentPage(v int32) *GetDingtalkMeetingMemberListShrinkRequest
	GetCurrentPage() *int32
	SetOrgId(v string) *GetDingtalkMeetingMemberListShrinkRequest
	GetOrgId() *string
	SetPageSize(v int32) *GetDingtalkMeetingMemberListShrinkRequest
	GetPageSize() *int32
}

type GetDingtalkMeetingMemberListShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// cluster-1
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s GetDingtalkMeetingMemberListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkMeetingMemberListShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetClusterName(v string) *GetDingtalkMeetingMemberListShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetConferenceId(v string) *GetDingtalkMeetingMemberListShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetCurrentPage(v int32) *GetDingtalkMeetingMemberListShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetOrgId(v string) *GetDingtalkMeetingMemberListShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) SetPageSize(v int32) *GetDingtalkMeetingMemberListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkMeetingMemberListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
