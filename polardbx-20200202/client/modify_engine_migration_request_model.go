// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEngineMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStrings(v string) *ModifyEngineMigrationRequest
	GetConnectionStrings() *string
	SetDBInstanceName(v string) *ModifyEngineMigrationRequest
	GetDBInstanceName() *string
	SetNewMasterDBInstanceName(v string) *ModifyEngineMigrationRequest
	GetNewMasterDBInstanceName() *string
	SetRegionId(v string) *ModifyEngineMigrationRequest
	GetRegionId() *string
	SetSourceDBInstanceName(v string) *ModifyEngineMigrationRequest
	GetSourceDBInstanceName() *string
	SetSwapConnectionString(v string) *ModifyEngineMigrationRequest
	GetSwapConnectionString() *string
}

type ModifyEngineMigrationRequest struct {
	// example:
	//
	// {\\"pc-bp1m9tt23o5kh834y.rwlb.rds.aliyuncs.com\\":\\"rm-bp1ycl5o6ojs957o0.mysql.rds.aliyuncs.com\\"}
	ConnectionStrings *string `json:"ConnectionStrings,omitempty" xml:"ConnectionStrings,omitempty"`
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// newmaster789
	NewMasterDBInstanceName *string `json:"NewMasterDBInstanceName,omitempty" xml:"NewMasterDBInstanceName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// source456
	SourceDBInstanceName *string `json:"SourceDBInstanceName,omitempty" xml:"SourceDBInstanceName,omitempty"`
	// example:
	//
	// true
	SwapConnectionString *string `json:"SwapConnectionString,omitempty" xml:"SwapConnectionString,omitempty"`
}

func (s ModifyEngineMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEngineMigrationRequest) GoString() string {
	return s.String()
}

func (s *ModifyEngineMigrationRequest) GetConnectionStrings() *string {
	return s.ConnectionStrings
}

func (s *ModifyEngineMigrationRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyEngineMigrationRequest) GetNewMasterDBInstanceName() *string {
	return s.NewMasterDBInstanceName
}

func (s *ModifyEngineMigrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyEngineMigrationRequest) GetSourceDBInstanceName() *string {
	return s.SourceDBInstanceName
}

func (s *ModifyEngineMigrationRequest) GetSwapConnectionString() *string {
	return s.SwapConnectionString
}

func (s *ModifyEngineMigrationRequest) SetConnectionStrings(v string) *ModifyEngineMigrationRequest {
	s.ConnectionStrings = &v
	return s
}

func (s *ModifyEngineMigrationRequest) SetDBInstanceName(v string) *ModifyEngineMigrationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyEngineMigrationRequest) SetNewMasterDBInstanceName(v string) *ModifyEngineMigrationRequest {
	s.NewMasterDBInstanceName = &v
	return s
}

func (s *ModifyEngineMigrationRequest) SetRegionId(v string) *ModifyEngineMigrationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyEngineMigrationRequest) SetSourceDBInstanceName(v string) *ModifyEngineMigrationRequest {
	s.SourceDBInstanceName = &v
	return s
}

func (s *ModifyEngineMigrationRequest) SetSwapConnectionString(v string) *ModifyEngineMigrationRequest {
	s.SwapConnectionString = &v
	return s
}

func (s *ModifyEngineMigrationRequest) Validate() error {
	return dara.Validate(s)
}
