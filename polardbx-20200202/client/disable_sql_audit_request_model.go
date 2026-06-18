// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditAccountName(v string) *DisableSqlAuditRequest
	GetAuditAccountName() *string
	SetAuditAccountPassword(v string) *DisableSqlAuditRequest
	GetAuditAccountPassword() *string
	SetDBInstanceId(v string) *DisableSqlAuditRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DisableSqlAuditRequest
	GetRegionId() *string
}

type DisableSqlAuditRequest struct {
	// The name of the audit administrator account. > If the three-role mode is enabled, this parameter is required. For more information about the three-role mode, see [Three-role mode](https://help.aliyun.com/document_detail/213824.html).
	//
	// example:
	//
	// test_daa
	AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	// The password of the audit administrator account. > If the three-role mode is enabled, this parameter is required. For more information about the three-role mode, see [Three-role mode](https://help.aliyun.com/document_detail/213824.html).
	//
	// example:
	//
	// Pw@11111
	AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
	// The instance ID. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID of the instance. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableSqlAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditRequest) GetAuditAccountName() *string {
	return s.AuditAccountName
}

func (s *DisableSqlAuditRequest) GetAuditAccountPassword() *string {
	return s.AuditAccountPassword
}

func (s *DisableSqlAuditRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DisableSqlAuditRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableSqlAuditRequest) SetAuditAccountName(v string) *DisableSqlAuditRequest {
	s.AuditAccountName = &v
	return s
}

func (s *DisableSqlAuditRequest) SetAuditAccountPassword(v string) *DisableSqlAuditRequest {
	s.AuditAccountPassword = &v
	return s
}

func (s *DisableSqlAuditRequest) SetDBInstanceId(v string) *DisableSqlAuditRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DisableSqlAuditRequest) SetRegionId(v string) *DisableSqlAuditRequest {
	s.RegionId = &v
	return s
}

func (s *DisableSqlAuditRequest) Validate() error {
	return dara.Validate(s)
}
