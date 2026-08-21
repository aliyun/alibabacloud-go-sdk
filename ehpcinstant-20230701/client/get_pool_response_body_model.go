// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPoolInfo(v *GetPoolResponseBodyPoolInfo) *GetPoolResponseBody
	GetPoolInfo() *GetPoolResponseBodyPoolInfo
	SetRequestId(v string) *GetPoolResponseBody
	GetRequestId() *string
}

type GetPoolResponseBody struct {
	// The resource pool information.
	PoolInfo *GetPoolResponseBodyPoolInfo `json:"PoolInfo,omitempty" xml:"PoolInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponseBody) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBody) GetPoolInfo() *GetPoolResponseBodyPoolInfo {
	return s.PoolInfo
}

func (s *GetPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPoolResponseBody) SetPoolInfo(v *GetPoolResponseBodyPoolInfo) *GetPoolResponseBody {
	s.PoolInfo = v
	return s
}

func (s *GetPoolResponseBody) SetRequestId(v string) *GetPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPoolResponseBody) Validate() error {
	if s.PoolInfo != nil {
		if err := s.PoolInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPoolResponseBodyPoolInfo struct {
	// The time when the resource pool was created.
	//
	// example:
	//
	// 2024-12-01 20:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the resource pool creator.
	//
	// example:
	//
	// 200428053788xxxx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The number of executor nodes that are currently running in the resource pool.
	//
	// example:
	//
	// 100
	ExecutorUsage *int32 `json:"ExecutorUsage,omitempty" xml:"ExecutorUsage,omitempty"`
	// Indicates whether the resource pool is the default resource pool. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The maximum number of executor nodes that can run concurrently in the resource pool.
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
	// - Jobs submitted to a resource pool with a higher priority value are scheduled before pending jobs in a resource pool with a lower priority value. The resource pool priority takes precedence over the job priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The error reason.
	//
	// example:
	//
	// Fails to **	- pool: ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The scheduling policy ID.
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
	// - Working: The resource pool is running.
	//
	// - Deleted: The resource pool has been deleted.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag information.
	Tags []*GetPoolResponseBodyPoolInfoTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The time when the resource pool was last updated.
	//
	// example:
	//
	// 2024-12-01 20:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetPoolResponseBodyPoolInfo) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponseBodyPoolInfo) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBodyPoolInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetPoolResponseBodyPoolInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetPoolResponseBodyPoolInfo) GetExecutorUsage() *int32 {
	return s.ExecutorUsage
}

func (s *GetPoolResponseBodyPoolInfo) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetPoolResponseBodyPoolInfo) GetMaxExecutorNum() *int32 {
	return s.MaxExecutorNum
}

func (s *GetPoolResponseBodyPoolInfo) GetPoolName() *string {
	return s.PoolName
}

func (s *GetPoolResponseBodyPoolInfo) GetPriority() *int32 {
	return s.Priority
}

func (s *GetPoolResponseBodyPoolInfo) GetReason() *string {
	return s.Reason
}

func (s *GetPoolResponseBodyPoolInfo) GetSchedulingPolicyId() *string {
	return s.SchedulingPolicyId
}

func (s *GetPoolResponseBodyPoolInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPoolResponseBodyPoolInfo) GetTags() []*GetPoolResponseBodyPoolInfoTags {
	return s.Tags
}

func (s *GetPoolResponseBodyPoolInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPoolResponseBodyPoolInfo) SetCreateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetCreator(v string) *GetPoolResponseBodyPoolInfo {
	s.Creator = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetExecutorUsage(v int32) *GetPoolResponseBodyPoolInfo {
	s.ExecutorUsage = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetIsDefault(v bool) *GetPoolResponseBodyPoolInfo {
	s.IsDefault = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetMaxExecutorNum(v int32) *GetPoolResponseBodyPoolInfo {
	s.MaxExecutorNum = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPoolName(v string) *GetPoolResponseBodyPoolInfo {
	s.PoolName = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetPriority(v int32) *GetPoolResponseBodyPoolInfo {
	s.Priority = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetReason(v string) *GetPoolResponseBodyPoolInfo {
	s.Reason = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetSchedulingPolicyId(v string) *GetPoolResponseBodyPoolInfo {
	s.SchedulingPolicyId = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetStatus(v string) *GetPoolResponseBodyPoolInfo {
	s.Status = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetTags(v []*GetPoolResponseBodyPoolInfoTags) *GetPoolResponseBodyPoolInfo {
	s.Tags = v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetUpdateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) Validate() error {
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

type GetPoolResponseBodyPoolInfoTags struct {
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

func (s GetPoolResponseBodyPoolInfoTags) String() string {
	return dara.Prettify(s)
}

func (s GetPoolResponseBodyPoolInfoTags) GoString() string {
	return s.String()
}

func (s *GetPoolResponseBodyPoolInfoTags) GetKey() *string {
	return s.Key
}

func (s *GetPoolResponseBodyPoolInfoTags) GetValue() *string {
	return s.Value
}

func (s *GetPoolResponseBodyPoolInfoTags) SetKey(v string) *GetPoolResponseBodyPoolInfoTags {
	s.Key = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfoTags) SetValue(v string) *GetPoolResponseBodyPoolInfoTags {
	s.Value = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfoTags) Validate() error {
	return dara.Validate(s)
}
