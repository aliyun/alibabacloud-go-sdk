// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupedVulResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeGroupedVulResponseBody
	GetCurrentPage() *int32
	SetGroupedVulItems(v []*DescribeGroupedVulResponseBodyGroupedVulItems) *DescribeGroupedVulResponseBody
	GetGroupedVulItems() []*DescribeGroupedVulResponseBodyGroupedVulItems
	SetPageSize(v int32) *DescribeGroupedVulResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGroupedVulResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeGroupedVulResponseBody
	GetTotalCount() *int32
}

type DescribeGroupedVulResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The information about the vulnerability.
	GroupedVulItems []*DescribeGroupedVulResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9BFA6D78-07EA-5C0A-9358-E4434573507B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeGroupedVulResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedVulResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeGroupedVulResponseBody) GetGroupedVulItems() []*DescribeGroupedVulResponseBodyGroupedVulItems {
	return s.GroupedVulItems
}

func (s *DescribeGroupedVulResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupedVulResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupedVulResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeGroupedVulResponseBody) SetCurrentPage(v int32) *DescribeGroupedVulResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetGroupedVulItems(v []*DescribeGroupedVulResponseBodyGroupedVulItems) *DescribeGroupedVulResponseBody {
	s.GroupedVulItems = v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetPageSize(v int32) *DescribeGroupedVulResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetRequestId(v string) *DescribeGroupedVulResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) SetTotalCount(v int32) *DescribeGroupedVulResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBody) Validate() error {
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

type DescribeGroupedVulResponseBodyGroupedVulItems struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2017:0184-Important: mysql security update
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The number of vulnerabilities that have the **high*	- priority.
	//
	// example:
	//
	// 0
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// The timestamp when the vulnerability was first detected. Unit: milliseconds.
	//
	// example:
	//
	// 1639371046000
	GmtFirst *int64 `json:"GmtFirst,omitempty" xml:"GmtFirst,omitempty"`
	// The timestamp when the vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1639371446000
	GmtLast *int64 `json:"GmtLast,omitempty" xml:"GmtLast,omitempty"`
	// The number of handled vulnerabilities.
	//
	// example:
	//
	// 0
	HandledCount *int32 `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	// The language type associated with the vulnerability. Valid values:
	//
	// 	- **java**
	//
	// 	- **php**
	//
	// >  This parameter is valid only for a vulnerability of the sca type.
	//
	// example:
	//
	// java
	LanguageType *string `json:"LanguageType,omitempty" xml:"LanguageType,omitempty"`
	// The number of vulnerabilities that have the **medium*	- priority.
	//
	// example:
	//
	// 0
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20170184
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of vulnerabilities that have the **low*	- priority.
	//
	// example:
	//
	// 59
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// Indicates whether the application protection feature is supported. Valid values:
	//
	// 	- **0**: not supported
	//
	// 	- **1**: supported
	//
	// >  If this parameter is not returned, the application protection feature is not supported.
	//
	// example:
	//
	// 1
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The IDs of the common vulnerabilities and exposures (CVEs) that are related to the vulnerability.
	//
	// example:
	//
	// CVE-2023-24881,CVE-2023-24898
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// The tag of the vulnerability. Valid values:
	//
	// 	- **Restart required**
	//
	// 	- **Remote utilization**
	//
	// 	- **EXP exists**
	//
	// 	- **Available**
	//
	// 	- **Elevation of Privilege**
	//
	// 	- **Code Execution**
	//
	// example:
	//
	// Code Execution
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The total number of fixed vulnerabilities.
	//
	// example:
	//
	// 0
	TotalFixCount *int64 `json:"TotalFixCount,omitempty" xml:"TotalFixCount,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// 	- **sca**: vulnerability that is detected by software component analysis
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupedVulResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetAsapCount() *int32 {
	return s.AsapCount
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetGmtFirst() *int64 {
	return s.GmtFirst
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetGmtLast() *int64 {
	return s.GmtLast
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetHandledCount() *int32 {
	return s.HandledCount
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetLanguageType() *string {
	return s.LanguageType
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetLaterCount() *int32 {
	return s.LaterCount
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetName() *string {
	return s.Name
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetNntfCount() *int32 {
	return s.NntfCount
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetRelated() *string {
	return s.Related
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetTags() *string {
	return s.Tags
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetTotalFixCount() *int64 {
	return s.TotalFixCount
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) GetType() *string {
	return s.Type
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetAsapCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.AsapCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetGmtFirst(v int64) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.GmtFirst = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetGmtLast(v int64) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.GmtLast = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetHandledCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.HandledCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetLanguageType(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.LanguageType = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetLaterCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.LaterCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetName(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetNntfCount(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.NntfCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetRaspDefend(v int32) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.RaspDefend = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetRelated(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Related = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetTags(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Tags = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetTotalFixCount(v int64) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.TotalFixCount = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) SetType(v string) *DescribeGroupedVulResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeGroupedVulResponseBodyGroupedVulItems) Validate() error {
	return dara.Validate(s)
}
