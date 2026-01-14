// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAcceleratorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListBasicAcceleratorsRequest
	GetAcceleratorId() *string
	SetPageNumber(v int32) *ListBasicAcceleratorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBasicAcceleratorsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListBasicAcceleratorsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListBasicAcceleratorsRequest
	GetResourceGroupId() *string
	SetState(v string) *ListBasicAcceleratorsRequest
	GetState() *string
	SetTag(v []*ListBasicAcceleratorsRequestTag) *ListBasicAcceleratorsRequest
	GetTag() []*ListBasicAcceleratorsRequestTag
}

type ListBasicAcceleratorsRequest struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region to which the basic GA instance belongs. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the basic GA instance belongs.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the basic GA instance. Valid values:
	//
	// 	- **init**: The basic GA instance is being initialized.
	//
	// 	- **active**: The basic GA instance is available.
	//
	// 	- **configuring**: The basic GA instance is being configured.
	//
	// 	- **binding**: The basic GA instance is being associated.
	//
	// 	- **unbinding**: The basic GA instance is being disassociated.
	//
	// 	- **deleting**: The basic GA instance is being deleted.
	//
	// 	- **finacialLocked**: The basic GA instance is locked due to overdue payments.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the GA instance.
	Tag []*ListBasicAcceleratorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListBasicAcceleratorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsRequest) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicAcceleratorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBasicAcceleratorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBasicAcceleratorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBasicAcceleratorsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBasicAcceleratorsRequest) GetState() *string {
	return s.State
}

func (s *ListBasicAcceleratorsRequest) GetTag() []*ListBasicAcceleratorsRequestTag {
	return s.Tag
}

func (s *ListBasicAcceleratorsRequest) SetAcceleratorId(v string) *ListBasicAcceleratorsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetPageNumber(v int32) *ListBasicAcceleratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetPageSize(v int32) *ListBasicAcceleratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetRegionId(v string) *ListBasicAcceleratorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetResourceGroupId(v string) *ListBasicAcceleratorsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetState(v string) *ListBasicAcceleratorsRequest {
	s.State = &v
	return s
}

func (s *ListBasicAcceleratorsRequest) SetTag(v []*ListBasicAcceleratorsRequestTag) *ListBasicAcceleratorsRequest {
	s.Tag = v
	return s
}

func (s *ListBasicAcceleratorsRequest) Validate() error {
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

type ListBasicAcceleratorsRequestTag struct {
	// The key of tag N of the basic GA instance. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// Keytest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the basic GA instance. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// Valuetest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBasicAcceleratorsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsRequestTag) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListBasicAcceleratorsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListBasicAcceleratorsRequestTag) SetKey(v string) *ListBasicAcceleratorsRequestTag {
	s.Key = &v
	return s
}

func (s *ListBasicAcceleratorsRequestTag) SetValue(v string) *ListBasicAcceleratorsRequestTag {
	s.Value = &v
	return s
}

func (s *ListBasicAcceleratorsRequestTag) Validate() error {
	return dara.Validate(s)
}
