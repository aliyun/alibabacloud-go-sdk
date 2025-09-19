// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPreCheckDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseType(v string) *StartPreCheckDatabaseRequest
	GetDatabaseType() *string
	SetInstanceUuid(v string) *StartPreCheckDatabaseRequest
	GetInstanceUuid() *string
	SetUniRegionId(v string) *StartPreCheckDatabaseRequest
	GetUniRegionId() *string
}

type StartPreCheckDatabaseRequest struct {
	// The type of the database. Valid values:
	//
	// 	- **MYSQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// This parameter is required.
	//
	// example:
	//
	// MYSQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The UUID of the agent that is used to back up the data of the database.
	//
	// > You can call the [DescribeUniBackupDatabase](~~DescribeUniBackupDatabase~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ec1c0ba21d2911ed800000163e0e****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The region ID of the server that hosts the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
}

func (s StartPreCheckDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s StartPreCheckDatabaseRequest) GoString() string {
	return s.String()
}

func (s *StartPreCheckDatabaseRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *StartPreCheckDatabaseRequest) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *StartPreCheckDatabaseRequest) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *StartPreCheckDatabaseRequest) SetDatabaseType(v string) *StartPreCheckDatabaseRequest {
	s.DatabaseType = &v
	return s
}

func (s *StartPreCheckDatabaseRequest) SetInstanceUuid(v string) *StartPreCheckDatabaseRequest {
	s.InstanceUuid = &v
	return s
}

func (s *StartPreCheckDatabaseRequest) SetUniRegionId(v string) *StartPreCheckDatabaseRequest {
	s.UniRegionId = &v
	return s
}

func (s *StartPreCheckDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
