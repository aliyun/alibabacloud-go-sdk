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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on each page. Maximum value: 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Queries the resource pool list.
	PoolList []*ListPoolsResponseBodyPoolList `json:"PoolList,omitempty" xml:"PoolList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of list entries.
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
	// The name of the resource pool.
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
}

func (s ListPoolsResponseBodyPoolList) String() string {
	return dara.Prettify(s)
}

func (s ListPoolsResponseBodyPoolList) GoString() string {
	return s.String()
}

func (s *ListPoolsResponseBodyPoolList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListPoolsResponseBodyPoolList) GetMaxExectorNum() *int32 {
	return s.MaxExectorNum
}

func (s *ListPoolsResponseBodyPoolList) GetPoolName() *string {
	return s.PoolName
}

func (s *ListPoolsResponseBodyPoolList) GetPriority() *int32 {
	return s.Priority
}

func (s *ListPoolsResponseBodyPoolList) GetStatus() *string {
	return s.Status
}

func (s *ListPoolsResponseBodyPoolList) SetIsDefault(v bool) *ListPoolsResponseBodyPoolList {
	s.IsDefault = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) SetMaxExectorNum(v int32) *ListPoolsResponseBodyPoolList {
	s.MaxExectorNum = &v
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

func (s *ListPoolsResponseBodyPoolList) SetStatus(v string) *ListPoolsResponseBodyPoolList {
	s.Status = &v
	return s
}

func (s *ListPoolsResponseBodyPoolList) Validate() error {
	return dara.Validate(s)
}
