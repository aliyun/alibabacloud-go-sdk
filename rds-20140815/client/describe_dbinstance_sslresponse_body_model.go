// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetACL(v string) *DescribeDBInstanceSSLResponseBody
	GetACL() *string
	SetCAType(v string) *DescribeDBInstanceSSLResponseBody
	GetCAType() *string
	SetClientCACert(v string) *DescribeDBInstanceSSLResponseBody
	GetClientCACert() *string
	SetClientCACertExpireTime(v string) *DescribeDBInstanceSSLResponseBody
	GetClientCACertExpireTime() *string
	SetClientCertRevocationList(v string) *DescribeDBInstanceSSLResponseBody
	GetClientCertRevocationList() *string
	SetConnectionString(v string) *DescribeDBInstanceSSLResponseBody
	GetConnectionString() *string
	SetForceEncryption(v string) *DescribeDBInstanceSSLResponseBody
	GetForceEncryption() *string
	SetLastModifyStatus(v string) *DescribeDBInstanceSSLResponseBody
	GetLastModifyStatus() *string
	SetModifyStatusReason(v string) *DescribeDBInstanceSSLResponseBody
	GetModifyStatusReason() *string
	SetReplicationACL(v string) *DescribeDBInstanceSSLResponseBody
	GetReplicationACL() *string
	SetRequestId(v string) *DescribeDBInstanceSSLResponseBody
	GetRequestId() *string
	SetRequireUpdate(v string) *DescribeDBInstanceSSLResponseBody
	GetRequireUpdate() *string
	SetRequireUpdateItem(v string) *DescribeDBInstanceSSLResponseBody
	GetRequireUpdateItem() *string
	SetRequireUpdateReason(v string) *DescribeDBInstanceSSLResponseBody
	GetRequireUpdateReason() *string
	SetSSLCreateTime(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLCreateTime() *string
	SetSSLEnabled(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLEnabled() *string
	SetSSLExpireTime(v string) *DescribeDBInstanceSSLResponseBody
	GetSSLExpireTime() *string
	SetServerCAUrl(v string) *DescribeDBInstanceSSLResponseBody
	GetServerCAUrl() *string
	SetServerCert(v string) *DescribeDBInstanceSSLResponseBody
	GetServerCert() *string
	SetServerKey(v string) *DescribeDBInstanceSSLResponseBody
	GetServerKey() *string
	SetTlsVersion(v string) *DescribeDBInstanceSSLResponseBody
	GetTlsVersion() *string
}

type DescribeDBInstanceSSLResponseBody struct {
	// The authentication method of the ApsaraDB RDS for PostgreSQL instance with cloud disks. Valid values:
	//
	// - **cert**
	//
	// - **prefer**
	//
	// - **verify-ca**
	//
	// - **verify-full*	- (supported by ApsaraDB RDS for PostgreSQL 12 and later)
	//
	// example:
	//
	// cert
	ACL *string `json:"ACL,omitempty" xml:"ACL,omitempty"`
	// The server certificate type of the ApsaraDB RDS for PostgreSQL instance with cloud disks. Valid values:
	//
	// - **aliyun**: The cloud certificate is used.
	//
	// - **custom**: A custom certificate is used.
	//
	// example:
	//
	// aliyun
	CAType *string `json:"CAType,omitempty" xml:"CAType,omitempty"`
	// The public key of the client certificate authority (CA) for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****viXk=-----END CERTIFICATE-----
	ClientCACert *string `json:"ClientCACert,omitempty" xml:"ClientCACert,omitempty"`
	// The expiration time of the public key of the client certificate authorization authority (CA) for the ApsaraDB RDS for PostgreSQL instance with cloud disks. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// This parameter is not supported. You can ignore this parameter.
	//
	// example:
	//
	// -
	ClientCACertExpireTime *string `json:"ClientCACertExpireTime,omitempty" xml:"ClientCACertExpireTime,omitempty"`
	// The client certificate revocation certificate file of the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -----BEGIN X509 CRL-----MIIB****19mg==-----END X509 CRL-----
	ClientCertRevocationList *string `json:"ClientCertRevocationList,omitempty" xml:"ClientCertRevocationList,omitempty"`
	// The endpoint that is protected by SSL.
	//
	// example:
	//
	// rm-bp162dfr55g47****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Indicates whether the [forced Secure Sockets Layer (SSL) encryption feature](https://help.aliyun.com/document_detail/95715.html) is enabled for the ApsaraDB RDS for SQL Server instance. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
	//
	// example:
	//
	// 1
	ForceEncryption *string `json:"ForceEncryption,omitempty" xml:"ForceEncryption,omitempty"`
	// The current SSL link configuration status of the ApsaraDB RDS for PostgreSQL instance with cloud disks. Valid values:
	//
	// - **success**: Successful.
	//
	// - **setting**: Being configured.
	//
	// - **failed**: Failed.
	//
	// example:
	//
	// setting
	LastModifyStatus *string `json:"LastModifyStatus,omitempty" xml:"LastModifyStatus,omitempty"`
	// The reason for the current SSL link configuration status of the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// Modify DB Instance SSL Config.
	ModifyStatusReason *string `json:"ModifyStatusReason,omitempty" xml:"ModifyStatusReason,omitempty"`
	// The authentication method for replication permissions of the ApsaraDB RDS for PostgreSQL instance with cloud disks. Valid values:
	//
	// - **cert**
	//
	// - **prefer**
	//
	// - **verify-ca**
	//
	// - **verify-full*	- (supported by ApsaraDB RDS for PostgreSQL 12 and later)
	//
	// example:
	//
	// cert
	ReplicationACL *string `json:"ReplicationACL,omitempty" xml:"ReplicationACL,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7705151C-E242-55AF-9929-2A3C39D979D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the SSL certificate needs to be updated. Valid values:
	//
	// > The SSL certificate is valid for one year. If the certificate is not renewed after it expires, client programs that use encrypted connections cannot connect to the instance.
	//
	// <details>
	//
	// <summary>MySQL and SQL Server</summary>
	//
	// - **No**: No update is required.
	//
	// - **Yes**: An update is required.
	//
	// </details>
	//
	// <details>
	//
	// <summary>PostgreSQL</summary>
	//
	// - **0**: No update is required.
	//
	// - **1**: An update is required.
	//
	// </details>
	//
	// example:
	//
	// Yes
	RequireUpdate *string `json:"RequireUpdate,omitempty" xml:"RequireUpdate,omitempty"`
	// The list of server certificates that need to be updated for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -
	RequireUpdateItem *string `json:"RequireUpdateItem,omitempty" xml:"RequireUpdateItem,omitempty"`
	// The reason why the certificates need to be updated for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -
	RequireUpdateReason *string `json:"RequireUpdateReason,omitempty" xml:"RequireUpdateReason,omitempty"`
	// The creation time of the server certificate for the ApsaraDB RDS for PostgreSQL instance with cloud disks. This parameter is valid only when CAType is set to aliyun.
	//
	// example:
	//
	// -
	SSLCreateTime *string `json:"SSLCreateTime,omitempty" xml:"SSLCreateTime,omitempty"`
	// The SSL encryption status. Valid values:
	//
	// <details>
	//
	// <summary>MySQL and SQL Server</summary>
	//
	// - **Yes**: Enabled.
	//
	// - **No**: Disabled.
	//
	// </details>
	//
	// <details>
	//
	// <summary>PostgreSQL</summary>
	//
	// - **on**: Enabled.
	//
	// - **off**: Disabled.
	//
	// </details>
	//
	// example:
	//
	// Yes
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The expiration time of the SSL certificate. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2025-06-16T08:16:43Z
	SSLExpireTime *string `json:"SSLExpireTime,omitempty" xml:"SSLExpireTime,omitempty"`
	// The URL of the CA certificate that is used to issue the server certificate for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -
	ServerCAUrl *string `json:"ServerCAUrl,omitempty" xml:"ServerCAUrl,omitempty"`
	// The content of the server certificate for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----MIID*****QqEP-----END CERTIFICATE-----
	ServerCert *string `json:"ServerCert,omitempty" xml:"ServerCert,omitempty"`
	// The private key of the server certificate for the ApsaraDB RDS for PostgreSQL instance with cloud disks.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----MIIE****ihfg==-----END PRIVATE KEY-----
	ServerKey *string `json:"ServerKey,omitempty" xml:"ServerKey,omitempty"`
	// The specified [minimum TLS version](https://help.aliyun.com/document_detail/95715.html) for the ApsaraDB RDS for SQL Server instance. Valid values: 1.0, 1.1, and 1.2.
	//
	// example:
	//
	// 1.1
	TlsVersion *string `json:"TlsVersion,omitempty" xml:"TlsVersion,omitempty"`
}

func (s DescribeDBInstanceSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponseBody) GetACL() *string {
	return s.ACL
}

func (s *DescribeDBInstanceSSLResponseBody) GetCAType() *string {
	return s.CAType
}

func (s *DescribeDBInstanceSSLResponseBody) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *DescribeDBInstanceSSLResponseBody) GetClientCACertExpireTime() *string {
	return s.ClientCACertExpireTime
}

func (s *DescribeDBInstanceSSLResponseBody) GetClientCertRevocationList() *string {
	return s.ClientCertRevocationList
}

func (s *DescribeDBInstanceSSLResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceSSLResponseBody) GetForceEncryption() *string {
	return s.ForceEncryption
}

func (s *DescribeDBInstanceSSLResponseBody) GetLastModifyStatus() *string {
	return s.LastModifyStatus
}

func (s *DescribeDBInstanceSSLResponseBody) GetModifyStatusReason() *string {
	return s.ModifyStatusReason
}

func (s *DescribeDBInstanceSSLResponseBody) GetReplicationACL() *string {
	return s.ReplicationACL
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequireUpdate() *string {
	return s.RequireUpdate
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequireUpdateItem() *string {
	return s.RequireUpdateItem
}

func (s *DescribeDBInstanceSSLResponseBody) GetRequireUpdateReason() *string {
	return s.RequireUpdateReason
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLCreateTime() *string {
	return s.SSLCreateTime
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeDBInstanceSSLResponseBody) GetSSLExpireTime() *string {
	return s.SSLExpireTime
}

func (s *DescribeDBInstanceSSLResponseBody) GetServerCAUrl() *string {
	return s.ServerCAUrl
}

func (s *DescribeDBInstanceSSLResponseBody) GetServerCert() *string {
	return s.ServerCert
}

func (s *DescribeDBInstanceSSLResponseBody) GetServerKey() *string {
	return s.ServerKey
}

func (s *DescribeDBInstanceSSLResponseBody) GetTlsVersion() *string {
	return s.TlsVersion
}

func (s *DescribeDBInstanceSSLResponseBody) SetACL(v string) *DescribeDBInstanceSSLResponseBody {
	s.ACL = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetCAType(v string) *DescribeDBInstanceSSLResponseBody {
	s.CAType = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetClientCACert(v string) *DescribeDBInstanceSSLResponseBody {
	s.ClientCACert = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetClientCACertExpireTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.ClientCACertExpireTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetClientCertRevocationList(v string) *DescribeDBInstanceSSLResponseBody {
	s.ClientCertRevocationList = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetConnectionString(v string) *DescribeDBInstanceSSLResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetForceEncryption(v string) *DescribeDBInstanceSSLResponseBody {
	s.ForceEncryption = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetLastModifyStatus(v string) *DescribeDBInstanceSSLResponseBody {
	s.LastModifyStatus = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetModifyStatusReason(v string) *DescribeDBInstanceSSLResponseBody {
	s.ModifyStatusReason = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetReplicationACL(v string) *DescribeDBInstanceSSLResponseBody {
	s.ReplicationACL = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequestId(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequireUpdate(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequireUpdate = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequireUpdateItem(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequireUpdateItem = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetRequireUpdateReason(v string) *DescribeDBInstanceSSLResponseBody {
	s.RequireUpdateReason = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLCreateTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLCreateTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLEnabled(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetSSLExpireTime(v string) *DescribeDBInstanceSSLResponseBody {
	s.SSLExpireTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetServerCAUrl(v string) *DescribeDBInstanceSSLResponseBody {
	s.ServerCAUrl = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetServerCert(v string) *DescribeDBInstanceSSLResponseBody {
	s.ServerCert = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetServerKey(v string) *DescribeDBInstanceSSLResponseBody {
	s.ServerKey = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) SetTlsVersion(v string) *DescribeDBInstanceSSLResponseBody {
	s.TlsVersion = &v
	return s
}

func (s *DescribeDBInstanceSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
