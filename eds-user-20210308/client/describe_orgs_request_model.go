// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrgsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *DescribeOrgsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeOrgsRequest
	GetNextToken() *string
	SetOrgName(v string) *DescribeOrgsRequest
	GetOrgName() *string
	SetParentOrgId(v string) *DescribeOrgsRequest
	GetParentOrgId() *string
}

type DescribeOrgsRequest struct {
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
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrgsRequest) GoString() string {
	return s.String()
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

func (s *DescribeOrgsRequest) Validate() error {
	return dara.Validate(s)
}
