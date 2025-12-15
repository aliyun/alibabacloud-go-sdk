// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoGroupingRemediationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAutoGroupingRemediationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoGroupingRemediationsResponseBody
	GetNextToken() *string
	SetRemediations(v []*ListAutoGroupingRemediationsResponseBodyRemediations) *ListAutoGroupingRemediationsResponseBody
	GetRemediations() []*ListAutoGroupingRemediationsResponseBodyRemediations
	SetRequestId(v string) *ListAutoGroupingRemediationsResponseBody
	GetRequestId() *string
}

type ListAutoGroupingRemediationsResponseBody struct {
	// The number of entries per page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The remediation records.
	Remediations []*ListAutoGroupingRemediationsResponseBodyRemediations `json:"Remediations,omitempty" xml:"Remediations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6F959E33-7B6D-5F58-BB0B-ED616DC7C70B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAutoGroupingRemediationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRemediationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoGroupingRemediationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoGroupingRemediationsResponseBody) GetRemediations() []*ListAutoGroupingRemediationsResponseBodyRemediations {
	return s.Remediations
}

func (s *ListAutoGroupingRemediationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoGroupingRemediationsResponseBody) SetMaxResults(v int32) *ListAutoGroupingRemediationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBody) SetNextToken(v string) *ListAutoGroupingRemediationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBody) SetRemediations(v []*ListAutoGroupingRemediationsResponseBodyRemediations) *ListAutoGroupingRemediationsResponseBody {
	s.Remediations = v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBody) SetRequestId(v string) *ListAutoGroupingRemediationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBody) Validate() error {
	if s.Remediations != nil {
		for _, item := range s.Remediations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAutoGroupingRemediationsResponseBodyRemediations struct {
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The remediation record ID.
	//
	// example:
	//
	// 0028d4****cfe94956ef6708a373f396fbc840e306f
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// The remediation time.
	//
	// example:
	//
	// 2022-01-01 00:00:00
	RemediationTime *string `json:"RemediationTime,omitempty" xml:"RemediationTime,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-uf664f66v1****drkea4
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// You can obtain the resource type from the **Resource type*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the Alibaba Cloud service.
	//
	// You can obtain the ID from the **Service code*	- column in [Services that work with Resource Group](https://help.aliyun.com/document_detail/94479.html).
	//
	// example:
	//
	// ecs
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The information about the new resource group.
	TargetResourceGroupInfo *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo `json:"TargetResourceGroupInfo,omitempty" xml:"TargetResourceGroupInfo,omitempty" type:"Struct"`
}

func (s ListAutoGroupingRemediationsResponseBodyRemediations) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRemediationsResponseBodyRemediations) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetRemediationId() *string {
	return s.RemediationId
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetRemediationTime() *string {
	return s.RemediationTime
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetService() *string {
	return s.Service
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) GetTargetResourceGroupInfo() *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo {
	return s.TargetResourceGroupInfo
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetRegionId(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.RegionId = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetRemediationId(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.RemediationId = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetRemediationTime(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.RemediationTime = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetResourceId(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.ResourceId = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetResourceType(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.ResourceType = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetService(v string) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.Service = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) SetTargetResourceGroupInfo(v *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) *ListAutoGroupingRemediationsResponseBodyRemediations {
	s.TargetResourceGroupInfo = v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediations) Validate() error {
	if s.TargetResourceGroupInfo != nil {
		if err := s.TargetResourceGroupInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo struct {
	// The resource group name.
	//
	// example:
	//
	// ProjectA
	ResourceGroupDisplayName *string `json:"ResourceGroupDisplayName,omitempty" xml:"ResourceGroupDisplayName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmygrk****wfa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) GoString() string {
	return s.String()
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) GetResourceGroupDisplayName() *string {
	return s.ResourceGroupDisplayName
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) SetResourceGroupDisplayName(v string) *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo {
	s.ResourceGroupDisplayName = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) SetResourceGroupId(v string) *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAutoGroupingRemediationsResponseBodyRemediationsTargetResourceGroupInfo) Validate() error {
	return dara.Validate(s)
}
