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
	// The backup method. Valid values:
	//
	// - **logical**: logical backup
	//
	// - **physical**: physical backup
	//
	// This parameter is required.
	//
	// example:
	//
	// logical
	BackupMethod *string `json:"BackupMethod,omitempty" xml:"BackupMethod,omitempty"`
	// A client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// KJSAHKJFHKJSHFKASHFKJADFHKDXXXX
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region of the database.
	//
	// example:
	//
	// cn-hangzhou
	DatabaseRegion *string `json:"DatabaseRegion,omitempty" xml:"DatabaseRegion,omitempty"`
	// The database type. Valid values:
	//
	// - **MySQL**
	//
	// - **MSSQL**
	//
	// - **Oracle**
	//
	// - **MariaDB**
	//
	// - **PostgreSQL**
	//
	// - **DRDS**
	//
	// - **MongoDB**
	//
	// - **Redis**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The source of the request. The default value is OpenAPI. You do not need to set this parameter.
	//
	// example:
	//
	// OpenAPI
	FromApp *string `json:"FromApp,omitempty" xml:"FromApp,omitempty"`
	// The instance class. Valid values:
	//
	// - **micro**: Entry
	//
	// - **small**: Basic
	//
	// - **medium**: Standard
	//
	// - **large**: Enhanced
	//
	// - **xlarge**: Enhanced (no traffic limit)
	//
	// > The higher the instance class, the better the performance of backup and recovery. For more information, see [Specifications](https://help.aliyun.com/document_detail/84372.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// micro
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The database instance type. Valid values:
	//
	// - **RDS**
	//
	// - **PolarDB**
	//
	// - **DDS**: Alibaba Cloud MongoDB
	//
	// - **Kvstore**: Alibaba Cloud Redis
	//
	// - **Other**: A database that is connected over the Internet.
	//
	// - **dg**: A self-managed database without a public IP address and port that is connected through Database Gateway (DG).
	//
	// example:
	//
	// RDS
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The payment method. Valid value:
	//
	// **prepay**: subscription
	//
	// example:
	//
	// prepay
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The billing cycle of the subscription instance. Valid values:
	//
	// - **Year**
	//
	// - **Month**
	//
	// example:
	//
	// Month
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the DBS instance. This parameter is required. Call the [DescribeRegions](https://help.aliyun.com/document_detail/2869853.html) operation to view the regions that DBS supports.
	//
	// > For more information, see [Endpoints](https://help.aliyun.com/document_detail/2869810.html).
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
	// The storage region.
	//
	// example:
	//
	// cn-hangzhou
	StorageRegion *string `json:"StorageRegion,omitempty" xml:"StorageRegion,omitempty"`
	// This parameter is not used.
	//
	// example:
	//
	// 无
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The subscription duration. Valid values:
	//
	// - If you set the **Period*	- parameter to **Year**, the value of **UsedTime*	- can be 1 to 5.
	//
	// - If you set the **Period*	- parameter to **Month**, the value of **UsedTime*	- can be 1 to 11.
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
