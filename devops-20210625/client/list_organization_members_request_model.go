// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrganizationMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainsExternInfo(v bool) *ListOrganizationMembersRequest
	GetContainsExternInfo() *bool
	SetExternUid(v string) *ListOrganizationMembersRequest
	GetExternUid() *string
	SetJoinTimeFrom(v int64) *ListOrganizationMembersRequest
	GetJoinTimeFrom() *int64
	SetJoinTimeTo(v int64) *ListOrganizationMembersRequest
	GetJoinTimeTo() *int64
	SetMaxResults(v int64) *ListOrganizationMembersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListOrganizationMembersRequest
	GetNextToken() *string
	SetOrganizationMemberName(v string) *ListOrganizationMembersRequest
	GetOrganizationMemberName() *string
	SetProvider(v string) *ListOrganizationMembersRequest
	GetProvider() *string
	SetState(v string) *ListOrganizationMembersRequest
	GetState() *string
}

type ListOrganizationMembersRequest struct {
	// 返回信息中是否包含第三方信息，默认不包含。
	ContainsExternInfo *bool `json:"containsExternInfo,omitempty" xml:"containsExternInfo,omitempty"`
	// example:
	//
	// 1236666
	ExternUid *string `json:"externUid,omitempty" xml:"externUid,omitempty"`
	// example:
	//
	// 1631845101798
	JoinTimeFrom *int64 `json:"joinTimeFrom,omitempty" xml:"joinTimeFrom,omitempty"`
	// example:
	//
	// 1631845101798
	JoinTimeTo *int64 `json:"joinTimeTo,omitempty" xml:"joinTimeTo,omitempty"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// ""
	NextToken              *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	OrganizationMemberName *string `json:"organizationMemberName,omitempty" xml:"organizationMemberName,omitempty"`
	// example:
	//
	// Dingtalk
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// example:
	//
	// normal
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s ListOrganizationMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOrganizationMembersRequest) GoString() string {
	return s.String()
}

func (s *ListOrganizationMembersRequest) GetContainsExternInfo() *bool {
	return s.ContainsExternInfo
}

func (s *ListOrganizationMembersRequest) GetExternUid() *string {
	return s.ExternUid
}

func (s *ListOrganizationMembersRequest) GetJoinTimeFrom() *int64 {
	return s.JoinTimeFrom
}

func (s *ListOrganizationMembersRequest) GetJoinTimeTo() *int64 {
	return s.JoinTimeTo
}

func (s *ListOrganizationMembersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListOrganizationMembersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrganizationMembersRequest) GetOrganizationMemberName() *string {
	return s.OrganizationMemberName
}

func (s *ListOrganizationMembersRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListOrganizationMembersRequest) GetState() *string {
	return s.State
}

func (s *ListOrganizationMembersRequest) SetContainsExternInfo(v bool) *ListOrganizationMembersRequest {
	s.ContainsExternInfo = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetExternUid(v string) *ListOrganizationMembersRequest {
	s.ExternUid = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetJoinTimeFrom(v int64) *ListOrganizationMembersRequest {
	s.JoinTimeFrom = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetJoinTimeTo(v int64) *ListOrganizationMembersRequest {
	s.JoinTimeTo = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetMaxResults(v int64) *ListOrganizationMembersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetNextToken(v string) *ListOrganizationMembersRequest {
	s.NextToken = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetOrganizationMemberName(v string) *ListOrganizationMembersRequest {
	s.OrganizationMemberName = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetProvider(v string) *ListOrganizationMembersRequest {
	s.Provider = &v
	return s
}

func (s *ListOrganizationMembersRequest) SetState(v string) *ListOrganizationMembersRequest {
	s.State = &v
	return s
}

func (s *ListOrganizationMembersRequest) Validate() error {
	return dara.Validate(s)
}
