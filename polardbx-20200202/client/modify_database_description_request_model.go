// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyDatabaseDescriptionRequest
	GetDBInstanceName() *string
	SetDbDescription(v string) *ModifyDatabaseDescriptionRequest
	GetDbDescription() *string
	SetDbName(v string) *ModifyDatabaseDescriptionRequest
	GetDbName() *string
	SetRegionId(v string) *ModifyDatabaseDescriptionRequest
	GetRegionId() *string
}

type ModifyDatabaseDescriptionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	DbDescription *string `json:"DbDescription,omitempty" xml:"DbDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDatabaseDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDatabaseDescriptionRequest) GetDbDescription() *string {
	return s.DbDescription
}

func (s *ModifyDatabaseDescriptionRequest) GetDbName() *string {
	return s.DbName
}

func (s *ModifyDatabaseDescriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabaseDescriptionRequest) SetDBInstanceName(v string) *ModifyDatabaseDescriptionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbDescription(v string) *ModifyDatabaseDescriptionRequest {
	s.DbDescription = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDbName(v string) *ModifyDatabaseDescriptionRequest {
	s.DbName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetRegionId(v string) *ModifyDatabaseDescriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
