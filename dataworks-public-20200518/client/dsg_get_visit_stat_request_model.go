// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitStatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *DsgGetVisitStatRequest
	GetBeginTime() *string
	SetEndTime(v string) *DsgGetVisitStatRequest
	GetEndTime() *string
	SetEngineName(v string) *DsgGetVisitStatRequest
	GetEngineName() *string
	SetNodeId(v string) *DsgGetVisitStatRequest
	GetNodeId() *string
	SetPageNo(v int64) *DsgGetVisitStatRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DsgGetVisitStatRequest
	GetPageSize() *int64
	SetProjectId(v string) *DsgGetVisitStatRequest
	GetProjectId() *string
	SetRuleName(v string) *DsgGetVisitStatRequest
	GetRuleName() *string
	SetSensLevel(v string) *DsgGetVisitStatRequest
	GetSensLevel() *string
}

type DsgGetVisitStatRequest struct {
	// The start time in the format of "2026-06-30 03:59:59".
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-06-30 03:59:59
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time in the format of "2026-06-30 23:59:59".
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
	// The node ID. You can call the [ListNodes](https://help.aliyun.com/document_detail/173979.html) operation to query the node ID.
	//
	// example:
	//
	// d0c72253-8eea-435b-91fc-163a90a54b33
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the project space. Example value: dsg_demo_gw.
	//
	// example:
	//
	// dsg_demo_gw
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the sensitive field.
	//
	// example:
	//
	// Name.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The classification level. Example value: 3.
	//
	// example:
	//
	// 3
	SensLevel *string `json:"SensLevel,omitempty" xml:"SensLevel,omitempty"`
}

func (s DsgGetVisitStatRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitStatRequest) GoString() string {
	return s.String()
}

func (s *DsgGetVisitStatRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DsgGetVisitStatRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DsgGetVisitStatRequest) GetEngineName() *string {
	return s.EngineName
}

func (s *DsgGetVisitStatRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DsgGetVisitStatRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DsgGetVisitStatRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DsgGetVisitStatRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DsgGetVisitStatRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DsgGetVisitStatRequest) GetSensLevel() *string {
	return s.SensLevel
}

func (s *DsgGetVisitStatRequest) SetBeginTime(v string) *DsgGetVisitStatRequest {
	s.BeginTime = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetEndTime(v string) *DsgGetVisitStatRequest {
	s.EndTime = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetEngineName(v string) *DsgGetVisitStatRequest {
	s.EngineName = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetNodeId(v string) *DsgGetVisitStatRequest {
	s.NodeId = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetPageNo(v int64) *DsgGetVisitStatRequest {
	s.PageNo = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetPageSize(v int64) *DsgGetVisitStatRequest {
	s.PageSize = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetProjectId(v string) *DsgGetVisitStatRequest {
	s.ProjectId = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetRuleName(v string) *DsgGetVisitStatRequest {
	s.RuleName = &v
	return s
}

func (s *DsgGetVisitStatRequest) SetSensLevel(v string) *DsgGetVisitStatRequest {
	s.SensLevel = &v
	return s
}

func (s *DsgGetVisitStatRequest) Validate() error {
	return dara.Validate(s)
}
