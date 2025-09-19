// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *DescribeCheckWarningsRequest
	GetCheckId() *int64
	SetCheckType(v string) *DescribeCheckWarningsRequest
	GetCheckType() *string
	SetContainerName(v string) *DescribeCheckWarningsRequest
	GetContainerName() *string
	SetCurrentPage(v int32) *DescribeCheckWarningsRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribeCheckWarningsRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCheckWarningsRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningsRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskId(v int64) *DescribeCheckWarningsRequest
	GetRiskId() *int64
	SetRiskStatus(v int32) *DescribeCheckWarningsRequest
	GetRiskStatus() *int32
	SetSourceIp(v string) *DescribeCheckWarningsRequest
	GetSourceIp() *string
	SetUuid(v string) *DescribeCheckWarningsRequest
	GetUuid() *string
}

type DescribeCheckWarningsRequest struct {
	// The ID of the check item.
	//
	// example:
	//
	// 2546
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The type of the check item. Valid values:
	//
	// 	- **hc.check.type.identity_auth**: identity authentication
	//
	// 	- **hc.check.type.access_control**: access control
	//
	// 	- **hc.check.type.network_service**: network and service
	//
	// 	- **hc.check.type.service_conf**: service configuration
	//
	// 	- **hc.check.type.file_rights**: file permission
	//
	// 	- **hc.check.type.security_audit**: security audit
	//
	// 	- **hc.check.type.attack_defense**: intrusion prevention
	//
	// 	- **hc.check.type.others**: others
	//
	// example:
	//
	// hc.check.type.attack_defense
	CheckType *string `json:"CheckType,omitempty" xml:"CheckType,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// /redis
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
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
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  To obtain the Alibaba Cloud account ID, call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The ID of the risk item. This parameter is required.
	//
	// >  To query the information about the risk items and check items of a server, you must specify the IDs of the risk items. You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of risk items.
	//
	// example:
	//
	// 10354
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **2**: verifying
	//
	// 	- **3**: passed
	//
	// 	- **5**: expired
	//
	// 	- **6**: ignored
	//
	// example:
	//
	// 1
	RiskStatus *int32 `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The UUID of the server on which the baseline check is performed.
	//
	// > To query specified risk items and the check items of a specified server, you must provide the ID of the server on which the baseline check is performed. You can call the [DescribeWarningMachines](~~DescribeWarningMachines~~) operation to query the IDs of servers.
	//
	// This parameter is required.
	//
	// example:
	//
	// d42f938c-d962-48a0-90f9-05****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeCheckWarningsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningsRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckWarningsRequest) GetCheckType() *string {
	return s.CheckType
}

func (s *DescribeCheckWarningsRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeCheckWarningsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCheckWarningsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCheckWarningsRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeCheckWarningsRequest) GetRiskStatus() *int32 {
	return s.RiskStatus
}

func (s *DescribeCheckWarningsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCheckWarningsRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeCheckWarningsRequest) SetCheckId(v int64) *DescribeCheckWarningsRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetCheckType(v string) *DescribeCheckWarningsRequest {
	s.CheckType = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetContainerName(v string) *DescribeCheckWarningsRequest {
	s.ContainerName = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetCurrentPage(v int32) *DescribeCheckWarningsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetLang(v string) *DescribeCheckWarningsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetPageSize(v int32) *DescribeCheckWarningsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetRiskId(v int64) *DescribeCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetRiskStatus(v int32) *DescribeCheckWarningsRequest {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetSourceIp(v string) *DescribeCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningsRequest) SetUuid(v string) *DescribeCheckWarningsRequest {
	s.Uuid = &v
	return s
}

func (s *DescribeCheckWarningsRequest) Validate() error {
	return dara.Validate(s)
}
