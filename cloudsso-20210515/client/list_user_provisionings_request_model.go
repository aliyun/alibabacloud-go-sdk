// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListUserProvisioningsRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListUserProvisioningsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserProvisioningsRequest
	GetNextToken() *string
	SetPrincipalId(v string) *ListUserProvisioningsRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *ListUserProvisioningsRequest
	GetPrincipalType() *string
	SetTargetId(v string) *ListUserProvisioningsRequest
	GetTargetId() *string
	SetTargetType(v string) *ListUserProvisioningsRequest
	GetTargetType() *string
}

type ListUserProvisioningsRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
	//
	// When you call this operation for the first time, if the total number of entries to return is larger than the value of `MaxResults`, the entries are truncated. The system returns entries based on the value of `MaxResults`, and does not return the excess entries. In this case, the value of the response parameter `IsTruncated` is `true`, and `NextToken` is returned. In the next call, you can use the value of `NextToken` and maintain the settings of the other request parameters to query the excess entries. You can repeat the call until the value of `IsTruncated` becomes `false`. This way, all entries are returned.
	//
	// example:
	//
	// 27EbL9j4ZgZjsMZFqbZFgbwQ1VXFU1Khcpx9e2vrW1zwzTBmTGWaM7ixHhRin8SCsxaJdazYVCzeKc2UF2QkyGb83cPhr8ZxrzoaiTd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The identity ID of the RAM user provisioning. Valid values:
	//
	// - If `Group` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user group (g-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// - If `User` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user (u-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// example:
	//
	// u-88d73u*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The identity type of the RAM user provisioning. Valid values:
	//
	// - User: The identity of the RAM user provisioning is a CloudSSO user.
	//
	// - Group: The identity of the RAM user provisioning is a CloudSSO user group.
	//
	// example:
	//
	// RD-Account
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the object for which you create the RAM user provisioning. The value is fixed as the ID of the member in the resource directory.
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The object for which you create the RAM user provisioning. The value is fixed as `RD-Account`.
	//
	// example:
	//
	// User
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ListUserProvisioningsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningsRequest) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserProvisioningsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserProvisioningsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserProvisioningsRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListUserProvisioningsRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListUserProvisioningsRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListUserProvisioningsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListUserProvisioningsRequest) SetDirectoryId(v string) *ListUserProvisioningsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetMaxResults(v int32) *ListUserProvisioningsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetNextToken(v string) *ListUserProvisioningsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetPrincipalId(v string) *ListUserProvisioningsRequest {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetPrincipalType(v string) *ListUserProvisioningsRequest {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetTargetId(v string) *ListUserProvisioningsRequest {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningsRequest) SetTargetType(v string) *ListUserProvisioningsRequest {
	s.TargetType = &v
	return s
}

func (s *ListUserProvisioningsRequest) Validate() error {
	return dara.Validate(s)
}
