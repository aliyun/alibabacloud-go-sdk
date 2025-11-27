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
	// The information about the resource pool.
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
	// The time when the resource pool is created.
	//
	// example:
	//
	// 2024-12-01 20:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The usage of execution nodes that are running in a resource pool.
	//
	// example:
	//
	// 1
	ExectorUsage *int32 `json:"ExectorUsage,omitempty" xml:"ExectorUsage,omitempty"`
	// Indices whether the resource pool is the default resource pool. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The maximum number of execution nodes that can run concurrently in a resource pool.
	//
	// example:
	//
	// 2000
	MaxExectorNum *int32 `json:"MaxExectorNum,omitempty" xml:"MaxExectorNum,omitempty"`
	// The name of the resource group.
	//
	// 	- The value can be up to 15 characters in length.
	//
	// 	- It can contain digits, uppercase letters, lowercase letters, underscores (_), and dots (.).
	//
	// example:
	//
	// PoolTest
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The priority of the resource pool.
	//
	// 	- You can set a priority in the range of 1 to 99. The default value is 1, which is the lowest priority.
	//
	// 	- Jobs submitted to a resource pool with a higher priority level value will be scheduled before pending jobs in a resource pool with a lower priority level value, and the priority level of the resource pool takes precedence over the priority of the job.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The cause of the error.
	//
	// example:
	//
	// Fails to **	- pool: ***.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the resource pool. Valid values:
	//
	// 	- Creating: The resource pool is being created.
	//
	// 	- Updating: The resource pool is being updated.
	//
	// 	- Deleting: The resource pool is being deleted.
	//
	// 	- Working: The resource pool is working.
	//
	// 	- Deleted: The resource pool is deleted.
	//
	// example:
	//
	// Working
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource pool was updated.
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

func (s *GetPoolResponseBodyPoolInfo) GetExectorUsage() *int32 {
	return s.ExectorUsage
}

func (s *GetPoolResponseBodyPoolInfo) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetPoolResponseBodyPoolInfo) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
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

func (s *GetPoolResponseBodyPoolInfo) GetStatus() *string {
	return s.Status
}

func (s *GetPoolResponseBodyPoolInfo) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetPoolResponseBodyPoolInfo) SetCreateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.CreateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetExectorUsage(v int32) *GetPoolResponseBodyPoolInfo {
	s.ExectorUsage = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetIsDefault(v bool) *GetPoolResponseBodyPoolInfo {
	s.IsDefault = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetMaxExectorNum(v int32) *GetPoolResponseBodyPoolInfo {
	s.MaxExectorNum = &v
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

func (s *GetPoolResponseBodyPoolInfo) SetStatus(v string) *GetPoolResponseBodyPoolInfo {
	s.Status = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) SetUpdateTime(v string) *GetPoolResponseBodyPoolInfo {
	s.UpdateTime = &v
	return s
}

func (s *GetPoolResponseBodyPoolInfo) Validate() error {
	return dara.Validate(s)
}
