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
	// This parameter is required.
	//
	// example:
	//
	// 200
	CacheSize *int32 `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// AB
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// selectdb-cn-7213c8y****-public.selectdbfe.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The specifications of the instance. Valid values:
	//
	// 	- **selectdb.xlarge**: 4 CPU cores and 32 GB of memory.
	//
	// 	- **selectdb.2xlarge**: 8 CPU cores and 64 GB of memory.
	//
	// 	- **selectdb.4xlarge**: 16 CPU cores and 128 GB of memory.
	//
	// 	- **selectdb.8xlarge**: 32 CPU cores and 256 GB of memory.
	//
	// 	- **selectdb.16xlarge**: 64 CPU cores and 512 GB of memory.
	//
	// 	- **selectdb.24xlarge**: 96 CPU cores and 768 GB of memory.
	//
	// 	- **selectdb.32xlarge**: 128 CPU cores and 1,024 GB of memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb.xlarge
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// SelectDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the database engine.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2.4
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 172.16.XX.XX/12,192.168.XX.XX/22
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The subscription duration of the instance. Valid values:
	//
	// 	- If Period is set to Year, valid values of UsedTime are 1, 2, 3, 4, and 5.
	//
	// 	- If Period is set to Month, valid values of UsedTime are 1 to 12.
	//
	// >  This parameter takes effect and is required only if ChargeType is set to Prepaid.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
