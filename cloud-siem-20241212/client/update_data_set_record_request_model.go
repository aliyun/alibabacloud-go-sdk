// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetFileName(v string) *UpdateDataSetRecordRequest
	GetDataSetFileName() *string
	SetDataSetId(v string) *UpdateDataSetRecordRequest
	GetDataSetId() *string
	SetDataSetRecords(v string) *UpdateDataSetRecordRequest
	GetDataSetRecords() *string
	SetLang(v string) *UpdateDataSetRecordRequest
	GetLang() *string
	SetRegionId(v string) *UpdateDataSetRecordRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataSetRecordRequest
	GetRoleFor() *int64
}

type UpdateDataSetRecordRequest struct {
	// example:
	//
	// cloudsiem-dataset/1358117679873357_174338773****.csv
	DataSetFileName *string `json:"DataSetFileName,omitempty" xml:"DataSetFileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dataset-10iy8mbifnb4gniv****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
	// example:
	//
	// [{\\"ip\\":\\"1.1.1.1\\",\\"userid\\":\\"1234\\",\\"name\\":\\"a12401\\"},
	//
	//  {\\"ip\\":\\"2.2.2.2\\",\\"userid\\":\\"33333\\",\\"name\\":\\"a12401\\"}]
	DataSetRecords *string `json:"DataSetRecords,omitempty" xml:"DataSetRecords,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataSetRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataSetRecordRequest) GetDataSetFileName() *string {
	return s.DataSetFileName
}

func (s *UpdateDataSetRecordRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *UpdateDataSetRecordRequest) GetDataSetRecords() *string {
	return s.DataSetRecords
}

func (s *UpdateDataSetRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataSetRecordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataSetRecordRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataSetRecordRequest) SetDataSetFileName(v string) *UpdateDataSetRecordRequest {
	s.DataSetFileName = &v
	return s
}

func (s *UpdateDataSetRecordRequest) SetDataSetId(v string) *UpdateDataSetRecordRequest {
	s.DataSetId = &v
	return s
}

func (s *UpdateDataSetRecordRequest) SetDataSetRecords(v string) *UpdateDataSetRecordRequest {
	s.DataSetRecords = &v
	return s
}

func (s *UpdateDataSetRecordRequest) SetLang(v string) *UpdateDataSetRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataSetRecordRequest) SetRegionId(v string) *UpdateDataSetRecordRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataSetRecordRequest) SetRoleFor(v int64) *UpdateDataSetRecordRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataSetRecordRequest) Validate() error {
	return dara.Validate(s)
}
