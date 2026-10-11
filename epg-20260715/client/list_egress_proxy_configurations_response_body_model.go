// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyConfigurations(v []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) *ListEgressProxyConfigurationsResponseBody
	GetEgressProxyConfigurations() []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations
	SetMaxResults(v int32) *ListEgressProxyConfigurationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyConfigurationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEgressProxyConfigurationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEgressProxyConfigurationsResponseBody
	GetTotalCount() *int32
}

type ListEgressProxyConfigurationsResponseBody struct {
	// The list of egress proxy configurations.
	EgressProxyConfigurations []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations `json:"EgressProxyConfigurations,omitempty" xml:"EgressProxyConfigurations,omitempty" type:"Repeated"`
	// The maximum number of entries returned for the current request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - If NextToken is empty, no next query is required.
	//
	// - If NextToken is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01A05ABD-A4C0-5F6B-A554-AEEEC73F9F70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of egress proxy configurations that meet the query conditions.
	//
	// example:
	//
	// 19
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEgressProxyConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsResponseBody) GetEgressProxyConfigurations() []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	return s.EgressProxyConfigurations
}

func (s *ListEgressProxyConfigurationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyConfigurationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEgressProxyConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEgressProxyConfigurationsResponseBody) SetEgressProxyConfigurations(v []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) *ListEgressProxyConfigurationsResponseBody {
	s.EgressProxyConfigurations = v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBody) SetMaxResults(v int32) *ListEgressProxyConfigurationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBody) SetNextToken(v string) *ListEgressProxyConfigurationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBody) SetRequestId(v string) *ListEgressProxyConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBody) SetTotalCount(v int32) *ListEgressProxyConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBody) Validate() error {
	if s.EgressProxyConfigurations != nil {
		for _, item := range s.EgressProxyConfigurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations struct {
	// The time when the egress proxy configuration was created. The time is displayed in UTC in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-07-02T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the egress proxy configuration. The description must be 2 to 256 characters in length, and can contain Chinese characters, letters, digits, hyphens (-), forward slashes (/), periods (.), and underscores (_).
	//
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the egress proxy configuration.
	//
	// example:
	//
	// epc-gw80d16a969933e7****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The name of the egress proxy configuration. The name must be 2 to 128 characters in length, and must start with a letter or a Chinese character. It can contain digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testEpg
	EgressProxyConfigurationName *string `json:"EgressProxyConfigurationName,omitempty" xml:"EgressProxyConfigurationName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek6jw4dpt****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the egress proxy configuration. Valid values:
	//
	// - **Provisioning**: The egress proxy configuration is being created.
	//
	// - **Available**: The egress proxy configuration is available.
	//
	// - **Configuring**: The specifications of the egress proxy configuration are being changed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetDescription() *string {
	return s.Description
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetEgressProxyConfigurationName() *string {
	return s.EgressProxyConfigurationName
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) GetTags() []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags {
	return s.Tags
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetCreateTime(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetDescription(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.Description = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetEgressProxyConfigurationId(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetEgressProxyConfigurationName(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.EgressProxyConfigurationName = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetResourceGroupId(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetStatus(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.Status = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) SetTags(v []*ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations {
	s.Tags = v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurations) Validate() error {
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

type ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags struct {
	// The tag key. The key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. The value can be up to 256 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) SetKey(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags {
	s.Key = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) SetValue(v string) *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags {
	s.Value = &v
	return s
}

func (s *ListEgressProxyConfigurationsResponseBodyEgressProxyConfigurationsTags) Validate() error {
	return dara.Validate(s)
}
