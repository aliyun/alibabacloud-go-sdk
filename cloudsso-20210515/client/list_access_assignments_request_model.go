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
	// The ID of the access configuration. This parameter is used as a filter condition for the query.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 20.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of query results. You do not need to specify `NextToken` when you call this operation for the first time.
	//
	// When you call this operation for the first time, if the total number of returned entries exceeds the value of `MaxResults`, the returned entries are truncated. Only the number of entries specified by `MaxResults` is returned, the `IsTruncated` parameter is set to `true`, and a `NextToken` value is returned. You can use the returned `NextToken` value to call this operation again to query the truncated entries while keeping other request parameters unchanged. You can repeat this process until `IsTruncated` is set to `false`, which indicates that all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the Alibaba Cloud account to which the authorization target belongs.
	//
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The ID of the CloudSSO identity, which is used as a filter condition for the query. Details:
	//
	// - If `PrincipalType` is set to User, `PrincipalId` is the ID of the CloudSSO user.
	//
	// - If `PrincipalType` is set to Group, `PrincipalId` is the ID of the CloudSSO group.
	//
	// > You must specify both `PrincipalId` and `PrincipalType` for the filter condition to take effect.
	//
	// example:
	//
	// u-00q8wbq42wiltcrk****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the CloudSSO identity, which is used as a filter condition for the query. Valid values:
	//
	// - User: user.
	//
	// - Group: group.
	//
	// > You must specify both `PrincipalId` and `PrincipalType` for the filter condition to take effect.
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the task target, which is used as a filter condition for the query.
	//
	// > You must specify both `TargetId` and `TargetType` for the filter condition to take effect.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the task target, which is used as a filter condition for the query.
	//
	// Valid values:
	//
	// - RD-Account: The task target is an account in a resource directory.
	//
	// > You must specify both `TargetId` and `TargetType` for the filter condition to take effect.
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
