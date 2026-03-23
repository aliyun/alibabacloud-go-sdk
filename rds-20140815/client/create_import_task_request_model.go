// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstanceId(v string) *CreateImportTaskRequest
	GetDbInstanceId() *string
	SetEstimatedSize(v int32) *CreateImportTaskRequest
	GetEstimatedSize() *int32
	SetHost(v string) *CreateImportTaskRequest
	GetHost() *string
	SetOwnerId(v int64) *CreateImportTaskRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateImportTaskRequest
	GetPassword() *string
	SetPort(v int32) *CreateImportTaskRequest
	GetPort() *int32
	SetRegionId(v string) *CreateImportTaskRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *CreateImportTaskRequest
	GetSourceInstanceId() *string
	SetSourcePlatform(v string) *CreateImportTaskRequest
	GetSourcePlatform() *string
	SetStreamPort(v int32) *CreateImportTaskRequest
	GetStreamPort() *int32
	SetUser(v string) *CreateImportTaskRequest
	GetUser() *string
	SetXtrabackupPath(v string) *CreateImportTaskRequest
	GetXtrabackupPath() *string
}

type CreateImportTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1u*****ggm7j9j
	DbInstanceId *string `json:"DbInstanceId,omitempty" xml:"DbInstanceId,omitempty"`
	// Estimated data space, in GB
	//
	// example:
	//
	// 1000
	EstimatedSize *int32 `json:"EstimatedSize,omitempty" xml:"EstimatedSize,omitempty"`
	// The source MySQL host IP address. RDS will access this IP address to retrieve the backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172.20.246.90
	Host    *string `json:"Host,omitempty" xml:"Host,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The Password of the source MySQL Account, which must be Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// OEF5JjVOM2pzZXFKRw==
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Source MySQL port
	//
	// This parameter is required.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The Region ID. You can invoke [DescribeRegions](https://help.aliyun.com/document_detail/610399.html) to obtain it.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source cloud instance ID.
	//
	// example:
	//
	// i-bp1fe296n52ub3chezpg
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// Source cloud instance type
	//
	// example:
	//
	// ECS
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// Stream port used for backup transmission
	//
	// This parameter is required.
	//
	// example:
	//
	// 9999
	StreamPort *int32 `json:"StreamPort,omitempty" xml:"StreamPort,omitempty"`
	// Source MySQL account, which must have permissions to create backups and set up replication. Refer to the following SQL for granting permissions:
	//
	// ```
	//
	// -- MySQL 5.7
	//
	// mysql> CREATE USER \\"myadmin\\"@\\"%\\" IDENTIFIED BY \\"s3cret\\";
	//
	// mysql> GRANT RELOAD, LOCK TABLES, PROCESS, REPLICATION CLIENT, REPLICATION SLAVE ON *.	- TO
	//
	//        \\"myadmin\\"@\\"%\\";
	//
	// mysql> FLUSH PRIVILEGES;
	//
	// -- MySQL 8.0
	//
	// mysql> CREATE USER \\"myadmin\\"@\\"%\\" IDENTIFIED BY \\"Test123!\\";
	//
	// mysql> GRANT BACKUP_ADMIN, PROCESS, RELOAD, LOCK TABLES, REPLICATION CLIENT, REPLICATION SLAVE ON *.	- TO \\"myadmin\\"@\\"%\\";
	//
	// mysql> GRANT SELECT ON performance_schema.log_status TO \\"myadmin\\"@\\"%\\";
	//
	// mysql> GRANT SELECT ON performance_schema.keyring_component_status TO myadmin@\\"%\\";
	//
	// mysql> GRANT SELECT ON performance_schema.replication_group_members TO myadmin@\\"%\\";
	//
	// mysql> FLUSH PRIVILEGES;
	//
	// ```
	//
	// This parameter is required.
	//
	// example:
	//
	// myadmin
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// Installation path of xtrabackup on the source
	//
	// example:
	//
	// /usr/bin/xtrabackup
	XtrabackupPath *string `json:"XtrabackupPath,omitempty" xml:"XtrabackupPath,omitempty"`
}

func (s CreateImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateImportTaskRequest) GetDbInstanceId() *string {
	return s.DbInstanceId
}

func (s *CreateImportTaskRequest) GetEstimatedSize() *int32 {
	return s.EstimatedSize
}

func (s *CreateImportTaskRequest) GetHost() *string {
	return s.Host
}

func (s *CreateImportTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImportTaskRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateImportTaskRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImportTaskRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *CreateImportTaskRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CreateImportTaskRequest) GetStreamPort() *int32 {
	return s.StreamPort
}

func (s *CreateImportTaskRequest) GetUser() *string {
	return s.User
}

func (s *CreateImportTaskRequest) GetXtrabackupPath() *string {
	return s.XtrabackupPath
}

func (s *CreateImportTaskRequest) SetDbInstanceId(v string) *CreateImportTaskRequest {
	s.DbInstanceId = &v
	return s
}

func (s *CreateImportTaskRequest) SetEstimatedSize(v int32) *CreateImportTaskRequest {
	s.EstimatedSize = &v
	return s
}

func (s *CreateImportTaskRequest) SetHost(v string) *CreateImportTaskRequest {
	s.Host = &v
	return s
}

func (s *CreateImportTaskRequest) SetOwnerId(v int64) *CreateImportTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImportTaskRequest) SetPassword(v string) *CreateImportTaskRequest {
	s.Password = &v
	return s
}

func (s *CreateImportTaskRequest) SetPort(v int32) *CreateImportTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateImportTaskRequest) SetRegionId(v string) *CreateImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImportTaskRequest) SetSourceInstanceId(v string) *CreateImportTaskRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateImportTaskRequest) SetSourcePlatform(v string) *CreateImportTaskRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CreateImportTaskRequest) SetStreamPort(v int32) *CreateImportTaskRequest {
	s.StreamPort = &v
	return s
}

func (s *CreateImportTaskRequest) SetUser(v string) *CreateImportTaskRequest {
	s.User = &v
	return s
}

func (s *CreateImportTaskRequest) SetXtrabackupPath(v string) *CreateImportTaskRequest {
	s.XtrabackupPath = &v
	return s
}

func (s *CreateImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
