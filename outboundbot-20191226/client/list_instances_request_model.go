// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListInstancesRequest
	GetName() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest
	GetTag() []*ListInstancesRequestTag
}

type ListInstancesRequest struct {
	// Name of the Intelligent outbound call instance
	//
	// example:
	//
	// xxx
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// > You can obtain this from the Resource Management document via the source API. For more information, see: https://api.aliyun.com/document/ResourceManager/2020-03-31/ListResourceGroups
	//
	// example:
	//
	// rg-acfm3iugit3uw7a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// List of tags of business instances.
	//
	// >You can obtain it by calling the ListResourceTags API.
	Tag []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetName() *string {
	return s.Name
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetTag() []*ListInstancesRequestTag {
	return s.Tag
}

func (s *ListInstancesRequest) SetName(v string) *ListInstancesRequest {
	s.Name = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesRequestTag struct {
	// Tag key of the instance.
	//
	// example:
	//
	// name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value of the instance.
	//
	// example:
	//
	// xxx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *ListInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
