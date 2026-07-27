// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v string) *DsgQueryDetailsRequest
	GetBeginTime() *string
	SetEndTime(v string) *DsgQueryDetailsRequest
	GetEndTime() *string
	SetEngineName(v string) *DsgQueryDetailsRequest
	GetEngineName() *string
	SetIp(v string) *DsgQueryDetailsRequest
	GetIp() *string
	SetIpAare(v string) *DsgQueryDetailsRequest
	GetIpAare() *string
	SetNodeId(v string) *DsgQueryDetailsRequest
	GetNodeId() *string
	SetPageNo(v int64) *DsgQueryDetailsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *DsgQueryDetailsRequest
	GetPageSize() *int64
	SetProjectId(v string) *DsgQueryDetailsRequest
	GetProjectId() *string
	SetRows(v int32) *DsgQueryDetailsRequest
	GetRows() *int32
	SetRuleType(v string) *DsgQueryDetailsRequest
	GetRuleType() *string
	SetSensLevel(v string) *DsgQueryDetailsRequest
	GetSensLevel() *string
	SetUser(v string) *DsgQueryDetailsRequest
	GetUser() *string
}

type DsgQueryDetailsRequest struct {
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
	// The internal IP address of the ECU.
	//
	// example:
	//
	// 203.107.80.20
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The region to which the IP address belongs. Example: China-Beijing-Beijing, or internal IP address.
	//
	// example:
	//
	// China-Beijing-Beijing.
	IpAare *string `json:"IpAare,omitempty" xml:"IpAare,omitempty"`
	// The node ID.
	//
	// example:
	//
	// 123541234
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project workspace name (essentially ProjectName). Example: dsg_demo_gw.
	//
	// example:
	//
	// dsg_demo_gw
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The minimum value of the export volume.
	//
	// example:
	//
	// 1
	Rows *int32 `json:"Rows,omitempty" xml:"Rows,omitempty"`
	// The type of triggered sensitive rule. Example: Name.
	//
	// example:
	//
	// Name.
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The classification level. Example: 3.
	//
	// example:
	//
	// 3
	SensLevel *string `json:"SensLevel,omitempty" xml:"SensLevel,omitempty"`
	// The operator account. Example: dsg_test.
	//
	// example:
	//
	// dsg_test
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DsgQueryDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDetailsRequest) GoString() string {
	return s.String()
}

func (s *DsgQueryDetailsRequest) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DsgQueryDetailsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DsgQueryDetailsRequest) GetEngineName() *string {
	return s.EngineName
}

func (s *DsgQueryDetailsRequest) GetIp() *string {
	return s.Ip
}

func (s *DsgQueryDetailsRequest) GetIpAare() *string {
	return s.IpAare
}

func (s *DsgQueryDetailsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DsgQueryDetailsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *DsgQueryDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DsgQueryDetailsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *DsgQueryDetailsRequest) GetRows() *int32 {
	return s.Rows
}

func (s *DsgQueryDetailsRequest) GetRuleType() *string {
	return s.RuleType
}

func (s *DsgQueryDetailsRequest) GetSensLevel() *string {
	return s.SensLevel
}

func (s *DsgQueryDetailsRequest) GetUser() *string {
	return s.User
}

func (s *DsgQueryDetailsRequest) SetBeginTime(v string) *DsgQueryDetailsRequest {
	s.BeginTime = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetEndTime(v string) *DsgQueryDetailsRequest {
	s.EndTime = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetEngineName(v string) *DsgQueryDetailsRequest {
	s.EngineName = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetIp(v string) *DsgQueryDetailsRequest {
	s.Ip = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetIpAare(v string) *DsgQueryDetailsRequest {
	s.IpAare = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetNodeId(v string) *DsgQueryDetailsRequest {
	s.NodeId = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetPageNo(v int64) *DsgQueryDetailsRequest {
	s.PageNo = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetPageSize(v int64) *DsgQueryDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetProjectId(v string) *DsgQueryDetailsRequest {
	s.ProjectId = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetRows(v int32) *DsgQueryDetailsRequest {
	s.Rows = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetRuleType(v string) *DsgQueryDetailsRequest {
	s.RuleType = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetSensLevel(v string) *DsgQueryDetailsRequest {
	s.SensLevel = &v
	return s
}

func (s *DsgQueryDetailsRequest) SetUser(v string) *DsgQueryDetailsRequest {
	s.User = &v
	return s
}

func (s *DsgQueryDetailsRequest) Validate() error {
	return dara.Validate(s)
}
