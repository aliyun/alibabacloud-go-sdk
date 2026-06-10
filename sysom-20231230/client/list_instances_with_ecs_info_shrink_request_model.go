// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesWithEcsInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListInstancesWithEcsInfoShrinkRequest
	GetCurrent() *int32
	SetHealthStatus(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetHealthStatus() *string
	SetInstanceId(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetInstanceId() *string
	SetInstanceIdName(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetInstanceIdName() *string
	SetInstanceName(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetInstanceName() *string
	SetInstanceTagShrink(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetInstanceTagShrink() *string
	SetIsManaged(v int32) *ListInstancesWithEcsInfoShrinkRequest
	GetIsManaged() *int32
	SetOsName(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetOsName() *string
	SetPageSize(v int32) *ListInstancesWithEcsInfoShrinkRequest
	GetPageSize() *int32
	SetPrivateIp(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetPrivateIp() *string
	SetPublicIp(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetPublicIp() *string
	SetRegion(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetRegion() *string
	SetResourceGroupId(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetResourceGroupId() *string
	SetResourceGroupIdName(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetResourceGroupIdName() *string
	SetResourceGroupName(v string) *ListInstancesWithEcsInfoShrinkRequest
	GetResourceGroupName() *string
}

type ListInstancesWithEcsInfoShrinkRequest struct {
	// This field exists when using paging and indicates the current page.
	//
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// Filter instances by health status.
	//
	// example:
	//
	// healthy
	HealthStatus *string `json:"health_status,omitempty" xml:"health_status,omitempty"`
	// If this field is specified, filter the Agent installation status for the specified instance.
	//
	// example:
	//
	// i-bp118piqcio9tiwgh84b
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// Filter by instance name or instance ID. Fuzzy query is supported.
	//
	// example:
	//
	// 84b
	InstanceIdName *string `json:"instance_id_name,omitempty" xml:"instance_id_name,omitempty"`
	// Widget instance name.
	//
	// example:
	//
	// block-load-balancer-hjdm9
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// Filter by tags of instances.
	InstanceTagShrink *string `json:"instance_tag,omitempty" xml:"instance_tag,omitempty"`
	// Filter instances by managed status.
	//
	// example:
	//
	// 1
	IsManaged *int32 `json:"is_managed,omitempty" xml:"is_managed,omitempty"`
	// Filter instances by operating system name.
	//
	// example:
	//
	// Alibaba Cloud Linux  3.2104 LTS 64bit
	OsName *string `json:"os_name,omitempty" xml:"os_name,omitempty"`
	// Page size. Default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Filter instances by private IP address.
	//
	// example:
	//
	// 1.1.1.1
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip,omitempty"`
	// Filter instances by public IP address.
	//
	// example:
	//
	// 1.1.1.1
	PublicIp *string `json:"public_ip,omitempty" xml:"public_ip,omitempty"`
	// Filter instances by region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Filter instances by resource group ID.
	//
	// example:
	//
	// rg-xxxxxxx
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
	// Filter by resource group name or resource group ID. Fuzzy query is supported.
	//
	// example:
	//
	// default
	ResourceGroupIdName *string `json:"resource_group_id_name,omitempty" xml:"resource_group_id_name,omitempty"`
	// Filter instances by resource group name.
	//
	// example:
	//
	// default resource group
	ResourceGroupName *string `json:"resource_group_name,omitempty" xml:"resource_group_name,omitempty"`
}

func (s ListInstancesWithEcsInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetInstanceIdName() *string {
	return s.InstanceIdName
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetInstanceTagShrink() *string {
	return s.InstanceTagShrink
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetIsManaged() *int32 {
	return s.IsManaged
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetOsName() *string {
	return s.OsName
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetResourceGroupIdName() *string {
	return s.ResourceGroupIdName
}

func (s *ListInstancesWithEcsInfoShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetCurrent(v int32) *ListInstancesWithEcsInfoShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetHealthStatus(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.HealthStatus = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetInstanceId(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetInstanceIdName(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.InstanceIdName = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetInstanceName(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetInstanceTagShrink(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.InstanceTagShrink = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetIsManaged(v int32) *ListInstancesWithEcsInfoShrinkRequest {
	s.IsManaged = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetOsName(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.OsName = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetPageSize(v int32) *ListInstancesWithEcsInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetPrivateIp(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.PrivateIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetPublicIp(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.PublicIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetRegion(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetResourceGroupId(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetResourceGroupIdName(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.ResourceGroupIdName = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) SetResourceGroupName(v string) *ListInstancesWithEcsInfoShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListInstancesWithEcsInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
