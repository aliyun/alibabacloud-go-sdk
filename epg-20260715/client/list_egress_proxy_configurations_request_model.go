// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyConfigurationIds(v []*string) *ListEgressProxyConfigurationsRequest
	GetEgressProxyConfigurationIds() []*string
	SetEgressProxyConfigurationNames(v []*string) *ListEgressProxyConfigurationsRequest
	GetEgressProxyConfigurationNames() []*string
	SetMaxResults(v int32) *ListEgressProxyConfigurationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyConfigurationsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListEgressProxyConfigurationsRequest
	GetResourceGroupId() *string
	SetSkip(v int32) *ListEgressProxyConfigurationsRequest
	GetSkip() *int32
	SetStatus(v string) *ListEgressProxyConfigurationsRequest
	GetStatus() *string
	SetTag(v []*ListEgressProxyConfigurationsRequestTag) *ListEgressProxyConfigurationsRequest
	GetTag() []*ListEgressProxyConfigurationsRequestTag
}

type ListEgressProxyConfigurationsRequest struct {
	// The list of egress proxy configuration IDs.
	EgressProxyConfigurationIds []*string `json:"EgressProxyConfigurationIds,omitempty" xml:"EgressProxyConfigurationIds,omitempty" type:"Repeated"`
	// The list of egress proxy configuration names.
	EgressProxyConfigurationNames []*string `json:"EgressProxyConfigurationNames,omitempty" xml:"EgressProxyConfigurationNames,omitempty" type:"Repeated"`
	// The maximum number of entries to return per page during paging in a paged query. Valid values: 0 to 1000. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query. Valid values:
	//
	// - Leave this parameter empty for the first query or when no next query is required.
	//
	// - If a next query is required, set this parameter to the value of NextToken returned from the previous API call.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4jcyx5e****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries to skip during paging in a paged query. The value must be greater than or equal to 0. Default value: 0.
	//
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// Filters egress proxy configurations by status. Valid values:
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
	Tag []*ListEgressProxyConfigurationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEgressProxyConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsRequest) GetEgressProxyConfigurationIds() []*string {
	return s.EgressProxyConfigurationIds
}

func (s *ListEgressProxyConfigurationsRequest) GetEgressProxyConfigurationNames() []*string {
	return s.EgressProxyConfigurationNames
}

func (s *ListEgressProxyConfigurationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyConfigurationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyConfigurationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyConfigurationsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListEgressProxyConfigurationsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyConfigurationsRequest) GetTag() []*ListEgressProxyConfigurationsRequestTag {
	return s.Tag
}

func (s *ListEgressProxyConfigurationsRequest) SetEgressProxyConfigurationIds(v []*string) *ListEgressProxyConfigurationsRequest {
	s.EgressProxyConfigurationIds = v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetEgressProxyConfigurationNames(v []*string) *ListEgressProxyConfigurationsRequest {
	s.EgressProxyConfigurationNames = v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetMaxResults(v int32) *ListEgressProxyConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetNextToken(v string) *ListEgressProxyConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetResourceGroupId(v string) *ListEgressProxyConfigurationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetSkip(v int32) *ListEgressProxyConfigurationsRequest {
	s.Skip = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetStatus(v string) *ListEgressProxyConfigurationsRequest {
	s.Status = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) SetTag(v []*ListEgressProxyConfigurationsRequestTag) *ListEgressProxyConfigurationsRequest {
	s.Tag = v
	return s
}

func (s *ListEgressProxyConfigurationsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyConfigurationsRequestTag struct {
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

func (s ListEgressProxyConfigurationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyConfigurationsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEgressProxyConfigurationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyConfigurationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyConfigurationsRequestTag) SetKey(v string) *ListEgressProxyConfigurationsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequestTag) SetValue(v string) *ListEgressProxyConfigurationsRequestTag {
	s.Value = &v
	return s
}

func (s *ListEgressProxyConfigurationsRequestTag) Validate() error {
	return dara.Validate(s)
}
