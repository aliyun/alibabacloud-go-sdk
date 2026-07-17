// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListAppGroupsRequest
	GetInstanceId() *string
	SetName(v string) *ListAppGroupsRequest
	GetName() *string
	SetPageNumber(v int32) *ListAppGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppGroupsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListAppGroupsRequest
	GetResourceGroupId() *string
	SetSortBy(v int32) *ListAppGroupsRequest
	GetSortBy() *int32
	SetTags(v []*ListAppGroupsRequestTags) *ListAppGroupsRequest
	GetTags() []*ListAppGroupsRequestTags
	SetType(v string) *ListAppGroupsRequest
	GetType() *string
}

type ListAppGroupsRequest struct {
	// The ID of the instance. An exact match is performed.
	//
	// example:
	//
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The application name.
	//
	// example:
	//
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// "110123123"
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The sort order. Valid values:
	//
	// - 0: Sorts applications by creation time in descending order.
	//
	// - 1: Sorts applications by modification time in descending order.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	SortBy *int32 `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// A list of tags.
	Tags []*ListAppGroupsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// - standard: a Standard Edition application.
	//
	// - enhanced: a Premium Edition application.
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAppGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAppGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListAppGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAppGroupsRequest) GetSortBy() *int32 {
	return s.SortBy
}

func (s *ListAppGroupsRequest) GetTags() []*ListAppGroupsRequestTags {
	return s.Tags
}

func (s *ListAppGroupsRequest) GetType() *string {
	return s.Type
}

func (s *ListAppGroupsRequest) SetInstanceId(v string) *ListAppGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsRequest) SetName(v string) *ListAppGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageNumber(v int32) *ListAppGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageSize(v int32) *ListAppGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupsRequest) SetResourceGroupId(v string) *ListAppGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsRequest) SetSortBy(v int32) *ListAppGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAppGroupsRequest) SetTags(v []*ListAppGroupsRequestTags) *ListAppGroupsRequest {
	s.Tags = v
	return s
}

func (s *ListAppGroupsRequest) SetType(v string) *ListAppGroupsRequest {
	s.Type = &v
	return s
}

func (s *ListAppGroupsRequest) Validate() error {
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

type ListAppGroupsRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bar
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListAppGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListAppGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListAppGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListAppGroupsRequestTags) SetKey(v string) *ListAppGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListAppGroupsRequestTags) SetValue(v string) *ListAppGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *ListAppGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
