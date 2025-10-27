// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageGroupedVulListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeImageGroupedVulListResponseBody
	GetCurrentPage() *int32
	SetGroupedVulItems(v []*DescribeImageGroupedVulListResponseBodyGroupedVulItems) *DescribeImageGroupedVulListResponseBody
	GetGroupedVulItems() []*DescribeImageGroupedVulListResponseBodyGroupedVulItems
	SetPageSize(v int32) *DescribeImageGroupedVulListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeImageGroupedVulListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageGroupedVulListResponseBody
	GetTotalCount() *int32
}

type DescribeImageGroupedVulListResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the image vulnerabilities.
	GroupedVulItems []*DescribeImageGroupedVulListResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 5E244439-UJND-8BF7-26F36E21B9AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of image system vulnerabilities.
	//
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageGroupedVulListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageGroupedVulListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageGroupedVulListResponseBody) GetGroupedVulItems() []*DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	return s.GroupedVulItems
}

func (s *DescribeImageGroupedVulListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageGroupedVulListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageGroupedVulListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageGroupedVulListResponseBody) SetCurrentPage(v int32) *DescribeImageGroupedVulListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetGroupedVulItems(v []*DescribeImageGroupedVulListResponseBodyGroupedVulItems) *DescribeImageGroupedVulListResponseBody {
	s.GroupedVulItems = v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetPageSize(v int32) *DescribeImageGroupedVulListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetRequestId(v string) *DescribeImageGroupedVulListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) SetTotalCount(v int32) *DescribeImageGroupedVulListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBody) Validate() error {
	if s.GroupedVulItems != nil {
		for _, item := range s.GroupedVulItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageGroupedVulListResponseBodyGroupedVulItems struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2017:3075-Important: wget security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The number of vulnerabilities that have the high priority.
	//
	// example:
	//
	// 26
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// Indicates whether the vulnerability can be fixed in the Security Center console. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	CanFix *string `json:"CanFix,omitempty" xml:"CanFix,omitempty"`
	// The timestamp when the first scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1611201274000
	GmtLast *int64 `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	// The timestamp when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1611201274000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// The number of vulnerabilities that have the medium priority.
	//
	// example:
	//
	// 26
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// debian:9:CVE-2019-3858
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of vulnerabilities that have the low priority.
	//
	// example:
	//
	// 29
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// The tag of this vulnerability. Valid values:
	//
	// 	- **AI**: AI-related components.
	//
	// example:
	//
	// AI
	RuleTag *string `json:"RuleTag,omitempty" xml:"RuleTag,omitempty"`
	// The status of the vulnerability. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: handled
	//
	// 	- **2**: verifying
	//
	// 	- **3**: added to the whitelist
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the vulnerability. Valid values:
	//
	// 	- Restart required
	//
	// 	- Remote exploitation
	//
	// 	- Exploit exists
	//
	// 	- Exploitable
	//
	// 	- Privilege escalation
	//
	// 	- Code execution
	//
	// example:
	//
	// EXP exists
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: image system vulnerability
	//
	// 	- **sca**: image application vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageGroupedVulListResponseBodyGroupedVulItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageGroupedVulListResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetAsapCount() *int32 {
	return s.AsapCount
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetCanFix() *string {
	return s.CanFix
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetGmtLast() *int64 {
	return s.GmtLast
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetLaterCount() *int32 {
	return s.LaterCount
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetName() *string {
	return s.Name
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetNntfCount() *int32 {
	return s.NntfCount
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetRuleTag() *string {
	return s.RuleTag
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetTags() *string {
	return s.Tags
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) GetType() *string {
	return s.Type
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetCanFix(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.CanFix = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetGmtLast(v int64) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetLastScanTime(v int64) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.LastScanTime = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetName(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetRuleTag(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.RuleTag = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetStatus(v int32) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Status = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetTags(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) SetType(v string) *DescribeImageGroupedVulListResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeImageGroupedVulListResponseBodyGroupedVulItems) Validate() error {
	return dara.Validate(s)
}
