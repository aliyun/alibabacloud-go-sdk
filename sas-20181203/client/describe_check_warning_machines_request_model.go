// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningMachinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *DescribeCheckWarningMachinesRequest
	GetCheckId() *int64
	SetCurrentPage(v int32) *DescribeCheckWarningMachinesRequest
	GetCurrentPage() *int32
	SetFilterUuid(v string) *DescribeCheckWarningMachinesRequest
	GetFilterUuid() *string
	SetInstanceId(v string) *DescribeCheckWarningMachinesRequest
	GetInstanceId() *string
	SetLang(v string) *DescribeCheckWarningMachinesRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCheckWarningMachinesRequest
	GetPageSize() *int32
	SetRemark(v string) *DescribeCheckWarningMachinesRequest
	GetRemark() *string
	SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningMachinesRequest
	GetResourceDirectoryAccountId() *int64
	SetRiskId(v int64) *DescribeCheckWarningMachinesRequest
	GetRiskId() *int64
	SetStatus(v int32) *DescribeCheckWarningMachinesRequest
	GetStatus() *int32
}

type DescribeCheckWarningMachinesRequest struct {
	// The ID of the check item.
	//
	// > You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 58
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The instance ID of the asset that you don\\"t want to query.
	//
	// example:
	//
	// 2f64e1a0f9316c48*******
	FilterUuid *string `json:"FilterUuid,omitempty" xml:"FilterUuid,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// cri-rv4nvbv8iju4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the server that you want to query. The value can be the name or the public IP address of the server.
	//
	// example:
	//
	// 1.2.XX.XX
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the IDs.
	//
	// example:
	//
	// 1232428423234****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The ID of the baseline.
	//
	// > You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of baselines.
	//
	// example:
	//
	// 43
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The risk status of the check item. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **3**: passed
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCheckWarningMachinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningMachinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningMachinesRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckWarningMachinesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningMachinesRequest) GetFilterUuid() *string {
	return s.FilterUuid
}

func (s *DescribeCheckWarningMachinesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCheckWarningMachinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCheckWarningMachinesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningMachinesRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCheckWarningMachinesRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeCheckWarningMachinesRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeCheckWarningMachinesRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCheckWarningMachinesRequest) SetCheckId(v int64) *DescribeCheckWarningMachinesRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetCurrentPage(v int32) *DescribeCheckWarningMachinesRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetFilterUuid(v string) *DescribeCheckWarningMachinesRequest {
	s.FilterUuid = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetInstanceId(v string) *DescribeCheckWarningMachinesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetLang(v string) *DescribeCheckWarningMachinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetPageSize(v int32) *DescribeCheckWarningMachinesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetRemark(v string) *DescribeCheckWarningMachinesRequest {
	s.Remark = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetResourceDirectoryAccountId(v int64) *DescribeCheckWarningMachinesRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetRiskId(v int64) *DescribeCheckWarningMachinesRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) SetStatus(v int32) *DescribeCheckWarningMachinesRequest {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningMachinesRequest) Validate() error {
	return dara.Validate(s)
}
