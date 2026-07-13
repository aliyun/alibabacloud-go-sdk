// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProtectionPoliciesResponseBodyData) *ListProtectionPoliciesResponseBody
	GetData() *ListProtectionPoliciesResponseBodyData
	SetRequestId(v string) *ListProtectionPoliciesResponseBody
	GetRequestId() *string
}

type ListProtectionPoliciesResponseBody struct {
	// The data returned.
	Data *ListProtectionPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5B2F09BF-CEBD-5A7E-AC01-E7F86169A5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProtectionPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBody) GetData() *ListProtectionPoliciesResponseBodyData {
	return s.Data
}

func (s *ListProtectionPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProtectionPoliciesResponseBody) SetData(v *ListProtectionPoliciesResponseBodyData) *ListProtectionPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ListProtectionPoliciesResponseBody) SetRequestId(v string) *ListProtectionPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectionPoliciesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProtectionPoliciesResponseBodyData struct {
	// The response parameters.
	Content []*ListProtectionPoliciesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of results requested.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAOTzWWYAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM4NzA3NTcwMzY2MjMwNzY2ODcyMzAzMTY2Nzg3ODY5MzY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProtectionPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyData) GetContent() []*ListProtectionPoliciesResponseBodyDataContent {
	return s.Content
}

func (s *ListProtectionPoliciesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectionPoliciesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectionPoliciesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProtectionPoliciesResponseBodyData) SetContent(v []*ListProtectionPoliciesResponseBodyDataContent) *ListProtectionPoliciesResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyData) SetMaxResults(v int32) *ListProtectionPoliciesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyData) SetNextToken(v string) *ListProtectionPoliciesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyData) SetTotalCount(v int64) *ListProtectionPoliciesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectionPoliciesResponseBodyDataContent struct {
	// The attached resource category IDs.
	BoundResourceCategoryIds []*string `json:"BoundResourceCategoryIds,omitempty" xml:"BoundResourceCategoryIds,omitempty" type:"Repeated"`
	// The summary of the latest application result.
	LatestApplySummary *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary `json:"LatestApplySummary,omitempty" xml:"LatestApplySummary,omitempty" type:"Struct"`
	// The time when the policy was last applied.
	//
	// example:
	//
	// 1742167218
	LatestApplyTime *int64 `json:"LatestApplyTime,omitempty" xml:"LatestApplyTime,omitempty"`
	// The task ID of the latest policy application.
	//
	// example:
	//
	// t-123***7890
	LatestTaskId *string `json:"LatestTaskId,omitempty" xml:"LatestTaskId,omitempty"`
	// The protection policy ID.
	//
	// example:
	//
	// p-123***7890
	ProtectionPolicyId *string `json:"ProtectionPolicyId,omitempty" xml:"ProtectionPolicyId,omitempty"`
	// The protection policy name.
	//
	// example:
	//
	// MyProtectionPolicy
	ProtectionPolicyName *string `json:"ProtectionPolicyName,omitempty" xml:"ProtectionPolicyName,omitempty"`
	// The region ID of the protection policy.
	//
	// example:
	//
	// cn-hangzhou
	ProtectionPolicyRegionId *string `json:"ProtectionPolicyRegionId,omitempty" xml:"ProtectionPolicyRegionId,omitempty"`
	// The sub-protection policies.
	SubProtectionPolicies []*ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies `json:"SubProtectionPolicies,omitempty" xml:"SubProtectionPolicies,omitempty" type:"Repeated"`
}

func (s ListProtectionPoliciesResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetBoundResourceCategoryIds() []*string {
	return s.BoundResourceCategoryIds
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetLatestApplySummary() *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary {
	return s.LatestApplySummary
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetLatestApplyTime() *int64 {
	return s.LatestApplyTime
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetLatestTaskId() *string {
	return s.LatestTaskId
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetProtectionPolicyId() *string {
	return s.ProtectionPolicyId
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetProtectionPolicyName() *string {
	return s.ProtectionPolicyName
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetProtectionPolicyRegionId() *string {
	return s.ProtectionPolicyRegionId
}

func (s *ListProtectionPoliciesResponseBodyDataContent) GetSubProtectionPolicies() []*ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies {
	return s.SubProtectionPolicies
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetBoundResourceCategoryIds(v []*string) *ListProtectionPoliciesResponseBodyDataContent {
	s.BoundResourceCategoryIds = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetLatestApplySummary(v *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) *ListProtectionPoliciesResponseBodyDataContent {
	s.LatestApplySummary = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetLatestApplyTime(v int64) *ListProtectionPoliciesResponseBodyDataContent {
	s.LatestApplyTime = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetLatestTaskId(v string) *ListProtectionPoliciesResponseBodyDataContent {
	s.LatestTaskId = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetProtectionPolicyId(v string) *ListProtectionPoliciesResponseBodyDataContent {
	s.ProtectionPolicyId = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetProtectionPolicyName(v string) *ListProtectionPoliciesResponseBodyDataContent {
	s.ProtectionPolicyName = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetProtectionPolicyRegionId(v string) *ListProtectionPoliciesResponseBodyDataContent {
	s.ProtectionPolicyRegionId = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) SetSubProtectionPolicies(v []*ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) *ListProtectionPoliciesResponseBodyDataContent {
	s.SubProtectionPolicies = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContent) Validate() error {
	if s.LatestApplySummary != nil {
		if err := s.LatestApplySummary.Validate(); err != nil {
			return err
		}
	}
	if s.SubProtectionPolicies != nil {
		for _, item := range s.SubProtectionPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectionPoliciesResponseBodyDataContentLatestApplySummary struct {
	// The count statistics of application status.
	ApplyStatusCount []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount `json:"ApplyStatusCount,omitempty" xml:"ApplyStatusCount,omitempty" type:"Repeated"`
	// The time when the task was completed. Unix timestamp format, in seconds.
	//
	// example:
	//
	// 1742167218
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// The count of resources by type.
	ResourceCount []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty" type:"Repeated"`
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) GetApplyStatusCount() []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount {
	return s.ApplyStatusCount
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) GetCompleteTime() *int64 {
	return s.CompleteTime
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) GetResourceCount() []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount {
	return s.ResourceCount
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) SetApplyStatusCount(v []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary {
	s.ApplyStatusCount = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) SetCompleteTime(v int64) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary {
	s.CompleteTime = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) SetResourceCount(v []*ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary {
	s.ResourceCount = v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummary) Validate() error {
	if s.ApplyStatusCount != nil {
		for _, item := range s.ApplyStatusCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResourceCount != nil {
		for _, item := range s.ResourceCount {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount struct {
	// The application status.
	//
	// example:
	//
	// FAILED
	ApplyStatus *string `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	// The count of resources by type.
	//
	// example:
	//
	// 3
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) GetApplyStatus() *string {
	return s.ApplyStatus
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) GetCount() *int64 {
	return s.Count
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) SetApplyStatus(v string) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount {
	s.ApplyStatus = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) SetCount(v int64) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount {
	s.Count = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryApplyStatusCount) Validate() error {
	return dara.Validate(s)
}

type ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount struct {
	// The count of resources by type.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::OTS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) GetCount() *int64 {
	return s.Count
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) SetCount(v int64) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount {
	s.Count = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) SetResourceType(v string) *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount {
	s.ResourceType = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentLatestApplySummaryResourceCount) Validate() error {
	return dara.Validate(s)
}

type ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies struct {
	// The sub-protection policy configuration.
	//
	// example:
	//
	// {\\"autoSnapshotPolicyId\\":\\"sp-123***7890\\"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The sub-protection policy type.
	//
	// example:
	//
	// ECS_AUTO_SNAPSHOT_POLICY
	SubProtectionPolicyType *string `json:"SubProtectionPolicyType,omitempty" xml:"SubProtectionPolicyType,omitempty"`
}

func (s ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) GoString() string {
	return s.String()
}

func (s *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) GetConfig() *string {
	return s.Config
}

func (s *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) SetConfig(v string) *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies {
	s.Config = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) SetSubProtectionPolicyType(v string) *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *ListProtectionPoliciesResponseBodyDataContentSubProtectionPolicies) Validate() error {
	return dara.Validate(s)
}
