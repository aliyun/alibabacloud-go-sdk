// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeCheckWarningSummaryResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *DescribeCheckWarningSummaryResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeCheckWarningSummaryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCheckWarningSummaryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCheckWarningSummaryResponseBody
	GetTotalCount() *int32
	SetWarningSummarys(v []*DescribeCheckWarningSummaryResponseBodyWarningSummarys) *DescribeCheckWarningSummaryResponseBody
	GetWarningSummarys() []*DescribeCheckWarningSummaryResponseBodyWarningSummarys
}

type DescribeCheckWarningSummaryResponseBody struct {
	// The number of check items returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 00BD7CE2-284A-4534-BD09-FB69836DD750
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of check items.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The statistics of check items.
	WarningSummarys []*DescribeCheckWarningSummaryResponseBodyWarningSummarys `json:"WarningSummarys,omitempty" xml:"WarningSummarys,omitempty" type:"Repeated"`
}

func (s DescribeCheckWarningSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCheckWarningSummaryResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningSummaryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCheckWarningSummaryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCheckWarningSummaryResponseBody) GetWarningSummarys() []*DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	return s.WarningSummarys
}

func (s *DescribeCheckWarningSummaryResponseBody) SetCount(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetCurrentPage(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetPageSize(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetRequestId(v string) *DescribeCheckWarningSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetTotalCount(v int32) *DescribeCheckWarningSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) SetWarningSummarys(v []*DescribeCheckWarningSummaryResponseBodyWarningSummarys) *DescribeCheckWarningSummaryResponseBody {
	s.WarningSummarys = v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCheckWarningSummaryResponseBodyWarningSummarys struct {
	// The number of check items.
	//
	// example:
	//
	// 10
	CheckCount *int32 `json:"CheckCount,omitempty" xml:"CheckCount,omitempty"`
	// Indicates whether the risk item can be exploited. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	CheckExploit *bool `json:"CheckExploit,omitempty" xml:"CheckExploit,omitempty"`
	// Indicates  whether the risk item is a container runtime risk item. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	ContainerRisk *bool `json:"ContainerRisk,omitempty" xml:"ContainerRisk,omitempty"`
	// Indicates whether the risk item is a database risk item. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	DatabaseRisk *bool `json:"DatabaseRisk,omitempty" xml:"DatabaseRisk,omitempty"`
	// The number of high-risk items.
	//
	// example:
	//
	// 1
	HighWarningCount *int32 `json:"HighWarningCount,omitempty" xml:"HighWarningCount,omitempty"`
	// The time when the last baseline check was performed.
	//
	// example:
	//
	// 2019-01-01 12:23:00
	LastFoundTime *string `json:"LastFoundTime,omitempty" xml:"LastFoundTime,omitempty"`
	// The risk level of the risk item. Valid values:
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
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The number of low-risk items.
	//
	// example:
	//
	// 3
	LowWarningCount *int32 `json:"LowWarningCount,omitempty" xml:"LowWarningCount,omitempty"`
	// The number of medium-risk items.
	//
	// example:
	//
	// 2
	MediumWarningCount *int32 `json:"MediumWarningCount,omitempty" xml:"MediumWarningCount,omitempty"`
	// The ID of the risk item.
	//
	// example:
	//
	// 118
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the risk item.
	//
	// example:
	//
	// Redis
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The level-2 type of the risk item.
	//
	// example:
	//
	// Redis
	SubTypeAlias *string `json:"SubTypeAlias,omitempty" xml:"SubTypeAlias,omitempty"`
	// The level-1 type of the check item. Examples: database, system, weak password, and middleware.
	//
	// example:
	//
	// databases
	TypeAlias *string `json:"TypeAlias,omitempty" xml:"TypeAlias,omitempty"`
	// The number of assets on which risk items are detected.
	//
	// example:
	//
	// 11
	WarningMachineCount *int32 `json:"WarningMachineCount,omitempty" xml:"WarningMachineCount,omitempty"`
}

func (s DescribeCheckWarningSummaryResponseBodyWarningSummarys) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningSummaryResponseBodyWarningSummarys) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetCheckCount() *int32 {
	return s.CheckCount
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetCheckExploit() *bool {
	return s.CheckExploit
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetContainerRisk() *bool {
	return s.ContainerRisk
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetDatabaseRisk() *bool {
	return s.DatabaseRisk
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetHighWarningCount() *int32 {
	return s.HighWarningCount
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetLastFoundTime() *string {
	return s.LastFoundTime
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetLevel() *string {
	return s.Level
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetLowWarningCount() *int32 {
	return s.LowWarningCount
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetMediumWarningCount() *int32 {
	return s.MediumWarningCount
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetRiskName() *string {
	return s.RiskName
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetSubTypeAlias() *string {
	return s.SubTypeAlias
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetTypeAlias() *string {
	return s.TypeAlias
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) GetWarningMachineCount() *int32 {
	return s.WarningMachineCount
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetCheckCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.CheckCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetCheckExploit(v bool) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.CheckExploit = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetContainerRisk(v bool) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.ContainerRisk = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetDatabaseRisk(v bool) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.DatabaseRisk = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetHighWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.HighWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLastFoundTime(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.LastFoundTime = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLevel(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.Level = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetLowWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.LowWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetMediumWarningCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.MediumWarningCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetRiskId(v int64) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetRiskName(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.RiskName = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetSubTypeAlias(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.SubTypeAlias = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetTypeAlias(v string) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.TypeAlias = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) SetWarningMachineCount(v int32) *DescribeCheckWarningSummaryResponseBodyWarningSummarys {
	s.WarningMachineCount = &v
	return s
}

func (s *DescribeCheckWarningSummaryResponseBodyWarningSummarys) Validate() error {
	return dara.Validate(s)
}
