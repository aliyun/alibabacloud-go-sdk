// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyRuleGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyRuleGroupIds(v []*string) *ListEgressProxyRuleGroupsRequest
	GetEgressProxyRuleGroupIds() []*string
	SetEgressProxyRuleGroupNames(v []*string) *ListEgressProxyRuleGroupsRequest
	GetEgressProxyRuleGroupNames() []*string
	SetMaxResults(v int32) *ListEgressProxyRuleGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyRuleGroupsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListEgressProxyRuleGroupsRequest
	GetResourceGroupId() *string
	SetSkip(v int32) *ListEgressProxyRuleGroupsRequest
	GetSkip() *int32
	SetStatus(v string) *ListEgressProxyRuleGroupsRequest
	GetStatus() *string
	SetTag(v []*ListEgressProxyRuleGroupsRequestTag) *ListEgressProxyRuleGroupsRequest
	GetTag() []*ListEgressProxyRuleGroupsRequestTag
}

type ListEgressProxyRuleGroupsRequest struct {
	EgressProxyRuleGroupIds   []*string `json:"EgressProxyRuleGroupIds,omitempty" xml:"EgressProxyRuleGroupIds,omitempty" type:"Repeated"`
	EgressProxyRuleGroupNames []*string `json:"EgressProxyRuleGroupNames,omitempty" xml:"EgressProxyRuleGroupNames,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rg-acfmwk2r5tq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 0
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// Available
	Status *string                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*ListEgressProxyRuleGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListEgressProxyRuleGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsRequest) GetEgressProxyRuleGroupIds() []*string {
	return s.EgressProxyRuleGroupIds
}

func (s *ListEgressProxyRuleGroupsRequest) GetEgressProxyRuleGroupNames() []*string {
	return s.EgressProxyRuleGroupNames
}

func (s *ListEgressProxyRuleGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyRuleGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyRuleGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyRuleGroupsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListEgressProxyRuleGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyRuleGroupsRequest) GetTag() []*ListEgressProxyRuleGroupsRequestTag {
	return s.Tag
}

func (s *ListEgressProxyRuleGroupsRequest) SetEgressProxyRuleGroupIds(v []*string) *ListEgressProxyRuleGroupsRequest {
	s.EgressProxyRuleGroupIds = v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetEgressProxyRuleGroupNames(v []*string) *ListEgressProxyRuleGroupsRequest {
	s.EgressProxyRuleGroupNames = v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetMaxResults(v int32) *ListEgressProxyRuleGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetNextToken(v string) *ListEgressProxyRuleGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetResourceGroupId(v string) *ListEgressProxyRuleGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetSkip(v int32) *ListEgressProxyRuleGroupsRequest {
	s.Skip = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetStatus(v string) *ListEgressProxyRuleGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) SetTag(v []*ListEgressProxyRuleGroupsRequestTag) *ListEgressProxyRuleGroupsRequest {
	s.Tag = v
	return s
}

func (s *ListEgressProxyRuleGroupsRequest) Validate() error {
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

type ListEgressProxyRuleGroupsRequestTag struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyRuleGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyRuleGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyRuleGroupsRequestTag) SetKey(v string) *ListEgressProxyRuleGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequestTag) SetValue(v string) *ListEgressProxyRuleGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *ListEgressProxyRuleGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
