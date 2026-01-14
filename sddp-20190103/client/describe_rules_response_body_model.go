// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeRulesResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeRulesResponseBodyItems) *DescribeRulesResponseBody
	GetItems() []*DescribeRulesResponseBodyItems
	SetPageSize(v int32) *DescribeRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRulesResponseBody
	GetTotalCount() *int32
}

type DescribeRulesResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The sensitive data detection rules.
	Items []*DescribeRulesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989C13
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRulesResponseBody) GetItems() []*DescribeRulesResponseBodyItems {
	return s.Items
}

func (s *DescribeRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRulesResponseBody) SetCurrentPage(v int32) *DescribeRulesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRulesResponseBody) SetItems(v []*DescribeRulesResponseBodyItems) *DescribeRulesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeRulesResponseBody) SetPageSize(v int32) *DescribeRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRulesResponseBody) SetTotalCount(v int32) *DescribeRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRulesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRulesResponseBodyItems struct {
	// example:
	//
	// 0
	AuditMode *int32 `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The content type of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: keyword
	//
	// 	- **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the content type of the sensitive data detection rule.
	//
	// example:
	//
	// Regular expression
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The content in the sensitive data detection rule.
	//
	// >  A built-in detection rule whose CustomType is 0 does not return the content of the rule.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The type of the content in the sensitive data detection rule. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates attempts to exploit SQL injections. The value 2 indicates bypass by using SQL injections. The value 3 indicates abuse of stored procedures. The value 4 indicates buffer overflow. The value 5 indicates SQL injections based on errors.
	//
	// example:
	//
	// 1
	ContentCategory *string `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The type of the sensitive data detection rule.
	//
	// 	- 0: built-in rule
	//
	// 	- 1: custom rule
	//
	// example:
	//
	// 1
	CustomType *int32 `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The description of the sensitive data detection rule.
	//
	// example:
	//
	// The sensitive data detection rule is used to detect IP addresses.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// ****test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the sensitive data detection rule is created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the sensitive data detection rule is modified. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The parent group type of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of times that the sensitive data detection rule is hit.
	//
	// example:
	//
	// 3
	HitTotalCount *int32 `json:"HitTotalCount,omitempty" xml:"HitTotalCount,omitempty"`
	// The ID of the sensitive data detection rule.
	//
	// example:
	//
	// 20000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The username of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The key of the primary dimension.
	//
	// example:
	//
	// key
	MajorKey *string `json:"MajorKey,omitempty" xml:"MajorKey,omitempty"`
	// The match type. Valid values:
	//
	// 	- **1**: rule-based match
	//
	// 	- **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The IDs of the models for sensitive data audit.
	//
	// example:
	//
	// 1452
	ModelRuleIds *string `json:"ModelRuleIds,omitempty" xml:"ModelRuleIds,omitempty"`
	// The name of the sensitive data detection rule.
	//
	// example:
	//
	// IP address
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the sensitive data detection rule is applied. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level of the sensitive data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **1**: N/A, which indicates that no sensitive data is detected.
	//
	// 	- **2**: S1, which indicates the low sensitivity level.
	//
	// 	- **3**: S2, which indicates the medium sensitivity level.
	//
	// 	- **4**: S3, which indicates the high sensitivity level.
	//
	// 	- **5**: S4, which indicates the highest sensitivity level.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The sensitivity level of data that hits the sensitive data detection rule. Valid values:
	//
	// 	- **N/A**: indicates that no sensitive data is detected.
	//
	// 	- **S1**: indicates the low sensitivity level.
	//
	// 	- **S2**: indicates the medium sensitivity level.
	//
	// 	- **S3**: indicates the high sensitivity level.
	//
	// 	- **S4**: indicates the highest sensitivity level.
	//
	// example:
	//
	// S2
	RiskLevelName *string `json:"RiskLevelName,omitempty" xml:"RiskLevelName,omitempty"`
	// The statistical expression.
	//
	// example:
	//
	// 1
	StatExpress *string `json:"StatExpress,omitempty" xml:"StatExpress,omitempty"`
	// The status of the sensitive data detection rule. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The data asset type that is supported by the sensitive data detection rule. Valid values:
	//
	// 	- **0**: all data assets
	//
	// 	- **1**: structured data assets
	//
	// 	- **2**: unstructured data assets
	//
	// example:
	//
	// 2
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The name of the service to which the data asset belongs. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The IDs of the templates that are used to audit sensitive data.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// example:
	//
	// 0
	ThreatAnalysisStatus *int32 `json:"ThreatAnalysisStatus,omitempty" xml:"ThreatAnalysisStatus,omitempty"`
	// The ID of the account that is used to create the sensitive data detection rule.
	//
	// example:
	//
	// 0
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The severity level. Valid values:
	//
	// 	- **1**: low
	//
	// 	- **2**: medium
	//
	// 	- **3**: high
	//
	// example:
	//
	// 2
	WarnLevel *int32 `json:"WarnLevel,omitempty" xml:"WarnLevel,omitempty"`
}

func (s DescribeRulesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeRulesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyItems) GetAuditMode() *int32 {
	return s.AuditMode
}

func (s *DescribeRulesResponseBodyItems) GetCategory() *int32 {
	return s.Category
}

func (s *DescribeRulesResponseBodyItems) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeRulesResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *DescribeRulesResponseBodyItems) GetContentCategory() *string {
	return s.ContentCategory
}

func (s *DescribeRulesResponseBodyItems) GetCustomType() *int32 {
	return s.CustomType
}

func (s *DescribeRulesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeRulesResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeRulesResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeRulesResponseBodyItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeRulesResponseBodyItems) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeRulesResponseBodyItems) GetHitTotalCount() *int32 {
	return s.HitTotalCount
}

func (s *DescribeRulesResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeRulesResponseBodyItems) GetLoginName() *string {
	return s.LoginName
}

func (s *DescribeRulesResponseBodyItems) GetMajorKey() *string {
	return s.MajorKey
}

func (s *DescribeRulesResponseBodyItems) GetMatchType() *int32 {
	return s.MatchType
}

func (s *DescribeRulesResponseBodyItems) GetModelRuleIds() *string {
	return s.ModelRuleIds
}

func (s *DescribeRulesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeRulesResponseBodyItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeRulesResponseBodyItems) GetProductId() *int64 {
	return s.ProductId
}

func (s *DescribeRulesResponseBodyItems) GetRiskLevelId() *int64 {
	return s.RiskLevelId
}

func (s *DescribeRulesResponseBodyItems) GetRiskLevelName() *string {
	return s.RiskLevelName
}

func (s *DescribeRulesResponseBodyItems) GetStatExpress() *string {
	return s.StatExpress
}

func (s *DescribeRulesResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeRulesResponseBodyItems) GetSupportForm() *int32 {
	return s.SupportForm
}

func (s *DescribeRulesResponseBodyItems) GetTarget() *string {
	return s.Target
}

func (s *DescribeRulesResponseBodyItems) GetTemplateRuleIds() *string {
	return s.TemplateRuleIds
}

func (s *DescribeRulesResponseBodyItems) GetThreatAnalysisStatus() *int32 {
	return s.ThreatAnalysisStatus
}

func (s *DescribeRulesResponseBodyItems) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeRulesResponseBodyItems) GetWarnLevel() *int32 {
	return s.WarnLevel
}

func (s *DescribeRulesResponseBodyItems) SetAuditMode(v int32) *DescribeRulesResponseBodyItems {
	s.AuditMode = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCategory(v int32) *DescribeRulesResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCategoryName(v string) *DescribeRulesResponseBodyItems {
	s.CategoryName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContent(v string) *DescribeRulesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetContentCategory(v string) *DescribeRulesResponseBodyItems {
	s.ContentCategory = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetCustomType(v int32) *DescribeRulesResponseBodyItems {
	s.CustomType = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetDescription(v string) *DescribeRulesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetDisplayName(v string) *DescribeRulesResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtCreate(v int64) *DescribeRulesResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGmtModified(v int64) *DescribeRulesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetGroupId(v string) *DescribeRulesResponseBodyItems {
	s.GroupId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetHitTotalCount(v int32) *DescribeRulesResponseBodyItems {
	s.HitTotalCount = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetId(v int64) *DescribeRulesResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetLoginName(v string) *DescribeRulesResponseBodyItems {
	s.LoginName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetMajorKey(v string) *DescribeRulesResponseBodyItems {
	s.MajorKey = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetMatchType(v int32) *DescribeRulesResponseBodyItems {
	s.MatchType = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetModelRuleIds(v string) *DescribeRulesResponseBodyItems {
	s.ModelRuleIds = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetName(v string) *DescribeRulesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductCode(v string) *DescribeRulesResponseBodyItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetProductId(v int64) *DescribeRulesResponseBodyItems {
	s.ProductId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelId(v int64) *DescribeRulesResponseBodyItems {
	s.RiskLevelId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetRiskLevelName(v string) *DescribeRulesResponseBodyItems {
	s.RiskLevelName = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatExpress(v string) *DescribeRulesResponseBodyItems {
	s.StatExpress = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetStatus(v int32) *DescribeRulesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetSupportForm(v int32) *DescribeRulesResponseBodyItems {
	s.SupportForm = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetTarget(v string) *DescribeRulesResponseBodyItems {
	s.Target = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetTemplateRuleIds(v string) *DescribeRulesResponseBodyItems {
	s.TemplateRuleIds = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetThreatAnalysisStatus(v int32) *DescribeRulesResponseBodyItems {
	s.ThreatAnalysisStatus = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetUserId(v int64) *DescribeRulesResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) SetWarnLevel(v int32) *DescribeRulesResponseBodyItems {
	s.WarnLevel = &v
	return s
}

func (s *DescribeRulesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
