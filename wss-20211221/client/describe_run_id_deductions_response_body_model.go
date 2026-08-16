// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRunIdDeductionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeductions(v []*DescribeRunIdDeductionsResponseBodyDeductions) *DescribeRunIdDeductionsResponseBody
	GetDeductions() []*DescribeRunIdDeductionsResponseBodyDeductions
	SetMaxResults(v int32) *DescribeRunIdDeductionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRunIdDeductionsResponseBody
	GetNextToken() *string
	SetPageNum(v int32) *DescribeRunIdDeductionsResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeRunIdDeductionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRunIdDeductionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeRunIdDeductionsResponseBody
	GetTotalCount() *int64
	SetTotalUsedTime(v int64) *DescribeRunIdDeductionsResponseBody
	GetTotalUsedTime() *int64
	SetTotalUsedTimeDecimal(v string) *DescribeRunIdDeductionsResponseBody
	GetTotalUsedTimeDecimal() *string
}

type DescribeRunIdDeductionsResponseBody struct {
	// The deduction details.
	Deductions []*DescribeRunIdDeductionsResponseBodyDeductions `json:"Deductions,omitempty" xml:"Deductions,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Leave this parameter empty for the first request. For subsequent requests, use the `nextToken` value from the previous response.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68BD3312-53D8-123E-BB32-1A9F25E07A03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of core-hour package deduction details in the query result.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total usage duration. Unit: seconds. Do not use this field for AI scenarios.
	//
	// example:
	//
	// 100000
	TotalUsedTime *int64 `json:"TotalUsedTime,omitempty" xml:"TotalUsedTime,omitempty"`
	// The total credits used that match the specified conditions.
	//
	// example:
	//
	// 1.23
	TotalUsedTimeDecimal *string `json:"TotalUsedTimeDecimal,omitempty" xml:"TotalUsedTimeDecimal,omitempty"`
}

func (s DescribeRunIdDeductionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunIdDeductionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRunIdDeductionsResponseBody) GetDeductions() []*DescribeRunIdDeductionsResponseBodyDeductions {
	return s.Deductions
}

func (s *DescribeRunIdDeductionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRunIdDeductionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRunIdDeductionsResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeRunIdDeductionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRunIdDeductionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRunIdDeductionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeRunIdDeductionsResponseBody) GetTotalUsedTime() *int64 {
	return s.TotalUsedTime
}

func (s *DescribeRunIdDeductionsResponseBody) GetTotalUsedTimeDecimal() *string {
	return s.TotalUsedTimeDecimal
}

func (s *DescribeRunIdDeductionsResponseBody) SetDeductions(v []*DescribeRunIdDeductionsResponseBodyDeductions) *DescribeRunIdDeductionsResponseBody {
	s.Deductions = v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetMaxResults(v int32) *DescribeRunIdDeductionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetNextToken(v string) *DescribeRunIdDeductionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetPageNum(v int32) *DescribeRunIdDeductionsResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetPageSize(v int32) *DescribeRunIdDeductionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetRequestId(v string) *DescribeRunIdDeductionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetTotalCount(v int64) *DescribeRunIdDeductionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetTotalUsedTime(v int64) *DescribeRunIdDeductionsResponseBody {
	s.TotalUsedTime = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) SetTotalUsedTimeDecimal(v string) *DescribeRunIdDeductionsResponseBody {
	s.TotalUsedTimeDecimal = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBody) Validate() error {
	if s.Deductions != nil {
		for _, item := range s.Deductions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRunIdDeductionsResponseBodyDeductions struct {
	// The agent type: `CREDIT_PACKAGE` / `JVS_CLAW` / `OPEN_CLAW` / `JVS_COPILOT`.
	//
	// example:
	//
	// OpenClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The end time of the period.
	//
	// example:
	//
	// 2024-07-31T03:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The group resource type.
	//
	// example:
	//
	// GROUP_CREDIT_PACKAGE、GROUP_BUSINESS_CREDIT_PACKAGE
	GroupResourceType *string `json:"GroupResourceType,omitempty" xml:"GroupResourceType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// jvs-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Model      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The credit or plan package ID.
	//
	// example:
	//
	// crp-xxx
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// sunwyic.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// CreditPackage、BusinessCreditPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The unique run ID.
	//
	// example:
	//
	// run-szwB1fYHCTocjGkFAIf6V8A
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2025-12-16T02:10:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The summary of the large language model call.
	//
	// example:
	//
	// How is the weather today
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The usage duration. Unit: seconds. Do not use this field for AI scenarios.
	//
	// example:
	//
	// 360000000
	UsedTime *int64 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The credits used.
	//
	// example:
	//
	// 1.23
	UsedTimeDecimal *string `json:"UsedTimeDecimal,omitempty" xml:"UsedTimeDecimal,omitempty"`
}

func (s DescribeRunIdDeductionsResponseBodyDeductions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRunIdDeductionsResponseBodyDeductions) GoString() string {
	return s.String()
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetGroupResourceType() *string {
	return s.GroupResourceType
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetModel() *string {
	return s.Model
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetPackageId() *string {
	return s.PackageId
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetRunId() *string {
	return s.RunId
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetSummary() *string {
	return s.Summary
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetUsedTime() *int64 {
	return s.UsedTime
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) GetUsedTimeDecimal() *string {
	return s.UsedTimeDecimal
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetAgentType(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.AgentType = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetEndTime(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.EndTime = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetGroupResourceType(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.GroupResourceType = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetInstanceId(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.InstanceId = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetModel(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.Model = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetPackageId(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.PackageId = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetResourceId(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.ResourceId = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetResourceType(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.ResourceType = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetRunId(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.RunId = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetStartTime(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.StartTime = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetSummary(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.Summary = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetUsedTime(v int64) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.UsedTime = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) SetUsedTimeDecimal(v string) *DescribeRunIdDeductionsResponseBodyDeductions {
	s.UsedTimeDecimal = &v
	return s
}

func (s *DescribeRunIdDeductionsResponseBodyDeductions) Validate() error {
	return dara.Validate(s)
}
