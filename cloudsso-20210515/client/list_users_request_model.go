// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListUsersRequest
	GetDirectoryId() *string
	SetFilter(v string) *ListUsersRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListUsersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUsersRequest
	GetNextToken() *string
	SetProvisionType(v string) *ListUsersRequest
	GetProvisionType() *string
	SetStatus(v string) *ListUsersRequest
	GetStatus() *string
	SetTags(v []*ListUsersRequestTags) *ListUsersRequest
	GetTags() []*ListUsersRequestTags
}

type ListUsersRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// Format: `<Attribute> <Operator> <Value>`. This value is case-insensitive. Currently, `<Attribute>` supports only `UserName`, and `Operator` supports only `eq` (Equals) and `sw` (Start With).
	//
	// Example: Filter = "UserName sw test" queries all users whose usernames start with test. Filter = "UserName eq testuser" queries the user whose username is `testuser`.
	//
	// example:
	//
	// UserName sw test
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
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
	// The token for the next page of results. You do not need to specify `NextToken` for the first API call.
	//
	// When you call the API for the first time, if the total number of entries exceeds the `MaxResults` limit, the data is truncated and only `MaxResults` entries are returned. In this case, the response parameter `IsTruncated` is `true` and a `NextToken` is returned. You can use the `NextToken` returned from the previous call to continue calling the API while keeping other request parameters unchanged to query the truncated data. You can repeat this process until `IsTruncated` is `false`, which indicates that all data has been retrieved.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The user type. This parameter is used as a filter condition. Valid values:
	//
	// - Manual: The user is manually created.
	//
	// - Synchronized: The user is synchronized from an external identity provider.
	//
	// example:
	//
	// Manual
	ProvisionType *string `json:"ProvisionType,omitempty" xml:"ProvisionType,omitempty"`
	// The user status. This parameter is used as a filter condition. Valid values:
	//
	// - Enabled: The user is enabled.
	//
	// - Disabled: The user is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*ListUsersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUsersRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListUsersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUsersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUsersRequest) GetProvisionType() *string {
	return s.ProvisionType
}

func (s *ListUsersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListUsersRequest) GetTags() []*ListUsersRequestTags {
	return s.Tags
}

func (s *ListUsersRequest) SetDirectoryId(v string) *ListUsersRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListUsersRequest) SetFilter(v string) *ListUsersRequest {
	s.Filter = &v
	return s
}

func (s *ListUsersRequest) SetMaxResults(v int32) *ListUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUsersRequest) SetNextToken(v string) *ListUsersRequest {
	s.NextToken = &v
	return s
}

func (s *ListUsersRequest) SetProvisionType(v string) *ListUsersRequest {
	s.ProvisionType = &v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

func (s *ListUsersRequest) SetTags(v []*ListUsersRequestTags) *ListUsersRequest {
	s.Tags = v
	return s
}

func (s *ListUsersRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUsersRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUsersRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListUsersRequestTags) GoString() string {
	return s.String()
}

func (s *ListUsersRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListUsersRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListUsersRequestTags) SetKey(v string) *ListUsersRequestTags {
	s.Key = &v
	return s
}

func (s *ListUsersRequestTags) SetValue(v string) *ListUsersRequestTags {
	s.Value = &v
	return s
}

func (s *ListUsersRequestTags) Validate() error {
	return dara.Validate(s)
}
