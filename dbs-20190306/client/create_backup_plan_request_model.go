// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupMethod(v string) *CreateBackupPlanRequest
	GetBackupMethod() *string
	SetClientToken(v string) *CreateBackupPlanRequest
	GetClientToken() *string
	SetDatabaseRegion(v string) *CreateBackupPlanRequest
	GetDatabaseRegion() *string
	SetDatabaseType(v string) *CreateBackupPlanRequest
	GetDatabaseType() *string
	SetFromApp(v string) *CreateBackupPlanRequest
	GetFromApp() *string
	SetInstanceClass(v string) *CreateBackupPlanRequest
	GetInstanceClass() *string
	SetInstanceType(v string) *CreateBackupPlanRequest
	GetInstanceType() *string
	SetOwnerId(v string) *CreateBackupPlanRequest
	GetOwnerId() *string
	SetPayType(v string) *CreateBackupPlanRequest
	GetPayType() *string
	SetPeriod(v string) *CreateBackupPlanRequest
	GetPeriod() *string
	SetRegion(v string) *CreateBackupPlanRequest
	GetRegion() *string
	SetResourceGroupId(v string) *CreateBackupPlanRequest
	GetResourceGroupId() *string
	SetStorageRegion(v string) *CreateBackupPlanRequest
	GetStorageRegion() *string
	SetStorageType(v string) *CreateBackupPlanRequest
	GetStorageType() *string
	SetUsedTime(v int32) *CreateBackupPlanRequest
	GetUsedTime() *int32
}

type CreateBackupPlanRequest struct {
	// The backup method of the backup schedule. Valid values:
	//
	// 	- **logical**: logical backup
	//
	// 	- **physical**: physical backup
	//
	// 	- **duplication**: dump backup
	//
	// This parameter is required.
	//
	// example:
	//
	// logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// KJSAHKJFHKJSHFKASHFKJADFHKDXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region in which the database you want to back up resides.
	//
	// > This parameter is required if the **PayType*	- parameter is set to **postpay**.
	//
	// example:
	//
	// cn-hangzhou
	DatabaseRegion *string `json:"DatabaseRegion,omitempty" xml:"DatabaseRegion,omitempty"`
	// The type of the source database. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// 	- **MariaDB**
	//
	// 	- **PostgreSQL**
	//
	// 	- **DRDS**
	//
	// 	- **MongoDB**
	//
	// 	- **Redis**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The source of the request. The default value is OpenAPI and cannot be changed.
	//
	// example:
	//
	// OpenAPI
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The type of the backup schedule. Valid values:
	//
	// 	- **micro**
	//
	// 	- **small**
	//
	// 	- **medium**
	//
	// 	- **large**
	//
	// 	- **xlarge**
	//
	// >  A backup schedule type with higher specifications offers higher backup and restoration performance. For more information, see [Select a backup schedule type](https://help.aliyun.com/document_detail/84372.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// micro
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The type of the source database instance. Valid values:
	//
	// 	- **RDS**: ApsaraDB RDS.
	//
	// 	- **PolarDB**: PolarDB.
	//
	// 	- **DDS**: ApsaraDB for MongoDB.
	//
	// 	- **Kvstore**: ApsaraDB for Redis.
	//
	// 	- **Other**: Database connected by using an IP address and a port number.
	//
	// 	- **dg**: Self-managed database that has no public IP address or port number and is connected over Database Gateway.
	//
	// >  If **PayType*	- is set to **postpay**, this parameter is required.
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the backup schedule. Valid values:
	//
	// 	- **postpay**: pay-as-you-go
	//
	// 	- **prepay**: subscription
	//
	// > The default value is **prepay**. If the **BackupMethod*	- parameter is set to **duplication**, **postpay*	- is supported.
	//
	// example:
	//
	// prepay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// 	- **Year**: yearly subscription
	//
	// 	- **Month**: monthly subscription
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the region in which you can activate Data Disaster Recovery. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2869853.html) operation to query the regions supported by Data Disaster Recovery.
	//
	// >  For more information, see [Endpoints](https://help.aliyun.com/document_detail/2869810.html).
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzecovzti****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The region in which you want to store the backup data.
	//
	// > This parameter is required if the **PayType*	- parameter is set to **postpay**.
	//
	// example:
	//
	// cn-hangzhou
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	// This parameter is unavailable.
	//
	// example:
	//
	// N/A
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The subscription period. Valid values:
	//
	// 	- If **Period*	- is set to **Year**, the valid values of **UsedTime*	- range from 1 to 5.
	//
	// 	- If **Period*	- is set to **Month**, the valid values of **UsedTime*	- range from 1 to 11.
	//
	// example:
	//
	// 1
	UsedTime *int32 `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateBackupPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanRequest) GetBackupMethod() *string {
	return s.BackupMethod
}

func (s *CreateBackupPlanRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateBackupPlanRequest) GetDatabaseRegion() *string {
	return s.DatabaseRegion
}

func (s *CreateBackupPlanRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateBackupPlanRequest) GetFromApp() *string {
	return s.FromApp
}

func (s *CreateBackupPlanRequest) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *CreateBackupPlanRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateBackupPlanRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CreateBackupPlanRequest) GetPayType() *string {
	return s.PayType
}

func (s *CreateBackupPlanRequest) GetPeriod() *string {
	return s.Period
}

func (s *CreateBackupPlanRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateBackupPlanRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateBackupPlanRequest) GetStorageRegion() *string {
	return s.StorageRegion
}

func (s *CreateBackupPlanRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateBackupPlanRequest) GetUsedTime() *int32 {
	return s.UsedTime
}

func (s *CreateBackupPlanRequest) SetBackupMethod(v string) *CreateBackupPlanRequest {
	s.BackupMethod = &v
	return s
}

func (s *CreateBackupPlanRequest) SetClientToken(v string) *CreateBackupPlanRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDatabaseRegion(v string) *CreateBackupPlanRequest {
	s.DatabaseRegion = &v
	return s
}

func (s *CreateBackupPlanRequest) SetDatabaseType(v string) *CreateBackupPlanRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetFromApp(v string) *CreateBackupPlanRequest {
	s.FromApp = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceClass(v string) *CreateBackupPlanRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateBackupPlanRequest) SetInstanceType(v string) *CreateBackupPlanRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetOwnerId(v string) *CreateBackupPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPayType(v string) *CreateBackupPlanRequest {
	s.PayType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetPeriod(v string) *CreateBackupPlanRequest {
	s.Period = &v
	return s
}

func (s *CreateBackupPlanRequest) SetRegion(v string) *CreateBackupPlanRequest {
	s.Region = &v
	return s
}

func (s *CreateBackupPlanRequest) SetResourceGroupId(v string) *CreateBackupPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateBackupPlanRequest) SetStorageRegion(v string) *CreateBackupPlanRequest {
	s.StorageRegion = &v
	return s
}

func (s *CreateBackupPlanRequest) SetStorageType(v string) *CreateBackupPlanRequest {
	s.StorageType = &v
	return s
}

func (s *CreateBackupPlanRequest) SetUsedTime(v int32) *CreateBackupPlanRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateBackupPlanRequest) Validate() error {
	return dara.Validate(s)
}
