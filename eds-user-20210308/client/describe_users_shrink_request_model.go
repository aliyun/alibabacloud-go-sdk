// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeUsersShrinkRequest
	GetBizType() *string
	SetEndUserIds(v []*string) *DescribeUsersShrinkRequest
	GetEndUserIds() []*string
	SetExcludeEndUserIds(v []*string) *DescribeUsersShrinkRequest
	GetExcludeEndUserIds() []*string
	SetExcludeGroupId(v string) *DescribeUsersShrinkRequest
	GetExcludeGroupId() *string
	SetFilter(v string) *DescribeUsersShrinkRequest
	GetFilter() *string
	SetFilterWithAssignedResourceShrink(v string) *DescribeUsersShrinkRequest
	GetFilterWithAssignedResourceShrink() *string
	SetFilterWithAssignedResourcesShrink(v string) *DescribeUsersShrinkRequest
	GetFilterWithAssignedResourcesShrink() *string
	SetGroupId(v string) *DescribeUsersShrinkRequest
	GetGroupId() *string
	SetIsQueryAllSubOrgs(v bool) *DescribeUsersShrinkRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *DescribeUsersShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeUsersShrinkRequest
	GetNextToken() *string
	SetOrgId(v string) *DescribeUsersShrinkRequest
	GetOrgId() *string
	SetShowExtrasShrink(v string) *DescribeUsersShrinkRequest
	GetShowExtrasShrink() *string
	SetSolutionId(v string) *DescribeUsersShrinkRequest
	GetSolutionId() *string
	SetStatus(v int32) *DescribeUsersShrinkRequest
	GetStatus() *int32
}

type DescribeUsersShrinkRequest struct {
	// example:
	//
	// null
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The usernames that must be exactly matched.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The usernames that must be exactly excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	ExcludeGroupId    *string   `json:"ExcludeGroupId,omitempty" xml:"ExcludeGroupId,omitempty"`
	// The string that is used for fuzzy search. You perform fuzzy search by username (EndUserId) and email address (Email). Wildcard characters (\\*) are supported. For example, if you set this parameter to `a*m`, usernames or email addresses that start with `a` and end with `m` are returned.
	//
	// example:
	//
	// a*m
	Filter                            *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	FilterWithAssignedResourceShrink  *string `json:"FilterWithAssignedResource,omitempty" xml:"FilterWithAssignedResource,omitempty"`
	FilterWithAssignedResourcesShrink *string `json:"FilterWithAssignedResources,omitempty" xml:"FilterWithAssignedResources,omitempty"`
	// The ID of the organization in which you want to query convenience users.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// true
	IsQueryAllSubOrgs *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The maximum number of entries per page.
	//
	// 	- Valid values: 1 to 500.
	//
	// 	- Default value: 500.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.\\
	//
	// If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the return value of NextToken to perform the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the organization in which you want to query users.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId            *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ShowExtrasShrink *string `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
	// example:
	//
	// null
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	// The status.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersShrinkRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeUsersShrinkRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeUsersShrinkRequest) GetExcludeEndUserIds() []*string {
	return s.ExcludeEndUserIds
}

func (s *DescribeUsersShrinkRequest) GetExcludeGroupId() *string {
	return s.ExcludeGroupId
}

func (s *DescribeUsersShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeUsersShrinkRequest) GetFilterWithAssignedResourceShrink() *string {
	return s.FilterWithAssignedResourceShrink
}

func (s *DescribeUsersShrinkRequest) GetFilterWithAssignedResourcesShrink() *string {
	return s.FilterWithAssignedResourcesShrink
}

func (s *DescribeUsersShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUsersShrinkRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *DescribeUsersShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeUsersShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUsersShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeUsersShrinkRequest) GetShowExtrasShrink() *string {
	return s.ShowExtrasShrink
}

func (s *DescribeUsersShrinkRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *DescribeUsersShrinkRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeUsersShrinkRequest) SetBizType(v string) *DescribeUsersShrinkRequest {
	s.BizType = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetEndUserIds(v []*string) *DescribeUsersShrinkRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersShrinkRequest) SetExcludeEndUserIds(v []*string) *DescribeUsersShrinkRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *DescribeUsersShrinkRequest) SetExcludeGroupId(v string) *DescribeUsersShrinkRequest {
	s.ExcludeGroupId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilter(v string) *DescribeUsersShrinkRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilterWithAssignedResourceShrink(v string) *DescribeUsersShrinkRequest {
	s.FilterWithAssignedResourceShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilterWithAssignedResourcesShrink(v string) *DescribeUsersShrinkRequest {
	s.FilterWithAssignedResourcesShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetGroupId(v string) *DescribeUsersShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetIsQueryAllSubOrgs(v bool) *DescribeUsersShrinkRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetMaxResults(v int64) *DescribeUsersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetNextToken(v string) *DescribeUsersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetOrgId(v string) *DescribeUsersShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetShowExtrasShrink(v string) *DescribeUsersShrinkRequest {
	s.ShowExtrasShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetSolutionId(v string) *DescribeUsersShrinkRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetStatus(v int32) *DescribeUsersShrinkRequest {
	s.Status = &v
	return s
}

func (s *DescribeUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
