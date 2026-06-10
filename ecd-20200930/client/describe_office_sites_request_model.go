// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSitesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountType(v string) *DescribeOfficeSitesRequest
	GetAccountType() *string
	SetMaxResults(v int32) *DescribeOfficeSitesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeOfficeSitesRequest
	GetNextToken() *string
	SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest
	GetOfficeSiteId() []*string
	SetOfficeSiteType(v string) *DescribeOfficeSitesRequest
	GetOfficeSiteType() *string
	SetRegionId(v string) *DescribeOfficeSitesRequest
	GetRegionId() *string
	SetSecurityProtection(v string) *DescribeOfficeSitesRequest
	GetSecurityProtection() *string
	SetStatus(v string) *DescribeOfficeSitesRequest
	GetStatus() *string
	SetVpcId(v string) *DescribeOfficeSitesRequest
	GetVpcId() *string
}

type DescribeOfficeSitesRequest struct {
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	// The number of entries to return on each page.
	//
	// - Maximum value: 100.
	//
	// - Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The office network IDs. You can specify 1 to 100 office network IDs.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId []*string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty" type:"Repeated"`
	// The type of account for the office network.
	//
	// example:
	//
	// SIMPLE
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security protection setting for the office network.
	//
	// example:
	//
	// SASE
	SecurityProtection *string `json:"SecurityProtection,omitempty" xml:"SecurityProtection,omitempty"`
	// The office network status.
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC).
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeOfficeSitesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSitesRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSitesRequest) GetAccountType() *string {
	return s.AccountType
}

func (s *DescribeOfficeSitesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeOfficeSitesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeOfficeSitesRequest) GetOfficeSiteId() []*string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSitesRequest) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *DescribeOfficeSitesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOfficeSitesRequest) GetSecurityProtection() *string {
	return s.SecurityProtection
}

func (s *DescribeOfficeSitesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeOfficeSitesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeOfficeSitesRequest) SetAccountType(v string) *DescribeOfficeSitesRequest {
	s.AccountType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetMaxResults(v int32) *DescribeOfficeSitesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetNextToken(v string) *DescribeOfficeSitesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteId(v []*string) *DescribeOfficeSitesRequest {
	s.OfficeSiteId = v
	return s
}

func (s *DescribeOfficeSitesRequest) SetOfficeSiteType(v string) *DescribeOfficeSitesRequest {
	s.OfficeSiteType = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetRegionId(v string) *DescribeOfficeSitesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetSecurityProtection(v string) *DescribeOfficeSitesRequest {
	s.SecurityProtection = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetStatus(v string) *DescribeOfficeSitesRequest {
	s.Status = &v
	return s
}

func (s *DescribeOfficeSitesRequest) SetVpcId(v string) *DescribeOfficeSitesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeOfficeSitesRequest) Validate() error {
	return dara.Validate(s)
}
