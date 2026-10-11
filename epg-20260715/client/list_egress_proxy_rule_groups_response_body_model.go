// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyRuleGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEgressProxyRuleGroups(v []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) *ListEgressProxyRuleGroupsResponseBody
	GetEgressProxyRuleGroups() []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups
	SetMaxResults(v int32) *ListEgressProxyRuleGroupsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEgressProxyRuleGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEgressProxyRuleGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListEgressProxyRuleGroupsResponseBody
	GetTotalCount() *int32
}

type ListEgressProxyRuleGroupsResponseBody struct {
	EgressProxyRuleGroups []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups `json:"EgressProxyRuleGroups,omitempty" xml:"EgressProxyRuleGroups,omitempty" type:"Repeated"`
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
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEgressProxyRuleGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsResponseBody) GetEgressProxyRuleGroups() []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	return s.EgressProxyRuleGroups
}

func (s *ListEgressProxyRuleGroupsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEgressProxyRuleGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEgressProxyRuleGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEgressProxyRuleGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListEgressProxyRuleGroupsResponseBody) SetEgressProxyRuleGroups(v []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) *ListEgressProxyRuleGroupsResponseBody {
	s.EgressProxyRuleGroups = v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBody) SetMaxResults(v int32) *ListEgressProxyRuleGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBody) SetNextToken(v string) *ListEgressProxyRuleGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBody) SetRequestId(v string) *ListEgressProxyRuleGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBody) SetTotalCount(v int32) *ListEgressProxyRuleGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBody) Validate() error {
	if s.EgressProxyRuleGroups != nil {
		for _, item := range s.EgressProxyRuleGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups struct {
	// example:
	//
	// 2026-08-06T03:18:12Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// testEpg
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// eppool-2ze06f6a9a2067ae****
	EgressProxyIpPoolId *string `json:"EgressProxyIpPoolId,omitempty" xml:"EgressProxyIpPoolId,omitempty"`
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// example:
	//
	// testEpg
	EgressProxyRuleGroupName *string `json:"EgressProxyRuleGroupName,omitempty" xml:"EgressProxyRuleGroupName,omitempty"`
	// example:
	//
	// rg-acfmwj4y6s7***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// Available
	Status *string                                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetDescription() *string {
	return s.Description
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetEgressProxyIpPoolId() *string {
	return s.EgressProxyIpPoolId
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetEgressProxyRuleGroupName() *string {
	return s.EgressProxyRuleGroupName
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetStatus() *string {
	return s.Status
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) GetTags() []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags {
	return s.Tags
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetCreateTime(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.CreateTime = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetDescription(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.Description = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetEgressProxyIpPoolId(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.EgressProxyIpPoolId = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetEgressProxyRuleGroupId(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetEgressProxyRuleGroupName(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.EgressProxyRuleGroupName = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetResourceGroupId(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetStatus(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.Status = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) SetTags(v []*ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups {
	s.Tags = v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroups) Validate() error {
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

type ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags struct {
	// example:
	//
	// epg
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) SetKey(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags {
	s.Key = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) SetValue(v string) *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags {
	s.Value = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponseBodyEgressProxyRuleGroupsTags) Validate() error {
	return dara.Validate(s)
}
