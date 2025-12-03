// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHighDefinationMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogProject(v string) *ModifyHighDefinationMonitorRequest
	GetLogProject() *string
	SetLogStore(v string) *ModifyHighDefinationMonitorRequest
	GetLogStore() *string
	SetOwnerAccount(v string) *ModifyHighDefinationMonitorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyHighDefinationMonitorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHighDefinationMonitorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHighDefinationMonitorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHighDefinationMonitorRequest
	GetResourceOwnerId() *int64
}

type ModifyHighDefinationMonitorRequest struct {
	// The new name of the project of Log Service. The name must be 4 to 63 characters in length, and can contain digits and lowercase letters. It must start and end with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-project
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The new name of the Logstore of Log Service. The name must be 2 to 64 characters in length and can contain digits, lowercase letters, hyphens (-) and underscores (_). It must start and end with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-log-store
	LogStore     *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyHighDefinationMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHighDefinationMonitorRequest) GoString() string {
	return s.String()
}

func (s *ModifyHighDefinationMonitorRequest) GetLogProject() *string {
	return s.LogProject
}

func (s *ModifyHighDefinationMonitorRequest) GetLogStore() *string {
	return s.LogStore
}

func (s *ModifyHighDefinationMonitorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyHighDefinationMonitorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHighDefinationMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHighDefinationMonitorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHighDefinationMonitorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHighDefinationMonitorRequest) SetLogProject(v string) *ModifyHighDefinationMonitorRequest {
	s.LogProject = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetLogStore(v string) *ModifyHighDefinationMonitorRequest {
	s.LogStore = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetOwnerAccount(v string) *ModifyHighDefinationMonitorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetOwnerId(v int64) *ModifyHighDefinationMonitorRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetRegionId(v string) *ModifyHighDefinationMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetResourceOwnerAccount(v string) *ModifyHighDefinationMonitorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) SetResourceOwnerId(v int64) *ModifyHighDefinationMonitorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHighDefinationMonitorRequest) Validate() error {
	return dara.Validate(s)
}
