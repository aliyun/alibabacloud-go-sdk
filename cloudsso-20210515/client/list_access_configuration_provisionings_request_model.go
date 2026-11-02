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
	// The access configuration ID used as a filter condition for the query.
	//
	// example:
	//
	// ac-00ccule7tadaijxc****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The directory ID.
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
	// The token used to return the next page of query results. You do not need to specify NextToken when you call this operation for the first time.
	//
	// When you call this operation for the first time, if the total number of returned entries exceeds the value of MaxResults, the returned results are truncated. Only the number of entries specified by MaxResults is returned, the IsTruncated parameter is set to true, and a NextToken is returned. You can use the returned NextToken to call this operation again to query the truncated entries while keeping other request parameters unchanged. You can repeat this process until IsTruncated is set to false, which indicates that all entries are returned.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The Alibaba Cloud account ID of the provisioning target.
	//
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The provisioning status used as a filter condition for the query. Valid values:
	//
	// - Provisioned: The provisioning succeeded.
	//
	// - ReprovisionRequired: Reprovisioning is required.
	//
	// - DeprovisionFailed: The provisioning failed.
	//
	// example:
	//
	// Provisioned
	ProvisioningStatus *string `json:"ProvisioningStatus,omitempty" xml:"ProvisioningStatus,omitempty"`
	// The target ID used as a filter condition for the query.
	//
	// > You must specify both TargetId and TargetType for them to take effect as filter conditions.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The target type used as a filter condition for the query.
	//
	// Valid values:
	//
	// - RD-Account: The target is an account in a resource directory.
	//
	// > You must specify both TargetId and TargetType for them to take effect as filter conditions.
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
