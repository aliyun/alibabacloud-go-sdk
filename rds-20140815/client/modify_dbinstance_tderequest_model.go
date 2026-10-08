// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *ModifyDBInstanceTDERequest
	GetCertificate() *string
	SetDBInstanceId(v string) *ModifyDBInstanceTDERequest
	GetDBInstanceId() *string
	SetDBName(v string) *ModifyDBInstanceTDERequest
	GetDBName() *string
	SetEncryptionKey(v string) *ModifyDBInstanceTDERequest
	GetEncryptionKey() *string
	SetIsRotate(v bool) *ModifyDBInstanceTDERequest
	GetIsRotate() *bool
	SetOwnerAccount(v string) *ModifyDBInstanceTDERequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceTDERequest
	GetOwnerId() *int64
	SetPassWord(v string) *ModifyDBInstanceTDERequest
	GetPassWord() *string
	SetPrivateKey(v string) *ModifyDBInstanceTDERequest
	GetPrivateKey() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceTDERequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceTDERequest
	GetResourceOwnerId() *int64
	SetRoleArn(v string) *ModifyDBInstanceTDERequest
	GetRoleArn() *string
	SetTDEStatus(v string) *ModifyDBInstanceTDERequest
	GetTDEStatus() *string
}

type ModifyDBInstanceTDERequest struct {
	// The certificate file.
	//
	// Format:
	//
	// - Public endpoint: `oss-<RegionId>.aliyuncs.com:<BucketName>:<CertificateFileName (with file extension)>`
	//
	// - Internal network endpoint: `oss-<RegionId>-internal.aliyuncs.com:<BucketName>:<CertificateFileName (with file extension)>`
	//
	// > - This parameter is active only for SQL Server 2019 Standard Edition, 2022 Standard Edition, 2025 Standard Edition, and SQL Server Enterprise instance instances.
	//
	// > - You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query active region IDs.
	//
	// example:
	//
	// oss-ap-southeast-1.aliyuncs.com:****:key.cer
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the database for which you want to enable TDE. You can specify multiple database names separated by commas (,). You can specify up to 50 database names.
	//
	// > This parameter is active and required only for SQL Server 2019 Standard Edition, 2022 Standard Edition, 2025 Standard Edition, and SQL Server Enterprise instance instances.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The custom key ID.
	//
	// > This parameter is available only for ApsaraDB RDS for MySQL and ApsaraDB RDS for PostgreSQL instances.
	//
	// example:
	//
	// 749c1df7-****-****-****-****
	EncryptionKey *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	// Specifies whether to rotate the key. Valid values:
	//
	// - **true**: Rotate the key.
	//
	// - **false*	- (default): Do not rotate the key.
	//
	// > This parameter is available only for ApsaraDB RDS for PostgreSQL instances.
	//
	// example:
	//
	// false
	IsRotate     *bool   `json:"IsRotate,omitempty" xml:"IsRotate,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The certificate password.
	//
	// > This parameter is active only for SQL Server 2019 Standard Edition, 2022 Standard Edition, 2025 Standard Edition, and SQL Server Enterprise instance instances.
	//
	// example:
	//
	// 1qaz@WSX
	PassWord *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	// The private key file.
	//
	// Format:
	//
	// - Public endpoint: `oss-<RegionId>.aliyuncs.com:<BucketName>:<PrivateKeyFileName (with file extension)>`
	//
	// - Internal network endpoint: `oss-<RegionId>-internal.aliyuncs.com:<BucketName>:<PrivateKeyFileName (with file extension)>`
	//
	// > - This parameter is active only for SQL Server 2019 Standard Edition, 2022 Standard Edition, 2025 Standard Edition, and SQL Server Enterprise instance instances.
	//
	// > - You can call [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) to query active region IDs.
	//
	// example:
	//
	// oss-ap-southeast-1.aliyuncs.com:****:key.pvk
	PrivateKey           *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The global resource descriptor of the RAM role. The resource descriptor is used to specify a RAM role. For details, see [RAM role overview](https://help.aliyun.com/document_detail/93689.html).
	//
	// > This parameter is available only for ApsaraDB RDS for MySQL and ApsaraDB RDS for PostgreSQL instances.
	//
	// example:
	//
	// acs:ram::1406926****:role/aliyunrdsinstanceencryptiondefaultrole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The TDE status. Valid values:
	//
	// - **Enabled*	-
	//
	// - **Disabled**
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	TDEStatus *string `json:"TDEStatus,omitempty" xml:"TDEStatus,omitempty"`
}

func (s ModifyDBInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceTDERequest) GetCertificate() *string {
	return s.Certificate
}

func (s *ModifyDBInstanceTDERequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceTDERequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyDBInstanceTDERequest) GetEncryptionKey() *string {
	return s.EncryptionKey
}

func (s *ModifyDBInstanceTDERequest) GetIsRotate() *bool {
	return s.IsRotate
}

func (s *ModifyDBInstanceTDERequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceTDERequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceTDERequest) GetPassWord() *string {
	return s.PassWord
}

func (s *ModifyDBInstanceTDERequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *ModifyDBInstanceTDERequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceTDERequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceTDERequest) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ModifyDBInstanceTDERequest) GetTDEStatus() *string {
	return s.TDEStatus
}

func (s *ModifyDBInstanceTDERequest) SetCertificate(v string) *ModifyDBInstanceTDERequest {
	s.Certificate = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetDBInstanceId(v string) *ModifyDBInstanceTDERequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetDBName(v string) *ModifyDBInstanceTDERequest {
	s.DBName = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetEncryptionKey(v string) *ModifyDBInstanceTDERequest {
	s.EncryptionKey = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetIsRotate(v bool) *ModifyDBInstanceTDERequest {
	s.IsRotate = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetPassWord(v string) *ModifyDBInstanceTDERequest {
	s.PassWord = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetPrivateKey(v string) *ModifyDBInstanceTDERequest {
	s.PrivateKey = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetResourceOwnerId(v int64) *ModifyDBInstanceTDERequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetRoleArn(v string) *ModifyDBInstanceTDERequest {
	s.RoleArn = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) SetTDEStatus(v string) *ModifyDBInstanceTDERequest {
	s.TDEStatus = &v
	return s
}

func (s *ModifyDBInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
