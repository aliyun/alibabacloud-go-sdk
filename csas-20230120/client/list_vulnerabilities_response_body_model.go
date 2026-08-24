// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulnerabilitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVulnerabilitiesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListVulnerabilitiesResponseBody
	GetTotalNum() *int64
	SetVulnerabilities(v []*ListVulnerabilitiesResponseBodyVulnerabilities) *ListVulnerabilitiesResponseBody
	GetVulnerabilities() []*ListVulnerabilitiesResponseBodyVulnerabilities
}

type ListVulnerabilitiesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities that match the query conditions.
	//
	// example:
	//
	// 37
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The list of vulnerabilities.
	Vulnerabilities []*ListVulnerabilitiesResponseBodyVulnerabilities `json:"Vulnerabilities,omitempty" xml:"Vulnerabilities,omitempty" type:"Repeated"`
}

func (s ListVulnerabilitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVulnerabilitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVulnerabilitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVulnerabilitiesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListVulnerabilitiesResponseBody) GetVulnerabilities() []*ListVulnerabilitiesResponseBodyVulnerabilities {
	return s.Vulnerabilities
}

func (s *ListVulnerabilitiesResponseBody) SetRequestId(v string) *ListVulnerabilitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVulnerabilitiesResponseBody) SetTotalNum(v int64) *ListVulnerabilitiesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListVulnerabilitiesResponseBody) SetVulnerabilities(v []*ListVulnerabilitiesResponseBodyVulnerabilities) *ListVulnerabilitiesResponseBody {
	s.Vulnerabilities = v
	return s
}

func (s *ListVulnerabilitiesResponseBody) Validate() error {
	if s.Vulnerabilities != nil {
		for _, item := range s.Vulnerabilities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVulnerabilitiesResponseBodyVulnerabilities struct {
	// The list of CVE IDs corresponding to the vulnerability. An empty list is returned if no CVE is associated.
	CveList []*string `json:"CveList,omitempty" xml:"CveList,omitempty" type:"Repeated"`
	// The English description of the vulnerability.
	//
	// example:
	//
	// This update fixes several remote code execution and privilege escalation vulnerabilities.
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The Chinese description of the vulnerability.
	//
	// example:
	//
	// 该更新修复了若干远程代码执行与权限提升漏洞
	DescriptionZh *string `json:"DescriptionZh,omitempty" xml:"DescriptionZh,omitempty"`
	// The list of Knowledge Base (KB) numbers corresponding to the vulnerability.
	Kbs []*string `json:"Kbs,omitempty" xml:"Kbs,omitempty" type:"Repeated"`
	// The name of the product affected by the vulnerability.
	//
	// example:
	//
	// Windows 11 Home
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The release time of the vulnerability, in seconds as a UNIX timestamp.
	//
	// example:
	//
	// 1786291200
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The English title of the vulnerability.
	//
	// example:
	//
	// 2026-08 Cumulative Update for Windows 11
	TitleEn *string `json:"TitleEn,omitempty" xml:"TitleEn,omitempty"`
	// The Chinese title of the vulnerability.
	//
	// example:
	//
	// 2026-08 适用于 Windows 11 的累积更新
	TitleZh *string `json:"TitleZh,omitempty" xml:"TitleZh,omitempty"`
	// The patch ID corresponding to the vulnerability. For Windows vulnerabilities, this is the Microsoft patch Update ID.
	//
	// example:
	//
	// 9f8c1d2e-4b7a-4c31-9e05-6d2f8a71****
	UpdateId *string `json:"UpdateId,omitempty" xml:"UpdateId,omitempty"`
	// The number of user endpoint devices affected by the vulnerability.
	//
	// example:
	//
	// 12
	VulDeviceCount *int64 `json:"VulDeviceCount,omitempty" xml:"VulDeviceCount,omitempty"`
	// The vulnerability risk level, mapped from the vendor risk level: Critical is mapped to High, Important is mapped to Mid, and others are mapped to Low. Valid values:
	//
	// - **High**: high risk.
	//
	// - **Mid**: medium risk.
	//
	// - **Low**: low risk.
	//
	// example:
	//
	// High
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
	// The vulnerability type. Valid values:
	//
	// - **windows**: Windows system vulnerability.
	//
	// - **ai_agent**: AI Agent vulnerability.
	//
	// example:
	//
	// windows
	VulType *string `json:"VulType,omitempty" xml:"VulType,omitempty"`
}

func (s ListVulnerabilitiesResponseBodyVulnerabilities) String() string {
	return dara.Prettify(s)
}

func (s ListVulnerabilitiesResponseBodyVulnerabilities) GoString() string {
	return s.String()
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetCveList() []*string {
	return s.CveList
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetDescriptionZh() *string {
	return s.DescriptionZh
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetKbs() []*string {
	return s.Kbs
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetProduct() *string {
	return s.Product
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetTitleEn() *string {
	return s.TitleEn
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetTitleZh() *string {
	return s.TitleZh
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetUpdateId() *string {
	return s.UpdateId
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetVulDeviceCount() *int64 {
	return s.VulDeviceCount
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetVulLevel() *string {
	return s.VulLevel
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) GetVulType() *string {
	return s.VulType
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetCveList(v []*string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.CveList = v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetDescriptionEn(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.DescriptionEn = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetDescriptionZh(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.DescriptionZh = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetKbs(v []*string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.Kbs = v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetProduct(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.Product = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetReleaseTime(v int64) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.ReleaseTime = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetTitleEn(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.TitleEn = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetTitleZh(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.TitleZh = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetUpdateId(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.UpdateId = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetVulDeviceCount(v int64) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.VulDeviceCount = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetVulLevel(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.VulLevel = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) SetVulType(v string) *ListVulnerabilitiesResponseBodyVulnerabilities {
	s.VulType = &v
	return s
}

func (s *ListVulnerabilitiesResponseBodyVulnerabilities) Validate() error {
	return dara.Validate(s)
}
