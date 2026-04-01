// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetACL(v string) *ModifyDBInstanceSSLRequest
	GetACL() *string
	SetCAType(v string) *ModifyDBInstanceSSLRequest
	GetCAType() *string
	SetCertificate(v string) *ModifyDBInstanceSSLRequest
	GetCertificate() *string
	SetClientCACert(v string) *ModifyDBInstanceSSLRequest
	GetClientCACert() *string
	SetClientCAEnabled(v int32) *ModifyDBInstanceSSLRequest
	GetClientCAEnabled() *int32
	SetClientCertRevocationList(v string) *ModifyDBInstanceSSLRequest
	GetClientCertRevocationList() *string
	SetClientCrlEnabled(v int32) *ModifyDBInstanceSSLRequest
	GetClientCrlEnabled() *int32
	SetConnectionString(v string) *ModifyDBInstanceSSLRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest
	GetDBInstanceId() *string
	SetForceEncryption(v string) *ModifyDBInstanceSSLRequest
	GetForceEncryption() *string
	SetOwnerAccount(v string) *ModifyDBInstanceSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceSSLRequest
	GetOwnerId() *int64
	SetPassWord(v string) *ModifyDBInstanceSSLRequest
	GetPassWord() *string
	SetReplicationACL(v string) *ModifyDBInstanceSSLRequest
	GetReplicationACL() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceSSLRequest
	GetResourceOwnerId() *int64
	SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest
	GetSSLEnabled() *int32
	SetServerCert(v string) *ModifyDBInstanceSSLRequest
	GetServerCert() *string
	SetServerKey(v string) *ModifyDBInstanceSSLRequest
	GetServerKey() *string
	SetTlsVersion(v string) *ModifyDBInstanceSSLRequest
	GetTlsVersion() *string
}

type ModifyDBInstanceSSLRequest struct {
	// The method that is used to verify the identities of clients. This parameter is supported only when the instance runs PostgreSQL with cloud disks. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. Valid values:
	//
	// 	- **cert**
	//
	// 	- **prefer**
	//
	// 	- **verify-ca**
	//
	// 	- **verify-full*	- (supported only when the instance runs PostgreSQL 12 or later)
	//
	// example:
	//
	// cert
	ACL *string `json:"ACL,omitempty" xml:"ACL,omitempty"`
	// The type of the server certificate. This parameter is supported only when the instance runs MySQL or PostgreSQL with cloud disks. If you set SSLEnabled to **1**, the default value of this parameter is **aliyun**. Valid values:
	//
	// 	- **aliyun**: a cloud certificate
	//
	// 	- **custom**: a custom certificate
	//
	// example:
	//
	// aliyun
	CAType *string `json:"CAType,omitempty" xml:"CAType,omitempty"`
	// The custom certificate. The custom certificate is in the `PFX` format.
	//
	// 	- Public endpoint: `oss-<The ID of the region>.aliyuncs.com:<The name of the bucket>:<The name of the certificate file (The file name contains the extension.)>`
	//
	// 	- Internal endpoint: `oss-<The ID of the region>-internal.aliyuncs.com:<The name of the bucket>:<The name of the certificate file (The file name contains the extension.)>`
	//
	// example:
	//
	// oss-cn-beijing-internal.aliyuncs.com:zhttest:test.pfx
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with cloud disks. This parameter must be specified when ClientCAEbabled is set to **1**.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****viXk=-----END CERTIFICATE-----
	ClientCACert *string `json:"ClientCACert,omitempty" xml:"ClientCACert,omitempty"`
	// Specifies whether to enable the public key of the CA that issues client certificates. This parameter is supported only when the instance runs PostgreSQL with cloud disks. Valid values:
	//
	// 	- **1**: enables the public key.
	//
	// 	- **0**: disables the public key.
	//
	// example:
	//
	// 1
	ClientCAEnabled *int32 `json:"ClientCAEnabled,omitempty" xml:"ClientCAEnabled,omitempty"`
	// The CRL that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with cloud disks. This parameter must be specified when ClientCrlEnabled is set to **1**.
	//
	// example:
	//
	// -----BEGIN X509 CRL-----MIIB****19mg==-----END X509 CRL-----
	ClientCertRevocationList *string `json:"ClientCertRevocationList,omitempty" xml:"ClientCertRevocationList,omitempty"`
	// Specifies whether to enable a certificate revocation list (CRL) that contains revoked client certificates. This parameter is supported only when the instance runs PostgreSQL with cloud disks. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. Valid values:
	//
	// 	- **1**: enables the CRL.
	//
	// 	- **0**: disables the CRL.
	//
	// example:
	//
	// 1
	ClientCrlEnabled *int32 `json:"ClientCrlEnabled,omitempty" xml:"ClientCrlEnabled,omitempty"`
	// The internal or public endpoint for which the server certificate needs to be created or updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxx.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to enable the forceful SSL encryption feature. This parameter is supported only for ApsaraDB RDS for SQL Server instances. For more information, see [Configure the SSL encryption feature](https://help.aliyun.com/document_detail/95715.html). Valid values:
	//
	// 	- **1**: enables the feature.
	//
	// 	- **0**: disables the feature.
	//
	// example:
	//
	// 1
	ForceEncryption *string `json:"ForceEncryption,omitempty" xml:"ForceEncryption,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the certificate.
	//
	// example:
	//
	// zht123456
	PassWord *string `json:"PassWord,omitempty" xml:"PassWord,omitempty"`
	// The method that is used to verify the replication permission. This parameter is supported only when the instance runs PostgreSQL with cloud disks. In addition, this parameter is available only when the public key of the CA that issues client certificates is enabled. Valid values:
	//
	// 	- **cert**
	//
	// 	- **prefer**
	//
	// 	- **verify-ca**
	//
	// 	- **verify-full*	- (supported only when the instance runs PostgreSQL 12 or later)
	//
	// example:
	//
	// cert
	ReplicationACL       *string `json:"ReplicationACL,omitempty" xml:"ReplicationACL,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable or disable the SSL encryption feature. Valid values:
	//
	// 	- **1**: enables the feature.
	//
	// 	- **0**: disables the feature.
	//
	// example:
	//
	// 1
	SSLEnabled *int32 `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The content of the server certificate. This parameter is supported only when the instance runs PostgreSQL with cloud disks. This parameter must be specified when CAType is set to **custom**.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****QqEP-----END CERTIFICATE-----
	ServerCert *string `json:"ServerCert,omitempty" xml:"ServerCert,omitempty"`
	// The private key of the server certificate. This parameter is supported only when the instance runs PostgreSQL with cloud disks. This parameter must be specified when CAType is set to **custom**.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----MIIE****ihfg==-----END PRIVATE KEY-----
	ServerKey *string `json:"ServerKey,omitempty" xml:"ServerKey,omitempty"`
	// The minimum Transport Layer Security (TLS) version. Valid values: 1.0, 1.1, and 1.2. This parameter is supported only for ApsaraDB RDS for SQL Server instances. For more information, see [Configure the SSL encryption feature](https://help.aliyun.com/document_detail/95715.html).
	//
	// example:
	//
	// 1.1
	TlsVersion *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) GetACL() *string {
	return s.ACL
}

func (s *ModifyDBInstanceSSLRequest) GetCAType() *string {
	return s.CAType
}

func (s *ModifyDBInstanceSSLRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *ModifyDBInstanceSSLRequest) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *ModifyDBInstanceSSLRequest) GetClientCAEnabled() *int32 {
	return s.ClientCAEnabled
}

func (s *ModifyDBInstanceSSLRequest) GetClientCertRevocationList() *string {
	return s.ClientCertRevocationList
}

func (s *ModifyDBInstanceSSLRequest) GetClientCrlEnabled() *int32 {
	return s.ClientCrlEnabled
}

func (s *ModifyDBInstanceSSLRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceSSLRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceSSLRequest) GetForceEncryption() *string {
	return s.ForceEncryption
}

func (s *ModifyDBInstanceSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceSSLRequest) GetPassWord() *string {
	return s.PassWord
}

func (s *ModifyDBInstanceSSLRequest) GetReplicationACL() *string {
	return s.ReplicationACL
}

func (s *ModifyDBInstanceSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceSSLRequest) GetSSLEnabled() *int32 {
	return s.SSLEnabled
}

func (s *ModifyDBInstanceSSLRequest) GetServerCert() *string {
	return s.ServerCert
}

func (s *ModifyDBInstanceSSLRequest) GetServerKey() *string {
	return s.ServerKey
}

func (s *ModifyDBInstanceSSLRequest) GetTlsVersion() *string {
	return s.TlsVersion
}

func (s *ModifyDBInstanceSSLRequest) SetACL(v string) *ModifyDBInstanceSSLRequest {
	s.ACL = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetCAType(v string) *ModifyDBInstanceSSLRequest {
	s.CAType = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetCertificate(v string) *ModifyDBInstanceSSLRequest {
	s.Certificate = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetClientCACert(v string) *ModifyDBInstanceSSLRequest {
	s.ClientCACert = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetClientCAEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.ClientCAEnabled = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetClientCertRevocationList(v string) *ModifyDBInstanceSSLRequest {
	s.ClientCertRevocationList = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetClientCrlEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.ClientCrlEnabled = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetForceEncryption(v string) *ModifyDBInstanceSSLRequest {
	s.ForceEncryption = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetPassWord(v string) *ModifyDBInstanceSSLRequest {
	s.PassWord = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetReplicationACL(v string) *ModifyDBInstanceSSLRequest {
	s.ReplicationACL = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int32) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetServerCert(v string) *ModifyDBInstanceSSLRequest {
	s.ServerCert = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetServerKey(v string) *ModifyDBInstanceSSLRequest {
	s.ServerKey = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetTlsVersion(v string) *ModifyDBInstanceSSLRequest {
	s.TlsVersion = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
