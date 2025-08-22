// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DescribeDcdnWafGroupResponseBody
	GetId() *int64
	SetName(v string) *DescribeDcdnWafGroupResponseBody
	GetName() *string
	SetPageNumber(v int32) *DescribeDcdnWafGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDcdnWafGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDcdnWafGroupResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeDcdnWafGroupResponseBodyRules) *DescribeDcdnWafGroupResponseBody
	GetRules() []*DescribeDcdnWafGroupResponseBodyRules
	SetSubscribe(v string) *DescribeDcdnWafGroupResponseBody
	GetSubscribe() *string
	SetTemplateId(v int64) *DescribeDcdnWafGroupResponseBody
	GetTemplateId() *int64
	SetTotalCount(v int32) *DescribeDcdnWafGroupResponseBody
	GetTotalCount() *int32
}

type DescribeDcdnWafGroupResponseBody struct {
	// The ID of the custom WAF rule group.
	//
	// example:
	//
	// 1012
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the WAF rule group.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations of the rule.
	Rules []*DescribeDcdnWafGroupResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// Indicates whether to enable subscription. Valid values:
	//
	// 	- **on:**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Subscribe *string `json:"Subscribe,omitempty" xml:"Subscribe,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 1012
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The total number of rules that are filtered out.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDcdnWafGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupResponseBody) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDcdnWafGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDcdnWafGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafGroupResponseBody) GetRules() []*DescribeDcdnWafGroupResponseBodyRules {
	return s.Rules
}

func (s *DescribeDcdnWafGroupResponseBody) GetSubscribe() *string {
	return s.Subscribe
}

func (s *DescribeDcdnWafGroupResponseBody) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeDcdnWafGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDcdnWafGroupResponseBody) SetId(v int64) *DescribeDcdnWafGroupResponseBody {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetName(v string) *DescribeDcdnWafGroupResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetPageNumber(v int32) *DescribeDcdnWafGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetPageSize(v int32) *DescribeDcdnWafGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetRequestId(v string) *DescribeDcdnWafGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetRules(v []*DescribeDcdnWafGroupResponseBodyRules) *DescribeDcdnWafGroupResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetSubscribe(v string) *DescribeDcdnWafGroupResponseBody {
	s.Subscribe = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetTemplateId(v int64) *DescribeDcdnWafGroupResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) SetTotalCount(v int32) *DescribeDcdnWafGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGroupResponseBodyRules struct {
	// The type of the application. Valid values:
	//
	// 	- **0**: Common
	//
	// 	- **1**: WordPress
	//
	// 	- **2**: DedeCMS
	//
	// 	- **3**: Discuz
	//
	// 	- **4**: PHP CMS
	//
	// 	- **5**: ECShop
	//
	// 	- **6**: ShopEX
	//
	// 	- **7**: Drupal
	//
	// 	- **8**: Joomla
	//
	// 	- **9**: MetInfo
	//
	// 	- **10**: Struts2
	//
	// 	- **11**: Spring Boot
	//
	// 	- **12**: JBoss
	//
	// 	- **13**: WebLogic
	//
	// 	- **14**: WebSphere
	//
	// 	- **15**: Tomcat
	//
	// 	- **16**: Elastic Search
	//
	// 	- **18**: ThinkPHP
	//
	// 	- **19**: Fastjson
	//
	// 	- **20**: ImageMagick
	//
	// 	- **21**: PHPWind
	//
	// 	- **22**: phpMyAdmin
	//
	// 	- **23**: Resin
	//
	// 	- **24**: IIS
	//
	// 	- **99**: Others
	//
	// example:
	//
	// 1
	ApplicationType *int32 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The Common Vulnerabilities and Exposures (CVE) ID of the related vulnerability.
	//
	// example:
	//
	// CVE-2021-22945
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The CVE link.
	//
	// example:
	//
	// https://nxx.nxxx.gov/vuln/detail/CVE-2022-XXXX
	CveUrl *string `json:"CveUrl,omitempty" xml:"CveUrl,omitempty"`
	// The description of the WAF rule.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the rule was modified.
	//
	// example:
	//
	// 2021-12-29T17:08:45Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the custom WAF rule.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the WAF rule.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Protection type Valid values:
	//
	// 	- **11**: SQL injection
	//
	// 	- **12**: cross-site scripting (XSS)
	//
	// 	- **13**: code execution
	//
	// 	- **14**: carriage return line feeds (CRLF)
	//
	// 	- **15**: local file inclusion
	//
	// 	- **16**: remote file inclusion
	//
	// 	- **17**: webshells
	//
	// 	- **19**: cross-site request forgery
	//
	// 	- **20**: others
	//
	// 	- **21**: SEMA
	//
	// example:
	//
	// 11
	ProtectionType *int32 `json:"ProtectionType,omitempty" xml:"ProtectionType,omitempty"`
	// The risk level of the resources that do not comply with the managed rule. Valid values:
	//
	// 	- **1**: high risk
	//
	// 	- **2**: medium risk
	//
	// 	- **3**: low risk
	//
	// example:
	//
	// 2
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeDcdnWafGroupResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGroupResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetApplicationType() *int32 {
	return s.ApplicationType
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetCveId() *string {
	return s.CveId
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetCveUrl() *string {
	return s.CveUrl
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetId() *int64 {
	return s.Id
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetProtectionType() *int32 {
	return s.ProtectionType
}

func (s *DescribeDcdnWafGroupResponseBodyRules) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetApplicationType(v int32) *DescribeDcdnWafGroupResponseBodyRules {
	s.ApplicationType = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetCveId(v string) *DescribeDcdnWafGroupResponseBodyRules {
	s.CveId = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetCveUrl(v string) *DescribeDcdnWafGroupResponseBodyRules {
	s.CveUrl = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetDescription(v string) *DescribeDcdnWafGroupResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetGmtModified(v string) *DescribeDcdnWafGroupResponseBodyRules {
	s.GmtModified = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetId(v int64) *DescribeDcdnWafGroupResponseBodyRules {
	s.Id = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetName(v string) *DescribeDcdnWafGroupResponseBodyRules {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetProtectionType(v int32) *DescribeDcdnWafGroupResponseBodyRules {
	s.ProtectionType = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) SetRiskLevel(v int32) *DescribeDcdnWafGroupResponseBodyRules {
	s.RiskLevel = &v
	return s
}

func (s *DescribeDcdnWafGroupResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
