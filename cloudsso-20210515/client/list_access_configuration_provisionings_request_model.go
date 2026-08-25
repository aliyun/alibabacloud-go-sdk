// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationProvisioningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurationId(v string) *ListAccessConfigurationProvisioningsRequest
	GetAccessConfigurationId() *string
	SetDirectoryId(v string) *ListAccessConfigurationProvisioningsRequest
	GetDirectoryId() *string
	SetMaxResults(v int32) *ListAccessConfigurationProvisioningsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessConfigurationProvisioningsRequest
	GetNextToken() *string
	SetOriginTargetId(v string) *ListAccessConfigurationProvisioningsRequest
	GetOriginTargetId() *string
	SetProvisioningStatus(v string) *ListAccessConfigurationProvisioningsRequest
	GetProvisioningStatus() *string
	SetTargetId(v string) *ListAccessConfigurationProvisioningsRequest
	GetTargetId() *string
	SetTargetType(v string) *ListAccessConfigurationProvisioningsRequest
	GetTargetType() *string
}

type ListAccessConfigurationProvisioningsRequest struct {
	// The ID of the access configuration. The ID can be used to filter access permissions.
	//
	// example:
	//
	// ac-00ccule7tadaijxc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The ID of the directory.
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
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. If this is your first time to call this operation, you do not need to specify the `NextToken` parameter.
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
	// The status of the access configuration. The value can be used to filter accounts. Valid values:
	//
	// - Provisioned: The access configuration is provisioned.
	//
	// - ReprovisionRequired: The access configuration needs to be re-provisioned.
	//
	// - DeprovisionFailed: The access configuration failed to be provisioned.
	//
	// example:
	//
	// Provisioned
	ProvisioningStatus *string `json:"ProvisioningStatus,omitempty" xml:"ProvisioningStatus,omitempty"`
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

func (s ListAccessConfigurationProvisioningsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationProvisioningsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationProvisioningsRequest) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListAccessConfigurationProvisioningsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAccessConfigurationProvisioningsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessConfigurationProvisioningsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessConfigurationProvisioningsRequest) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ListAccessConfigurationProvisioningsRequest) GetProvisioningStatus() *string {
	return s.ProvisioningStatus
}

func (s *ListAccessConfigurationProvisioningsRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *ListAccessConfigurationProvisioningsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListAccessConfigurationProvisioningsRequest) SetAccessConfigurationId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetDirectoryId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetMaxResults(v int32) *ListAccessConfigurationProvisioningsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetNextToken(v string) *ListAccessConfigurationProvisioningsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetOriginTargetId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.OriginTargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetProvisioningStatus(v string) *ListAccessConfigurationProvisioningsRequest {
	s.ProvisioningStatus = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetTargetId(v string) *ListAccessConfigurationProvisioningsRequest {
	s.TargetId = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) SetTargetType(v string) *ListAccessConfigurationProvisioningsRequest {
	s.TargetType = &v
	return s
}

func (s *ListAccessConfigurationProvisioningsRequest) Validate() error {
	return dara.Validate(s)
}
