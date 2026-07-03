// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetId(v string) *DeleteDataSetRecordRequest
	GetDataSetId() *string
	SetDataSetRecordIds(v string) *DeleteDataSetRecordRequest
	GetDataSetRecordIds() *string
	SetLang(v string) *DeleteDataSetRecordRequest
	GetLang() *string
	SetRegionId(v string) *DeleteDataSetRecordRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteDataSetRecordRequest
	GetRoleFor() *int64
}

type DeleteDataSetRecordRequest struct {
	// The ID of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataset-10iy8mbifnb4gniv****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// A list of dataset record IDs in a JSON array format.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2,3,4]
	DataSetRecordIds *string `json:"DataSetRecordIds,omitempty" xml:"DataSetRecordIds,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region of the Data Management center for Threat Analysis. Select a region based on where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to operate as the member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteDataSetRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetRecordRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRecordRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *DeleteDataSetRecordRequest) GetDataSetRecordIds() *string {
	return s.DataSetRecordIds
}

func (s *DeleteDataSetRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataSetRecordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataSetRecordRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteDataSetRecordRequest) SetDataSetId(v string) *DeleteDataSetRecordRequest {
	s.DataSetId = &v
	return s
}

func (s *DeleteDataSetRecordRequest) SetDataSetRecordIds(v string) *DeleteDataSetRecordRequest {
	s.DataSetRecordIds = &v
	return s
}

func (s *DeleteDataSetRecordRequest) SetLang(v string) *DeleteDataSetRecordRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataSetRecordRequest) SetRegionId(v string) *DeleteDataSetRecordRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSetRecordRequest) SetRoleFor(v int64) *DeleteDataSetRecordRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteDataSetRecordRequest) Validate() error {
	return dara.Validate(s)
}
