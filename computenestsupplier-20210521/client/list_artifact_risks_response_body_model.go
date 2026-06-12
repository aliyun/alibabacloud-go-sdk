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
	// The list of artifact risks.
	ArtifactRiskList []*ListArtifactRisksResponseBodyArtifactRiskList `json:"ArtifactRiskList,omitempty" xml:"ArtifactRiskList,omitempty" type:"Repeated"`
	// The request ID.
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
	if s.ArtifactRiskList != nil {
		for _, item := range s.ArtifactRiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactRisksResponseBodyArtifactRiskList struct {
	// The CVE ID.
	//
	// example:
	//
	// CVE-2023-4911
	CveNos *string `json:"CveNos,omitempty" xml:"CveNos,omitempty"`
	// The extended information in JSON format. Parse this information based on the risk type.
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
	// The risk level.
	//
	// - high: High
	//
	// example:
	//
	// high
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The name of the risk.
	//
	// example:
	//
	// USN-3686-1: file vulnerabilities
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The risk type. Valid values:
	//
	// - AcrCve: system vulnerabilities in a container image
	//
	// - AcrSca: application vulnerabilities in a container image
	//
	// - EcsVulnerability: ECS image vulnerabilities
	//
	// - EcsAlarm: ECS image security alerts
	//
	// - EcsBaseline: ECS image baseline checks
	//
	// example:
	//
	// AcrCve
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The name of the risk type.
	//
	// example:
	//
	// Container image system vulnerability
	RiskTypeName *string `json:"RiskTypeName,omitempty" xml:"RiskTypeName,omitempty"`
	// The solution to the risk.
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
