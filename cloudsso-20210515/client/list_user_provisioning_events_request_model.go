// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListUserProvisioningEventsRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListUserProvisioningEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserProvisioningEventsRequest
	GetNextToken() *string
	SetUserProvisioningId(v string) *ListUserProvisioningEventsRequest
	GetUserProvisioningId() *string
}

type ListUserProvisioningEventsRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-00xz91nf****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return is larger than the value of `MaxResults`, the entries are truncated. The system returns entries based on the value of `MaxResults`, and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s ListUserProvisioningEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningEventsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserProvisioningEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserProvisioningEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserProvisioningEventsRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *ListUserProvisioningEventsRequest) SetDirectoryId(v string) *ListUserProvisioningEventsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetMaxResults(v int32) *ListUserProvisioningEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetNextToken(v string) *ListUserProvisioningEventsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) SetUserProvisioningId(v string) *ListUserProvisioningEventsRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *ListUserProvisioningEventsRequest) Validate() error {
	return dara.Validate(s)
}
