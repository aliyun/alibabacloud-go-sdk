// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRiskCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *DescribeRiskCheckResultRequest
	GetAssetType() *string
	SetCurrentPage(v int32) *DescribeRiskCheckResultRequest
	GetCurrentPage() *int32
	SetGroupId(v int64) *DescribeRiskCheckResultRequest
	GetGroupId() *int64
	SetItemIds(v []*string) *DescribeRiskCheckResultRequest
	GetItemIds() []*string
	SetLang(v string) *DescribeRiskCheckResultRequest
	GetLang() *string
	SetName(v string) *DescribeRiskCheckResultRequest
	GetName() *string
	SetPageSize(v int32) *DescribeRiskCheckResultRequest
	GetPageSize() *int32
	SetQueryFlag(v string) *DescribeRiskCheckResultRequest
	GetQueryFlag() *string
	SetResourceOwnerId(v int64) *DescribeRiskCheckResultRequest
	GetResourceOwnerId() *int64
	SetRiskLevel(v string) *DescribeRiskCheckResultRequest
	GetRiskLevel() *string
	SetSourceIp(v string) *DescribeRiskCheckResultRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeRiskCheckResultRequest
	GetStatus() *string
}

type DescribeRiskCheckResultRequest struct {
	// The cloud service whose configuration check results you want to query. For more information about the check items for the cloud service, see the check item table in the "Response parameters" section of this topic.
	//
	// example:
	//
	// RDS
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the check item that you want to query. Valid values:
	//
	// 	- **1**: identity authentication and permissions
	//
	// 	- **2**: network access control
	//
	// 	- **3**: log audit
	//
	// 	- **4**: data security
	//
	// 	- **5**: monitoring and alerting
	//
	// 	- **6**: basic security protection
	//
	// > If you do not specify this parameter, all types of check items are queried.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// An array that consists of the IDs of check items. For more information about the check item, see the check item table in the "Response parameters" section of this topic.
	//
	// example:
	//
	// 1
	ItemIds []*string `json:"ItemIds,omitempty" xml:"ItemIds,omitempty" type:"Repeated"`
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
	// The name of the check item. For more information about the check item, see the check item table in the "Response parameters" section of this topic.
	//
	// example:
	//
	// ALB_NetWorkAccessControl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Specifies whether the check item is supported by the edition of Security Center that you purchase. Valid values:
	//
	// 	- **enabled**: yes
	//
	// 	- **disabled**: no
	//
	// example:
	//
	// enabled
	QueryFlag       *string `json:"QueryFlag,omitempty" xml:"QueryFlag,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The risk level of the check item that you want to query. Valid values:
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
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 1.2.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The status of the check results. Valid values:
	//
	// 	- **pass**
	//
	// 	- **failed**
	//
	// 	- **running**
	//
	// 	- **waiting**
	//
	// 	- **ignored**
	//
	// 	- **falsePositive**
	//
	// example:
	//
	// pass
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRiskCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRiskCheckResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeRiskCheckResultRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeRiskCheckResultRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeRiskCheckResultRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeRiskCheckResultRequest) GetItemIds() []*string {
	return s.ItemIds
}

func (s *DescribeRiskCheckResultRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRiskCheckResultRequest) GetName() *string {
	return s.Name
}

func (s *DescribeRiskCheckResultRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRiskCheckResultRequest) GetQueryFlag() *string {
	return s.QueryFlag
}

func (s *DescribeRiskCheckResultRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRiskCheckResultRequest) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeRiskCheckResultRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeRiskCheckResultRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeRiskCheckResultRequest) SetAssetType(v string) *DescribeRiskCheckResultRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetCurrentPage(v int32) *DescribeRiskCheckResultRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetGroupId(v int64) *DescribeRiskCheckResultRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetItemIds(v []*string) *DescribeRiskCheckResultRequest {
	s.ItemIds = v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetLang(v string) *DescribeRiskCheckResultRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetName(v string) *DescribeRiskCheckResultRequest {
	s.Name = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetPageSize(v int32) *DescribeRiskCheckResultRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetQueryFlag(v string) *DescribeRiskCheckResultRequest {
	s.QueryFlag = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetResourceOwnerId(v int64) *DescribeRiskCheckResultRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetRiskLevel(v string) *DescribeRiskCheckResultRequest {
	s.RiskLevel = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetSourceIp(v string) *DescribeRiskCheckResultRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) SetStatus(v string) *DescribeRiskCheckResultRequest {
	s.Status = &v
	return s
}

func (s *DescribeRiskCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
