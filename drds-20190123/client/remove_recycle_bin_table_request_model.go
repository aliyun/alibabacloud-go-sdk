// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRecycleBinTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *RemoveRecycleBinTableRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *RemoveRecycleBinTableRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *RemoveRecycleBinTableRequest
	GetRegionId() *string
	SetTableName(v string) *RemoveRecycleBinTableRequest
	GetTableName() *string
}

type RemoveRecycleBinTableRequest struct {
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
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the logical table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RemoveRecycleBinTableRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *RemoveRecycleBinTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *RemoveRecycleBinTableRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *RemoveRecycleBinTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveRecycleBinTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *RemoveRecycleBinTableRequest) SetDbName(v string) *RemoveRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetDrdsInstanceId(v string) *RemoveRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetRegionId(v string) *RemoveRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) SetTableName(v string) *RemoveRecycleBinTableRequest {
	s.TableName = &v
	return s
}

func (s *RemoveRecycleBinTableRequest) Validate() error {
	return dara.Validate(s)
}
