// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UpgradeKernelVersionRequest
	GetDBClusterId() *string
	SetDBVersion(v string) *UpgradeKernelVersionRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *UpgradeKernelVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpgradeKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchMode(v int32) *UpgradeKernelVersionRequest
	GetSwitchMode() *int32
}

type UpgradeKernelVersionRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the information about all AnalyticDB for MySQL Data Warehouse Edition clusters within a region, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The minor version to which you want to update.
	//
	// >  You can call the **DescribeKernelVersion*	- operation to query the supported minor versions.
	//
	// example:
	//
	// 3.1.9
	DBVersion            *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when to perform the update. Valid values:
	//
	// 	- **0*	- (default): immediately performs the update.
	//
	// 	- **1**: performs the update during the maintenance window.
	//
	// >  You can call the [ModifyDBClusterMaintainTime](https://help.aliyun.com/document_detail/612236.html) operation to modify the maintenance window of an AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// 0
	SwitchMode *int32 `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s UpgradeKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeKernelVersionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpgradeKernelVersionRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *UpgradeKernelVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpgradeKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeKernelVersionRequest) GetSwitchMode() *int32 {
	return s.SwitchMode
}

func (s *UpgradeKernelVersionRequest) SetDBClusterId(v string) *UpgradeKernelVersionRequest {
	s.DBClusterId = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetDBVersion(v string) *UpgradeKernelVersionRequest {
	s.DBVersion = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetOwnerAccount(v string) *UpgradeKernelVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetOwnerId(v int64) *UpgradeKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeKernelVersionRequest) SetSwitchMode(v int32) *UpgradeKernelVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
