// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcceleratorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListAcceleratorsRequest
	GetAcceleratorId() *string
	SetPageNumber(v int32) *ListAcceleratorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAcceleratorsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListAcceleratorsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListAcceleratorsRequest
	GetResourceGroupId() *string
	SetState(v string) *ListAcceleratorsRequest
	GetState() *string
	SetTag(v []*ListAcceleratorsRequestTag) *ListAcceleratorsRequest
	GetTag() []*ListAcceleratorsRequestTag
}

type ListAcceleratorsRequest struct {
	// The ID of the GA instance.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Deprecated
	//
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekztkx4zwc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the GA instance. Valid values:
	//
	// 	- **init**: The GA instance is being initialized.
	//
	// 	- **active**: The GA instance is available.
	//
	// 	- **configuring**: The GA instance is being configured.
	//
	// 	- **binding**: The GA instance is being associated.
	//
	// 	- **unbinding**: The GA instance is being disassociated.
	//
	// 	- **deleting**: The GA instance is being deleted.
	//
	// 	- **finacialLocked**: The GA instance is locked due to overdue payments.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the GA instance.
	Tag []*ListAcceleratorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAcceleratorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsRequest) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListAcceleratorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAcceleratorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAcceleratorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAcceleratorsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListAcceleratorsRequest) GetState() *string {
	return s.State
}

func (s *ListAcceleratorsRequest) GetTag() []*ListAcceleratorsRequestTag {
	return s.Tag
}

func (s *ListAcceleratorsRequest) SetAcceleratorId(v string) *ListAcceleratorsRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetPageNumber(v int32) *ListAcceleratorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAcceleratorsRequest) SetPageSize(v int32) *ListAcceleratorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAcceleratorsRequest) SetRegionId(v string) *ListAcceleratorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetResourceGroupId(v string) *ListAcceleratorsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAcceleratorsRequest) SetState(v string) *ListAcceleratorsRequest {
	s.State = &v
	return s
}

func (s *ListAcceleratorsRequest) SetTag(v []*ListAcceleratorsRequestTag) *ListAcceleratorsRequest {
	s.Tag = v
	return s
}

func (s *ListAcceleratorsRequest) Validate() error {
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

type ListAcceleratorsRequestTag struct {
	// The tag key of the GA resource. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the GA resource. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// You can specify up to 20 tag values.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAcceleratorsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAcceleratorsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAcceleratorsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAcceleratorsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAcceleratorsRequestTag) SetKey(v string) *ListAcceleratorsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAcceleratorsRequestTag) SetValue(v string) *ListAcceleratorsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAcceleratorsRequestTag) Validate() error {
	return dara.Validate(s)
}
