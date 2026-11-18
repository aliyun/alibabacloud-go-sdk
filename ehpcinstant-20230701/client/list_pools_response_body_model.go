// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPoolsResponseBody
	GetPageSize() *int32
	SetPoolList(v []*ListPoolsResponseBodyPoolList) *ListPoolsResponseBody
	GetPoolList() []*ListPoolsResponseBodyPoolList
	SetRequestId(v string) *ListPoolsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListPoolsResponseBody
	GetTotalCount() *int32
}

type ListPoolsResponseBody struct {
	// The current page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paging query. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of resource pools.
	PoolList []*ListPoolsResponseBodyPoolList `json:"PoolList,omitempty" xml:"PoolList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoolsResponseBody) GetPoolList() []*ListPoolsResponseBodyPoolList {
	return s.PoolList
}

func (s *ListPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoolsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPoolsResponseBody) SetPageNumber(v int32) *ListPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPoolsResponseBody) SetPageSize(v int32) *ListPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPoolsResponseBody) SetPoolList(v []*ListPoolsResponseBodyPoolList) *ListPoolsResponseBody {
	s.PoolList = v
	return s
}

func (s *ListPoolsResponseBody) SetRequestId(v string) *ListPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoolsResponseBody) SetTotalCount(v int32) *ListPoolsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPoolsResponseBody) Validate() error {
	if s.PoolList != nil {
		for _, item := range s.PoolList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPoolsResponseBodyPoolList struct {
	// The time when the resource pool was created. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Specifies whether the resource pool is the default resource pool. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The number of running executor nodes in a resource pool.
	//
	// example:
	//
	// 100
	MaxExecutorNum *int32 `json:"MaxExecutorNum,omitempty" xml:"MaxExecutorNum,omitempty"`
	// The resource pool name.
	//
	// - The name can be up to 15 characters in length.
	//
	// - The name can contain digits, uppercase letters, lowercase letters, underscores (_), and periods (.).
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// - Valid values: 1 to 99. Default value: 1, which indicates the lowest priority.
	//
	// - Jobs submitted to a resource pool with a higher priority level are scheduled before pending jobs in a resource pool with a lower priority level. The priority level of a resource pool takes precedence over the priority of a job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the scheduling policy.
	//
	// example:
	//
	// policy-xxx
	SchedulingPolicyId *string `json:"SchedulingPolicyId,omitempty" xml:"SchedulingPolicyId,omitempty"`
	// The resource pool status. Valid values:
	//
	// - Creating: The resource pool is being created.
	//
	// - Updating: The resource pool is being updated.
	//
	// - Deleting: The resource pool is being deleted.
	//
	// - Working: The resource pool is working.
	//
	// - Deleted: The resource pool is deleted.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	Tags []*ListPoolsResponseBodyPoolListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the resource pool was updated. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2026-04-20 11:09:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPoolsResponseBodyPoolList) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBodyPoolList) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBodyPoolList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPoolsResponseBodyPoolList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListPoolsResponseBodyPoolList) GetMaxExecutorNum() *int32 {
	return s.MaxExecutorNum
}

func (s *ListPoolsResponseBodyPoolList) GetPoolName() *string {
	return s.PoolName
}

func (s *ListPoolsResponseBodyPoolList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPoolsResponseBodyPoolList) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *ListPoolsResponseBodyPoolList) GetStatus() *string {
	return s.Status
}

func (s *ListPoolsResponseBodyPoolList) GetTags() []*ListPoolsResponseBodyPoolListTags {
	return s.Tags
}

func (s *ListPoolsResponseBodyPoolList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPoolsResponseBodyPoolList) SetCreateTime(v string) *ListPoolsResponseBodyPoolList {
	s.CreateTime = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetIsDefault(v bool) *ListPoolsResponseBodyPoolList {
	s.IsDefault = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetMaxExecutorNum(v int32) *ListPoolsResponseBodyPoolList {
	s.MaxExecutorNum = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPoolName(v string) *ListPoolsResponseBodyPoolList {
	s.PoolName = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetPriority(v int32) *ListPoolsResponseBodyPoolList {
	s.Priority = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetSchedulingPolicyId(v string) *ListPoolsResponseBodyPoolList {
	s.SchedulingPolicyId = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetStatus(v string) *ListPoolsResponseBodyPoolList {
	s.Status = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetTags(v []*ListPoolsResponseBodyPoolListTags) *ListPoolsResponseBodyPoolList {
	s.Tags = v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetUpdateTime(v string) *ListPoolsResponseBodyPoolList {
	s.UpdateTime = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) Validate() error {
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

type ListPoolsResponseBodyPoolListTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPoolsResponseBodyPoolListTags) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBodyPoolListTags) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBodyPoolListTags) GetKey() *string {
	return s.Key
}

func (s *ListPoolsResponseBodyPoolListTags) GetValue() *string {
	return s.Value
}

func (s *ListPoolsResponseBodyPoolListTags) SetKey(v string) *ListPoolsResponseBodyPoolListTags {
	s.Key = &v
	return s
}

func (s *ListPoolsResponseBodyPoolListTags) SetValue(v string) *ListPoolsResponseBodyPoolListTags {
	s.Value = &v
	return s
}

func (s *ListPoolsResponseBodyPoolListTags) Validate() error {
	return dara.Validate(s)
}
