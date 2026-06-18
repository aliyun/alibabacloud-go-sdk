// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseEngineMigrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinueEnableBinlog(v string) *CloseEngineMigrationRequest
	GetContinueEnableBinlog() *string
	SetDBInstanceName(v string) *CloseEngineMigrationRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *CloseEngineMigrationRequest
	GetRegionId() *string
}

type CloseEngineMigrationRequest struct {
	// Specifies whether to keep binlog enabled. Valid values:
	//
	// 	- **true**: Keeps binlog enabled.
	//
	// 	- **false**: Disables binlog.
	//
	// Default value: **true**.
	//
	// > Disabling binlog restarts your PolarDB cluster.
	//
	// example:
	//
	// true
	ContinueEnableBinlog *string `json:"ContinueEnableBinlog,omitempty" xml:"ContinueEnableBinlog,omitempty"`
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-htri0****r4k9p
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CloseEngineMigrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseEngineMigrationRequest) GoString() string {
	return s.String()
}

func (s *CloseEngineMigrationRequest) GetContinueEnableBinlog() *string {
	return s.ContinueEnableBinlog
}

func (s *CloseEngineMigrationRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CloseEngineMigrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloseEngineMigrationRequest) SetContinueEnableBinlog(v string) *CloseEngineMigrationRequest {
	s.ContinueEnableBinlog = &v
	return s
}

func (s *CloseEngineMigrationRequest) SetDBInstanceName(v string) *CloseEngineMigrationRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CloseEngineMigrationRequest) SetRegionId(v string) *CloseEngineMigrationRequest {
	s.RegionId = &v
	return s
}

func (s *CloseEngineMigrationRequest) Validate() error {
	return dara.Validate(s)
}
