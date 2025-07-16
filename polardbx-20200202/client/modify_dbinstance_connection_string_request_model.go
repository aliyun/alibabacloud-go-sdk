// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetConnectionString() *string
	SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceName() *string
	SetNewPort(v string) *ModifyDBInstanceConnectionStringRequest
	GetNewPort() *string
	SetNewPrefix(v string) *ModifyDBInstanceConnectionStringRequest
	GetNewPrefix() *string
	SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest
	GetRegionId() *string
}

type ModifyDBInstanceConnectionStringRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8.polarx.huhehaote.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-unrf5ssig0ecg8
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3300
	NewPort *string `json:"NewPort,omitempty" xml:"NewPort,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test2
	NewPrefix *string `json:"NewPrefix,omitempty" xml:"NewPrefix,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNewPort() *string {
	return s.NewPort
}

func (s *ModifyDBInstanceConnectionStringRequest) GetNewPrefix() *string {
	return s.NewPrefix
}

func (s *ModifyDBInstanceConnectionStringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceName(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewPort = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetNewPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.NewPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
