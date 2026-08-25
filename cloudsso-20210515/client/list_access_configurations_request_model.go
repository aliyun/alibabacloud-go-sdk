// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *ListAccessConfigurationsRequest
	GetDirectoryId() *string
	SetFilter(v string) *ListAccessConfigurationsRequest
	GetFilter() *string
	SetMaxResults(v int32) *ListAccessConfigurationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessConfigurationsRequest
	GetNextToken() *string
	SetStatusNotifications(v string) *ListAccessConfigurationsRequest
	GetStatusNotifications() *string
	SetTags(v []*ListAccessConfigurationsRequestTags) *ListAccessConfigurationsRequest
	GetTags() []*ListAccessConfigurationsRequestTags
}

type ListAccessConfigurationsRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The filter condition.
	//
	// Format: <Attribute> <Operator> <Value>. The filter is case-insensitive. Currently, <Attribute> supports only AccessConfigurationName, and <Operator> supports only eq (Equals) and sw (Start With).
	//
	// Example: Filter = "AccessConfigurationName sw test" queries all access configurations whose names start with test. Filter = "AccessConfigurationName eq TestAccessConfiguration" queries the access configuration named TestAccessConfiguration.
	//
	// example:
	//
	// AccessConfigurationName sw test
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
	// When you call this API operation for the first time, if the total number of results exceeds the `MaxResults` limit, the results are truncated and only `MaxResults` entries are returned. In this case, the `IsTruncated` parameter is set to `true` and a `NextToken` is returned. You can use the `NextToken` returned from the previous call to continue calling this API operation while keeping other request parameters unchanged to query the truncated results. You can repeat this process until `IsTruncated` is `false`, which indicates that all data has been retrieved.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status notification information, which is used as a filter condition for the query.
	//
	// Valid values: ReprovisionRequired, which queries access configurations that need to be reprovisioned.
	//
	// example:
	//
	// ReprovisionRequired
	StatusNotifications *string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty"`
	// The list of tags.
	Tags []*ListAccessConfigurationsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAccessConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListAccessConfigurationsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListAccessConfigurationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessConfigurationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessConfigurationsRequest) GetStatusNotifications() *string {
	return s.StatusNotifications
}

func (s *ListAccessConfigurationsRequest) GetTags() []*ListAccessConfigurationsRequestTags {
	return s.Tags
}

func (s *ListAccessConfigurationsRequest) SetDirectoryId(v string) *ListAccessConfigurationsRequest {
	s.DirectoryId = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetFilter(v string) *ListAccessConfigurationsRequest {
	s.Filter = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetMaxResults(v int32) *ListAccessConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetNextToken(v string) *ListAccessConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetStatusNotifications(v string) *ListAccessConfigurationsRequest {
	s.StatusNotifications = &v
	return s
}

func (s *ListAccessConfigurationsRequest) SetTags(v []*ListAccessConfigurationsRequestTags) *ListAccessConfigurationsRequest {
	s.Tags = v
	return s
}

func (s *ListAccessConfigurationsRequest) Validate() error {
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

type ListAccessConfigurationsRequestTags struct {
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

func (s ListAccessConfigurationsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsRequestTags) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListAccessConfigurationsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListAccessConfigurationsRequestTags) SetKey(v string) *ListAccessConfigurationsRequestTags {
	s.Key = &v
	return s
}

func (s *ListAccessConfigurationsRequestTags) SetValue(v string) *ListAccessConfigurationsRequestTags {
	s.Value = &v
	return s
}

func (s *ListAccessConfigurationsRequestTags) Validate() error {
	return dara.Validate(s)
}
