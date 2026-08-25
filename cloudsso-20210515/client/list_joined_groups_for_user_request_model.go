// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJoinedGroupsForUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListJoinedGroupsForUserRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListJoinedGroupsForUserRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListJoinedGroupsForUserRequest
	GetNextToken() *string
	SetUserId(v string) *ListJoinedGroupsForUserRequest
	GetUserId() *string
}

type ListJoinedGroupsForUserRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
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
	// The ID of the user.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListJoinedGroupsForUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJoinedGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListJoinedGroupsForUserRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListJoinedGroupsForUserRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJoinedGroupsForUserRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJoinedGroupsForUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *ListJoinedGroupsForUserRequest) SetDirectoryId(v string) *ListJoinedGroupsForUserRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetMaxResults(v int32) *ListJoinedGroupsForUserRequest {
	s.MaxResults = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetNextToken(v string) *ListJoinedGroupsForUserRequest {
	s.NextToken = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) SetUserId(v string) *ListJoinedGroupsForUserRequest {
	s.UserId = &v
	return s
}

func (s *ListJoinedGroupsForUserRequest) Validate() error {
	return dara.Validate(s)
}
