// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListGroupsRequest
	GetDirectoryId() *string
	SetFilter(v string) *ListGroupsRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListGroupsRequest
	GetNextToken() *string
	SetProvisionType(v string) *ListGroupsRequest
	GetProvisionType() *string
}

type ListGroupsRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// You must specify the value in the `<Attribute> <Operator> <Value>` format. The value is not case-sensitive. You can set `<Attribute>` only to `GroupName` and `<Operator>` only to `eq` or `sw`. The value eq indicates Equals, and the value sw indicates Start With.
	//
	// For example, if you set Filter to GroupName sw test, the operation queries the groups whose names start with test. If you set Filter to GroupName eq testgroup, the operation queries the group whose name is testgroup.
	//
	// example:
	//
	// GroupName eq testgroup
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// When you call this operation for the first time, if the total number of entries to return is larger than the value of the `MaxResults` parameter, the entries are truncated. The system returns entries based on the value of the `MaxResults` parameter, and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and the `NextToken` parameter is returned. In the next call, you can use the value of the `NextToken` parameter and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the group. The type can be used to filter groups. Valid values:
	//
	// 	- Manual: The group is manually created.
	//
	// 	- Synchronized: The group is synchronized from an external IdP.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListGroupsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsRequest) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListGroupsRequest) SetDirectoryId(v string) *ListGroupsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListGroupsRequest) SetFilter(v string) *ListGroupsRequest {
	s.Filter = &v
	return s
}

func (s *ListGroupsRequest) SetMaxResults(v int32) *ListGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsRequest) SetNextToken(v string) *ListGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListGroupsRequest) SetProvisionType(v string) *ListGroupsRequest {
	s.ProvisionType = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
