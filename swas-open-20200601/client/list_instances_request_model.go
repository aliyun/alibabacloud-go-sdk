// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *ListInstancesRequest
	GetChargeType() *string
	SetInstanceIds(v string) *ListInstancesRequest
	GetInstanceIds() *string
	SetInstanceName(v string) *ListInstancesRequest
	GetInstanceName() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetPlanType(v string) *ListInstancesRequest
	GetPlanType() *string
	SetPublicIpAddresses(v string) *ListInstancesRequest
	GetPublicIpAddresses() *string
	SetRegionId(v string) *ListInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListInstancesRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListInstancesRequest
	GetStatus() *string
	SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest
	GetTag() []*ListInstancesRequestTag
}

type ListInstancesRequest struct {
	// The billing method of the instance. Valid values:
	//
	// - PrePaid: subscription. Currently, only the subscription billing method is supported.
	//
	// Default value: PrePaid.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The instance IDs of the simple application servers. The value can be a JSON array that consists of up to 100 instance IDs. Separate multiple IDs with commas (,).
	//
	// > If you specify both the InstanceIds and PublicIpAddresses parameters, ensure that the specified instance IDs and public IP addresses belong to the same simple application server. Otherwise, an empty response is returned.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The name of the simple application server. You can use the asterisk (*) wildcard for fuzzy search.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The page number of the simple application server list.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The plan type. Valid values:
	//
	// - NORMAL: general-purpose.
	//
	// - CAPACITY: storage-optimized.
	//
	// - PREVIOUS: previous generation.
	//
	// - MULTI_IP: multi-public-IP.
	//
	// - INTERNATIONAL: international.
	//
	// > For more information, see [Instance families](https://help.aliyun.com/document_detail/2861773.html).
	//
	// example:
	//
	// NORMAL
	PlanType *string `json:"PlanType,omitempty" xml:"PlanType,omitempty"`
	// The public IP addresses of the simple application servers. The value can be a JSON array that consists of up to 100 public IP addresses. Separate multiple IP addresses with commas (,).
	//
	// > If you specify both the InstanceIds and PublicIpAddresses parameters, ensure that the specified instance IDs and public IP addresses belong to the same simple application server. Otherwise, an empty response is returned.
	//
	// example:
	//
	// ["``42.1.**.**``", "``42.2.**.**``"]
	PublicIpAddresses *string `json:"PublicIpAddresses,omitempty" xml:"PublicIpAddresses,omitempty"`
	// The region ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the simple application server belongs.
	//
	// example:
	//
	// rg-aek2bti7cf7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status. Valid values:
	//
	// - Pending: The instance is being prepared.
	//
	// - Starting: The instance is starting.
	//
	// - Running: The instance is running.
	//
	// - Stopping: The instance is stopping.
	//
	// - Stopped: The instance is stopped.
	//
	// - Resetting: The instance is being reset.
	//
	// - Upgrading: The instance is being upgraded.
	//
	// - Disabled: The instance is unavailable.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the instance.
	Tag []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ListInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetPlanType() *string {
	return s.PlanType
}

func (s *ListInstancesRequest) GetPublicIpAddresses() *string {
	return s.PublicIpAddresses
}

func (s *ListInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesRequest) GetTag() []*ListInstancesRequestTag {
	return s.Tag
}

func (s *ListInstancesRequest) SetChargeType(v string) *ListInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceIds(v string) *ListInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
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

func (s *ListInstancesRequest) SetPlanType(v string) *ListInstancesRequest {
	s.PlanType = &v
	return s
}

func (s *ListInstancesRequest) SetPublicIpAddresses(v string) *ListInstancesRequest {
	s.PublicIpAddresses = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetResourceGroupId(v string) *ListInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
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
	// The tag key of the instance. The tag key must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the instance. The tag value must be 1 to 64 characters in length. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
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
