// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupBroadcastTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *SetupBroadcastTablesRequest
	GetActive() *bool
	SetDbName(v string) *SetupBroadcastTablesRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SetupBroadcastTablesRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *SetupBroadcastTablesRequest
	GetRegionId() *string
	SetTableName(v []*string) *SetupBroadcastTablesRequest
	GetTableName() []*string
}

type SetupBroadcastTablesRequest struct {
	// Specifies whether to activate a broadcast table for the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The name of the database for which you want to configure a broadcast table.
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
	// The ID of the region in which the PolarDB-X 1.0 instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupBroadcastTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetupBroadcastTablesRequest) GoString() string {
	return s.String()
}

func (s *SetupBroadcastTablesRequest) GetActive() *bool {
	return s.Active
}

func (s *SetupBroadcastTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *SetupBroadcastTablesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetupBroadcastTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetupBroadcastTablesRequest) GetTableName() []*string {
	return s.TableName
}

func (s *SetupBroadcastTablesRequest) SetActive(v bool) *SetupBroadcastTablesRequest {
	s.Active = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDbName(v string) *SetupBroadcastTablesRequest {
	s.DbName = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetDrdsInstanceId(v string) *SetupBroadcastTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetRegionId(v string) *SetupBroadcastTablesRequest {
	s.RegionId = &v
	return s
}

func (s *SetupBroadcastTablesRequest) SetTableName(v []*string) *SetupBroadcastTablesRequest {
	s.TableName = v
	return s
}

func (s *SetupBroadcastTablesRequest) Validate() error {
	return dara.Validate(s)
}
