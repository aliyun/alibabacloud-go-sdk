// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListStackInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStackInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStackInstancesResponseBody
	GetRequestId() *string
	SetStackInstances(v []*ListStackInstancesResponseBodyStackInstances) *ListStackInstancesResponseBody
	GetStackInstances() []*ListStackInstancesResponseBodyStackInstances
	SetTotalCount(v int32) *ListStackInstancesResponseBody
	GetTotalCount() *int32
}

type ListStackInstancesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 85DE34BD-7FF9-480F-8C21-556E9DA93ACD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stacks.
	StackInstances []*ListStackInstancesResponseBodyStackInstances `json:"StackInstances,omitempty" xml:"StackInstances,omitempty" type:"Repeated"`
	// The total number of stacks.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStackInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStackInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackInstancesResponseBody) GetStackInstances() []*ListStackInstancesResponseBodyStackInstances {
	return s.StackInstances
}

func (s *ListStackInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackInstancesResponseBody) SetPageNumber(v int32) *ListStackInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetPageSize(v int32) *ListStackInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetRequestId(v string) *ListStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetStackInstances(v []*ListStackInstancesResponseBodyStackInstances) *ListStackInstancesResponseBody {
	s.StackInstances = v
	return s
}

func (s *ListStackInstancesResponseBody) SetTotalCount(v int32) *ListStackInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackInstancesResponseBody) Validate() error {
	if s.StackInstances != nil {
		for _, item := range s.StackInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackInstancesResponseBodyStackInstances struct {
	// The ID of the destination account to which the stack belongs.
	//
	// example:
	//
	// 156552876021****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the last successful drift detection was performed on the stack.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The ID of the folder in the resource directory.
	//
	// > This parameter is returned only if the stack group is granted service-managed permissions.
	//
	// example:
	//
	// fd-4PvlVLOL8v
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the stack when the last successful drift detection was performed on the stack group.
	//
	// Valid values:
	//
	// 	- DRIFTED: The stack has drifted.
	//
	// 	- NOT_CHECKED: No successful drift detection is performed on the stack.
	//
	// 	- IN_SYNC: The stack is being synchronized.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	//
	// example:
	//
	// IN_SYNC
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack group.
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// > This parameter is returned only if the stack is in the CURRENT state.
	//
	// example:
	//
	// 35ad60e3-6a92-42d8-8812-f0700d45****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The state of the stack.
	//
	// Valid values:
	//
	// 	- CURRENT: The stack is up-to-date with the stack group.
	//
	// 	- OUTDATED: The stack is not up-to-date with the stack group. Stacks are in the OUTDATED state due to the following possible reasons:
	//
	//     	- When the CreateStackInstances operation is called to create stacks, the stacks fail to be created.
	//
	//     	- When the UpdateStackInstances or UpdateStackGroup operation is called to update stacks, the stacks fail to be updated, or only specific stacks are updated.
	//
	//     	- The creation or update operation is not complete.
	//
	// example:
	//
	// CURRENT
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack instance is in the OUTDATED state.
	//
	// > This parameter is returned only if the stack instance is in the OUTDATED state.
	//
	// example:
	//
	// User initiated stop
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackInstancesResponseBodyStackInstances) String() string {
	return dara.Prettify(s)
}

func (s ListStackInstancesResponseBodyStackInstances) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBodyStackInstances) GetAccountId() *string {
	return s.AccountId
}

func (s *ListStackInstancesResponseBodyStackInstances) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *ListStackInstancesResponseBodyStackInstances) GetRdFolderId() *string {
	return s.RdFolderId
}

func (s *ListStackInstancesResponseBodyStackInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStackDriftStatus() *string {
	return s.StackDriftStatus
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStackId() *string {
	return s.StackId
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStatus() *string {
	return s.Status
}

func (s *ListStackInstancesResponseBodyStackInstances) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListStackInstancesResponseBodyStackInstances) SetAccountId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.AccountId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetDriftDetectionTime(v string) *ListStackInstancesResponseBodyStackInstances {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetRdFolderId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.RdFolderId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetRegionId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.RegionId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackDriftStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupName(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.Status = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatusReason(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StatusReason = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) Validate() error {
	return dara.Validate(s)
}
