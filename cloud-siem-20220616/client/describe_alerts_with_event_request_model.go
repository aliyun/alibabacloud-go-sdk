// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertName(v string) *DescribeAlertsWithEventRequest
	GetAlertName() *string
	SetAlertTitle(v string) *DescribeAlertsWithEventRequest
	GetAlertTitle() *string
	SetAlertType(v string) *DescribeAlertsWithEventRequest
	GetAlertType() *string
	SetAssetId(v string) *DescribeAlertsWithEventRequest
	GetAssetId() *string
	SetAssetName(v string) *DescribeAlertsWithEventRequest
	GetAssetName() *string
	SetCurrentPage(v int32) *DescribeAlertsWithEventRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *DescribeAlertsWithEventRequest
	GetEndTime() *int64
	SetEntityId(v string) *DescribeAlertsWithEventRequest
	GetEntityId() *string
	SetEntityName(v string) *DescribeAlertsWithEventRequest
	GetEntityName() *string
	SetIncidentUuid(v string) *DescribeAlertsWithEventRequest
	GetIncidentUuid() *string
	SetIsDefend(v string) *DescribeAlertsWithEventRequest
	GetIsDefend() *string
	SetLevel(v []*string) *DescribeAlertsWithEventRequest
	GetLevel() []*string
	SetPageSize(v int32) *DescribeAlertsWithEventRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAlertsWithEventRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DescribeAlertsWithEventRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *DescribeAlertsWithEventRequest
	GetRoleType() *int32
	SetSource(v string) *DescribeAlertsWithEventRequest
	GetSource() *string
	SetStartTime(v int64) *DescribeAlertsWithEventRequest
	GetStartTime() *int64
	SetSubUserId(v int64) *DescribeAlertsWithEventRequest
	GetSubUserId() *int64
}

type DescribeAlertsWithEventRequest struct {
	// The alert name.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The alert title.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The alert type.
	//
	// example:
	//
	// Scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The asset ID.
	//
	// example:
	//
	// 09414e9ebaa9c19b84d851abb91d****
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The asset name.
	//
	// example:
	//
	// hostname-****
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The current page number. The value must be greater than or equal to 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end time of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// 21034e803f492b926cea9e5beab4****
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The entity name.
	//
	// example:
	//
	// launch-advisor-*****
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// The incident ID.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Specifies whether the alert has been defended. Valid values:
	//
	// - 0: detected
	//
	// - 1: blocked
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The threat level in JSON array format. Valid values:
	//
	// - serious: High.
	//
	// - suspicious: Medium.
	//
	// - remind: Low.
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
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets reside in Chinese mainland or China (Hong Kong).
	//
	// - ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The member account ID of the resource directory.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The alert data source.
	//
	// example:
	//
	// sas
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The start time of the query. Unit: milliseconds.
	//
	// example:
	//
	// 1577808000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the account associated with the alert.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeAlertsWithEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertsWithEventRequest) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DescribeAlertsWithEventRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertsWithEventRequest) GetAssetId() *string {
	return s.AssetId
}

func (s *DescribeAlertsWithEventRequest) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeAlertsWithEventRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsWithEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeAlertsWithEventRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DescribeAlertsWithEventRequest) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeAlertsWithEventRequest) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertsWithEventRequest) GetIsDefend() *string {
	return s.IsDefend
}

func (s *DescribeAlertsWithEventRequest) GetLevel() []*string {
	return s.Level
}

func (s *DescribeAlertsWithEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsWithEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertsWithEventRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DescribeAlertsWithEventRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *DescribeAlertsWithEventRequest) GetSource() *string {
	return s.Source
}

func (s *DescribeAlertsWithEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeAlertsWithEventRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeAlertsWithEventRequest) SetAlertName(v string) *DescribeAlertsWithEventRequest {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAlertTitle(v string) *DescribeAlertsWithEventRequest {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAlertType(v string) *DescribeAlertsWithEventRequest {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAssetId(v string) *DescribeAlertsWithEventRequest {
	s.AssetId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetAssetName(v string) *DescribeAlertsWithEventRequest {
	s.AssetName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetCurrentPage(v int32) *DescribeAlertsWithEventRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEndTime(v int64) *DescribeAlertsWithEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEntityId(v string) *DescribeAlertsWithEventRequest {
	s.EntityId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetEntityName(v string) *DescribeAlertsWithEventRequest {
	s.EntityName = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetIncidentUuid(v string) *DescribeAlertsWithEventRequest {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetIsDefend(v string) *DescribeAlertsWithEventRequest {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetLevel(v []*string) *DescribeAlertsWithEventRequest {
	s.Level = v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetPageSize(v int32) *DescribeAlertsWithEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRegionId(v string) *DescribeAlertsWithEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRoleFor(v int64) *DescribeAlertsWithEventRequest {
	s.RoleFor = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetRoleType(v int32) *DescribeAlertsWithEventRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetSource(v string) *DescribeAlertsWithEventRequest {
	s.Source = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetStartTime(v int64) *DescribeAlertsWithEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) SetSubUserId(v int64) *DescribeAlertsWithEventRequest {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsWithEventRequest) Validate() error {
	return dara.Validate(s)
}
