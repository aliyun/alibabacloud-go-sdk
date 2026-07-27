// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *DsgGetVisitDetailRequest
	GetBeginTime() *string
	SetEndTime(v string) *DsgGetVisitDetailRequest
	GetEndTime() *string
	SetEngineName(v string) *DsgGetVisitDetailRequest
	GetEngineName() *string
	SetKeyword(v string) *DsgGetVisitDetailRequest
	GetKeyword() *string
	SetPageNo(v int64) *DsgGetVisitDetailRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DsgGetVisitDetailRequest
	GetPageSize() *int64
	SetProjectId(v string) *DsgGetVisitDetailRequest
	GetProjectId() *string
	SetRuleName(v string) *DsgGetVisitDetailRequest
	GetRuleName() *string
	SetSensLevel(v string) *DsgGetVisitDetailRequest
	GetSensLevel() *string
}

type DsgGetVisitDetailRequest struct {
	// The start time of the query range. Example: "2026-06-26 00:00:00".
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-26 00:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the query range. Example: "2026-06-30 23:59:59".
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-30 23:59:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The engine type. Valid values:
	//
	// - ODPS.ODPS
	//
	// - EMR
	//
	// - HOLO.POSTGRES
	//
	// This parameter is required.
	//
	// example:
	//
	// ODPS.ODPS
	EngineName *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	// The keyword of the table or project name. DataWorks supports fuzzy match. You can enter a keyword to query table or project names that contain the keyword.
	//
	// example:
	//
	// ods
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project name (ProjectName is easier to understand). Example: dsg_demo_gw.
	//
	// This parameter is required.
	//
	// example:
	//
	// dsg_demo_gw
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the sensitive field.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The sensitivity level. Example: 3.
	//
	// example:
	//
	// 3
	SensLevel *string `json:"SensLevel,omitempty" xml:"SensLevel,omitempty"`
}

func (s DsgGetVisitDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitDetailRequest) GoString() string {
	return s.String()
}

func (s *DsgGetVisitDetailRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DsgGetVisitDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DsgGetVisitDetailRequest) GetEngineName() *string {
	return s.EngineName
}

func (s *DsgGetVisitDetailRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DsgGetVisitDetailRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DsgGetVisitDetailRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DsgGetVisitDetailRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DsgGetVisitDetailRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgGetVisitDetailRequest) GetSensLevel() *string {
	return s.SensLevel
}

func (s *DsgGetVisitDetailRequest) SetBeginTime(v string) *DsgGetVisitDetailRequest {
	s.BeginTime = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetEndTime(v string) *DsgGetVisitDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetEngineName(v string) *DsgGetVisitDetailRequest {
	s.EngineName = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetKeyword(v string) *DsgGetVisitDetailRequest {
	s.Keyword = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetPageNo(v int64) *DsgGetVisitDetailRequest {
	s.PageNo = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetPageSize(v int64) *DsgGetVisitDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetProjectId(v string) *DsgGetVisitDetailRequest {
	s.ProjectId = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetRuleName(v string) *DsgGetVisitDetailRequest {
	s.RuleName = &v
	return s
}

func (s *DsgGetVisitDetailRequest) SetSensLevel(v string) *DsgGetVisitDetailRequest {
	s.SensLevel = &v
	return s
}

func (s *DsgGetVisitDetailRequest) Validate() error {
	return dara.Validate(s)
}
