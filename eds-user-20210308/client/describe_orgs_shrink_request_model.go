// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// The maximum number of entries to return. Valid values: 1 to 100.\\
	//
	// Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. The return value is the value of the NextToken response parameter that was returned last time the DescribeOrgs operation was called.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// org****
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// org-****
	ParentOrgId      *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
	ShowExtrasShrink *string `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
}

func (s DescribeOrgsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsShrinkRequest) GoString() string {
	return s.String()
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
