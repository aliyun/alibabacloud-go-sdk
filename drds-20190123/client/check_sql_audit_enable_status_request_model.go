// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlAuditEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *CheckSqlAuditEnableStatusRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *CheckSqlAuditEnableStatusRequest
	GetDrdsInstanceId() *string
}

type CheckSqlAuditEnableStatusRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
}

func (s CheckSqlAuditEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlAuditEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckSqlAuditEnableStatusRequest) GetDbName() *string {
	return s.DbName
}

func (s *CheckSqlAuditEnableStatusRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *CheckSqlAuditEnableStatusRequest) SetDbName(v string) *CheckSqlAuditEnableStatusRequest {
	s.DbName = &v
	return s
}

func (s *CheckSqlAuditEnableStatusRequest) SetDrdsInstanceId(v string) *CheckSqlAuditEnableStatusRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *CheckSqlAuditEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
