// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBProxyInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetDBProxyEngineType() *string
	SetOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetSwitchTime() *string
	SetTargetMinorVersion(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetTargetMinorVersion() *string
	SetUpgradeTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest
	GetUpgradeTime() *string
}

type UpgradeDBProxyInstanceKernelVersionRequest struct {
	// Instance ID of the instance. You can call the DescribeDBInstances operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType    *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The specified time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// >If you set **UpgradeTime*	- to **SpecifyTime**, you must specify this parameter.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The target minor engine version.
	//
	// example:
	//
	// 2.25.9_20251028
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The time when you want to perform the upgrade. Valid values:
	//
	// 	- **MaintainTime*	- (default): The upgrade is performed within the [maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **Immediate**: The upgrade is performed immediately.
	//
	// 	- **SpecifyTime**: The upgrade is performed at a specified time.
	//
	// example:
	//
	// MaintainTime
	UpgradeTime *string `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty"`
}

func (s UpgradeDBProxyInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBProxyInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetTargetMinorVersion() *string {
	return s.TargetMinorVersion
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) GetUpgradeTime() *string {
	return s.UpgradeTime
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetDBInstanceId(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetDBProxyEngineType(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetResourceOwnerAccount(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetResourceOwnerId(v int64) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetSwitchTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.SwitchTime = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetTargetMinorVersion(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.TargetMinorVersion = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) SetUpgradeTime(v string) *UpgradeDBProxyInstanceKernelVersionRequest {
	s.UpgradeTime = &v
	return s
}

func (s *UpgradeDBProxyInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
