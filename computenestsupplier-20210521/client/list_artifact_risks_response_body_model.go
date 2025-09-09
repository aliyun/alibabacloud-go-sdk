// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactRisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactRiskList(v []*ListArtifactRisksResponseBodyArtifactRiskList) *ListArtifactRisksResponseBody
	GetArtifactRiskList() []*ListArtifactRisksResponseBodyArtifactRiskList
	SetRequestId(v string) *ListArtifactRisksResponseBody
	GetRequestId() *string
}

type ListArtifactRisksResponseBody struct {
	// List of artifact risks
	ArtifactRiskList []*ListArtifactRisksResponseBodyArtifactRiskList `json:"ArtifactRiskList,omitempty" xml:"ArtifactRiskList,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 52919DB1-03A0-55F5-BDD4-DB6DEBB8267A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListArtifactRisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactRisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponseBody) GetArtifactRiskList() []*ListArtifactRisksResponseBodyArtifactRiskList {
	return s.ArtifactRiskList
}

func (s *ListArtifactRisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactRisksResponseBody) SetArtifactRiskList(v []*ListArtifactRisksResponseBodyArtifactRiskList) *ListArtifactRisksResponseBody {
	s.ArtifactRiskList = v
	return s
}

func (s *ListArtifactRisksResponseBody) SetRequestId(v string) *ListArtifactRisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactRisksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListArtifactRisksResponseBodyArtifactRiskList struct {
	// CVE numbers
	//
	// example:
	//
	// CVE-2023-4911
	CveNos *string `json:"CveNos,omitempty" xml:"CveNos,omitempty"`
	// Extended information, in JSON format, to be parsed according to the risk category
	//
	// example:
	//
	// {
	//
	//   "feature": "ntpdate",
	//
	//   "version": "4.2.6",
	//
	//   "cveLocation": "/usr/lib"
	//
	// }
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// Risk level:
	//
	// - high represents high
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Risk name.
	//
	// example:
	//
	// USN-3686-1: file vulnerabilities
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// Risk type. Values:
	//
	// - AcrCve  Container image system vulnerability
	//
	// - AcrSca  Container image application vulnerability
	//
	// - EcsVulnerability  ECS image vulnerability information
	//
	// - EcsAlarm  ECS image security alarm
	//
	// - EcsBaseline  ECS image baseline check
	//
	// example:
	//
	// AcrCve
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// Risk Type name
	//
	// example:
	//
	// Container System Vulner
	RiskTypeName *string `json:"RiskTypeName,omitempty" xml:"RiskTypeName,omitempty"`
	// Solution for the risk item.
	//
	// example:
	//
	// apt-get update && apt-get install ntpdate --only-upgrade
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
}

func (s ListArtifactRisksResponseBodyArtifactRiskList) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactRisksResponseBodyArtifactRiskList) GoString() string {
	return s.String()
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetCveNos() *string {
	return s.CveNos
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetLevel() *string {
	return s.Level
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetRiskName() *string {
	return s.RiskName
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetRiskType() *string {
	return s.RiskType
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetRiskTypeName() *string {
	return s.RiskTypeName
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) GetSolution() *string {
	return s.Solution
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetCveNos(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.CveNos = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetExtendInfo(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.ExtendInfo = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetLevel(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.Level = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskName(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskName = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskType(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskType = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetRiskTypeName(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.RiskTypeName = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) SetSolution(v string) *ListArtifactRisksResponseBodyArtifactRiskList {
	s.Solution = &v
	return s
}

func (s *ListArtifactRisksResponseBodyArtifactRiskList) Validate() error {
	return dara.Validate(s)
}
