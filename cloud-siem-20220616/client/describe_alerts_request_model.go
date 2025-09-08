// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *DescribeAlertsRequest
	GetAlertName() *string
	SetAlertTitle(v string) *DescribeAlertsRequest
	GetAlertTitle() *string
	SetAlertType(v string) *DescribeAlertsRequest
	GetAlertType() *string
	SetAlertUuid(v string) *DescribeAlertsRequest
	GetAlertUuid() *string
	SetAssetId(v string) *DescribeAlertsRequest
	GetAssetId() *string
	SetAssetName(v string) *DescribeAlertsRequest
	GetAssetName() *string
	SetCurrentPage(v int32) *DescribeAlertsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeAlertsRequest
	GetEndTime() *int64
	SetEntityId(v string) *DescribeAlertsRequest
	GetEntityId() *string
	SetEntityName(v string) *DescribeAlertsRequest
	GetEntityName() *string
	SetIsDefend(v string) *DescribeAlertsRequest
	GetIsDefend() *string
	SetLabelType(v string) *DescribeAlertsRequest
	GetLabelType() *string
	SetLevel(v []*string) *DescribeAlertsRequest
	GetLevel() []*string
	SetPageSize(v int32) *DescribeAlertsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlertsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertsRequest
	GetRoleType() *int32
	SetSource(v string) *DescribeAlertsRequest
	GetSource() *string
	SetStartTime(v int64) *DescribeAlertsRequest
	GetStartTime() *int64
	SetSubUserId(v string) *DescribeAlertsRequest
	GetSubUserId() *string
}

type DescribeAlertsRequest struct {
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Unusual Logon-login_common_account
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The page number. Pages start from page 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 异常连接-TFTP恶意扫描
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// Specifies whether an attack is defended. Valid values:
	//
	// 	- 0: detected.
	//
	// 	- 1: blocked.
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// example:
	//
	// 176555323***
	LabelType *string `json:"LabelType,omitempty" xml:"LabelType,omitempty"`
	// The risk level. The value is a JSON array. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// ["serious","suspicious","remind"]
	Level []*string `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the account that you switch from the management account.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The type of the view. Valid values:
	//
	// - 0: the current Alibaba Cloud account
	//
	// - 1: the global account
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertsRequest) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DescribeAlertsRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertsRequest) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeAlertsRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *DescribeAlertsRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeAlertsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertsRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DescribeAlertsRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeAlertsRequest) GetIsDefend() *string {
	return s.IsDefend
}

func (s *DescribeAlertsRequest) GetLabelType() *string {
	return s.LabelType
}

func (s *DescribeAlertsRequest) GetLevel() []*string {
	return s.Level
}

func (s *DescribeAlertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertsRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeAlertsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertsRequest) GetSubUserId() *string {
	return s.SubUserId
}

func (s *DescribeAlertsRequest) SetAlertName(v string) *DescribeAlertsRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertTitle(v string) *DescribeAlertsRequest {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertType(v string) *DescribeAlertsRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsRequest) SetAlertUuid(v string) *DescribeAlertsRequest {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsRequest) SetAssetId(v string) *DescribeAlertsRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeAlertsRequest) SetAssetName(v string) *DescribeAlertsRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeAlertsRequest) SetCurrentPage(v int32) *DescribeAlertsRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsRequest) SetEndTime(v int64) *DescribeAlertsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsRequest) SetEntityId(v string) *DescribeAlertsRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsRequest) SetEntityName(v string) *DescribeAlertsRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeAlertsRequest) SetIsDefend(v string) *DescribeAlertsRequest {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsRequest) SetLabelType(v string) *DescribeAlertsRequest {
	s.LabelType = &v
	return s
}

func (s *DescribeAlertsRequest) SetLevel(v []*string) *DescribeAlertsRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertsRequest) SetPageSize(v int32) *DescribeAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsRequest) SetRegionId(v string) *DescribeAlertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsRequest) SetRoleFor(v int64) *DescribeAlertsRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsRequest) SetRoleType(v int32) *DescribeAlertsRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsRequest) SetSource(v string) *DescribeAlertsRequest {
	s.Source = &v
	return s
}

func (s *DescribeAlertsRequest) SetStartTime(v int64) *DescribeAlertsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsRequest) SetSubUserId(v string) *DescribeAlertsRequest {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsRequest) Validate() error {
	return dara.Validate(s)
}
