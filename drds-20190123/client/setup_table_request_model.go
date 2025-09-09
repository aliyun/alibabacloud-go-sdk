// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowFullTableScan(v bool) *SetupTableRequest
	GetAllowFullTableScan() *bool
	SetDbName(v string) *SetupTableRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *SetupTableRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *SetupTableRequest
	GetRegionId() *string
	SetTableName(v []*string) *SetupTableRequest
	GetTableName() []*string
}

type SetupTableRequest struct {
	// Specifies whether to enable full table scan.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AllowFullTableScan *bool `json:"AllowFullTableScan,omitempty" xml:"AllowFullTableScan,omitempty"`
	// The name of the database in which the table resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the DRDS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the streaming domain resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName []*string `json:"TableName,omitempty" xml:"TableName,omitempty" type:"Repeated"`
}

func (s SetupTableRequest) String() string {
	return dara.Prettify(s)
}

func (s SetupTableRequest) GoString() string {
	return s.String()
}

func (s *SetupTableRequest) GetAllowFullTableScan() *bool {
	return s.AllowFullTableScan
}

func (s *SetupTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *SetupTableRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetupTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetupTableRequest) GetTableName() []*string {
	return s.TableName
}

func (s *SetupTableRequest) SetAllowFullTableScan(v bool) *SetupTableRequest {
	s.AllowFullTableScan = &v
	return s
}

func (s *SetupTableRequest) SetDbName(v string) *SetupTableRequest {
	s.DbName = &v
	return s
}

func (s *SetupTableRequest) SetDrdsInstanceId(v string) *SetupTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupTableRequest) SetRegionId(v string) *SetupTableRequest {
	s.RegionId = &v
	return s
}

func (s *SetupTableRequest) SetTableName(v []*string) *SetupTableRequest {
	s.TableName = v
	return s
}

func (s *SetupTableRequest) Validate() error {
	return dara.Validate(s)
}
