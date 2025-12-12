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
	// The check method. Valid values:
	//
	// 	- **0**: proof of concept (POC) verification
	//
	// 	- **1**: version comparison
	//
	// example:
	//
	// 0
	CheckType *int32 `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether the vulnerability poses risks.\\
	//
	// If you do not specify this parameter, all vulnerabilities are queried regardless of whether the vulnerabilities pose risks. Valid values:
	//
	// 	- **y**: yes
	//
	// 	- **n**: no
	//
	// example:
	//
	// y
	RiskStatus *string `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The method that is used to detect the vulnerability.\\
	//
	// If you do not specify this parameter, all vulnerabilities are queried regardless of which method is used. Valid values:
	//
	// 	- **python**: The Version method is used. Security Center checks the software versions of your server to check whether disclosed vulnerabilities exist on your server.
	//
	// 	- **scan**: The Network Scan method is used. Security Center analyzes the access traffic to your server over the Internet to check whether vulnerabilities exist on your server.
	//
	// example:
	//
	// python
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The name of the urgent vulnerability.
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
