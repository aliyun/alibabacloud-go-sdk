// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessConfigurations(v []*ListAccessConfigurationsResponseBodyAccessConfigurations) *ListAccessConfigurationsResponseBody
	GetAccessConfigurations() []*ListAccessConfigurationsResponseBodyAccessConfigurations
	SetIsTruncated(v bool) *ListAccessConfigurationsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListAccessConfigurationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccessConfigurationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessConfigurationsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListAccessConfigurationsResponseBody
	GetTotalCounts() *int32
}

type ListAccessConfigurationsResponseBody struct {
	// The list of access configurations.
	AccessConfigurations []*ListAccessConfigurationsResponseBodyAccessConfigurations `json:"AccessConfigurations,omitempty" xml:"AccessConfigurations,omitempty" type:"Repeated"`
	// Indicates whether the results are truncated. Valid values:
	//
	// - true: The results are truncated.
	//
	// - false: The results are not truncated.
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// > This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// K1c3o9K7pFxoTtxH1Nm7MMLb7zrDGvftYBQBPDVv7AD3a8yhRb3Mk8L9ivmN6bFSjfkZNTAg3h4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2BC0CBAC-45E1-5BD3-BF6E-F69D1D5391C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that match the request parameters.
	//
	// example:
	//
	// 2
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListAccessConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponseBody) GetAccessConfigurations() []*ListAccessConfigurationsResponseBodyAccessConfigurations {
	return s.AccessConfigurations
}

func (s *ListAccessConfigurationsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListAccessConfigurationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccessConfigurationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessConfigurationsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListAccessConfigurationsResponseBody) SetAccessConfigurations(v []*ListAccessConfigurationsResponseBodyAccessConfigurations) *ListAccessConfigurationsResponseBody {
	s.AccessConfigurations = v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetIsTruncated(v bool) *ListAccessConfigurationsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetMaxResults(v int32) *ListAccessConfigurationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetNextToken(v string) *ListAccessConfigurationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetRequestId(v string) *ListAccessConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) SetTotalCounts(v int32) *ListAccessConfigurationsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListAccessConfigurationsResponseBody) Validate() error {
	if s.AccessConfigurations != nil {
		for _, item := range s.AccessConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessConfigurationsResponseBodyAccessConfigurations struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// The time when the access configuration was created.
	//
	// example:
	//
	// 2021-11-02T08:44:23Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the access configuration.
	//
	// example:
	//
	// This is an access configuration.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The initial access page.
	//
	// The URL of the initial page that is displayed when a CloudSSO user uses the access configuration to access an account in a resource directory.
	//
	// example:
	//
	// https://cloudsso.console.aliyun.com
	RelayState *string `json:"RelayState,omitempty" xml:"RelayState,omitempty"`
	// The session duration.
	//
	// The maximum duration of a session when a CloudSSO user uses the access configuration to access an account in a resource directory.
	//
	// Unit: seconds.
	//
	// example:
	//
	// 900
	SessionDuration *int32 `json:"SessionDuration,omitempty" xml:"SessionDuration,omitempty"`
	// The status notification information.
	StatusNotifications []*string `json:"StatusNotifications,omitempty" xml:"StatusNotifications,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*ListAccessConfigurationsResponseBodyAccessConfigurationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the access configuration was last modified.
	//
	// example:
	//
	// 2021-11-02T08:44:23Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListAccessConfigurationsResponseBodyAccessConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsResponseBodyAccessConfigurations) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetDescription() *string {
	return s.Description
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetRelayState() *string {
	return s.RelayState
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetSessionDuration() *int32 {
	return s.SessionDuration
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetStatusNotifications() []*string {
	return s.StatusNotifications
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetTags() []*ListAccessConfigurationsResponseBodyAccessConfigurationsTags {
	return s.Tags
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetAccessConfigurationId(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.AccessConfigurationId = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetAccessConfigurationName(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.AccessConfigurationName = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetCreateTime(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetDescription(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.Description = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetRelayState(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.RelayState = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetSessionDuration(v int32) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.SessionDuration = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetStatusNotifications(v []*string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.StatusNotifications = v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetTags(v []*ListAccessConfigurationsResponseBodyAccessConfigurationsTags) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.Tags = v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) SetUpdateTime(v string) *ListAccessConfigurationsResponseBodyAccessConfigurations {
	s.UpdateTime = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurations) Validate() error {
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

type ListAccessConfigurationsResponseBodyAccessConfigurationsTags struct {
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

func (s ListAccessConfigurationsResponseBodyAccessConfigurationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAccessConfigurationsResponseBodyAccessConfigurationsTags) GoString() string {
	return s.String()
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurationsTags) GetKey() *string {
	return s.Key
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurationsTags) GetValue() *string {
	return s.Value
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurationsTags) SetKey(v string) *ListAccessConfigurationsResponseBodyAccessConfigurationsTags {
	s.Key = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurationsTags) SetValue(v string) *ListAccessConfigurationsResponseBodyAccessConfigurationsTags {
	s.Value = &v
	return s
}

func (s *ListAccessConfigurationsResponseBodyAccessConfigurationsTags) Validate() error {
	return dara.Validate(s)
}
