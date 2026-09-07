// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DescribeOrgsShrinkRequest
	GetBusinessChannel() *string
	SetIncludeOrgIds(v []*string) *DescribeOrgsShrinkRequest
	GetIncludeOrgIds() []*string
	SetIsQueryAllSubOrgs(v bool) *DescribeOrgsShrinkRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *DescribeOrgsShrinkRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeOrgsShrinkRequest
	GetNextToken() *string
	SetOrgName(v string) *DescribeOrgsShrinkRequest
	GetOrgName() *string
	SetParentOrgId(v string) *DescribeOrgsShrinkRequest
	GetParentOrgId() *string
	SetShowExtrasShrink(v string) *DescribeOrgsShrinkRequest
	GetShowExtrasShrink() *string
}

type DescribeOrgsShrinkRequest struct {
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string   `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	IncludeOrgIds   []*string `json:"IncludeOrgIds,omitempty" xml:"IncludeOrgIds,omitempty" type:"Repeated"`
	// Specifies whether to query all subordinate organizations when a parent organization is specified.
	IsQueryAllSubOrgs *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken that was returned in the previous API call.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organization name.
	//
	// example:
	//
	// ProductDepartment
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// org-11fs****
	ParentOrgId      *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
	ShowExtrasShrink *string `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
}

func (s DescribeOrgsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgsShrinkRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeOrgsShrinkRequest) GetIncludeOrgIds() []*string {
	return s.IncludeOrgIds
}

func (s *DescribeOrgsShrinkRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *DescribeOrgsShrinkRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeOrgsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOrgsShrinkRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeOrgsShrinkRequest) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeOrgsShrinkRequest) GetShowExtrasShrink() *string {
	return s.ShowExtrasShrink
}

func (s *DescribeOrgsShrinkRequest) SetBusinessChannel(v string) *DescribeOrgsShrinkRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetIncludeOrgIds(v []*string) *DescribeOrgsShrinkRequest {
	s.IncludeOrgIds = v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetIsQueryAllSubOrgs(v bool) *DescribeOrgsShrinkRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetMaxResults(v int64) *DescribeOrgsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetNextToken(v string) *DescribeOrgsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetOrgName(v string) *DescribeOrgsShrinkRequest {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetParentOrgId(v string) *DescribeOrgsShrinkRequest {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) SetShowExtrasShrink(v string) *DescribeOrgsShrinkRequest {
	s.ShowExtrasShrink = &v
	return s
}

func (s *DescribeOrgsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
