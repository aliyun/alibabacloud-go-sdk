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
	// example:
	//
	// test_daa
	AuditAccountName *string `json:"AuditAccountName,omitempty" xml:"AuditAccountName,omitempty"`
	// example:
	//
	// Pw@11111
	AuditAccountPassword *string `json:"AuditAccountPassword,omitempty" xml:"AuditAccountPassword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
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
