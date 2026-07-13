// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPolicyApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProtectionPolicyApplicationsResponseBodyData) *ListProtectionPolicyApplicationsResponseBody
	GetData() *ListProtectionPolicyApplicationsResponseBodyData
	SetRequestId(v string) *ListProtectionPolicyApplicationsResponseBody
	GetRequestId() *string
}

type ListProtectionPolicyApplicationsResponseBody struct {
	// The response data.
	Data *ListProtectionPolicyApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 34081B20-C4C0-514F-93F6-8EEC3D1A587E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProtectionPolicyApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPolicyApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProtectionPolicyApplicationsResponseBody) GetData() *ListProtectionPolicyApplicationsResponseBodyData {
	return s.Data
}

func (s *ListProtectionPolicyApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProtectionPolicyApplicationsResponseBody) SetData(v *ListProtectionPolicyApplicationsResponseBodyData) *ListProtectionPolicyApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBody) SetRequestId(v string) *ListProtectionPolicyApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProtectionPolicyApplicationsResponseBodyData struct {
	// The response content.
	Content []*ListProtectionPolicyApplicationsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The maximum number of results to return.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for retrieving the next page of results. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// fb836242f4225fa0f0e0257362dfc6dd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProtectionPolicyApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPolicyApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) GetContent() []*ListProtectionPolicyApplicationsResponseBodyDataContent {
	return s.Content
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) SetContent(v []*ListProtectionPolicyApplicationsResponseBodyDataContent) *ListProtectionPolicyApplicationsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) SetMaxResults(v int32) *ListProtectionPolicyApplicationsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) SetNextToken(v string) *ListProtectionPolicyApplicationsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) SetTotalCount(v int64) *ListProtectionPolicyApplicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyData) Validate() error {
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

type ListProtectionPolicyApplicationsResponseBodyDataContent struct {
	// The error details, returned when the application fails.
	//
	// example:
	//
	// {"errorCode":"HbrPolicyNotFound","errorMessage":"po-123***7890"}
	ApplyDetail *string `json:"ApplyDetail,omitempty" xml:"ApplyDetail,omitempty"`
	// The policy application status.
	//
	// example:
	//
	// FAILED
	ApplyStatus *string `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	// The time the policy was applied.
	//
	// example:
	//
	// 2024-11-19T02:19:06Z
	ApplyTime *int64 `json:"ApplyTime,omitempty" xml:"ApplyTime,omitempty"`
	// The product type.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The protection policy ID.
	//
	// example:
	//
	// p-123***7890
	ProtectionPolicyId *string `json:"ProtectionPolicyId,omitempty" xml:"ProtectionPolicyId,omitempty"`
	// The resource ARN.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 123***7890
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The sub-protection policy type.
	//
	// example:
	//
	// ECS_AUTO_SNAPSHOT_POLICY
	SubProtectionPolicyType *string `json:"SubProtectionPolicyType,omitempty" xml:"SubProtectionPolicyType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListProtectionPolicyApplicationsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPolicyApplicationsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetApplyDetail() *string {
	return s.ApplyDetail
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetApplyStatus() *string {
	return s.ApplyStatus
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetProductType() *string {
	return s.ProductType
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetProtectionPolicyId() *string {
	return s.ProtectionPolicyId
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetResourceArn() *string {
	return s.ResourceArn
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) GetTaskId() *string {
	return s.TaskId
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetApplyDetail(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ApplyDetail = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetApplyStatus(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ApplyStatus = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetApplyTime(v int64) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ApplyTime = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetProductType(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetProtectionPolicyId(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ProtectionPolicyId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetResourceArn(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetResourceId(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetResourceOwnerId(v int64) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetResourceType(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetSubProtectionPolicyType(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) SetTaskId(v string) *ListProtectionPolicyApplicationsResponseBodyDataContent {
	s.TaskId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
