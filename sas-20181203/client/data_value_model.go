// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetRiskMachine(v int32) *DataValue
	GetRiskMachine() *int32
	SetScanMachine(v int32) *DataValue
	GetScanMachine() *int32
	SetMaliciousFile(v int32) *DataValue
	GetMaliciousFile() *int32
	SetVulnerability(v int32) *DataValue
	GetVulnerability() *int32
	SetLastTaskTime(v int64) *DataValue
	GetLastTaskTime() *int64
	SetBaselineCheckCount(v int32) *DataValue
	GetBaselineCheckCount() *int32
	SetScaVulCount(v int32) *DataValue
	GetScaVulCount() *int32
	SetCveVulCount(v int32) *DataValue
	GetCveVulCount() *int32
	SetSysVulCount(v int32) *DataValue
	GetSysVulCount() *int32
	SetSensitiveFileCount(v int32) *DataValue
	GetSensitiveFileCount() *int32
	SetEstimateUsedSize(v int64) *DataValue
	GetEstimateUsedSize() *int64
	SetCveNum(v int32) *DataValue
	GetCveNum() *int32
	SetEmgNum(v int32) *DataValue
	GetEmgNum() *int32
	SetSysNum(v int32) *DataValue
	GetSysNum() *int32
	SetCmsNum(v int32) *DataValue
	GetCmsNum() *int32
	SetAppNum(v int32) *DataValue
	GetAppNum() *int32
	SetScaNum(v int32) *DataValue
	GetScaNum() *int32
	SetVulAsapSum(v int32) *DataValue
	GetVulAsapSum() *int32
	SetVulLaterSum(v int32) *DataValue
	GetVulLaterSum() *int32
	SetVulNntfSum(v int32) *DataValue
	GetVulNntfSum() *int32
	SetSysAsapNum(v int32) *DataValue
	GetSysAsapNum() *int32
}

type DataValue struct {
	// The number of risky hosts.
	//
	// example:
	//
	// 1
	RiskMachine *int32 `json:"RiskMachine,omitempty" xml:"RiskMachine,omitempty"`
	// The number of scanned hosts.
	//
	// example:
	//
	// 1
	ScanMachine *int32 `json:"ScanMachine,omitempty" xml:"ScanMachine,omitempty"`
	// The total number of malicious sample files.
	//
	// example:
	//
	// 1
	MaliciousFile *int32 `json:"MaliciousFile,omitempty" xml:"MaliciousFile,omitempty"`
	// The number of vulnerability risks.
	//
	// example:
	//
	// 1
	Vulnerability *int32 `json:"Vulnerability,omitempty" xml:"Vulnerability,omitempty"`
	// The timestamp of the last scan time. Unit: milliseconds.
	//
	// example:
	//
	// 1682577532318
	LastTaskTime *int64 `json:"LastTaskTime,omitempty" xml:"LastTaskTime,omitempty"`
	// The total number of baseline check items.
	//
	// example:
	//
	// 1
	BaselineCheckCount *int32 `json:"BaselineCheckCount,omitempty" xml:"BaselineCheckCount,omitempty"`
	// The total number of application vulnerabilities.
	//
	// example:
	//
	// 1
	ScaVulCount *int32 `json:"ScaVulCount,omitempty" xml:"ScaVulCount,omitempty"`
	// The total number of system vulnerabilities.
	//
	// example:
	//
	// 1
	CveVulCount *int32 `json:"CveVulCount,omitempty" xml:"CveVulCount,omitempty"`
	// The total number of Windows system vulnerabilities.
	//
	// example:
	//
	// 1
	SysVulCount *int32 `json:"SysVulCount,omitempty" xml:"SysVulCount,omitempty"`
	// The total number of sensitive files.
	//
	// example:
	//
	// 1
	SensitiveFileCount *int32 `json:"SensitiveFileCount,omitempty" xml:"SensitiveFileCount,omitempty"`
	// The estimated detection volume. Unit: GB. This field is not returned by the batch statistics operation.
	//
	// example:
	//
	// 10
	EstimateUsedSize *int64 `json:"EstimateUsedSize,omitempty" xml:"EstimateUsedSize,omitempty"`
	// The number of Linux software vulnerabilities.
	//
	// example:
	//
	// 1
	CveNum *int32 `json:"CveNum,omitempty" xml:"CveNum,omitempty"`
	// The number of emergency vulnerabilities. This field is 0 when ImageVul is set to true.
	//
	// example:
	//
	// 0
	EmgNum *int32 `json:"EmgNum,omitempty" xml:"EmgNum,omitempty"`
	// The number of Windows system vulnerabilities. This field is 0 when ImageVul is set to true.
	//
	// example:
	//
	// 0
	SysNum *int32 `json:"SysNum,omitempty" xml:"SysNum,omitempty"`
	// The number of Web-CMS vulnerabilities. This field is 0 when ImageVul is set to true.
	//
	// example:
	//
	// 0
	CmsNum *int32 `json:"CmsNum,omitempty" xml:"CmsNum,omitempty"`
	// The number of application vulnerabilities. This field is 0 when ImageVul is set to true.
	//
	// example:
	//
	// 0
	AppNum *int32 `json:"AppNum,omitempty" xml:"AppNum,omitempty"`
	// The number of software composition analysis (SCA) vulnerabilities.
	//
	// example:
	//
	// 2
	ScaNum *int32 `json:"ScaNum,omitempty" xml:"ScaNum,omitempty"`
	// The number of high-priority vulnerabilities.
	//
	// example:
	//
	// 1
	VulAsapSum *int32 `json:"VulAsapSum,omitempty" xml:"VulAsapSum,omitempty"`
	// The number of medium-priority vulnerabilities.
	//
	// example:
	//
	// 1
	VulLaterSum *int32 `json:"VulLaterSum,omitempty" xml:"VulLaterSum,omitempty"`
	// The number of low-priority vulnerabilities.
	//
	// example:
	//
	// 1
	VulNntfSum *int32 `json:"VulNntfSum,omitempty" xml:"VulNntfSum,omitempty"`
	// The number of high-priority system vulnerabilities among Linux software vulnerabilities and Windows system vulnerabilities.
	//
	// example:
	//
	// 1
	SysAsapNum *int32 `json:"SysAsapNum,omitempty" xml:"SysAsapNum,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetRiskMachine() *int32 {
	return s.RiskMachine
}

func (s *DataValue) GetScanMachine() *int32 {
	return s.ScanMachine
}

func (s *DataValue) GetMaliciousFile() *int32 {
	return s.MaliciousFile
}

func (s *DataValue) GetVulnerability() *int32 {
	return s.Vulnerability
}

func (s *DataValue) GetLastTaskTime() *int64 {
	return s.LastTaskTime
}

func (s *DataValue) GetBaselineCheckCount() *int32 {
	return s.BaselineCheckCount
}

func (s *DataValue) GetScaVulCount() *int32 {
	return s.ScaVulCount
}

func (s *DataValue) GetCveVulCount() *int32 {
	return s.CveVulCount
}

func (s *DataValue) GetSysVulCount() *int32 {
	return s.SysVulCount
}

func (s *DataValue) GetSensitiveFileCount() *int32 {
	return s.SensitiveFileCount
}

func (s *DataValue) GetEstimateUsedSize() *int64 {
	return s.EstimateUsedSize
}

func (s *DataValue) GetCveNum() *int32 {
	return s.CveNum
}

func (s *DataValue) GetEmgNum() *int32 {
	return s.EmgNum
}

func (s *DataValue) GetSysNum() *int32 {
	return s.SysNum
}

func (s *DataValue) GetCmsNum() *int32 {
	return s.CmsNum
}

func (s *DataValue) GetAppNum() *int32 {
	return s.AppNum
}

func (s *DataValue) GetScaNum() *int32 {
	return s.ScaNum
}

func (s *DataValue) GetVulAsapSum() *int32 {
	return s.VulAsapSum
}

func (s *DataValue) GetVulLaterSum() *int32 {
	return s.VulLaterSum
}

func (s *DataValue) GetVulNntfSum() *int32 {
	return s.VulNntfSum
}

func (s *DataValue) GetSysAsapNum() *int32 {
	return s.SysAsapNum
}

func (s *DataValue) SetRiskMachine(v int32) *DataValue {
	s.RiskMachine = &v
	return s
}

func (s *DataValue) SetScanMachine(v int32) *DataValue {
	s.ScanMachine = &v
	return s
}

func (s *DataValue) SetMaliciousFile(v int32) *DataValue {
	s.MaliciousFile = &v
	return s
}

func (s *DataValue) SetVulnerability(v int32) *DataValue {
	s.Vulnerability = &v
	return s
}

func (s *DataValue) SetLastTaskTime(v int64) *DataValue {
	s.LastTaskTime = &v
	return s
}

func (s *DataValue) SetBaselineCheckCount(v int32) *DataValue {
	s.BaselineCheckCount = &v
	return s
}

func (s *DataValue) SetScaVulCount(v int32) *DataValue {
	s.ScaVulCount = &v
	return s
}

func (s *DataValue) SetCveVulCount(v int32) *DataValue {
	s.CveVulCount = &v
	return s
}

func (s *DataValue) SetSysVulCount(v int32) *DataValue {
	s.SysVulCount = &v
	return s
}

func (s *DataValue) SetSensitiveFileCount(v int32) *DataValue {
	s.SensitiveFileCount = &v
	return s
}

func (s *DataValue) SetEstimateUsedSize(v int64) *DataValue {
	s.EstimateUsedSize = &v
	return s
}

func (s *DataValue) SetCveNum(v int32) *DataValue {
	s.CveNum = &v
	return s
}

func (s *DataValue) SetEmgNum(v int32) *DataValue {
	s.EmgNum = &v
	return s
}

func (s *DataValue) SetSysNum(v int32) *DataValue {
	s.SysNum = &v
	return s
}

func (s *DataValue) SetCmsNum(v int32) *DataValue {
	s.CmsNum = &v
	return s
}

func (s *DataValue) SetAppNum(v int32) *DataValue {
	s.AppNum = &v
	return s
}

func (s *DataValue) SetScaNum(v int32) *DataValue {
	s.ScaNum = &v
	return s
}

func (s *DataValue) SetVulAsapSum(v int32) *DataValue {
	s.VulAsapSum = &v
	return s
}

func (s *DataValue) SetVulLaterSum(v int32) *DataValue {
	s.VulLaterSum = &v
	return s
}

func (s *DataValue) SetVulNntfSum(v int32) *DataValue {
	s.VulNntfSum = &v
	return s
}

func (s *DataValue) SetSysAsapNum(v int32) *DataValue {
	s.SysAsapNum = &v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
