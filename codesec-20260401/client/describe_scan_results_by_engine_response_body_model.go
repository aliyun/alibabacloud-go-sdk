// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanResultsByEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeScanResultsByEngineResponseBody
	GetEngine() *string
	SetItems(v []*DescribeScanResultsByEngineResponseBodyItems) *DescribeScanResultsByEngineResponseBody
	GetItems() []*DescribeScanResultsByEngineResponseBodyItems
	SetMaxResults(v int64) *DescribeScanResultsByEngineResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *DescribeScanResultsByEngineResponseBody
	GetNextToken() *string
	SetProjectId(v int64) *DescribeScanResultsByEngineResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *DescribeScanResultsByEngineResponseBody
	GetRequestId() *string
	SetScanId(v int64) *DescribeScanResultsByEngineResponseBody
	GetScanId() *int64
	SetTotalCount(v int64) *DescribeScanResultsByEngineResponseBody
	GetTotalCount() *int64
}

type DescribeScanResultsByEngineResponseBody struct {
	// The engine type. Valid values:
	//
	// 	- sast
	//
	// 	- sca
	//
	// example:
	//
	// sca
	Engine *string `json:"engine,omitempty" xml:"engine,omitempty"`
	// The result list.
	Items []*DescribeScanResultsByEngineResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not pass nextToken or pass an empty string for the first page. To retrieve the next page, pass the nextToken value from the previous response without any modification. When the nextToken in the response is empty, you have reached the last page.
	//
	// example:
	//
	// eyJ0IjoiMjAyNi0wNy0xNlQwNzo1MzozOC4wMjFaIiwiaSI6MTAwMDQ0OH0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 76851f2b5bf0187fbc29e8bca4
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 3088795
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s DescribeScanResultsByEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeScanResultsByEngineResponseBody) GetItems() []*DescribeScanResultsByEngineResponseBodyItems {
	return s.Items
}

func (s *DescribeScanResultsByEngineResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *DescribeScanResultsByEngineResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScanResultsByEngineResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DescribeScanResultsByEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScanResultsByEngineResponseBody) GetScanId() *int64 {
	return s.ScanId
}

func (s *DescribeScanResultsByEngineResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeScanResultsByEngineResponseBody) SetEngine(v string) *DescribeScanResultsByEngineResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetItems(v []*DescribeScanResultsByEngineResponseBodyItems) *DescribeScanResultsByEngineResponseBody {
	s.Items = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetMaxResults(v int64) *DescribeScanResultsByEngineResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetNextToken(v string) *DescribeScanResultsByEngineResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetProjectId(v int64) *DescribeScanResultsByEngineResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetRequestId(v string) *DescribeScanResultsByEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetScanId(v int64) *DescribeScanResultsByEngineResponseBody {
	s.ScanId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) SetTotalCount(v int64) *DescribeScanResultsByEngineResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBody) Validate() error {
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

type DescribeScanResultsByEngineResponseBodyItems struct {
	// Filters results by incremental scan baseline status. Valid values: new, unchanged, absent, updated.
	//
	// example:
	//
	// new
	BaselineState *string `json:"baselineState,omitempty" xml:"baselineState,omitempty"`
	// The category. The system classifies files based on file name extensions and MIME types. Common categories include doc, image, audio, and video.
	//
	// example:
	//
	// SQL Injection
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The code snippet near the primary location (SAST).
	//
	// example:
	//
	// 1
	CodeSnippet *string `json:"codeSnippet,omitempty" xml:"codeSnippet,omitempty"`
	// The rule confidence level, ranging from 0 to 1. This field is common in SAST results and is omitted if not applicable.
	//
	// example:
	//
	// 1
	Confidence *float64 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// The time when the finding record was created (RFC 3339 format).
	//
	// example:
	//
	// 2026-07-28T03:36:31.573Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The associated CWE ID.
	//
	// example:
	//
	// CWE-601
	CweId *string `json:"cweId,omitempty" xml:"cweId,omitempty"`
	// The issue description.
	//
	// example:
	//
	// User input is directly concatenated into an SQL statement, which may lead to SQL injection.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The end line number.
	//
	// example:
	//
	// 1
	EndLine *int64 `json:"endLine,omitempty" xml:"endLine,omitempty"`
	// The file path.
	//
	// example:
	//
	// aone2/2247063/1767838276141/Artifacts_EBookingApps_Pipeline_V2_Staging_131.tgz
	FilePath *string `json:"filePath,omitempty" xml:"filePath,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 934
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The brief summary of the finding. Unlike description, this field is more of a conclusion statement.
	//
	// example:
	//
	// User input is used to construct SQL queries without sanitization
	ItemSummary *string `json:"itemSummary,omitempty" xml:"itemSummary,omitempty"`
	// The OWASP category.
	//
	// example:
	//
	// A01:2021-Broken Access Control
	OwaspCategory *string `json:"owaspCategory,omitempty" xml:"owaspCategory,omitempty"`
	// The project name.
	//
	// example:
	//
	// cdp2cdl_cdp_sha_nprd
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// The remediation code example.
	//
	// example:
	//
	// String sql = \\"SELECT 	- FROM users WHERE name = ?\\";\\nPreparedStatement ps = conn.prepareStatement(sql);\\nps.setString(1, username);\\nResultSet rs = ps.executeQuery();
	RemediationCodeExample *string `json:"remediationCodeExample,omitempty" xml:"remediationCodeExample,omitempty"`
	// The remediation suggestion.
	//
	// example:
	//
	// Use PreparedStatement with parameterized queries instead of string concatenation
	RemediationSuggestion *string `json:"remediationSuggestion,omitempty" xml:"remediationSuggestion,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// java-sqli-001
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The SCA component information. This field is returned only when engine is set to sca.
	ScaComponent *DescribeScanResultsByEngineResponseBodyItemsScaComponent `json:"scaComponent,omitempty" xml:"scaComponent,omitempty" type:"Struct"`
	// The task ID.
	//
	// example:
	//
	// 3176
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// The severity level. Valid values:
	//
	// 	- critical
	//
	// 	- high
	//
	// 	- medium
	//
	// 	- low
	//
	// example:
	//
	// high
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
	// The source.
	//
	// example:
	//
	// todo_open_dingoj06pvqfeayy3lkr
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The start line number.
	//
	// example:
	//
	// 1
	StartLine *int64 `json:"startLine,omitempty" xml:"startLine,omitempty"`
	// The status. Valid values:
	//
	// 	- running: Running.
	//
	// 	- completed: Completed.
	//
	// 	- failed: Failed.
	//
	// example:
	//
	// completed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The SAST taint analysis call chain that describes the complete propagation path of sensitive data from the taint source to the dangerous sink. This field is returned only when engine is set to sast.
	TaintFlow []*DescribeScanResultsByEngineResponseBodyItemsTaintFlow `json:"taintFlow,omitempty" xml:"taintFlow,omitempty" type:"Repeated"`
	// The text summary of the taint call chain. This field is returned only when engine is set to sast.
	//
	// example:
	//
	// User input flows from HTTP parameter into SQL query executio
	TaintFlowSummary *string `json:"taintFlowSummary,omitempty" xml:"taintFlowSummary,omitempty"`
	// The issue title.
	//
	// example:
	//
	// SQL Injection Risk
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s DescribeScanResultsByEngineResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetBaselineState() *string {
	return s.BaselineState
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetCategory() *string {
	return s.Category
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetCodeSnippet() *string {
	return s.CodeSnippet
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetConfidence() *float64 {
	return s.Confidence
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetCweId() *string {
	return s.CweId
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetEndLine() *int64 {
	return s.EndLine
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetItemSummary() *string {
	return s.ItemSummary
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetOwaspCategory() *string {
	return s.OwaspCategory
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetRemediationCodeExample() *string {
	return s.RemediationCodeExample
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetRemediationSuggestion() *string {
	return s.RemediationSuggestion
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetScaComponent() *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	return s.ScaComponent
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetScanId() *int64 {
	return s.ScanId
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetSource() *string {
	return s.Source
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetStartLine() *int64 {
	return s.StartLine
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetTaintFlow() []*DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	return s.TaintFlow
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetTaintFlowSummary() *string {
	return s.TaintFlowSummary
}

func (s *DescribeScanResultsByEngineResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetBaselineState(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.BaselineState = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetCategory(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Category = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetCodeSnippet(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.CodeSnippet = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetConfidence(v float64) *DescribeScanResultsByEngineResponseBodyItems {
	s.Confidence = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetCreatedAt(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetCweId(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.CweId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetDescription(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetEndLine(v int64) *DescribeScanResultsByEngineResponseBodyItems {
	s.EndLine = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetFilePath(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.FilePath = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetId(v int64) *DescribeScanResultsByEngineResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetItemSummary(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.ItemSummary = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetOwaspCategory(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.OwaspCategory = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetProjectName(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.ProjectName = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetRemediationCodeExample(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.RemediationCodeExample = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetRemediationSuggestion(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.RemediationSuggestion = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetRuleId(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.RuleId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetScaComponent(v *DescribeScanResultsByEngineResponseBodyItemsScaComponent) *DescribeScanResultsByEngineResponseBodyItems {
	s.ScaComponent = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetScanId(v int64) *DescribeScanResultsByEngineResponseBodyItems {
	s.ScanId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetSeverity(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Severity = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetSource(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Source = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetStartLine(v int64) *DescribeScanResultsByEngineResponseBodyItems {
	s.StartLine = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetStatus(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetTaintFlow(v []*DescribeScanResultsByEngineResponseBodyItemsTaintFlow) *DescribeScanResultsByEngineResponseBodyItems {
	s.TaintFlow = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetTaintFlowSummary(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.TaintFlowSummary = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) SetTitle(v string) *DescribeScanResultsByEngineResponseBodyItems {
	s.Title = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItems) Validate() error {
	if s.ScaComponent != nil {
		if err := s.ScaComponent.Validate(); err != nil {
			return err
		}
	}
	if s.TaintFlow != nil {
		for _, item := range s.TaintFlow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScanResultsByEngineResponseBodyItemsScaComponent struct {
	// The number of CVEs.
	//
	// example:
	//
	// 1
	CveCount *int64 `json:"cveCount,omitempty" xml:"cveCount,omitempty"`
	// The list of CVE details associated with a component in the SCA finding.
	CveDetails []*DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails `json:"cveDetails,omitempty" xml:"cveDetails,omitempty" type:"Repeated"`
	// The list of dependency introduction paths in the SCA component information. This field is returned only when engine is set to sca.
	IntroPaths []*string `json:"introPaths,omitempty" xml:"introPaths,omitempty" type:"Repeated"`
	// Indicates whether the component is a direct dependency.
	//
	// example:
	//
	// false
	IsDirect *bool `json:"isDirect,omitempty" xml:"isDirect,omitempty"`
	// The component coordinate.
	//
	// example:
	//
	// org.apache.logging.log4j:log4j-core
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
	// The component-level remediation suggestion.
	//
	// example:
	//
	// Upgrade log4j-core to version 2.17.1 or later
	Remediation *string `json:"remediation,omitempty" xml:"remediation,omitempty"`
	// The component version.
	//
	// example:
	//
	// 2.14.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeScanResultsByEngineResponseBodyItemsScaComponent) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponseBodyItemsScaComponent) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetCveCount() *int64 {
	return s.CveCount
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetCveDetails() []*DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	return s.CveDetails
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetIntroPaths() []*string {
	return s.IntroPaths
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetIsDirect() *bool {
	return s.IsDirect
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetPackageName() *string {
	return s.PackageName
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetRemediation() *string {
	return s.Remediation
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) GetVersion() *string {
	return s.Version
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetCveCount(v int64) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.CveCount = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetCveDetails(v []*DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.CveDetails = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetIntroPaths(v []*string) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.IntroPaths = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetIsDirect(v bool) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.IsDirect = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetPackageName(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.PackageName = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetRemediation(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.Remediation = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) SetVersion(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponent {
	s.Version = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponent) Validate() error {
	if s.CveDetails != nil {
		for _, item := range s.CveDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails struct {
	// The associated CWE ID.
	//
	// example:
	//
	// CVE-2021-44228
	CveId *string `json:"cveId,omitempty" xml:"cveId,omitempty"`
	// The CVSS score.
	//
	// example:
	//
	// 1
	Cvss *float64 `json:"cvss,omitempty" xml:"cvss,omitempty"`
	// The CVSS version.
	//
	// example:
	//
	// v2.0
	CvssVersion *string `json:"cvssVersion,omitempty" xml:"cvssVersion,omitempty"`
	// The description.
	//
	// example:
	//
	// Apache Log4j2 JNDI features do not protect against attacker-controlled LDAP and other JNDI-related endpoints.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The reference information.
	References []*string `json:"references,omitempty" xml:"references,omitempty" type:"Repeated"`
	// The severity level. Valid values:
	//
	// 	- critical
	//
	// 	- high
	//
	// 	- medium
	//
	// 	- low
	//
	// example:
	//
	// critical
	Severity *string `json:"severity,omitempty" xml:"severity,omitempty"`
}

func (s DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetCveId() *string {
	return s.CveId
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetCvss() *float64 {
	return s.Cvss
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetCvssVersion() *string {
	return s.CvssVersion
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetDescription() *string {
	return s.Description
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetReferences() []*string {
	return s.References
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetCveId(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.CveId = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetCvss(v float64) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.Cvss = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetCvssVersion(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.CvssVersion = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetDescription(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.Description = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetReferences(v []*string) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.References = v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) SetSeverity(v string) *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails {
	s.Severity = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsScaComponentCveDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeScanResultsByEngineResponseBodyItemsTaintFlow struct {
	// The code.
	//
	// example:
	//
	// stmt.executeQuery(sql);
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The file path.
	//
	// example:
	//
	// src/main/java/com/example/UserService.java
	File *string `json:"file,omitempty" xml:"file,omitempty"`
	// The role type in the taint propagation chain. Valid values:
	//
	// 	- source: taint source.
	//
	// 	- propagator: propagation node.
	//
	// 	- validation: validation or scrubbing center.
	//
	// 	- sink: dangerous sink.
	//
	// example:
	//
	// source
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The line number.
	//
	// example:
	//
	// 111
	Line *int32 `json:"line,omitempty" xml:"line,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Sink: SQL query executed
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// The step number, starting from 0 and incrementing.
	//
	// example:
	//
	// 1
	Step *int32 `json:"step,omitempty" xml:"step,omitempty"`
}

func (s DescribeScanResultsByEngineResponseBodyItemsTaintFlow) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GoString() string {
	return s.String()
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetCode() *string {
	return s.Code
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetFile() *string {
	return s.File
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetKind() *string {
	return s.Kind
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetLine() *int32 {
	return s.Line
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetNote() *string {
	return s.Note
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) GetStep() *int32 {
	return s.Step
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetCode(v string) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.Code = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetFile(v string) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.File = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetKind(v string) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.Kind = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetLine(v int32) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.Line = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetNote(v string) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.Note = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) SetStep(v int32) *DescribeScanResultsByEngineResponseBodyItemsTaintFlow {
	s.Step = &v
	return s
}

func (s *DescribeScanResultsByEngineResponseBodyItemsTaintFlow) Validate() error {
	return dara.Validate(s)
}
