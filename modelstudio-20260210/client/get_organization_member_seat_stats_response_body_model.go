// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationMemberSeatStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetAdminRoleUserCount() *int32
	SetMemberRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetMemberRoleUserCount() *int32
	SetOrgId(v string) *GetOrganizationMemberSeatStatsResponseBody
	GetOrgId() *string
	SetOwnerRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetOwnerRoleUserCount() *int32
	SetSeatedMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetSeatedMemberCount() *int32
	SetTotalMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetTotalMemberCount() *int32
	SetUnseatedMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody
	GetUnseatedMemberCount() *int32
}

type GetOrganizationMemberSeatStatsResponseBody struct {
	// The number of administrators (ORG_ADMIN role).
	//
	// example:
	//
	// 3
	AdminRoleUserCount *int32 `json:"AdminRoleUserCount,omitempty" xml:"AdminRoleUserCount,omitempty"`
	// The number of regular members (ORG_MEMBER role).
	//
	// example:
	//
	// 8
	MemberRoleUserCount *int32 `json:"MemberRoleUserCount,omitempty" xml:"MemberRoleUserCount,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// 5ffd468b1e45db3c1cc26ad6
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The number of owner accounts (ORG_OWNER role).
	//
	// example:
	//
	// 1
	OwnerRoleUserCount *int32 `json:"OwnerRoleUserCount,omitempty" xml:"OwnerRoleUserCount,omitempty"`
	// The number of members with assigned seats.
	//
	// example:
	//
	// 2
	SeatedMemberCount *int32 `json:"SeatedMemberCount,omitempty" xml:"SeatedMemberCount,omitempty"`
	// The total number of members.
	//
	// example:
	//
	// 12
	TotalMemberCount *int32 `json:"TotalMemberCount,omitempty" xml:"TotalMemberCount,omitempty"`
	// The number of members without assigned seats.
	//
	// example:
	//
	// 10
	UnseatedMemberCount *int32 `json:"UnseatedMemberCount,omitempty" xml:"UnseatedMemberCount,omitempty"`
}

func (s GetOrganizationMemberSeatStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberSeatStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetAdminRoleUserCount() *int32 {
	return s.AdminRoleUserCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetMemberRoleUserCount() *int32 {
	return s.MemberRoleUserCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetOrgId() *string {
	return s.OrgId
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetOwnerRoleUserCount() *int32 {
	return s.OwnerRoleUserCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetSeatedMemberCount() *int32 {
	return s.SeatedMemberCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetTotalMemberCount() *int32 {
	return s.TotalMemberCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) GetUnseatedMemberCount() *int32 {
	return s.UnseatedMemberCount
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetAdminRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.AdminRoleUserCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetMemberRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.MemberRoleUserCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetOrgId(v string) *GetOrganizationMemberSeatStatsResponseBody {
	s.OrgId = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetOwnerRoleUserCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.OwnerRoleUserCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetSeatedMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.SeatedMemberCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetTotalMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.TotalMemberCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) SetUnseatedMemberCount(v int32) *GetOrganizationMemberSeatStatsResponseBody {
	s.UnseatedMemberCount = &v
	return s
}

func (s *GetOrganizationMemberSeatStatsResponseBody) Validate() error {
	return dara.Validate(s)
}
