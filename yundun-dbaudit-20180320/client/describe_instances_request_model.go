// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeInstancesRequest
	GetCurrentPage() *int32
	SetInstanceId(v []*string) *DescribeInstancesRequest
	GetInstanceId() []*string
	SetInstanceStatus(v string) *DescribeInstancesRequest
	GetInstanceStatus() *string
	SetPageNo(v int32) *DescribeInstancesRequest
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
}

type DescribeInstancesRequest struct {
	CurrentPage     *int32                         `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceId      []*string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	InstanceStatus  *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	PageNo          *int32                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId        *string                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeInstancesRequest) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeInstancesRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequest) GetTag() []*DescribeInstancesRequestTag {
	return s.Tag
}

func (s *DescribeInstancesRequest) SetCurrentPage(v int32) *DescribeInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceId(v []*string) *DescribeInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceStatus(v string) *DescribeInstancesRequest {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNo(v int32) *DescribeInstancesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRegionId(v string) *DescribeInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
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

type DescribeInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInstancesRequestTag) SetKey(v string) *DescribeInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInstancesRequestTag) SetValue(v string) *DescribeInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
