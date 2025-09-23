// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v int32) *DescribeInstancesRequest
	GetEdition() *int32
	SetEnabled(v int32) *DescribeInstancesRequest
	GetEnabled() *int32
	SetExpireEndTime(v int64) *DescribeInstancesRequest
	GetExpireEndTime() *int64
	SetExpireStartTime(v int64) *DescribeInstancesRequest
	GetExpireStartTime() *int64
	SetInstanceIds(v string) *DescribeInstancesRequest
	GetInstanceIds() *string
	SetIp(v string) *DescribeInstancesRequest
	GetIp() *string
	SetPageNo(v string) *DescribeInstancesRequest
	GetPageNo() *string
	SetPageSize(v string) *DescribeInstancesRequest
	GetPageSize() *string
	SetRemark(v string) *DescribeInstancesRequest
	GetRemark() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *DescribeInstancesRequest
	GetSourceIp() *string
	SetStatus(v []*int32) *DescribeInstancesRequest
	GetStatus() []*int32
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
}

type DescribeInstancesRequest struct {
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// example:
	//
	// 1578931200000
	ExpireEndTime *int64 `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	// example:
	//
	// 1578931200000
	ExpireStartTime *int64 `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	// example:
	//
	// ["ddoscoo-cn-XXXXX"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 1.1.1.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// testRemark
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// example:
	//
	// 1
	Status []*int32                       `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Tag    []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) GetEdition() *int32 {
	return s.Edition
}

func (s *DescribeInstancesRequest) GetEnabled() *int32 {
	return s.Enabled
}

func (s *DescribeInstancesRequest) GetExpireEndTime() *int64 {
	return s.ExpireEndTime
}

func (s *DescribeInstancesRequest) GetExpireStartTime() *int64 {
	return s.ExpireStartTime
}

func (s *DescribeInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *DescribeInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstancesRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInstancesRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeInstancesRequest) GetStatus() []*int32 {
	return s.Status
}

func (s *DescribeInstancesRequest) GetTag() []*DescribeInstancesRequestTag {
	return s.Tag
}

func (s *DescribeInstancesRequest) SetEdition(v int32) *DescribeInstancesRequest {
	s.Edition = &v
	return s
}

func (s *DescribeInstancesRequest) SetEnabled(v int32) *DescribeInstancesRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireEndTime(v int64) *DescribeInstancesRequest {
	s.ExpireEndTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetExpireStartTime(v int64) *DescribeInstancesRequest {
	s.ExpireStartTime = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceIds(v string) *DescribeInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNo(v string) *DescribeInstancesRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v string) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetRemark(v string) *DescribeInstancesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSourceIp(v string) *DescribeInstancesRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeInstancesRequest) SetStatus(v []*int32) *DescribeInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeInstancesRequest) SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesRequestTag struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
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
