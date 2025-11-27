// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeUsersRequest
	GetBizType() *string
	SetEndUserIds(v []*string) *DescribeUsersRequest
	GetEndUserIds() []*string
	SetExcludeEndUserIds(v []*string) *DescribeUsersRequest
	GetExcludeEndUserIds() []*string
	SetExcludeGroupId(v string) *DescribeUsersRequest
	GetExcludeGroupId() *string
	SetFilter(v string) *DescribeUsersRequest
	GetFilter() *string
	SetFilterWithAssignedResource(v map[string]*string) *DescribeUsersRequest
	GetFilterWithAssignedResource() map[string]*string
	SetFilterWithAssignedResources(v map[string]*bool) *DescribeUsersRequest
	GetFilterWithAssignedResources() map[string]*bool
	SetGroupId(v string) *DescribeUsersRequest
	GetGroupId() *string
	SetIsQueryAllSubOrgs(v bool) *DescribeUsersRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *DescribeUsersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeUsersRequest
	GetNextToken() *string
	SetOrgId(v string) *DescribeUsersRequest
	GetOrgId() *string
	SetShowExtras(v map[string]interface{}) *DescribeUsersRequest
	GetShowExtras() map[string]interface{}
	SetSolutionId(v string) *DescribeUsersRequest
	GetSolutionId() *string
	SetStatus(v int32) *DescribeUsersRequest
	GetStatus() *int32
}

type DescribeUsersRequest struct {
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
	Filter                      *string            `json:"Filter,omitempty" xml:"Filter,omitempty"`
	FilterWithAssignedResource  map[string]*string `json:"FilterWithAssignedResource,omitempty" xml:"FilterWithAssignedResource,omitempty"`
	FilterWithAssignedResources map[string]*bool   `json:"FilterWithAssignedResources,omitempty" xml:"FilterWithAssignedResources,omitempty"`
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
	OrgId      *string                `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ShowExtras map[string]interface{} `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
	// example:
	//
	// null
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	// The status.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeUsersRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeUsersRequest) GetExcludeEndUserIds() []*string {
	return s.ExcludeEndUserIds
}

func (s *DescribeUsersRequest) GetExcludeGroupId() *string {
	return s.ExcludeGroupId
}

func (s *DescribeUsersRequest) GetFilter() *string {
	return s.Filter
}

func (s *DescribeUsersRequest) GetFilterWithAssignedResource() map[string]*string {
	return s.FilterWithAssignedResource
}

func (s *DescribeUsersRequest) GetFilterWithAssignedResources() map[string]*bool {
	return s.FilterWithAssignedResources
}

func (s *DescribeUsersRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeUsersRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *DescribeUsersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUsersRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *DescribeUsersRequest) GetShowExtras() map[string]interface{} {
	return s.ShowExtras
}

func (s *DescribeUsersRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *DescribeUsersRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeUsersRequest) SetBizType(v string) *DescribeUsersRequest {
	s.BizType = &v
	return s
}

func (s *DescribeUsersRequest) SetEndUserIds(v []*string) *DescribeUsersRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetExcludeEndUserIds(v []*string) *DescribeUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetExcludeGroupId(v string) *DescribeUsersRequest {
	s.ExcludeGroupId = &v
	return s
}

func (s *DescribeUsersRequest) SetFilter(v string) *DescribeUsersRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersRequest) SetFilterWithAssignedResource(v map[string]*string) *DescribeUsersRequest {
	s.FilterWithAssignedResource = v
	return s
}

func (s *DescribeUsersRequest) SetFilterWithAssignedResources(v map[string]*bool) *DescribeUsersRequest {
	s.FilterWithAssignedResources = v
	return s
}

func (s *DescribeUsersRequest) SetGroupId(v string) *DescribeUsersRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersRequest) SetIsQueryAllSubOrgs(v bool) *DescribeUsersRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeUsersRequest) SetMaxResults(v int64) *DescribeUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersRequest) SetNextToken(v string) *DescribeUsersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersRequest) SetOrgId(v string) *DescribeUsersRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersRequest) SetShowExtras(v map[string]interface{}) *DescribeUsersRequest {
	s.ShowExtras = v
	return s
}

func (s *DescribeUsersRequest) SetSolutionId(v string) *DescribeUsersRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeUsersRequest) SetStatus(v int32) *DescribeUsersRequest {
	s.Status = &v
	return s
}

func (s *DescribeUsersRequest) Validate() error {
	return dara.Validate(s)
}
