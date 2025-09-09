// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSqlAuditRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DisableSqlAuditRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DisableSqlAuditRequest
	GetDrdsInstanceId() *string
}

type DisableSqlAuditRequest struct {
	// The name of the database for which you want to disable the SQL audit feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s DisableSqlAuditRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSqlAuditRequest) GoString() string {
	return s.String()
}

func (s *DisableSqlAuditRequest) GetDbName() *string {
	return s.DbName
}

func (s *DisableSqlAuditRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DisableSqlAuditRequest) SetDbName(v string) *DisableSqlAuditRequest {
	s.DbName = &v
	return s
}

func (s *DisableSqlAuditRequest) SetDrdsInstanceId(v string) *DisableSqlAuditRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DisableSqlAuditRequest) Validate() error {
	return dara.Validate(s)
}
