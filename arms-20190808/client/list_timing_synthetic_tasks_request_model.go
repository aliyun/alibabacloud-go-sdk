// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTimingSyntheticTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListTimingSyntheticTasksRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListTimingSyntheticTasksRequest
	GetResourceGroupId() *string
	SetSearch(v *ListTimingSyntheticTasksRequestSearch) *ListTimingSyntheticTasksRequest
	GetSearch() *ListTimingSyntheticTasksRequestSearch
	SetTags(v []*ListTimingSyntheticTasksRequestTags) *ListTimingSyntheticTasksRequest
	GetTags() []*ListTimingSyntheticTasksRequestTags
}

type ListTimingSyntheticTasksRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The search keyword.
	Search *ListTimingSyntheticTasksRequestSearch `json:"Search,omitempty" xml:"Search,omitempty" type:"Struct"`
	// The tags.
	Tags []*ListTimingSyntheticTasksRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTimingSyntheticTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTimingSyntheticTasksRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTimingSyntheticTasksRequest) GetSearch() *ListTimingSyntheticTasksRequestSearch {
	return s.Search
}

func (s *ListTimingSyntheticTasksRequest) GetTags() []*ListTimingSyntheticTasksRequestTags {
	return s.Tags
}

func (s *ListTimingSyntheticTasksRequest) SetRegionId(v string) *ListTimingSyntheticTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListTimingSyntheticTasksRequest) SetResourceGroupId(v string) *ListTimingSyntheticTasksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTimingSyntheticTasksRequest) SetSearch(v *ListTimingSyntheticTasksRequestSearch) *ListTimingSyntheticTasksRequest {
	s.Search = v
	return s
}

func (s *ListTimingSyntheticTasksRequest) SetTags(v []*ListTimingSyntheticTasksRequestTags) *ListTimingSyntheticTasksRequest {
	s.Tags = v
	return s
}

func (s *ListTimingSyntheticTasksRequest) Validate() error {
	if s.Search != nil {
		if err := s.Search.Validate(); err != nil {
			return err
		}
	}
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

type ListTimingSyntheticTasksRequestSearch struct {
	// The task name.
	//
	// example:
	//
	// AlibabaCloud DNS Task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order by which tasks are sorted. 1: ascending order. -1: descending order.
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// The condition by which tasks are sorted. You can sort tasks by gmtCreate, gmtModified, status, or monitorCount.
	//
	// example:
	//
	// status
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The page number. This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task status. CREATING: The task is being created. RUNNING: The task is running. PARTIAL_RUNNING: The task is partially running. STOP: The task is stopped. LIMIT_STOP: The task is stopped due to quota limit. EXCEPTION: The task is abnormal. DELETE: The task is deleted. DELETE_EXCEPTION: An exception occurs while deleting the task.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task IDs.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
	// The task types.
	TaskTypes []*int32 `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListTimingSyntheticTasksRequestSearch) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksRequestSearch) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksRequestSearch) GetName() *string {
	return s.Name
}

func (s *ListTimingSyntheticTasksRequestSearch) GetOrder() *int32 {
	return s.Order
}

func (s *ListTimingSyntheticTasksRequestSearch) GetOrderField() *string {
	return s.OrderField
}

func (s *ListTimingSyntheticTasksRequestSearch) GetPage() *int32 {
	return s.Page
}

func (s *ListTimingSyntheticTasksRequestSearch) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTimingSyntheticTasksRequestSearch) GetStatus() *string {
	return s.Status
}

func (s *ListTimingSyntheticTasksRequestSearch) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *ListTimingSyntheticTasksRequestSearch) GetTaskTypes() []*int32 {
	return s.TaskTypes
}

func (s *ListTimingSyntheticTasksRequestSearch) SetName(v string) *ListTimingSyntheticTasksRequestSearch {
	s.Name = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetOrder(v int32) *ListTimingSyntheticTasksRequestSearch {
	s.Order = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetOrderField(v string) *ListTimingSyntheticTasksRequestSearch {
	s.OrderField = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetPage(v int32) *ListTimingSyntheticTasksRequestSearch {
	s.Page = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetPageSize(v int32) *ListTimingSyntheticTasksRequestSearch {
	s.PageSize = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetStatus(v string) *ListTimingSyntheticTasksRequestSearch {
	s.Status = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetTaskIds(v []*string) *ListTimingSyntheticTasksRequestSearch {
	s.TaskIds = v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) SetTaskTypes(v []*int32) *ListTimingSyntheticTasksRequestSearch {
	s.TaskTypes = v
	return s
}

func (s *ListTimingSyntheticTasksRequestSearch) Validate() error {
	return dara.Validate(s)
}

type ListTimingSyntheticTasksRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// mark
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1111
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTimingSyntheticTasksRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTimingSyntheticTasksRequestTags) GoString() string {
	return s.String()
}

func (s *ListTimingSyntheticTasksRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTimingSyntheticTasksRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTimingSyntheticTasksRequestTags) SetKey(v string) *ListTimingSyntheticTasksRequestTags {
	s.Key = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestTags) SetValue(v string) *ListTimingSyntheticTasksRequestTags {
	s.Value = &v
	return s
}

func (s *ListTimingSyntheticTasksRequestTags) Validate() error {
	return dara.Validate(s)
}
