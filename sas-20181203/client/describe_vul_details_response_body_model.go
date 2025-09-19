// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCves(v []*DescribeVulDetailsResponseBodyCves) *DescribeVulDetailsResponseBody
	GetCves() []*DescribeVulDetailsResponseBodyCves
	SetRequestId(v string) *DescribeVulDetailsResponseBody
	GetRequestId() *string
}

type DescribeVulDetailsResponseBody struct {
	// The details of the vulnerability.
	Cves []*DescribeVulDetailsResponseBodyCves `json:"Cves,omitempty" xml:"Cves,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// EDA40EA3-6265-5900-AD99-C83E4F109CA8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVulDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBody) GetCves() []*DescribeVulDetailsResponseBodyCves {
	return s.Cves
}

func (s *DescribeVulDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulDetailsResponseBody) SetCves(v []*DescribeVulDetailsResponseBodyCves) *DescribeVulDetailsResponseBody {
	s.Cves = v
	return s
}

func (s *DescribeVulDetailsResponseBody) SetRequestId(v string) *DescribeVulDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVulDetailsResponseBodyCves struct {
	// The type of the vulnerability.
	//
	// example:
	//
	// remote_code_execution
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The vulnerability types.
	Classifys []*DescribeVulDetailsResponseBodyCvesClassifys `json:"Classifys,omitempty" xml:"Classifys,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The China National Vulnerability Database (CNVD) ID.
	//
	// example:
	//
	// CNVD-2019-9167
	CnvdId *string `json:"CnvdId,omitempty" xml:"CnvdId,omitempty"`
	// Deprecated
	//
	// The difficulty level of exploiting the vulnerability. Valid values:
	//
	// 	- **LOW**
	//
	// 	- **MEDIUM**
	//
	// 	- **HIGH**
	//
	// example:
	//
	// LOW
	Complexity *string `json:"Complexity,omitempty" xml:"Complexity,omitempty"`
	// Deprecated
	//
	// The CVE content.
	//
	// example:
	//
	// Apache Shiro is a user authentication and authorization framework for a wide range of rights management applications.↵Recently, Apache Shiro released version 1.7.0, which fixes the Apache Shiro authentication bypass vulnerability (CVE-2020-17510).↵Attackers can bypass Shiro\\"s authentication using malicious requests containing payloads.↵↵Related bugs:↵CVE-2020-17510 Shiro < 1.7.0 Validation Bypass Vulnerability↵CVE-2020-13933 Shiro < 1.6.0 Validation Bypass Vulnerability↵CVE-2020-11989 Shiro < 1.5.3 Validation Bypass Vulnerability↵CVE-2020-1957 Shiro < 1.5.2 Validation Bypass Vulnerability↵CVE-2016-6802 Shiro < 1.3.2 Validation Bypass Vulnerability
	//
	// Check whether the fastjson version currently running on the system is in the affected version and whether safeMode is configured to disable autoType. If it is in the affected version and safeMode is not configured to disable autoType, the vulnerability is considered to exist.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID.
	//
	// example:
	//
	// CVE-2019-9167
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The link to the CVE details.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/CVE-2022-1184
	CveLink *string `json:"CveLink,omitempty" xml:"CveLink,omitempty"`
	// The Common Vulnerability Scoring System (CVSS) score of the vulnerability in the Alibaba Cloud vulnerability library.
	//
	// example:
	//
	// 10.0
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	// The vector that is used to calculate the CVSS score.
	//
	// example:
	//
	// AV:N/AC:L/Au:N/C:C/I:C/A:C
	CvssVector *string `json:"CvssVector,omitempty" xml:"CvssVector,omitempty"`
	// Deprecated
	//
	// The name of the instance.
	//
	// >  This parameter is deprecated. You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the instance that is affected by the vulnerability.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Deprecated
	//
	// The public IP address of the server.
	//
	// >  This parameter is deprecated. You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the instance that is affected by the vulnerability.
	//
	// example:
	//
	// 47.114.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// Deprecated
	//
	// The private IP address of the server.
	//
	// >  This parameter is deprecated. You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the instance that is affected by the vulnerability.
	//
	// example:
	//
	// 172.19.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The ID of the vulnerability.
	//
	// example:
	//
	// CVE-2020-8597
	OtherId *string `json:"OtherId,omitempty" xml:"OtherId,omitempty"`
	// Deprecated
	//
	// The POC content.
	Poc *string `json:"Poc,omitempty" xml:"Poc,omitempty"`
	// Deprecated
	//
	// The UNIX timestamp when the proof of concept (POC) was created. Unit: milliseconds.
	//
	// example:
	//
	// 1554189334000
	PocCreateTime *int64 `json:"PocCreateTime,omitempty" xml:"PocCreateTime,omitempty"`
	// Deprecated
	//
	// The UNIX timestamp when the POC was disclosed. Unit: milliseconds.
	//
	// example:
	//
	// 1554189334000
	PocDisclosureTime *int64 `json:"PocDisclosureTime,omitempty" xml:"PocDisclosureTime,omitempty"`
	// Deprecated
	//
	// The service that is affected by the vulnerability.
	//
	// example:
	//
	// Log4j2
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The reference of the vulnerability in the Alibaba Cloud vulnerability library. The value is a URL.
	//
	// example:
	//
	// https://example.com
	Reference *string `json:"Reference,omitempty" xml:"Reference,omitempty"`
	// The disclosure time that is displayed for the vulnerability in the Alibaba Cloud vulnerability library. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1554189334000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The fixing suggestions of the vulnerability.
	//
	// example:
	//
	// <p>At present, Chanjet has urgently released a vulnerability patch to fix the vulnerability. CNVD recommends affected units and users to upgrade to the latest version immediately:</p>↵<p>https://www.chanjetvip.com/product/goods/goods-detail?id=53aaa40295d458e44f5d3ce5</p>↵<p>At the same time, organizations and users affected by the vulnerability are requested to immediately follow the steps below to conduct self-inspection and repair work:</p>↵<ol>↵<li><p>User self-check steps:↵<br  />Check whether website/bin/load.aspx.cdcab7d2.compiled, website/bin/App_Web_load.aspx.cdcab7d2.dll, and tplus/Load.aspx files exist locally. If they exist, it means that they have been poisoned, and you must reinstall the system and install the product. patch.</p>↵</li>↵<li><p>Non-poisoned users please:↵<br  />1) Update the latest product patch.↵<br  />2) Install anti-virus software and update the virus database in time.↵<br  />3) Upgrade the lower version of IIS and Nginx to IIS10.0 and Windows 2016.↵<br  />4) Local installation customers need to confirm whether the backup file is complete as soon as possible, and do off-site backup. Customers on the cloud should enable the mirroring function in time.↵<br  />5) Users who fail to update the patch in time can contact Chanjet technical support and take temporary preventive measures such as deleting files.</p>↵</li>↵<li><p>Poisoned users please:↵<br  />1) Check whether the server has taken regular snapshots or backups. If so, you can restore data through snapshots or backups.↵<br  />2) Contact Chanjet technical support to confirm whether it has the conditions and operation methods to restore data from backup files.</p>↵</li>↵</ol>↵<p>If you have any technical problems, please contact Chanjet technical support: 4006600566-9</p>
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The introduction to the vulnerability.
	//
	// example:
	//
	// Chanjet T-Plus is an Internet business management software. There is an unauthorized access vulnerability in one of its interfaces disclosed on the Internet. Attackers can construct malicious requests to upload malicious files to execute arbitrary code and control the server.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Deprecated
	//
	// The ID of the asset that is scanned.
	//
	// >  This parameter is deprecated. You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the instance that is affected by the vulnerability.
	//
	// example:
	//
	// m-bp17m0pc0xprzbwo****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// Deprecated
	//
	// The name of the asset that is scanned.
	//
	// >  This parameter is deprecated. You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the instance that is affected by the vulnerability.
	//
	// example:
	//
	// frontend
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The title of the vulnerability announcement.
	//
	// example:
	//
	// Chanjet T-Plus SetupAccount/Upload. Aspx file upload vulnerability(CNVD-2022-60632)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Deprecated
	//
	// The vendor that disclosed the vulnerability.
	//
	// example:
	//
	// Apache
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The severity of the vulnerability. Valid values:
	//
	// 	- **serious**
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// serious
	VulLevel *string `json:"VulLevel,omitempty" xml:"VulLevel,omitempty"`
}

func (s DescribeVulDetailsResponseBodyCves) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCves) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCves) GetClassify() *string {
	return s.Classify
}

func (s *DescribeVulDetailsResponseBodyCves) GetClassifys() []*DescribeVulDetailsResponseBodyCvesClassifys {
	return s.Classifys
}

func (s *DescribeVulDetailsResponseBodyCves) GetCnvdId() *string {
	return s.CnvdId
}

func (s *DescribeVulDetailsResponseBodyCves) GetComplexity() *string {
	return s.Complexity
}

func (s *DescribeVulDetailsResponseBodyCves) GetContent() *string {
	return s.Content
}

func (s *DescribeVulDetailsResponseBodyCves) GetCveId() *string {
	return s.CveId
}

func (s *DescribeVulDetailsResponseBodyCves) GetCveLink() *string {
	return s.CveLink
}

func (s *DescribeVulDetailsResponseBodyCves) GetCvssScore() *string {
	return s.CvssScore
}

func (s *DescribeVulDetailsResponseBodyCves) GetCvssVector() *string {
	return s.CvssVector
}

func (s *DescribeVulDetailsResponseBodyCves) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeVulDetailsResponseBodyCves) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeVulDetailsResponseBodyCves) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeVulDetailsResponseBodyCves) GetOtherId() *string {
	return s.OtherId
}

func (s *DescribeVulDetailsResponseBodyCves) GetPoc() *string {
	return s.Poc
}

func (s *DescribeVulDetailsResponseBodyCves) GetPocCreateTime() *int64 {
	return s.PocCreateTime
}

func (s *DescribeVulDetailsResponseBodyCves) GetPocDisclosureTime() *int64 {
	return s.PocDisclosureTime
}

func (s *DescribeVulDetailsResponseBodyCves) GetProduct() *string {
	return s.Product
}

func (s *DescribeVulDetailsResponseBodyCves) GetReference() *string {
	return s.Reference
}

func (s *DescribeVulDetailsResponseBodyCves) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *DescribeVulDetailsResponseBodyCves) GetSolution() *string {
	return s.Solution
}

func (s *DescribeVulDetailsResponseBodyCves) GetSummary() *string {
	return s.Summary
}

func (s *DescribeVulDetailsResponseBodyCves) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeVulDetailsResponseBodyCves) GetTargetName() *string {
	return s.TargetName
}

func (s *DescribeVulDetailsResponseBodyCves) GetTitle() *string {
	return s.Title
}

func (s *DescribeVulDetailsResponseBodyCves) GetVendor() *string {
	return s.Vendor
}

func (s *DescribeVulDetailsResponseBodyCves) GetVulLevel() *string {
	return s.VulLevel
}

func (s *DescribeVulDetailsResponseBodyCves) SetClassify(v string) *DescribeVulDetailsResponseBodyCves {
	s.Classify = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetClassifys(v []*DescribeVulDetailsResponseBodyCvesClassifys) *DescribeVulDetailsResponseBodyCves {
	s.Classifys = v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCnvdId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CnvdId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetComplexity(v string) *DescribeVulDetailsResponseBodyCves {
	s.Complexity = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetContent(v string) *DescribeVulDetailsResponseBodyCves {
	s.Content = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCveId(v string) *DescribeVulDetailsResponseBodyCves {
	s.CveId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCveLink(v string) *DescribeVulDetailsResponseBodyCves {
	s.CveLink = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssScore(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssScore = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetCvssVector(v string) *DescribeVulDetailsResponseBodyCves {
	s.CvssVector = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetInstanceName(v string) *DescribeVulDetailsResponseBodyCves {
	s.InstanceName = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetInternetIp(v string) *DescribeVulDetailsResponseBodyCves {
	s.InternetIp = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetIntranetIp(v string) *DescribeVulDetailsResponseBodyCves {
	s.IntranetIp = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetOtherId(v string) *DescribeVulDetailsResponseBodyCves {
	s.OtherId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPoc(v string) *DescribeVulDetailsResponseBodyCves {
	s.Poc = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPocCreateTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.PocCreateTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetPocDisclosureTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.PocDisclosureTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetProduct(v string) *DescribeVulDetailsResponseBodyCves {
	s.Product = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetReference(v string) *DescribeVulDetailsResponseBodyCves {
	s.Reference = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetReleaseTime(v int64) *DescribeVulDetailsResponseBodyCves {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSolution(v string) *DescribeVulDetailsResponseBodyCves {
	s.Solution = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetSummary(v string) *DescribeVulDetailsResponseBodyCves {
	s.Summary = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTargetId(v string) *DescribeVulDetailsResponseBodyCves {
	s.TargetId = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTargetName(v string) *DescribeVulDetailsResponseBodyCves {
	s.TargetName = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetTitle(v string) *DescribeVulDetailsResponseBodyCves {
	s.Title = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetVendor(v string) *DescribeVulDetailsResponseBodyCves {
	s.Vendor = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) SetVulLevel(v string) *DescribeVulDetailsResponseBodyCves {
	s.VulLevel = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCves) Validate() error {
	return dara.Validate(s)
}

type DescribeVulDetailsResponseBodyCvesClassifys struct {
	// The type of the vulnerability.
	//
	// example:
	//
	// remote_code_execution
	Classify *string `json:"Classify,omitempty" xml:"Classify,omitempty"`
	// The URL of the demo video for the vulnerability.
	//
	// example:
	//
	// https://example.com
	DemoVideoUrl *string `json:"DemoVideoUrl,omitempty" xml:"DemoVideoUrl,omitempty"`
	// The description of the vulnerability type.
	//
	// example:
	//
	// Remote code execution
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s DescribeVulDetailsResponseBodyCvesClassifys) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulDetailsResponseBodyCvesClassifys) GoString() string {
	return s.String()
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) GetClassify() *string {
	return s.Classify
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) GetDemoVideoUrl() *string {
	return s.DemoVideoUrl
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) GetDescription() *string {
	return s.Description
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetClassify(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.Classify = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetDemoVideoUrl(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.DemoVideoUrl = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) SetDescription(v string) *DescribeVulDetailsResponseBodyCvesClassifys {
	s.Description = &v
	return s
}

func (s *DescribeVulDetailsResponseBodyCvesClassifys) Validate() error {
	return dara.Validate(s)
}
