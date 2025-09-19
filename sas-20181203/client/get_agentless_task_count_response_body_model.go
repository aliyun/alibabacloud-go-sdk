// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentlessTaskCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBaselineCheckCount(v int32) *GetAgentlessTaskCountResponseBody
	GetBaselineCheckCount() *int32
	SetCveVulCount(v int32) *GetAgentlessTaskCountResponseBody
	GetCveVulCount() *int32
	SetLastTaskTime(v int64) *GetAgentlessTaskCountResponseBody
	GetLastTaskTime() *int64
	SetMaliciousFile(v int32) *GetAgentlessTaskCountResponseBody
	GetMaliciousFile() *int32
	SetRequestId(v string) *GetAgentlessTaskCountResponseBody
	GetRequestId() *string
	SetRiskMachine(v int32) *GetAgentlessTaskCountResponseBody
	GetRiskMachine() *int32
	SetScaVulCount(v int32) *GetAgentlessTaskCountResponseBody
	GetScaVulCount() *int32
	SetScanMachine(v int32) *GetAgentlessTaskCountResponseBody
	GetScanMachine() *int32
	SetSensitiveFileCount(v int32) *GetAgentlessTaskCountResponseBody
	GetSensitiveFileCount() *int32
	SetSysVulCount(v string) *GetAgentlessTaskCountResponseBody
	GetSysVulCount() *string
	SetVulnerability(v int32) *GetAgentlessTaskCountResponseBody
	GetVulnerability() *int32
}

type GetAgentlessTaskCountResponseBody struct {
	// The number of baseline checks.
	//
	// example:
	//
	// 1
	BaselineCheckCount *int32 `json:"BaselineCheckCount,omitempty" xml:"BaselineCheckCount,omitempty"`
	// The number of system vulnerabilities.
	//
	// example:
	//
	// 1
	CveVulCount *int32 `json:"CveVulCount,omitempty" xml:"CveVulCount,omitempty"`
	// The timestamp generated when the last detection is performed.
	//
	// example:
	//
	// 1682577532318
	LastTaskTime *int64 `json:"LastTaskTime,omitempty" xml:"LastTaskTime,omitempty"`
	// The number of malicious files.
	//
	// example:
	//
	// 1
	MaliciousFile *int32 `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of risky hosts.
	//
	// example:
	//
	// 1
	RiskMachine *int32 `json:"RiskMachine,omitempty" xml:"RiskMachine,omitempty"`
	// The number of application vulnerabilities.
	//
	// example:
	//
	// 1
	ScaVulCount *int32 `json:"ScaVulCount,omitempty" xml:"ScaVulCount,omitempty"`
	// The number of hosts that are scanned.
	//
	// example:
	//
	// 1
	ScanMachine *int32 `json:"ScanMachine,omitempty" xml:"ScanMachine,omitempty"`
	// The total number of sensitive files.
	//
	// example:
	//
	// 1
	SensitiveFileCount *int32 `json:"SensitiveFileCount,omitempty" xml:"SensitiveFileCount,omitempty"`
	// The total number of Windows system vulnerabilities.
	//
	// example:
	//
	// 1
	SysVulCount *string `json:"SysVulCount,omitempty" xml:"SysVulCount,omitempty"`
	// The number of vulnerabilities.
	//
	// example:
	//
	// 1
	Vulnerability *int32 `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty"`
}

func (s GetAgentlessTaskCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentlessTaskCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentlessTaskCountResponseBody) GetBaselineCheckCount() *int32 {
	return s.BaselineCheckCount
}

func (s *GetAgentlessTaskCountResponseBody) GetCveVulCount() *int32 {
	return s.CveVulCount
}

func (s *GetAgentlessTaskCountResponseBody) GetLastTaskTime() *int64 {
	return s.LastTaskTime
}

func (s *GetAgentlessTaskCountResponseBody) GetMaliciousFile() *int32 {
	return s.MaliciousFile
}

func (s *GetAgentlessTaskCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentlessTaskCountResponseBody) GetRiskMachine() *int32 {
	return s.RiskMachine
}

func (s *GetAgentlessTaskCountResponseBody) GetScaVulCount() *int32 {
	return s.ScaVulCount
}

func (s *GetAgentlessTaskCountResponseBody) GetScanMachine() *int32 {
	return s.ScanMachine
}

func (s *GetAgentlessTaskCountResponseBody) GetSensitiveFileCount() *int32 {
	return s.SensitiveFileCount
}

func (s *GetAgentlessTaskCountResponseBody) GetSysVulCount() *string {
	return s.SysVulCount
}

func (s *GetAgentlessTaskCountResponseBody) GetVulnerability() *int32 {
	return s.Vulnerability
}

func (s *GetAgentlessTaskCountResponseBody) SetBaselineCheckCount(v int32) *GetAgentlessTaskCountResponseBody {
	s.BaselineCheckCount = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetCveVulCount(v int32) *GetAgentlessTaskCountResponseBody {
	s.CveVulCount = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetLastTaskTime(v int64) *GetAgentlessTaskCountResponseBody {
	s.LastTaskTime = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetMaliciousFile(v int32) *GetAgentlessTaskCountResponseBody {
	s.MaliciousFile = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetRequestId(v string) *GetAgentlessTaskCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetRiskMachine(v int32) *GetAgentlessTaskCountResponseBody {
	s.RiskMachine = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetScaVulCount(v int32) *GetAgentlessTaskCountResponseBody {
	s.ScaVulCount = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetScanMachine(v int32) *GetAgentlessTaskCountResponseBody {
	s.ScanMachine = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetSensitiveFileCount(v int32) *GetAgentlessTaskCountResponseBody {
	s.SensitiveFileCount = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetSysVulCount(v string) *GetAgentlessTaskCountResponseBody {
	s.SysVulCount = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) SetVulnerability(v int32) *GetAgentlessTaskCountResponseBody {
	s.Vulnerability = &v
	return s
}

func (s *GetAgentlessTaskCountResponseBody) Validate() error {
	return dara.Validate(s)
}
