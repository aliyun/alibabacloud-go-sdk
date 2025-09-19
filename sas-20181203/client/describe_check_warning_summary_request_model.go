// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeCheckWarningSummaryRequest
	GetClusterId() *string
	SetContainerFieldName(v string) *DescribeCheckWarningSummaryRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeCheckWarningSummaryRequest
	GetContainerFieldValue() *string
	SetCurrentPage(v int32) *DescribeCheckWarningSummaryRequest
	GetCurrentPage() *int32
	SetGroupId(v int64) *DescribeCheckWarningSummaryRequest
	GetGroupId() *int64
	SetLang(v string) *DescribeCheckWarningSummaryRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeCheckWarningSummaryRequest
	GetPageSize() *int32
	SetRiskName(v string) *DescribeCheckWarningSummaryRequest
	GetRiskName() *string
	SetRiskStatus(v int32) *DescribeCheckWarningSummaryRequest
	GetRiskStatus() *int32
	SetSourceIp(v string) *DescribeCheckWarningSummaryRequest
	GetSourceIp() *string
	SetStatus(v string) *DescribeCheckWarningSummaryRequest
	GetStatus() *string
	SetStrategyId(v int64) *DescribeCheckWarningSummaryRequest
	GetStrategyId() *int64
	SetTargetType(v string) *DescribeCheckWarningSummaryRequest
	GetTargetType() *string
	SetTypeName(v string) *DescribeCheckWarningSummaryRequest
	GetTypeName() *string
	SetUuids(v string) *DescribeCheckWarningSummaryRequest
	GetUuids() *string
}

type DescribeCheckWarningSummaryRequest struct {
	// The ID of the container cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// example:
	//
	// c80dae73bd1be442699766b14ffd0****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the container field. Valid values:
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// 	- **image**: the name of the image
	//
	// 	- **imageId**: the ID of the image
	//
	// 	- **namespace**: the namespace
	//
	// example:
	//
	// namespace
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container field.
	//
	// example:
	//
	// c819391d2d520485fa3e81e2dc2ea****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the asset group.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of asset groups.
	//
	// example:
	//
	// 123
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The language of the content within the request and the response. Valid values:
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
	// The name of the risk item.
	//
	// example:
	//
	// Redis
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The status of the baseline check. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **3**: passed
	//
	// example:
	//
	// 1
	RiskStatus *int32 `json:"RiskStatus,omitempty" xml:"RiskStatus,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 219.133.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
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
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the baseline check policy.
	//
	// example:
	//
	// 1
	StrategyId *int64 `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The type of the query condition. Valid values:
	//
	// 	- **uuid**: the ID of an asset
	//
	// example:
	//
	// uuid
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The level-1 type of check items.
	//
	// >  You can call the [DescribeRiskType](~~DescribeRiskType~~) operation to query the level-1 types of check items.
	//
	// example:
	//
	// database
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The UUID of the asset.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of assets.
	//
	// example:
	//
	// f03259d8-1e81-4fae-bcbb-275fb5****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeCheckWarningSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCheckWarningSummaryRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeCheckWarningSummaryRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeCheckWarningSummaryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCheckWarningSummaryRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeCheckWarningSummaryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCheckWarningSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCheckWarningSummaryRequest) GetRiskName() *string {
	return s.RiskName
}

func (s *DescribeCheckWarningSummaryRequest) GetRiskStatus() *int32 {
	return s.RiskStatus
}

func (s *DescribeCheckWarningSummaryRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeCheckWarningSummaryRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCheckWarningSummaryRequest) GetStrategyId() *int64 {
	return s.StrategyId
}

func (s *DescribeCheckWarningSummaryRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *DescribeCheckWarningSummaryRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeCheckWarningSummaryRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeCheckWarningSummaryRequest) SetClusterId(v string) *DescribeCheckWarningSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetContainerFieldName(v string) *DescribeCheckWarningSummaryRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetContainerFieldValue(v string) *DescribeCheckWarningSummaryRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetCurrentPage(v int32) *DescribeCheckWarningSummaryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetGroupId(v int64) *DescribeCheckWarningSummaryRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetLang(v string) *DescribeCheckWarningSummaryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetPageSize(v int32) *DescribeCheckWarningSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetRiskName(v string) *DescribeCheckWarningSummaryRequest {
	s.RiskName = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetRiskStatus(v int32) *DescribeCheckWarningSummaryRequest {
	s.RiskStatus = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetSourceIp(v string) *DescribeCheckWarningSummaryRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetStatus(v string) *DescribeCheckWarningSummaryRequest {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetStrategyId(v int64) *DescribeCheckWarningSummaryRequest {
	s.StrategyId = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetTargetType(v string) *DescribeCheckWarningSummaryRequest {
	s.TargetType = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetTypeName(v string) *DescribeCheckWarningSummaryRequest {
	s.TypeName = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) SetUuids(v string) *DescribeCheckWarningSummaryRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeCheckWarningSummaryRequest) Validate() error {
	return dara.Validate(s)
}
