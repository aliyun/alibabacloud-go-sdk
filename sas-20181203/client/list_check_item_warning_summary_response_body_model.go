// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListCheckItemWarningSummaryResponseBodyList) *ListCheckItemWarningSummaryResponseBody
	GetList() []*ListCheckItemWarningSummaryResponseBodyList
	SetPageInfo(v *ListCheckItemWarningSummaryResponseBodyPageInfo) *ListCheckItemWarningSummaryResponseBody
	GetPageInfo() *ListCheckItemWarningSummaryResponseBodyPageInfo
	SetRequestId(v string) *ListCheckItemWarningSummaryResponseBody
	GetRequestId() *string
}

type ListCheckItemWarningSummaryResponseBody struct {
	// List of check item risk statistics.
	List []*ListCheckItemWarningSummaryResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCheckItemWarningSummaryResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DC97C9EC-4B7D-5EFF-8A5E-A5CCC9ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckItemWarningSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningSummaryResponseBody) GetList() []*ListCheckItemWarningSummaryResponseBodyList {
	return s.List
}

func (s *ListCheckItemWarningSummaryResponseBody) GetPageInfo() *ListCheckItemWarningSummaryResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCheckItemWarningSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckItemWarningSummaryResponseBody) SetList(v []*ListCheckItemWarningSummaryResponseBodyList) *ListCheckItemWarningSummaryResponseBody {
	s.List = v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBody) SetPageInfo(v *ListCheckItemWarningSummaryResponseBodyPageInfo) *ListCheckItemWarningSummaryResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBody) SetRequestId(v string) *ListCheckItemWarningSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningSummaryResponseBodyList struct {
	// The suggestion on the check item.
	//
	// example:
	//
	// In the Administrative Tools window, double-click Local Security Policy. In the Local Security Policy window that appears, choose Security Settings\\\\Local Policies\\\\Audit Policy, configure all audit policies as: `Success, Failure`.
	Advice *string `json:"Advice,omitempty" xml:"Advice,omitempty"`
	// The types of the baselines to which the check item belongs.
	AffiliatedRiskTypes []*string `json:"AffiliatedRiskTypes,omitempty" xml:"AffiliatedRiskTypes,omitempty" type:"Repeated"`
	// The baselines to which the check item belongs.
	AffiliatedRisks []*string `json:"AffiliatedRisks,omitempty" xml:"AffiliatedRisks,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The alias of the baseline type.
	//
	// example:
	//
	// week_pa****
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The ID of the check item.
	//
	// example:
	//
	// 696
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// Config the Event Audit policys
	CheckItem *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	// The risk level of the check item. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// high
	CheckLevel *string `json:"CheckLevel,omitempty" xml:"CheckLevel,omitempty"`
	// The type of the check item.
	//
	// example:
	//
	// Security audit
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// Indicates whether the check item belongs to the container runtime type. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ContainerCheckItem *bool `json:"ContainerCheckItem,omitempty" xml:"ContainerCheckItem,omitempty"`
	// The description of the check item.
	//
	// example:
	//
	// Config the Event Audit policys
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The baselines in which the check item is enabled.
	EnableRisks []*string `json:"EnableRisks,omitempty" xml:"EnableRisks,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The type of the baseline.
	//
	// example:
	//
	// weak_password
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// Risk status of check items. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **3**: passed
	//
	// 	- **6**: whitelisted
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of servers that are affected by the check item.
	//
	// example:
	//
	// 20
	WarningMachineCount *int32 `json:"WarningMachineCount,omitempty" xml:"WarningMachineCount,omitempty"`
}

func (s ListCheckItemWarningSummaryResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningSummaryResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetAdvice() *string {
	return s.Advice
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetAffiliatedRiskTypes() []*string {
	return s.AffiliatedRiskTypes
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetAffiliatedRisks() []*string {
	return s.AffiliatedRisks
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetAlias() *string {
	return s.Alias
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetCheckItem() *string {
	return s.CheckItem
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetCheckLevel() *string {
	return s.CheckLevel
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetCheckType() *string {
	return s.CheckType
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetContainerCheckItem() *bool {
	return s.ContainerCheckItem
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetDescription() *string {
	return s.Description
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetEnableRisks() []*string {
	return s.EnableRisks
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetRiskType() *string {
	return s.RiskType
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListCheckItemWarningSummaryResponseBodyList) GetWarningMachineCount() *int32 {
	return s.WarningMachineCount
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetAdvice(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.Advice = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetAffiliatedRiskTypes(v []*string) *ListCheckItemWarningSummaryResponseBodyList {
	s.AffiliatedRiskTypes = v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetAffiliatedRisks(v []*string) *ListCheckItemWarningSummaryResponseBodyList {
	s.AffiliatedRisks = v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetAlias(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.Alias = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetCheckId(v int64) *ListCheckItemWarningSummaryResponseBodyList {
	s.CheckId = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetCheckItem(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.CheckItem = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetCheckLevel(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.CheckLevel = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetCheckType(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.CheckType = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetContainerCheckItem(v bool) *ListCheckItemWarningSummaryResponseBodyList {
	s.ContainerCheckItem = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetDescription(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetEnableRisks(v []*string) *ListCheckItemWarningSummaryResponseBodyList {
	s.EnableRisks = v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetRiskType(v string) *ListCheckItemWarningSummaryResponseBodyList {
	s.RiskType = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetStatus(v int32) *ListCheckItemWarningSummaryResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) SetWarningMachineCount(v int32) *ListCheckItemWarningSummaryResponseBodyList {
	s.WarningMachineCount = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningSummaryResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckItemWarningSummaryResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningSummaryResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) SetCount(v int32) *ListCheckItemWarningSummaryResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) SetCurrentPage(v int32) *ListCheckItemWarningSummaryResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) SetPageSize(v int32) *ListCheckItemWarningSummaryResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) SetTotalCount(v int32) *ListCheckItemWarningSummaryResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
