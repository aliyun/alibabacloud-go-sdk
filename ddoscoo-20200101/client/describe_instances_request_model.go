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
	SetInstanceIds(v []*string) *DescribeInstancesRequest
	GetInstanceIds() []*string
	SetIp(v string) *DescribeInstancesRequest
	GetIp() *string
	SetPageNumber(v string) *DescribeInstancesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeInstancesRequest
	GetPageSize() *string
	SetRemark(v string) *DescribeInstancesRequest
	GetRemark() *string
	SetResourceGroupId(v string) *DescribeInstancesRequest
	GetResourceGroupId() *string
	SetStatus(v []*int32) *DescribeInstancesRequest
	GetStatus() []*int32
	SetTag(v []*DescribeInstancesRequestTag) *DescribeInstancesRequest
	GetTag() []*DescribeInstancesRequestTag
}

type DescribeInstancesRequest struct {
	// The mitigation plan edition of the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Valid values:
	//
	// - **0**: Anti-DDoS Pro and Anti-DDoS Premium (outside the Chinese mainland) Insurance edition.
	//
	// - **1**: Anti-DDoS Pro and Anti-DDoS Premium (outside the Chinese mainland) Unlimited edition.
	//
	// - **2**: Anti-DDoS Pro and Anti-DDoS Premium (outside the Chinese mainland) acceleration line.
	//
	// - **9**: Anti-DDoS Pro and Anti-DDoS Premium (the Chinese mainland) Professional Edition.
	//
	// example:
	//
	// 9
	Edition *int32 `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The service traffic forwarding status of the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Valid values:
	//
	// - **0**: Service traffic forwarding is stopped.
	//
	// - **1**: Service traffic forwarding is normal.
	//
	// example:
	//
	// 1
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The latest expiration time of the Anti-DDoS Pro or Anti-DDoS Premium instances to query. This means that only instances whose expiration time is earlier than the value of **ExpireEndTime*	- are queried. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640361700000
	ExpireEndTime *int64 `json:"ExpireEndTime,omitempty" xml:"ExpireEndTime,omitempty"`
	// The earliest expiration time of the Anti-DDoS Pro or Anti-DDoS Premium instances to query. This means that only instances whose expiration time is later than the value of **ExpireStartTime*	- are queried. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640361500000
	ExpireStartTime *int64 `json:"ExpireStartTime,omitempty" xml:"ExpireStartTime,omitempty"`
	// The IDs of the Anti-DDoS Pro or Anti-DDoS Premium instances to query. You can specify up to 200 Anti-DDoS Pro or Anti-DDoS Premium instances.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance to query.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The page number for paging.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances per page for paging. Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The remark of the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Fuzzy match is supported.
	//
	// example:
	//
	// doc-test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management.
	//
	// If you do not specify this parameter, the default resource group is used.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The statuses of the Anti-DDoS Pro or Anti-DDoS Premium instances to query. You can specify up to 2 Anti-DDoS Pro or Anti-DDoS Premium instance statuses.
	Status []*int32 `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags attached to the Anti-DDoS Pro or Anti-DDoS Premium instances to query.
	Tag []*DescribeInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *DescribeInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeInstancesRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribeInstancesRequest) GetPageNumber() *string {
	return s.PageNumber
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

func (s *DescribeInstancesRequest) SetInstanceIds(v []*string) *DescribeInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeInstancesRequest) SetIp(v string) *DescribeInstancesRequest {
	s.Ip = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v string) *DescribeInstancesRequest {
	s.PageNumber = &v
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

func (s *DescribeInstancesRequest) SetStatus(v []*int32) *DescribeInstancesRequest {
	s.Status = v
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
	// The tag key attached to the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Maximum value of N: 200, which means you can specify up to 200 tag keys. The configuration rules are as follows:
	//
	// - Each tag consists of a tag key (**Key**) and a tag value (**Value**), separated by a comma (,).
	//
	// - Tags are separated by commas (,).
	//
	// > The tag key (**Key**) and tag value (**Value**) must appear in pairs.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value attached to the Anti-DDoS Pro or Anti-DDoS Premium instance to query. Maximum value of N: 200, which means you can specify up to 200 tag values. The configuration rules are as follows:
	//
	// - Each tag consists of a tag key (**Key**) and a tag value (**Value**), separated by a comma (,).
	//
	// - Tags are separated by commas (,).
	//
	// > The tag key (**Key**) and tag value (**Value**) must appear in pairs.
	//
	// example:
	//
	// test-value
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
