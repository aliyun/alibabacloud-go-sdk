// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUpgradeReportForSyncCloneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationCategory(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetCreationCategory() *string
	SetCreationOption(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetCreationOption() *string
	SetDBName(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetDBName() *string
	SetDBType(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetDBType() *string
	SetDBVersion(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GenerateUpgradeReportForSyncCloneRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetRegionId() *string
	SetReserve(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetReserve() *string
	SetResourceOwnerAccount(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GenerateUpgradeReportForSyncCloneRequest
	GetResourceOwnerId() *int64
	SetSourceDBClusterId(v string) *GenerateUpgradeReportForSyncCloneRequest
	GetSourceDBClusterId() *string
}

type GenerateUpgradeReportForSyncCloneRequest struct {
	// The edition of the cluster. Valid values:
	//
	// 	- **Normal**: Cluster Edition (default)
	//
	// 	- **SENormal**: Standard Edition
	//
	// For more information about cluster editions, see [Product editions](https://help.aliyun.com/document_detail/183258.html).
	//
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// The creation pattern. Valid values:
	//
	// 	- **MigrationFromRDS**: migrates data from an existing ApsaraDB RDS instance to a new PolarDB cluster. The created PolarDB cluster is in read-only mode and has binary logging enabled by default. For more information about the console operation, see [Upgrade an ApsaraDB RDS for MySQL instance to a PolarDB for MySQL cluster](https://help.aliyun.com/document_detail/121582.html).
	//
	// 	- **UpgradeFromPolarDB**: performs instance migration from PolarDB. For more information about the console operation, see [Major engine version upgrade](https://help.aliyun.com/document_detail/459712.html).
	//
	// example:
	//
	// MigrationFromRDS
	CreationOption *string `json:"CreationOption,omitempty" xml:"CreationOption,omitempty"`
	// The name of the database. You can specify only one database name.
	//
	// > Only PolarDB for PostgreSQL (Compatible with Oracle) clusters support this parameter.
	//
	// example:
	//
	// testDB
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The type of the database engine. Valid values:
	//
	// - **MySQL**
	//
	// - **Oracle**
	//
	// example:
	//
	// PostgreSQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version number of the destination database DPI engine.
	//
	// 	- Valid values for MySQL:
	//
	//     	- **5.6**
	//
	//     	- **5.7**
	//
	//     	- **8.0**
	//
	// 	- Valid value for Oracle: **14**.
	//
	// example:
	//
	// 5.6
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The reserved parameter. The value is a JSON string.
	//
	// example:
	//
	// {\\"targetTableMode\\":2}
	Reserve              *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source instance ID.
	//
	// example:
	//
	// pc-k2j96w169uhu868l8
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
}

func (s GenerateUpgradeReportForSyncCloneRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUpgradeReportForSyncCloneRequest) GoString() string {
	return s.String()
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetCreationOption() *string {
	return s.CreationOption
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetDBName() *string {
	return s.DBName
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetDBType() *string {
	return s.DBType
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetReserve() *string {
	return s.Reserve
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GenerateUpgradeReportForSyncCloneRequest) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetCreationCategory(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.CreationCategory = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetCreationOption(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.CreationOption = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetDBName(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.DBName = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetDBType(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.DBType = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetDBVersion(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.DBVersion = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetOwnerAccount(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetOwnerId(v int64) *GenerateUpgradeReportForSyncCloneRequest {
	s.OwnerId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetRegionId(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetReserve(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.Reserve = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetResourceOwnerAccount(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetResourceOwnerId(v int64) *GenerateUpgradeReportForSyncCloneRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) SetSourceDBClusterId(v string) *GenerateUpgradeReportForSyncCloneRequest {
	s.SourceDBClusterId = &v
	return s
}

func (s *GenerateUpgradeReportForSyncCloneRequest) Validate() error {
	return dara.Validate(s)
}
