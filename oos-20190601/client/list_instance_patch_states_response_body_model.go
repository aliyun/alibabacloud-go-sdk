// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancePatchStatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstancePatchStates(v []*ListInstancePatchStatesResponseBodyInstancePatchStates) *ListInstancePatchStatesResponseBody
	GetInstancePatchStates() []*ListInstancePatchStatesResponseBodyInstancePatchStates
	SetMaxResults(v int32) *ListInstancePatchStatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListInstancePatchStatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstancePatchStatesResponseBody
	GetRequestId() *string
}

type ListInstancePatchStatesResponseBody struct {
	// The details of patches of the instance.
	InstancePatchStates []*ListInstancePatchStatesResponseBodyInstancePatchStates `json:"InstancePatchStates,omitempty" xml:"InstancePatchStates,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9A47C086-E64D-52EE-8B2C-EFD23877C55E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePatchStatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBody) GetInstancePatchStates() []*ListInstancePatchStatesResponseBodyInstancePatchStates {
	return s.InstancePatchStates
}

func (s *ListInstancePatchStatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstancePatchStatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstancePatchStatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancePatchStatesResponseBody) SetInstancePatchStates(v []*ListInstancePatchStatesResponseBodyInstancePatchStates) *ListInstancePatchStatesResponseBody {
	s.InstancePatchStates = v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetMaxResults(v int32) *ListInstancePatchStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetNextToken(v string) *ListInstancePatchStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetRequestId(v string) *ListInstancePatchStatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) Validate() error {
	if s.InstancePatchStates != nil {
		for _, item := range s.InstancePatchStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancePatchStatesResponseBodyInstancePatchStates struct {
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-f9393021b7a049e1b34e
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The number of patches that failed to be installed.
	//
	// example:
	//
	// 0
	FailedCount *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The number of installed patches.
	//
	// example:
	//
	// 0
	InstalledCount *string `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
	// The number of patches that do not meet the baseline.
	//
	// example:
	//
	// 0
	InstalledOtherCount *string `json:"InstalledOtherCount,omitempty" xml:"InstalledOtherCount,omitempty"`
	// The number of patches that have been installed but require a restart to take effect.
	//
	// example:
	//
	// 0
	InstalledPendingRebootCount *string `json:"InstalledPendingRebootCount,omitempty" xml:"InstalledPendingRebootCount,omitempty"`
	// The number of patches that are rejected by the user.
	//
	// example:
	//
	// 0
	InstalledRejectedCount *string `json:"InstalledRejectedCount,omitempty" xml:"InstalledRejectedCount,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of patches that are not installed.
	//
	// example:
	//
	// 0
	MissingCount *string `json:"MissingCount,omitempty" xml:"MissingCount,omitempty"`
	// The time when the operation ended.
	//
	// example:
	//
	// 2021-09-10T11:42:22Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The time when the operation was initiated.
	//
	// example:
	//
	// 2021-09-10T11:42:22Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The operation type.
	//
	// example:
	//
	// scan
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The information about the user.
	//
	// example:
	//
	// “”
	OwnerInformation *string `json:"OwnerInformation,omitempty" xml:"OwnerInformation,omitempty"`
	// The patch group.
	//
	// example:
	//
	// null
	PatchGroup *string `json:"PatchGroup,omitempty" xml:"PatchGroup,omitempty"`
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) String() string {
	return dara.Prettify(s)
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetBaselineId() *string {
	return s.BaselineId
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetFailedCount() *string {
	return s.FailedCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetInstalledCount() *string {
	return s.InstalledCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetInstalledOtherCount() *string {
	return s.InstalledOtherCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetInstalledPendingRebootCount() *string {
	return s.InstalledPendingRebootCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetInstalledRejectedCount() *string {
	return s.InstalledRejectedCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetMissingCount() *string {
	return s.MissingCount
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetOperationEndTime() *string {
	return s.OperationEndTime
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetOperationStartTime() *string {
	return s.OperationStartTime
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetOperationType() *string {
	return s.OperationType
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetOwnerInformation() *string {
	return s.OwnerInformation
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) GetPatchGroup() *string {
	return s.PatchGroup
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetBaselineId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.BaselineId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetFailedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.FailedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledOtherCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledOtherCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledPendingRebootCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledPendingRebootCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledRejectedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledRejectedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstanceId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetMissingCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.MissingCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationEndTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationEndTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationStartTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationStartTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationType(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationType = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOwnerInformation(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OwnerInformation = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetPatchGroup(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.PatchGroup = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) Validate() error {
	return dara.Validate(s)
}
