// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasProServiceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetDasProServiceUsageResponseBody
	GetCode() *int64
	SetData(v *GetDasProServiceUsageResponseBodyData) *GetDasProServiceUsageResponseBody
	GetData() *GetDasProServiceUsageResponseBodyData
	SetMessage(v string) *GetDasProServiceUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDasProServiceUsageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDasProServiceUsageResponseBody
	GetSuccess() *bool
}

type GetDasProServiceUsageResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {         "storageFreeQuotaInMB": 5120,         "ip": "rm-2ze8g2am97624****.mysql.****.com",         "custinsId": 12448331,         "userId": "196278346919****",         "uuid": "hdm_b0ae36343407609bf3e8df8709d8****",         "expireTime": 1924963200000,         "instanceId": "rm-2ze8g2am97624****",         "storageUsed": 10773752667393,         "engine": "MySQL",         "instanceAlias": "TESTDB01_PROD",         "port": 3310,         "vpcId": "hdm_****",         "commodityInstanceId": "daspro-****",         "startTime": 1606381940000,         "isSpare": false,         "region": "cn-shanghai",         "serviceUnitId": "5",         "sqlRetention": 30     }
	Data *GetDasProServiceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDasProServiceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDasProServiceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetDasProServiceUsageResponseBody) GetData() *GetDasProServiceUsageResponseBodyData {
	return s.Data
}

func (s *GetDasProServiceUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDasProServiceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDasProServiceUsageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDasProServiceUsageResponseBody) SetCode(v int64) *GetDasProServiceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetData(v *GetDasProServiceUsageResponseBodyData) *GetDasProServiceUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetMessage(v string) *GetDasProServiceUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetRequestId(v string) *GetDasProServiceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) SetSuccess(v bool) *GetDasProServiceUsageResponseBody {
	s.Success = &v
	return s
}

func (s *GetDasProServiceUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDasProServiceUsageResponseBodyData struct {
	// The ID of the DAS Enterprise Edition instance.
	//
	// example:
	//
	// daspro-cn-v0h1l6i****
	CommodityInstanceId *string `json:"commodityInstanceId,omitempty" xml:"commodityInstanceId,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The point of time when DAS Enterprise Edition for the database instance expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1648742400000
	ExpireTime *int64 `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The name of the database instance.
	//
	// example:
	//
	// TESTDB01
	InstanceAlias *string `json:"instanceAlias,omitempty" xml:"instanceAlias,omitempty"`
	// The database instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The endpoint of the database instance.
	//
	// example:
	//
	// rm-2ze8g2am97624****.mysql.****.com
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
	// Indicates whether DAS Enterprise Edition for the database instance has expired. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IsSpare *bool `json:"isSpare,omitempty" xml:"isSpare,omitempty"`
	// The estimated remaining time for migrating the data generated by the SQL Explorer and Audit feature from the previous version to the new version. Unit: milliseconds.
	//
	// >  This parameter is returned only when the SQL Explorer and Audit feature is migrated from the previous version to the new version.
	//
	// example:
	//
	// 60000
	MigrationPredictRemainingTime *int64 `json:"migrationPredictRemainingTime,omitempty" xml:"migrationPredictRemainingTime,omitempty"`
	// The port number that is used to connect to the database instance.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The region in which the database instance resides.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The service unit ID.
	//
	// example:
	//
	// 4
	ServiceUnitId *string `json:"serviceUnitId,omitempty" xml:"serviceUnitId,omitempty"`
	// The storage duration of SQL Explorer data. Unit: days.
	//
	// example:
	//
	// 180
	SqlRetention *string `json:"sqlRetention,omitempty" xml:"sqlRetention,omitempty"`
	// The time when DAS Enterprise Edition was enabled for the database instance. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1646100892000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The SQL Explorer storage space that is offered free-of-charge. Unit: MB.
	//
	// example:
	//
	// 5120
	StorageFreeQuotaInMB *float64 `json:"storageFreeQuotaInMB,omitempty" xml:"storageFreeQuotaInMB,omitempty"`
	// The storage usage of SQL Explorer of the database instance. Unit: bytes.
	//
	// example:
	//
	// 35903498
	StorageUsed *int64 `json:"storageUsed,omitempty" xml:"storageUsed,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-2zentqj1sk4qmolci****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetDasProServiceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDasProServiceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageResponseBodyData) GetCommodityInstanceId() *string {
	return s.CommodityInstanceId
}

func (s *GetDasProServiceUsageResponseBodyData) GetEngine() *string {
	return s.Engine
}

func (s *GetDasProServiceUsageResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetDasProServiceUsageResponseBodyData) GetInstanceAlias() *string {
	return s.InstanceAlias
}

func (s *GetDasProServiceUsageResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDasProServiceUsageResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *GetDasProServiceUsageResponseBodyData) GetIsSpare() *bool {
	return s.IsSpare
}

func (s *GetDasProServiceUsageResponseBodyData) GetMigrationPredictRemainingTime() *int64 {
	return s.MigrationPredictRemainingTime
}

func (s *GetDasProServiceUsageResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *GetDasProServiceUsageResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetDasProServiceUsageResponseBodyData) GetServiceUnitId() *string {
	return s.ServiceUnitId
}

func (s *GetDasProServiceUsageResponseBodyData) GetSqlRetention() *string {
	return s.SqlRetention
}

func (s *GetDasProServiceUsageResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDasProServiceUsageResponseBodyData) GetStorageFreeQuotaInMB() *float64 {
	return s.StorageFreeQuotaInMB
}

func (s *GetDasProServiceUsageResponseBodyData) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *GetDasProServiceUsageResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetDasProServiceUsageResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetDasProServiceUsageResponseBodyData) SetCommodityInstanceId(v string) *GetDasProServiceUsageResponseBodyData {
	s.CommodityInstanceId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetEngine(v string) *GetDasProServiceUsageResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetExpireTime(v int64) *GetDasProServiceUsageResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetInstanceAlias(v string) *GetDasProServiceUsageResponseBodyData {
	s.InstanceAlias = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetInstanceId(v string) *GetDasProServiceUsageResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetIp(v string) *GetDasProServiceUsageResponseBodyData {
	s.Ip = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetIsSpare(v bool) *GetDasProServiceUsageResponseBodyData {
	s.IsSpare = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetMigrationPredictRemainingTime(v int64) *GetDasProServiceUsageResponseBodyData {
	s.MigrationPredictRemainingTime = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetPort(v int32) *GetDasProServiceUsageResponseBodyData {
	s.Port = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetRegion(v string) *GetDasProServiceUsageResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetServiceUnitId(v string) *GetDasProServiceUsageResponseBodyData {
	s.ServiceUnitId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetSqlRetention(v string) *GetDasProServiceUsageResponseBodyData {
	s.SqlRetention = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetStartTime(v int64) *GetDasProServiceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetStorageFreeQuotaInMB(v float64) *GetDasProServiceUsageResponseBodyData {
	s.StorageFreeQuotaInMB = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetStorageUsed(v int64) *GetDasProServiceUsageResponseBodyData {
	s.StorageUsed = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetUserId(v string) *GetDasProServiceUsageResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) SetVpcId(v string) *GetDasProServiceUsageResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetDasProServiceUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
