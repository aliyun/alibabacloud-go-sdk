// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportId(v int64) *GetExportTaskRequest
	GetExportId() *int64
	SetLang(v string) *GetExportTaskRequest
	GetLang() *string
	SetRegionId(v string) *GetExportTaskRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetExportTaskRequest
	GetRoleFor() *int64
}

type GetExportTaskRequest struct {
	// example:
	//
	// 200013
	ExportId *int64 `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
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

func (s GetExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExportTaskRequest) GoString() string {
	return s.String()
}

func (s *GetExportTaskRequest) GetExportId() *int64 {
	return s.ExportId
}

func (s *GetExportTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *GetExportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExportTaskRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetExportTaskRequest) SetExportId(v int64) *GetExportTaskRequest {
	s.ExportId = &v
	return s
}

func (s *GetExportTaskRequest) SetLang(v string) *GetExportTaskRequest {
	s.Lang = &v
	return s
}

func (s *GetExportTaskRequest) SetRegionId(v string) *GetExportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetExportTaskRequest) SetRoleFor(v int64) *GetExportTaskRequest {
	s.RoleFor = &v
	return s
}

func (s *GetExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
