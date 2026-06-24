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
	// A list of sensitive data detection rules.
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
	// The audit mode.
	//
	// example:
	//
	// 0
	AuditMode *int32 `json:"AuditMode,omitempty" xml:"AuditMode,omitempty"`
	// The type of content in the sensitive data detection rule. Valid values:
	//
	// - **0**: keyword
	//
	// - **2**: regular expression
	//
	// example:
	//
	// 2
	Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
	// The name of the content type for the sensitive data detection rule.
	//
	// example:
	//
	// Regular expression
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// The content of the sensitive data detection rule.
	//
	// > The content of a built-in rule, for which CustomType is 0, is not returned.
	//
	// example:
	//
	// (?:\\\\D|^)((?:(?:25[0-4]|2[0-4]\\\\d|1\\\\d{2}|[1-9]\\\\d{1})\\\\.)(?:(?:25[0-5]|2[0-4]\\\\d|[01]?\\\\d?\\\\d)\\\\.){2}(?:25[0-5]|2[0-4]\\\\d|1[0-9]\\\\d|[1-9]\\\\d|[1-9]))(?:\\\\D|$)
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type. Valid values:
	//
	// - **1**: SQL injection exploits
	//
	// - **2**: SQL injection bypass attempts
	//
	// - **3**: stored procedure abuse
	//
	// - **4**: buffer overflows
	//
	// - **5**: error-based SQL injections
	//
	// example:
	//
	// 1
	ContentCategory *string `json:"ContentCategory,omitempty" xml:"ContentCategory,omitempty"`
	// The type of the sensitive data detection rule.
	//
	// - 0: built-in
	//
	// - 1: custom
	//
	// example:
	//
	// 1
	CustomType *int32 `json:"CustomType,omitempty" xml:"CustomType,omitempty"`
	// The description of the sensitive data detection rule.
	//
	// example:
	//
	// Used to identify IP addresses
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the user who created the sensitive data detection rule.
	//
	// example:
	//
	// ****test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The time when the sensitive data detection rule was created. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the sensitive data detection rule was last modified. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1545277010000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The parent group of the rule.
	//
	// example:
	//
	// 4_1
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The number of times the rule was hit.
	//
	// example:
	//
	// 3
	HitTotalCount *int32 `json:"HitTotalCount,omitempty" xml:"HitTotalCount,omitempty"`
	// The unique ID of the sensitive data detection rule.
	//
	// example:
	//
	// 20000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The logon name of the user who created the sensitive data detection rule.
	//
	// example:
	//
	// det1111
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// The primary dimension key.
	//
	// example:
	//
	// key
	MajorKey *string `json:"MajorKey,omitempty" xml:"MajorKey,omitempty"`
	// The match type. Valid values:
	//
	// - **1**: rule-based match
	//
	// - **2**: dictionary-based match
	//
	// example:
	//
	// 1
	MatchType *int32 `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// A collection of model IDs for sensitive data auditing.
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
	// The name of the service to which the data asset belongs. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADS**
	//
	// - **OTS**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// example:
	//
	// MaxCompute
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the service to which the data asset belongs. Valid values:
	//
	// - **1**: MaxCompute
	//
	// - **2**: OSS
	//
	// - **3**: ADS
	//
	// - **4**: OTS
	//
	// - **5**: RDS
	//
	// - **6**: SELF_DB
	//
	// example:
	//
	// 2
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The sensitivity level ID of the sensitive data detection rule. Valid values:
	//
	// - **1**: N/A. No sensitive data is detected.
	//
	// - **2**: S1. Level 1 sensitive data.
	//
	// - **3**: S2. Level 2 sensitive data.
	//
	// - **4**: S3. Level 3 sensitive data.
	//
	// - **5**: S4. Level 4 sensitive data.
	//
	// example:
	//
	// 2
	RiskLevelId *int64 `json:"RiskLevelId,omitempty" xml:"RiskLevelId,omitempty"`
	// The name of the sensitivity level for the sensitive data detection rule. Valid values:
	//
	// - **N/A**: No sensitive data is detected.
	//
	// - **S1**: Level 1 sensitive data.
	//
	// - **S2**: Level 2 sensitive data.
	//
	// - **S3**: Level 3 sensitive data.
	//
	// - **S4**: Level 4 sensitive data.
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
	// The detection status of the sensitive data detection rule. Valid values:
	//
	// - **0**: disabled
	//
	// - **1**: enabled
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of data asset that the rule supports. Valid values:
	//
	// - **0**: all assets
	//
	// - **1**: structured assets
	//
	// - **2**: unstructured assets
	//
	// example:
	//
	// 2
	SupportForm *int32 `json:"SupportForm,omitempty" xml:"SupportForm,omitempty"`
	// The name of the service to which the data asset belongs. Valid values:
	//
	// - **MaxCompute**
	//
	// - **OSS**
	//
	// - **ADS**
	//
	// - **OTS**
	//
	// - **RDS**
	//
	// - **SELF_DB**
	//
	// example:
	//
	// MaxCompute
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// A collection of template IDs for sensitive data auditing.
	//
	// example:
	//
	// 1
	TemplateRuleIds *string `json:"TemplateRuleIds,omitempty" xml:"TemplateRuleIds,omitempty"`
	// The threat analysis mode status. Valid values: 0 (disabled), 1 (enabled).
	//
	// example:
	//
	// 0
	ThreatAnalysisStatus *int32 `json:"ThreatAnalysisStatus,omitempty" xml:"ThreatAnalysisStatus,omitempty"`
	// The ID of the user who created the sensitive data detection rule.
	//
	// example:
	//
	// 0
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The risk level.
	//
	// - **1**: Low
	//
	// - **2**: Medium
	//
	// - **3**: High
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
