// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest
	GetAuditLogStatus() *string
	SetDBClusterId(v string) *ModifyAuditLogConfigRequest
	GetDBClusterId() *string
	SetEngineType(v string) *ModifyAuditLogConfigRequest
	GetEngineType() *string
	SetOwnerAccount(v string) *ModifyAuditLogConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAuditLogConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyAuditLogConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyAuditLogConfigRequest struct {
	// The status of SQL audit logging. Valid values:
	//
	// - **on**: Enables SQL audit logging.
	//
	// - **off**: Disables SQL audit logging.
	//
	// > After SQL audit logging is disabled, all SQL audit logs are deleted. Query and export the SQL audit logs before disabling SQL audit logging. For more information, see [DescribeAuditLogRecords](https://help.aliyun.com/document_detail/612426.html). When SQL audit logging is enabled again, audit logs are displayed starting from the most recent time that audit logging was enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	AuditLogStatus *string `json:"AuditLogStatus,omitempty" xml:"AuditLogStatus,omitempty"`
	// <props="china">The ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The ID of the Data Lakehouse Edition cluster.
	//
	// > You can call [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) to query the IDs of all clusters in a specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-t4nj8619bz2w3****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of the compute engine. Valid values:
	//
	// - XIHE (**default**): Xihe compute engine.
	//
	// - SPARK: Spark compute engine.
	//
	// example:
	//
	// XIHE
	EngineType   *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// > You can call [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) to query the region ID of a specified cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAuditLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogConfigRequest) GetAuditLogStatus() *string {
	return s.AuditLogStatus
}

func (s *ModifyAuditLogConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAuditLogConfigRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *ModifyAuditLogConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAuditLogConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAuditLogConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAuditLogConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAuditLogConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAuditLogConfigRequest) SetAuditLogStatus(v string) *ModifyAuditLogConfigRequest {
	s.AuditLogStatus = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetDBClusterId(v string) *ModifyAuditLogConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetEngineType(v string) *ModifyAuditLogConfigRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetRegionId(v string) *ModifyAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) SetResourceOwnerId(v int64) *ModifyAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
