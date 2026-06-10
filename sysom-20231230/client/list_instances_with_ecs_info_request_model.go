// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesWithEcsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int32) *ListInstancesWithEcsInfoRequest
	GetCurrent() *int32
	SetHealthStatus(v string) *ListInstancesWithEcsInfoRequest
	GetHealthStatus() *string
	SetInstanceId(v string) *ListInstancesWithEcsInfoRequest
	GetInstanceId() *string
	SetInstanceIdName(v string) *ListInstancesWithEcsInfoRequest
	GetInstanceIdName() *string
	SetInstanceName(v string) *ListInstancesWithEcsInfoRequest
	GetInstanceName() *string
	SetInstanceTag(v *ListInstancesWithEcsInfoRequestInstanceTag) *ListInstancesWithEcsInfoRequest
	GetInstanceTag() *ListInstancesWithEcsInfoRequestInstanceTag
	SetIsManaged(v int32) *ListInstancesWithEcsInfoRequest
	GetIsManaged() *int32
	SetOsName(v string) *ListInstancesWithEcsInfoRequest
	GetOsName() *string
	SetPageSize(v int32) *ListInstancesWithEcsInfoRequest
	GetPageSize() *int32
	SetPrivateIp(v string) *ListInstancesWithEcsInfoRequest
	GetPrivateIp() *string
	SetPublicIp(v string) *ListInstancesWithEcsInfoRequest
	GetPublicIp() *string
	SetRegion(v string) *ListInstancesWithEcsInfoRequest
	GetRegion() *string
	SetResourceGroupId(v string) *ListInstancesWithEcsInfoRequest
	GetResourceGroupId() *string
	SetResourceGroupIdName(v string) *ListInstancesWithEcsInfoRequest
	GetResourceGroupIdName() *string
	SetResourceGroupName(v string) *ListInstancesWithEcsInfoRequest
	GetResourceGroupName() *string
}

type ListInstancesWithEcsInfoRequest struct {
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
	InstanceTag *ListInstancesWithEcsInfoRequestInstanceTag `json:"instance_tag,omitempty" xml:"instance_tag,omitempty" type:"Struct"`
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

func (s ListInstancesWithEcsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListInstancesWithEcsInfoRequest) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListInstancesWithEcsInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesWithEcsInfoRequest) GetInstanceIdName() *string {
	return s.InstanceIdName
}

func (s *ListInstancesWithEcsInfoRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListInstancesWithEcsInfoRequest) GetInstanceTag() *ListInstancesWithEcsInfoRequestInstanceTag {
	return s.InstanceTag
}

func (s *ListInstancesWithEcsInfoRequest) GetIsManaged() *int32 {
	return s.IsManaged
}

func (s *ListInstancesWithEcsInfoRequest) GetOsName() *string {
	return s.OsName
}

func (s *ListInstancesWithEcsInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesWithEcsInfoRequest) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *ListInstancesWithEcsInfoRequest) GetPublicIp() *string {
	return s.PublicIp
}

func (s *ListInstancesWithEcsInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *ListInstancesWithEcsInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesWithEcsInfoRequest) GetResourceGroupIdName() *string {
	return s.ResourceGroupIdName
}

func (s *ListInstancesWithEcsInfoRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListInstancesWithEcsInfoRequest) SetCurrent(v int32) *ListInstancesWithEcsInfoRequest {
	s.Current = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetHealthStatus(v string) *ListInstancesWithEcsInfoRequest {
	s.HealthStatus = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetInstanceId(v string) *ListInstancesWithEcsInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetInstanceIdName(v string) *ListInstancesWithEcsInfoRequest {
	s.InstanceIdName = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetInstanceName(v string) *ListInstancesWithEcsInfoRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetInstanceTag(v *ListInstancesWithEcsInfoRequestInstanceTag) *ListInstancesWithEcsInfoRequest {
	s.InstanceTag = v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetIsManaged(v int32) *ListInstancesWithEcsInfoRequest {
	s.IsManaged = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetOsName(v string) *ListInstancesWithEcsInfoRequest {
	s.OsName = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetPageSize(v int32) *ListInstancesWithEcsInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetPrivateIp(v string) *ListInstancesWithEcsInfoRequest {
	s.PrivateIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetPublicIp(v string) *ListInstancesWithEcsInfoRequest {
	s.PublicIp = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetRegion(v string) *ListInstancesWithEcsInfoRequest {
	s.Region = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetResourceGroupId(v string) *ListInstancesWithEcsInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetResourceGroupIdName(v string) *ListInstancesWithEcsInfoRequest {
	s.ResourceGroupIdName = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) SetResourceGroupName(v string) *ListInstancesWithEcsInfoRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequest) Validate() error {
	if s.InstanceTag != nil {
		if err := s.InstanceTag.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstancesWithEcsInfoRequestInstanceTag struct {
	// Tag key.
	//
	// example:
	//
	// feature_dim_radar_chart
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Tag value.
	//
	// example:
	//
	// test
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesWithEcsInfoRequestInstanceTag) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoRequestInstanceTag) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoRequestInstanceTag) GetKey() *string {
	return s.Key
}

func (s *ListInstancesWithEcsInfoRequestInstanceTag) GetValue() *string {
	return s.Value
}

func (s *ListInstancesWithEcsInfoRequestInstanceTag) SetKey(v string) *ListInstancesWithEcsInfoRequestInstanceTag {
	s.Key = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequestInstanceTag) SetValue(v string) *ListInstancesWithEcsInfoRequestInstanceTag {
	s.Value = &v
	return s
}

func (s *ListInstancesWithEcsInfoRequestInstanceTag) Validate() error {
	return dara.Validate(s)
}
