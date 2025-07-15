// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHighDefinitionMonitorLogStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetInstanceId() *string
	SetInstanceType(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetInstanceType() *string
	SetLogProject(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetLogProject() *string
	SetLogStore(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetLogStore() *string
	SetOwnerAccount(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetHighDefinitionMonitorLogStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetHighDefinitionMonitorLogStatusRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *SetHighDefinitionMonitorLogStatusRequest
	GetStatus() *string
}

type SetHighDefinitionMonitorLogStatusRequest struct {
	// The ID of the instance for which you want to configure fine-grained monitoring.
	//
	// This parameter is required.
	//
	// example:
	//
	// eip-wz9fi6qboho9fwgx7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance type. Set the value to **EIP**.
	//
	// example:
	//
	// EIP
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the Simple Log Service (SLS) project.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdmonitor-cn-shenzhen
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The name of the Logstore.
	//
	// This parameter is required.
	//
	// example:
	//
	// hdmonitor
	LogStore     *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of fine-grained monitoring. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetHighDefinitionMonitorLogStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetHighDefinitionMonitorLogStatusRequest) GoString() string {
	return s.String()
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetLogProject() *string {
	return s.LogProject
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetHighDefinitionMonitorLogStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetInstanceId(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetInstanceType(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.InstanceType = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetLogProject(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.LogProject = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetLogStore(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.LogStore = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetOwnerAccount(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetOwnerId(v int64) *SetHighDefinitionMonitorLogStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetRegionId(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetResourceOwnerAccount(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetResourceOwnerId(v int64) *SetHighDefinitionMonitorLogStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) SetStatus(v string) *SetHighDefinitionMonitorLogStatusRequest {
	s.Status = &v
	return s
}

func (s *SetHighDefinitionMonitorLogStatusRequest) Validate() error {
	return dara.Validate(s)
}
