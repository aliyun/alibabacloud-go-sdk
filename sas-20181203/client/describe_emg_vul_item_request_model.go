// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEmgVulItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *DescribeEmgVulItemRequest
	GetCheckType() *int32
	SetCurrentPage(v int32) *DescribeEmgVulItemRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeEmgVulItemRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeEmgVulItemRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeEmgVulItemRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskStatus(v string) *DescribeEmgVulItemRequest
	GetRiskStatus() *string
	SetScanType(v string) *DescribeEmgVulItemRequest
	GetScanType() *string
	SetVulName(v string) *DescribeEmgVulItemRequest
	GetVulName() *string
}

type DescribeEmgVulItemRequest struct {
	// The check type. Valid values:
	//
	// - **0**: POC verification
	//
	// - **1**: version comparison.
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The page number of the first page to return. Default value: **1**, which indicates that query results are displayed starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page in a paged query. Default value: **10**, which indicates that 10 emergency vulnerability entries are displayed per page. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the member accounts in the resource directory (Alibaba Cloud account).
	//
	// > Invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The risk status of the vulnerabilities to query. If this parameter is not specified, vulnerabilities of all risk statuses are returned, including those with risks and those without risks. Valid values:
	//
	// - **y**: at risk
	//
	// - **n**: not at risk.
	//
	// example:
	//
	// y
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The detection method of the vulnerabilities to query. If this parameter is not specified, vulnerabilities detected by all methods are returned by default, including version detection and network scanning. Valid values:
	//
	// - **python**: version detection (server software version detection). Detects whether your server has disclosed software vulnerabilities.
	//
	// - **scan**: network scanning (network traffic detection). Detects whether your public assets (Internet-accessible servers) have vulnerabilities.
	//
	// example:
	//
	// python
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The name of the emergency vulnerability to query.
	//
	// example:
	//
	// Changjietong T + SetupAccount/Upload.aspx file Upload vulnerability (CNVD-2022-60632)
	VulName *string `json:"VulName,omitempty" xml:"VulName,omitempty"`
}

func (s DescribeEmgVulItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEmgVulItemRequest) GoString() string {
	return s.String()
}

func (s *DescribeEmgVulItemRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *DescribeEmgVulItemRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeEmgVulItemRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEmgVulItemRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEmgVulItemRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeEmgVulItemRequest) GetRiskStatus() *string {
	return s.RiskStatus
}

func (s *DescribeEmgVulItemRequest) GetScanType() *string {
	return s.ScanType
}

func (s *DescribeEmgVulItemRequest) GetVulName() *string {
	return s.VulName
}

func (s *DescribeEmgVulItemRequest) SetCheckType(v int32) *DescribeEmgVulItemRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetCurrentPage(v int32) *DescribeEmgVulItemRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetLang(v string) *DescribeEmgVulItemRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetPageSize(v int32) *DescribeEmgVulItemRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetResourceDirectoryAccountId(v int64) *DescribeEmgVulItemRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetRiskStatus(v string) *DescribeEmgVulItemRequest {
	s.RiskStatus = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetScanType(v string) *DescribeEmgVulItemRequest {
	s.ScanType = &v
	return s
}

func (s *DescribeEmgVulItemRequest) SetVulName(v string) *DescribeEmgVulItemRequest {
	s.VulName = &v
	return s
}

func (s *DescribeEmgVulItemRequest) Validate() error {
	return dara.Validate(s)
}
