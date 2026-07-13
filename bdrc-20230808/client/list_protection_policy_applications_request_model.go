// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProtectionPolicyApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyStatus(v string) *ListProtectionPolicyApplicationsRequest
	GetApplyStatus() *string
	SetMaxResults(v int32) *ListProtectionPolicyApplicationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListProtectionPolicyApplicationsRequest
	GetNextToken() *string
	SetResourceType(v string) *ListProtectionPolicyApplicationsRequest
	GetResourceType() *string
	SetSubProtectionPolicyType(v string) *ListProtectionPolicyApplicationsRequest
	GetSubProtectionPolicyType() *string
	SetTaskId(v string) *ListProtectionPolicyApplicationsRequest
	GetTaskId() *string
}

type ListProtectionPolicyApplicationsRequest struct {
	// The application status.
	//
	// example:
	//
	// FAILED
	ApplyStatus *string `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	// The maximum number of results to return in a single page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. The response returns a `NextToken` value only when more results are available. To retrieve the next page, include the `NextToken` from the previous response in your request. If the response does not include a `NextToken` value, all results have been retrieved.
	//
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// The task ID. You can call the DescribeTasks operation to query task IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-0004d9ctt1ii********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ListProtectionPolicyApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProtectionPolicyApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListProtectionPolicyApplicationsRequest) GetApplyStatus() *string {
	return s.ApplyStatus
}

func (s *ListProtectionPolicyApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProtectionPolicyApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProtectionPolicyApplicationsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProtectionPolicyApplicationsRequest) GetSubProtectionPolicyType() *string {
	return s.SubProtectionPolicyType
}

func (s *ListProtectionPolicyApplicationsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListProtectionPolicyApplicationsRequest) SetApplyStatus(v string) *ListProtectionPolicyApplicationsRequest {
	s.ApplyStatus = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) SetMaxResults(v int32) *ListProtectionPolicyApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) SetNextToken(v string) *ListProtectionPolicyApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) SetResourceType(v string) *ListProtectionPolicyApplicationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) SetSubProtectionPolicyType(v string) *ListProtectionPolicyApplicationsRequest {
	s.SubProtectionPolicyType = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) SetTaskId(v string) *ListProtectionPolicyApplicationsRequest {
	s.TaskId = &v
	return s
}

func (s *ListProtectionPolicyApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
