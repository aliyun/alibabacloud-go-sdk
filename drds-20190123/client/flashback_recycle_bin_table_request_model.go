// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlashbackRecycleBinTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *FlashbackRecycleBinTableRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *FlashbackRecycleBinTableRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *FlashbackRecycleBinTableRequest
	GetRegionId() *string
	SetTableName(v string) *FlashbackRecycleBinTableRequest
	GetTableName() *string
}

type FlashbackRecycleBinTableRequest struct {
	// The name of the database to which the table belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance to which the table belongs.
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
	// The name of the logical table to be restored.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s FlashbackRecycleBinTableRequest) String() string {
	return dara.Prettify(s)
}

func (s FlashbackRecycleBinTableRequest) GoString() string {
	return s.String()
}

func (s *FlashbackRecycleBinTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *FlashbackRecycleBinTableRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *FlashbackRecycleBinTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FlashbackRecycleBinTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *FlashbackRecycleBinTableRequest) SetDbName(v string) *FlashbackRecycleBinTableRequest {
	s.DbName = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetDrdsInstanceId(v string) *FlashbackRecycleBinTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetRegionId(v string) *FlashbackRecycleBinTableRequest {
	s.RegionId = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) SetTableName(v string) *FlashbackRecycleBinTableRequest {
	s.TableName = &v
	return s
}

func (s *FlashbackRecycleBinTableRequest) Validate() error {
	return dara.Validate(s)
}
