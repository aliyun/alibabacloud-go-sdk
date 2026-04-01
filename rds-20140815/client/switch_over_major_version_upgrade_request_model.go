// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverMajorVersionUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SwitchOverMajorVersionUpgradeRequest
	GetClientToken() *string
	SetDBInstanceName(v string) *SwitchOverMajorVersionUpgradeRequest
	GetDBInstanceName() *string
	SetOwnerAccount(v string) *SwitchOverMajorVersionUpgradeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchOverMajorVersionUpgradeRequest
	GetOwnerId() *int64
	SetRegionId(v []byte) *SwitchOverMajorVersionUpgradeRequest
	GetRegionId() []byte
	SetResourceGroupId(v string) *SwitchOverMajorVersionUpgradeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *SwitchOverMajorVersionUpgradeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchOverMajorVersionUpgradeRequest
	GetResourceOwnerId() *int64
	SetSwitchoverTimeout(v int32) *SwitchOverMajorVersionUpgradeRequest
	GetSwitchoverTimeout() *int32
	SetType(v string) *SwitchOverMajorVersionUpgradeRequest
	GetType() *string
}

type SwitchOverMajorVersionUpgradeRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pgm-m5e4gegx63fh92bn
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             []byte  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The timeout period for the switchover operation. The operation is canceled after it has been performed for a time period that exceeds the value. Unit: seconds. Valid value: 10 to 3600.
	//
	// example:
	//
	// 10
	SwitchoverTimeout *int32 `json:"SwitchoverTimeout,omitempty" xml:"SwitchoverTimeout,omitempty"`
	// The type of the switchover operation. Valid values:
	//
	// 	- switch
	//
	// 	- cancel
	//
	// 	- interrupt
	//
	// example:
	//
	// switch
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SwitchOverMajorVersionUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverMajorVersionUpgradeRequest) GoString() string {
	return s.String()
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetRegionId() []byte {
	return s.RegionId
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetSwitchoverTimeout() *int32 {
	return s.SwitchoverTimeout
}

func (s *SwitchOverMajorVersionUpgradeRequest) GetType() *string {
	return s.Type
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetClientToken(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.ClientToken = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetDBInstanceName(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetOwnerAccount(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetOwnerId(v int64) *SwitchOverMajorVersionUpgradeRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetRegionId(v []byte) *SwitchOverMajorVersionUpgradeRequest {
	s.RegionId = v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetResourceGroupId(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetResourceOwnerAccount(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetResourceOwnerId(v int64) *SwitchOverMajorVersionUpgradeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetSwitchoverTimeout(v int32) *SwitchOverMajorVersionUpgradeRequest {
	s.SwitchoverTimeout = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) SetType(v string) *SwitchOverMajorVersionUpgradeRequest {
	s.Type = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
