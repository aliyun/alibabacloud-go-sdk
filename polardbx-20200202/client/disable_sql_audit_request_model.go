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
