// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmgVulItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeEmgVulItemResponseBody
	GetCurrentPage() *int32
	SetGroupedVulItems(v []*DescribeEmgVulItemResponseBodyGroupedVulItems) *DescribeEmgVulItemResponseBody
	GetGroupedVulItems() []*DescribeEmgVulItemResponseBodyGroupedVulItems
	SetPageSize(v int32) *DescribeEmgVulItemResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeEmgVulItemResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeEmgVulItemResponseBody
	GetTotalCount() *int32
}

type DescribeEmgVulItemResponseBody struct {
	// The page number of the returned page. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of the urgent vulnerabilities returned.
	GroupedVulItems []*DescribeEmgVulItemResponseBodyGroupedVulItems `json:"GroupedVulItems,omitempty" xml:"GroupedVulItems,omitempty" type:"Repeated"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC1868ED-A0E1-4D1C-BF7E-10DC0C34B3C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the urgent vulnerabilities returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEmgVulItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgVulItemResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulItemResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEmgVulItemResponseBody) GetGroupedVulItems() []*DescribeEmgVulItemResponseBodyGroupedVulItems {
	return s.GroupedVulItems
}

func (s *DescribeEmgVulItemResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEmgVulItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEmgVulItemResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeEmgVulItemResponseBody) SetCurrentPage(v int32) *DescribeEmgVulItemResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEmgVulItemResponseBody) SetGroupedVulItems(v []*DescribeEmgVulItemResponseBodyGroupedVulItems) *DescribeEmgVulItemResponseBody {
	s.GroupedVulItems = v
	return s
}

func (s *DescribeEmgVulItemResponseBody) SetPageSize(v int32) *DescribeEmgVulItemResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeEmgVulItemResponseBody) SetRequestId(v string) *DescribeEmgVulItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEmgVulItemResponseBody) SetTotalCount(v int32) *DescribeEmgVulItemResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeEmgVulItemResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEmgVulItemResponseBodyGroupedVulItems struct {
	// The name of the urgent vulnerability.
	//
	// example:
	//
	// Changjietong T + SetupAccount/Upload.aspx file Upload vulnerability (CNVD-2022-60632)
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The check method.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The introduction to the vulnerability.
	//
	// example:
	//
	// Chanjet T-Plus is an Internet business management software. There is an unauthorized access vulnerability in one of its interfaces disclosed on the Internet. Attackers can construct malicious requests to upload malicious files to execute arbitrary code and control the server.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the urgent vulnerability was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1619286031000
	GmtLastCheck *int64 `json:"GmtLastCheck,omitempty" xml:"GmtLastCheck,omitempty"`
	// The timestamp when the urgent vulnerability was last disclosed. Unit: milliseconds.
	//
	// example:
	//
	// 1618887687000
	GmtPublish *int64 `json:"GmtPublish,omitempty" xml:"GmtPublish,omitempty"`
	// The name of the detection rule.
	//
	// example:
	//
	// scan:AVD-2021-179344
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of unhandled urgent vulnerabilities.
	//
	// example:
	//
	// 0
	PendingCount *int32 `json:"PendingCount,omitempty" xml:"PendingCount,omitempty"`
	// The progress of the urgent vulnerability detection task. Valid values: 0 to 100.
	//
	// >  This parameter is returned only when an urgent vulnerability is being detected.
	//
	// example:
	//
	// 50
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Indicates whether the application protection feature is supported. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// >  If this parameter is not returned, the application protection is not supported.
	//
	// example:
	//
	// 1
	RaspDefend *int32 `json:"RaspDefend,omitempty" xml:"RaspDefend,omitempty"`
	// The detection status of the urgent vulnerability. Valid values:
	//
	// 	- **10**: The urgent vulnerability is not detected.
	//
	// 	- **20**: The urgent vulnerability is being detected.
	//
	// 	- **30**: The urgent vulnerability detection is complete.
	//
	// example:
	//
	// 30
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The method that is used to detect the urgent vulnerability. Valid values:
	//
	// 	- **python**: The Version method is used. Security Center checks the software versions of your server to check whether disclosed vulnerabilities exist on your server.
	//
	// 	- **scan**: The Network Scan method is used. Security Center analyzes the access traffic to your server over the Internet to check whether vulnerabilities exist on your server.
	//
	// example:
	//
	// scan
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeEmgVulItemResponseBodyGroupedVulItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgVulItemResponseBodyGroupedVulItems) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetGmtLastCheck() *int64 {
	return s.GmtLastCheck
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetGmtPublish() *int64 {
	return s.GmtPublish
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetName() *string {
	return s.Name
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetPendingCount() *int32 {
	return s.PendingCount
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetRaspDefend() *int32 {
	return s.RaspDefend
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) GetType() *string {
	return s.Type
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetAliasName(v string) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.AliasName = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetCheckType(v int32) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.CheckType = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetDescription(v string) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.Description = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetGmtLastCheck(v int64) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.GmtLastCheck = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetGmtPublish(v int64) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.GmtPublish = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetName(v string) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.Name = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetPendingCount(v int32) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.PendingCount = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetProgress(v int32) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.Progress = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetRaspDefend(v int32) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.RaspDefend = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetStatus(v int32) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.Status = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) SetType(v string) *DescribeEmgVulItemResponseBodyGroupedVulItems {
	s.Type = &v
	return s
}

func (s *DescribeEmgVulItemResponseBodyGroupedVulItems) Validate() error {
	return dara.Validate(s)
}
