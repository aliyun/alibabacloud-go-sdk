// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChangeSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSets(v []*ListChangeSetsResponseBodyChangeSets) *ListChangeSetsResponseBody
	GetChangeSets() []*ListChangeSetsResponseBodyChangeSets
	SetPageNumber(v int32) *ListChangeSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListChangeSetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListChangeSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListChangeSetsResponseBody
	GetTotalCount() *int32
}

type ListChangeSetsResponseBody struct {
	// The change sets.
	ChangeSets []*ListChangeSetsResponseBodyChangeSets `json:"ChangeSets,omitempty" xml:"ChangeSets,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of change sets returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChangeSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChangeSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBody) GetChangeSets() []*ListChangeSetsResponseBodyChangeSets {
	return s.ChangeSets
}

func (s *ListChangeSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChangeSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChangeSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChangeSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListChangeSetsResponseBody) SetChangeSets(v []*ListChangeSetsResponseBodyChangeSets) *ListChangeSetsResponseBody {
	s.ChangeSets = v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageNumber(v int32) *ListChangeSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageSize(v int32) *ListChangeSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetRequestId(v string) *ListChangeSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetTotalCount(v int32) *ListChangeSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListChangeSetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListChangeSetsResponseBodyChangeSets struct {
	// The ID of the change set.
	//
	// example:
	//
	// 1f6521a4-05af-4975-afe9-bc4b45ad****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of the change set.
	//
	// example:
	//
	// MyChangeSet
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set.
	//
	// example:
	//
	// UPDATE
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The time when the change set was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-08-01T05:16:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the change set.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution status of the change set.
	//
	// example:
	//
	// AVAILABLE
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The region ID of the change set.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the stack with which the change set is associated.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack with which the change set is associated.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The status of the change set.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the change set is in its current state.
	//
	// example:
	//
	// too many changes
	StatusReason *string                                     `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	Tags         []*ListChangeSetsResponseBodyChangeSetsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListChangeSetsResponseBodyChangeSets) String() string {
	return dara.Prettify(s)
}

func (s ListChangeSetsResponseBodyChangeSets) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBodyChangeSets) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *ListChangeSetsResponseBodyChangeSets) GetChangeSetName() *string {
	return s.ChangeSetName
}

func (s *ListChangeSetsResponseBodyChangeSets) GetChangeSetType() *string {
	return s.ChangeSetType
}

func (s *ListChangeSetsResponseBodyChangeSets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListChangeSetsResponseBodyChangeSets) GetDescription() *string {
	return s.Description
}

func (s *ListChangeSetsResponseBodyChangeSets) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListChangeSetsResponseBodyChangeSets) GetRegionId() *string {
	return s.RegionId
}

func (s *ListChangeSetsResponseBodyChangeSets) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListChangeSetsResponseBodyChangeSets) GetStackId() *string {
	return s.StackId
}

func (s *ListChangeSetsResponseBodyChangeSets) GetStackName() *string {
	return s.StackName
}

func (s *ListChangeSetsResponseBodyChangeSets) GetStatus() *string {
	return s.Status
}

func (s *ListChangeSetsResponseBodyChangeSets) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListChangeSetsResponseBodyChangeSets) GetTags() []*ListChangeSetsResponseBodyChangeSetsTags {
	return s.Tags
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetType(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetType = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetCreateTime(v string) *ListChangeSetsResponseBodyChangeSets {
	s.CreateTime = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetDescription(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Description = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetExecutionStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ExecutionStatus = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetRegionId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.RegionId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetResourceGroupId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ResourceGroupId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Status = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatusReason(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StatusReason = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetTags(v []*ListChangeSetsResponseBodyChangeSetsTags) *ListChangeSetsResponseBodyChangeSets {
	s.Tags = v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) Validate() error {
	return dara.Validate(s)
}

type ListChangeSetsResponseBodyChangeSetsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListChangeSetsResponseBodyChangeSetsTags) String() string {
	return dara.Prettify(s)
}

func (s ListChangeSetsResponseBodyChangeSetsTags) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBodyChangeSetsTags) GetKey() *string {
	return s.Key
}

func (s *ListChangeSetsResponseBodyChangeSetsTags) GetValue() *string {
	return s.Value
}

func (s *ListChangeSetsResponseBodyChangeSetsTags) SetKey(v string) *ListChangeSetsResponseBodyChangeSetsTags {
	s.Key = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSetsTags) SetValue(v string) *ListChangeSetsResponseBodyChangeSetsTags {
	s.Value = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSetsTags) Validate() error {
	return dara.Validate(s)
}
