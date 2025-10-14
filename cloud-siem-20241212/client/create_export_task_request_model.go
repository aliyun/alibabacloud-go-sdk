// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportTaskParameter(v string) *CreateExportTaskRequest
	GetExportTaskParameter() *string
	SetExportTaskType(v string) *CreateExportTaskRequest
	GetExportTaskType() *string
	SetLang(v string) *CreateExportTaskRequest
	GetLang() *string
	SetRegionId(v string) *CreateExportTaskRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateExportTaskRequest
	GetRoleFor() *int64
}

type CreateExportTaskRequest struct {
	// example:
	//
	// {
	//
	//     "IncidentName": "name",
	//
	//     "IncidentStatus": "1",
	//
	//     "ThreatLevel":["1","2"],
	//
	//     "IncidentUuids":"315EE627-39DC-560A-8DFD-CBF66AD9****"
	//
	// }
	ExportTaskParameter *string `json:"ExportTaskParameter,omitempty" xml:"ExportTaskParameter,omitempty"`
	// example:
	//
	// incident_list
	ExportTaskType *string `json:"ExportTaskType,omitempty" xml:"ExportTaskType,omitempty"`
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

func (s CreateExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateExportTaskRequest) GetExportTaskParameter() *string {
	return s.ExportTaskParameter
}

func (s *CreateExportTaskRequest) GetExportTaskType() *string {
	return s.ExportTaskType
}

func (s *CreateExportTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateExportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExportTaskRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateExportTaskRequest) SetExportTaskParameter(v string) *CreateExportTaskRequest {
	s.ExportTaskParameter = &v
	return s
}

func (s *CreateExportTaskRequest) SetExportTaskType(v string) *CreateExportTaskRequest {
	s.ExportTaskType = &v
	return s
}

func (s *CreateExportTaskRequest) SetLang(v string) *CreateExportTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateExportTaskRequest) SetRegionId(v string) *CreateExportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExportTaskRequest) SetRoleFor(v int64) *CreateExportTaskRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
