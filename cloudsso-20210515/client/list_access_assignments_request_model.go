// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessAssignmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *ListAccessAssignmentsRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *ListAccessAssignmentsRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListAccessAssignmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessAssignmentsRequest
	GetNextToken() *string
	SetOriginTargetId(v string) *ListAccessAssignmentsRequest
	GetOriginTargetId() *string
	SetPrincipalId(v string) *ListAccessAssignmentsRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *ListAccessAssignmentsRequest
	GetPrincipalType() *string
	SetTargetId(v string) *ListAccessAssignmentsRequest
	GetTargetId() *string
	SetTargetType(v string) *ListAccessAssignmentsRequest
	GetTargetType() *string
}

type ListAccessAssignmentsRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return exceeds the value of `MaxResults`, the entries are truncated. Only the entries that match the value of `MaxResults` are returned, and the excess entries are not returned. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the CloudSSO identity. The ID can be used to filter access permissions.
	//
	// - If you set `PrincipalType` to User, set `PrincipalId` to the ID of the CloudSSO user.
	//
	// - If you set `PrincipalType` to Group, set `PrincipalId` to the ID of the CloudSSO group.
	//
	// > You can use the type to filter access permissions only if you specify both PrincipalId and `PrincipalType`.\\`\\`
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity. The type can be used to filter access permissions. Valid values:
	//
	// - User
	//
	// - Group
	//
	// > You can use the type to filter access permissions only if you specify both PrincipalId and `PrincipalType`.\\`\\`
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task object. The ID can be used to filter access permissions.
	//
	// > You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task object. The type can be used to filter access permissions.
	//
	// Set the value to RD-Account, which specifies the accounts in the resource directory.
	//
	// > You can use the type to filter access permissions only if you specify both `TargetId` and `TargetType`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListAccessAssignmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessAssignmentsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessAssignmentsRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListAccessAssignmentsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAccessAssignmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessAssignmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessAssignmentsRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ListAccessAssignmentsRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListAccessAssignmentsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListAccessAssignmentsRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAccessAssignmentsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAccessAssignmentsRequest) SetAccessConfigurationId(v string) *ListAccessAssignmentsRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetDirectoryId(v string) *ListAccessAssignmentsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetMaxResults(v int32) *ListAccessAssignmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetNextToken(v string) *ListAccessAssignmentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetOriginTargetId(v string) *ListAccessAssignmentsRequest {
	s.OriginTargetId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetPrincipalId(v string) *ListAccessAssignmentsRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetPrincipalType(v string) *ListAccessAssignmentsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetTargetId(v string) *ListAccessAssignmentsRequest {
	s.TargetId = &v
	return s
}

func (s *ListAccessAssignmentsRequest) SetTargetType(v string) *ListAccessAssignmentsRequest {
	s.TargetType = &v
	return s
}

func (s *ListAccessAssignmentsRequest) Validate() error {
	return dara.Validate(s)
}
