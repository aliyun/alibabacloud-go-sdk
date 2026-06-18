// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlAuditInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditAccountName(v string) *DescribeSqlAuditInfoRequest
	GetAuditAccountName() *string
	SetAuditAccountPassword(v string) *DescribeSqlAuditInfoRequest
	GetAuditAccountPassword() *string
	SetDBInstanceId(v string) *DescribeSqlAuditInfoRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeSqlAuditInfoRequest
	GetRegionId() *string
}

type DescribeSqlAuditInfoRequest struct {
	// The name of the audit administrator account. > This parameter is required if the three-role mode is enabled. For more information, see [Three-role mode](https://help.aliyun.com/document_detail/213824.html).
	//
	// example:
	//
	// test_daa
	AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	// The password of the audit administrator account. > This parameter is required if the three-role mode is enabled. For more information about the three-role mode, see [Three-role mode](https://help.aliyun.com/document_detail/213824.html).
	//
	// example:
	//
	// ******
	AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
	// The instance ID. > You can call [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) to query the details of all instances in the specified region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the region where the instance resides. > You can call [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSqlAuditInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlAuditInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSqlAuditInfoRequest) GetAuditAccountName() *string {
	return s.AuditAccountName
}

func (s *DescribeSqlAuditInfoRequest) GetAuditAccountPassword() *string {
	return s.AuditAccountPassword
}

func (s *DescribeSqlAuditInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSqlAuditInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSqlAuditInfoRequest) SetAuditAccountName(v string) *DescribeSqlAuditInfoRequest {
	s.AuditAccountName = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetAuditAccountPassword(v string) *DescribeSqlAuditInfoRequest {
	s.AuditAccountPassword = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetDBInstanceId(v string) *DescribeSqlAuditInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) SetRegionId(v string) *DescribeSqlAuditInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSqlAuditInfoRequest) Validate() error {
	return dara.Validate(s)
}
