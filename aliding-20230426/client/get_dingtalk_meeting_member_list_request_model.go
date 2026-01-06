// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkMeetingMemberListRequestTenantContext) *GetDingtalkMeetingMemberListRequest
	GetTenantContext() *GetDingtalkMeetingMemberListRequestTenantContext
	SetClusterName(v string) *GetDingtalkMeetingMemberListRequest
	GetClusterName() *string
	SetConferenceId(v string) *GetDingtalkMeetingMemberListRequest
	GetConferenceId() *string
	SetCurrentPage(v int32) *GetDingtalkMeetingMemberListRequest
	GetCurrentPage() *int32
	SetOrgId(v string) *GetDingtalkMeetingMemberListRequest
	GetOrgId() *string
	SetPageSize(v int32) *GetDingtalkMeetingMemberListRequest
	GetPageSize() *int32
}

type GetDingtalkMeetingMemberListRequest struct {
	TenantContext *GetDingtalkMeetingMemberListRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s GetDingtalkMeetingMemberListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListRequest) GetTenantContext() *GetDingtalkMeetingMemberListRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkMeetingMemberListRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetDingtalkMeetingMemberListRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMemberListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDingtalkMeetingMemberListRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMemberListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDingtalkMeetingMemberListRequest) SetTenantContext(v *GetDingtalkMeetingMemberListRequestTenantContext) *GetDingtalkMeetingMemberListRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) SetClusterName(v string) *GetDingtalkMeetingMemberListRequest {
	s.ClusterName = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) SetConferenceId(v string) *GetDingtalkMeetingMemberListRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) SetCurrentPage(v int32) *GetDingtalkMeetingMemberListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) SetOrgId(v string) *GetDingtalkMeetingMemberListRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) SetPageSize(v int32) *GetDingtalkMeetingMemberListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingMemberListRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkMeetingMemberListRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkMeetingMemberListRequestTenantContext) SetTenantId(v string) *GetDingtalkMeetingMemberListRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
