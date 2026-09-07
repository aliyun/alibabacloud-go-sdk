// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessChannel(v string) *DescribeOrgsRequest
	GetBusinessChannel() *string
	SetIncludeOrgIds(v []*string) *DescribeOrgsRequest
	GetIncludeOrgIds() []*string
	SetIsQueryAllSubOrgs(v bool) *DescribeOrgsRequest
	GetIsQueryAllSubOrgs() *bool
	SetMaxResults(v int64) *DescribeOrgsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeOrgsRequest
	GetNextToken() *string
	SetOrgName(v string) *DescribeOrgsRequest
	GetOrgName() *string
	SetParentOrgId(v string) *DescribeOrgsRequest
	GetParentOrgId() *string
	SetShowExtras(v map[string]interface{}) *DescribeOrgsRequest
	GetShowExtras() map[string]interface{}
}

type DescribeOrgsRequest struct {
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
	ParentOrgId *string                `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
	ShowExtras  map[string]interface{} `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
}

func (s DescribeOrgsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgsRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *DescribeOrgsRequest) GetIncludeOrgIds() []*string {
	return s.IncludeOrgIds
}

func (s *DescribeOrgsRequest) GetIsQueryAllSubOrgs() *bool {
	return s.IsQueryAllSubOrgs
}

func (s *DescribeOrgsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeOrgsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOrgsRequest) GetOrgName() *string {
	return s.OrgName
}

func (s *DescribeOrgsRequest) GetParentOrgId() *string {
	return s.ParentOrgId
}

func (s *DescribeOrgsRequest) GetShowExtras() map[string]interface{} {
	return s.ShowExtras
}

func (s *DescribeOrgsRequest) SetBusinessChannel(v string) *DescribeOrgsRequest {
	s.BusinessChannel = &v
	return s
}

func (s *DescribeOrgsRequest) SetIncludeOrgIds(v []*string) *DescribeOrgsRequest {
	s.IncludeOrgIds = v
	return s
}

func (s *DescribeOrgsRequest) SetIsQueryAllSubOrgs(v bool) *DescribeOrgsRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeOrgsRequest) SetMaxResults(v int64) *DescribeOrgsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOrgsRequest) SetNextToken(v string) *DescribeOrgsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOrgsRequest) SetOrgName(v string) *DescribeOrgsRequest {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgsRequest) SetParentOrgId(v string) *DescribeOrgsRequest {
	s.ParentOrgId = &v
	return s
}

func (s *DescribeOrgsRequest) SetShowExtras(v map[string]interface{}) *DescribeOrgsRequest {
	s.ShowExtras = v
	return s
}

func (s *DescribeOrgsRequest) Validate() error {
	return dara.Validate(s)
}
