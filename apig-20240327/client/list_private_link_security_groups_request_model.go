// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateLinkSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPrivateLinkSecurityGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPrivateLinkSecurityGroupsRequest
	GetNextToken() *string
	SetVpcId(v string) *ListPrivateLinkSecurityGroupsRequest
	GetVpcId() *string
}

type ListPrivateLinkSecurityGroupsRequest struct {
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the VPC owned by the authorized user. The candidate security groups must belong to this VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListPrivateLinkSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateLinkSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateLinkSecurityGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPrivateLinkSecurityGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPrivateLinkSecurityGroupsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ListPrivateLinkSecurityGroupsRequest) SetMaxResults(v int32) *ListPrivateLinkSecurityGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsRequest) SetNextToken(v string) *ListPrivateLinkSecurityGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsRequest) SetVpcId(v string) *ListPrivateLinkSecurityGroupsRequest {
	s.VpcId = &v
	return s
}

func (s *ListPrivateLinkSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
