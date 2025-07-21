// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportAppOtaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseVersion(v string) *ReportAppOtaInfoRequest
	GetBaseVersion() *string
	SetClientType(v int32) *ReportAppOtaInfoRequest
	GetClientType() *int32
	SetClientUid(v string) *ReportAppOtaInfoRequest
	GetClientUid() *string
	SetNote(v string) *ReportAppOtaInfoRequest
	GetNote() *string
	SetOsType(v string) *ReportAppOtaInfoRequest
	GetOsType() *string
	SetProject(v string) *ReportAppOtaInfoRequest
	GetProject() *string
	SetStatus(v int32) *ReportAppOtaInfoRequest
	GetStatus() *int32
	SetTargetVersion(v string) *ReportAppOtaInfoRequest
	GetTargetVersion() *string
	SetTaskUid(v string) *ReportAppOtaInfoRequest
	GetTaskUid() *string
}

type ReportAppOtaInfoRequest struct {
	BaseVersion   *string `json:"BaseVersion,omitempty" xml:"BaseVersion,omitempty"`
	ClientType    *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	ClientUid     *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Note          *string `json:"Note,omitempty" xml:"Note,omitempty"`
	OsType        *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	Project       *string `json:"Project,omitempty" xml:"Project,omitempty"`
	Status        *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
	TaskUid       *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s ReportAppOtaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportAppOtaInfoRequest) GoString() string {
	return s.String()
}

func (s *ReportAppOtaInfoRequest) GetBaseVersion() *string {
	return s.BaseVersion
}

func (s *ReportAppOtaInfoRequest) GetClientType() *int32 {
	return s.ClientType
}

func (s *ReportAppOtaInfoRequest) GetClientUid() *string {
	return s.ClientUid
}

func (s *ReportAppOtaInfoRequest) GetNote() *string {
	return s.Note
}

func (s *ReportAppOtaInfoRequest) GetOsType() *string {
	return s.OsType
}

func (s *ReportAppOtaInfoRequest) GetProject() *string {
	return s.Project
}

func (s *ReportAppOtaInfoRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ReportAppOtaInfoRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *ReportAppOtaInfoRequest) GetTaskUid() *string {
	return s.TaskUid
}

func (s *ReportAppOtaInfoRequest) SetBaseVersion(v string) *ReportAppOtaInfoRequest {
	s.BaseVersion = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetClientType(v int32) *ReportAppOtaInfoRequest {
	s.ClientType = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetClientUid(v string) *ReportAppOtaInfoRequest {
	s.ClientUid = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetNote(v string) *ReportAppOtaInfoRequest {
	s.Note = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetOsType(v string) *ReportAppOtaInfoRequest {
	s.OsType = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetProject(v string) *ReportAppOtaInfoRequest {
	s.Project = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetStatus(v int32) *ReportAppOtaInfoRequest {
	s.Status = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetTargetVersion(v string) *ReportAppOtaInfoRequest {
	s.TargetVersion = &v
	return s
}

func (s *ReportAppOtaInfoRequest) SetTaskUid(v string) *ReportAppOtaInfoRequest {
	s.TaskUid = &v
	return s
}

func (s *ReportAppOtaInfoRequest) Validate() error {
	return dara.Validate(s)
}
