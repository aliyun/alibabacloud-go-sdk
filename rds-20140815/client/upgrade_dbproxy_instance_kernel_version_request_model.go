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
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
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
	// The specific point in time when you want to perform the upgrade. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	//
	// >  If you set **UpgradeTime*	- to **SpecifyTime**, you must specify SwitchTime.
	//
	// example:
	//
	// 2019-07-10T13:15:12Z
	SwitchTime         *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	TargetMinorVersion *string `json:"TargetMinorVersion,omitempty" xml:"TargetMinorVersion,omitempty"`
	// The time when you want to upgrade the database proxy version of the instance. Valid values:
	//
	// 	- **MaintainTime*	- (default): performs the upgrade during the maintenance window that you specified. For more information, see [Modify the maintenance window](https://help.aliyun.com/document_detail/610402.html).
	//
	// 	- **Immediate**: performs the upgrade immediately.
	//
	// 	- **SpecifyTime**: performs the upgrade at a specified point in time.
	//
	// > 	- **If the instance runs MySQL, you can set this parameter to **MaintainTime**, **Immediate**, or SpecifyTime**.
	//
	// > 	- If the instance runs PostgreSQL, you can set this parameter to **MaintainTime*	- or **Immediate**.
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
