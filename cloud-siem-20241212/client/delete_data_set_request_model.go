// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSetId(v string) *DeleteDataSetRequest
	GetDataSetId() *string
	SetLang(v string) *DeleteDataSetRequest
	GetLang() *string
	SetRegionId(v string) *DeleteDataSetRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteDataSetRequest
	GetRoleFor() *int64
}

type DeleteDataSetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dataset-10iy8mbifnb4gniv****
	DataSetId *string `json:"DataSetId,omitempty" xml:"DataSetId,omitempty"`
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

func (s DeleteDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataSetRequest) GetDataSetId() *string {
	return s.DataSetId
}

func (s *DeleteDataSetRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteDataSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataSetRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteDataSetRequest) SetDataSetId(v string) *DeleteDataSetRequest {
	s.DataSetId = &v
	return s
}

func (s *DeleteDataSetRequest) SetLang(v string) *DeleteDataSetRequest {
	s.Lang = &v
	return s
}

func (s *DeleteDataSetRequest) SetRegionId(v string) *DeleteDataSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataSetRequest) SetRoleFor(v int64) *DeleteDataSetRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteDataSetRequest) Validate() error {
	return dara.Validate(s)
}
