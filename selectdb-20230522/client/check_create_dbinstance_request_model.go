// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCreateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v int32) *CheckCreateDBInstanceRequest
	GetCacheSize() *int32
	SetChargeType(v string) *CheckCreateDBInstanceRequest
	GetChargeType() *string
	SetClientToken(v string) *CheckCreateDBInstanceRequest
	GetClientToken() *string
	SetConnectionString(v string) *CheckCreateDBInstanceRequest
	GetConnectionString() *string
	SetDBInstanceClass(v string) *CheckCreateDBInstanceRequest
	GetDBInstanceClass() *string
	SetDBInstanceDescription(v string) *CheckCreateDBInstanceRequest
	GetDBInstanceDescription() *string
	SetEngine(v string) *CheckCreateDBInstanceRequest
	GetEngine() *string
	SetEngineVersion(v string) *CheckCreateDBInstanceRequest
	GetEngineVersion() *string
	SetPeriod(v string) *CheckCreateDBInstanceRequest
	GetPeriod() *string
	SetRegionId(v string) *CheckCreateDBInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CheckCreateDBInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *CheckCreateDBInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityIPList(v string) *CheckCreateDBInstanceRequest
	GetSecurityIPList() *string
	SetUsedTime(v int32) *CheckCreateDBInstanceRequest
	GetUsedTime() *int32
	SetVSwitchId(v string) *CheckCreateDBInstanceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CheckCreateDBInstanceRequest
	GetVpcId() *string
	SetZoneId(v string) *CheckCreateDBInstanceRequest
	GetZoneId() *string
}

type CheckCreateDBInstanceRequest struct {
	// The reserved cache size, in GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. The token must be unique among different requests. The token can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The connection string of the instance.
	//
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance type. Valid values:
	//
	// - **selectdb.xlarge**: 4 cores, 32 GB.
	//
	// - **selectdb.2xlarge**: 8 cores, 64 GB.
	//
	// - **selectdb.4xlarge**: 16 cores, 128 GB.
	//
	// - **selectdb.8xlarge**: 32 cores, 256 GB.
	//
	// - **selectdb.16xlarge**: 64 cores, 512 GB.
	//
	// - **selectdb.24xlarge**: 96 cores, 768 GB.
	//
	// - **selectdb.32xlarge**: 128 cores, 1024 GB.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb.xlarge
	DBInstanceClass *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// New instance test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// SelectDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The database engine version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// - **Year**
	//
	// - **Month**
	//
	// > This parameter is required and takes effect only when **ChargeType*	- is set to **Prepaid**.
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address whitelist for the instance. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 172.16.XX.XX/12,192.168.XX.XX/22
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The subscription duration. Valid values:
	//
	// - If Period is set to Year, valid values are 1, 2, 3, 4, and 5.
	//
	// - If Period is set to Month, valid values are integers from 1 to 12.
	//
	// > This parameter is required and takes effect only when ChargeType is set to Prepaid.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// The vSwitch ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckCreateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CheckCreateDBInstanceRequest) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *CheckCreateDBInstanceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CheckCreateDBInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CheckCreateDBInstanceRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CheckCreateDBInstanceRequest) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *CheckCreateDBInstanceRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *CheckCreateDBInstanceRequest) GetEngine() *string {
	return s.Engine
}

func (s *CheckCreateDBInstanceRequest) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *CheckCreateDBInstanceRequest) GetPeriod() *string {
	return s.Period
}

func (s *CheckCreateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCreateDBInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CheckCreateDBInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CheckCreateDBInstanceRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CheckCreateDBInstanceRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CheckCreateDBInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CheckCreateDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CheckCreateDBInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckCreateDBInstanceRequest) SetCacheSize(v int32) *CheckCreateDBInstanceRequest {
	s.CacheSize = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetChargeType(v string) *CheckCreateDBInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetClientToken(v string) *CheckCreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetConnectionString(v string) *CheckCreateDBInstanceRequest {
	s.ConnectionString = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetDBInstanceClass(v string) *CheckCreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetDBInstanceDescription(v string) *CheckCreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetEngine(v string) *CheckCreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetEngineVersion(v string) *CheckCreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetPeriod(v string) *CheckCreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetRegionId(v string) *CheckCreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetResourceGroupId(v string) *CheckCreateDBInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetResourceOwnerId(v int64) *CheckCreateDBInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetSecurityIPList(v string) *CheckCreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetUsedTime(v int32) *CheckCreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetVSwitchId(v string) *CheckCreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetVpcId(v string) *CheckCreateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) SetZoneId(v string) *CheckCreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CheckCreateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
