// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListSupabaseProjectsResponseBodyItems) *ListSupabaseProjectsResponseBody
	GetItems() []*ListSupabaseProjectsResponseBodyItems
	SetMaxResults(v int32) *ListSupabaseProjectsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupabaseProjectsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListSupabaseProjectsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListSupabaseProjectsResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListSupabaseProjectsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *ListSupabaseProjectsResponseBody
	GetTotalRecordCount() *int32
}

type ListSupabaseProjectsResponseBody struct {
	Items []*ListSupabaseProjectsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListSupabaseProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectsResponseBody) GetItems() []*ListSupabaseProjectsResponseBodyItems {
	return s.Items
}

func (s *ListSupabaseProjectsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupabaseProjectsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseProjectsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSupabaseProjectsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListSupabaseProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupabaseProjectsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListSupabaseProjectsResponseBody) SetItems(v []*ListSupabaseProjectsResponseBodyItems) *ListSupabaseProjectsResponseBody {
	s.Items = v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetMaxResults(v int32) *ListSupabaseProjectsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetNextToken(v string) *ListSupabaseProjectsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetPageNumber(v int32) *ListSupabaseProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetPageRecordCount(v int32) *ListSupabaseProjectsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetRequestId(v string) *ListSupabaseProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) SetTotalRecordCount(v int32) *ListSupabaseProjectsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListSupabaseProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSupabaseProjectsResponseBodyItems struct {
	// example:
	//
	// 2021-10-09T04:54:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// xxpassword
	DashboardPassword *string `json:"DashboardPassword,omitempty" xml:"DashboardPassword,omitempty"`
	// example:
	//
	// null
	DashboardUserName *string `json:"DashboardUserName,omitempty" xml:"DashboardUserName,omitempty"`
	// example:
	//
	// PL0
	DiskPerformanceLevel *string `json:"DiskPerformanceLevel,omitempty" xml:"DiskPerformanceLevel,omitempty"`
	// example:
	//
	// gpdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 6.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 192.168.0.1
	PrivateConnectUrl *string `json:"PrivateConnectUrl,omitempty" xml:"PrivateConnectUrl,omitempty"`
	// example:
	//
	// sbp-12***
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// supabase_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1C1G
	ProjectSpec *string `json:"ProjectSpec,omitempty" xml:"ProjectSpec,omitempty"`
	// example:
	//
	// 10.154.11.10
	PublicConnectUrl *string `json:"PublicConnectUrl,omitempty" xml:"PublicConnectUrl,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 127.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	StorageSize *int64 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	// vSwitch ID。
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-bp19ame5m1r3oejns****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListSupabaseProjectsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectsResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSupabaseProjectsResponseBodyItems) GetDashboardPassword() *string {
	return s.DashboardPassword
}

func (s *ListSupabaseProjectsResponseBodyItems) GetDashboardUserName() *string {
	return s.DashboardUserName
}

func (s *ListSupabaseProjectsResponseBodyItems) GetDiskPerformanceLevel() *string {
	return s.DiskPerformanceLevel
}

func (s *ListSupabaseProjectsResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *ListSupabaseProjectsResponseBodyItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ListSupabaseProjectsResponseBodyItems) GetPayType() *string {
	return s.PayType
}

func (s *ListSupabaseProjectsResponseBodyItems) GetPrivateConnectUrl() *string {
	return s.PrivateConnectUrl
}

func (s *ListSupabaseProjectsResponseBodyItems) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListSupabaseProjectsResponseBodyItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListSupabaseProjectsResponseBodyItems) GetProjectSpec() *string {
	return s.ProjectSpec
}

func (s *ListSupabaseProjectsResponseBodyItems) GetPublicConnectUrl() *string {
	return s.PublicConnectUrl
}

func (s *ListSupabaseProjectsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSupabaseProjectsResponseBodyItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ListSupabaseProjectsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListSupabaseProjectsResponseBodyItems) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *ListSupabaseProjectsResponseBodyItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListSupabaseProjectsResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *ListSupabaseProjectsResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListSupabaseProjectsResponseBodyItems) SetCreateTime(v string) *ListSupabaseProjectsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetDashboardPassword(v string) *ListSupabaseProjectsResponseBodyItems {
	s.DashboardPassword = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetDashboardUserName(v string) *ListSupabaseProjectsResponseBodyItems {
	s.DashboardUserName = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetDiskPerformanceLevel(v string) *ListSupabaseProjectsResponseBodyItems {
	s.DiskPerformanceLevel = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetEngine(v string) *ListSupabaseProjectsResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetEngineVersion(v string) *ListSupabaseProjectsResponseBodyItems {
	s.EngineVersion = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetPayType(v string) *ListSupabaseProjectsResponseBodyItems {
	s.PayType = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetPrivateConnectUrl(v string) *ListSupabaseProjectsResponseBodyItems {
	s.PrivateConnectUrl = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetProjectId(v string) *ListSupabaseProjectsResponseBodyItems {
	s.ProjectId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetProjectName(v string) *ListSupabaseProjectsResponseBodyItems {
	s.ProjectName = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetProjectSpec(v string) *ListSupabaseProjectsResponseBodyItems {
	s.ProjectSpec = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetPublicConnectUrl(v string) *ListSupabaseProjectsResponseBodyItems {
	s.PublicConnectUrl = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetRegionId(v string) *ListSupabaseProjectsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetSecurityIPList(v string) *ListSupabaseProjectsResponseBodyItems {
	s.SecurityIPList = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetStatus(v string) *ListSupabaseProjectsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetStorageSize(v int64) *ListSupabaseProjectsResponseBodyItems {
	s.StorageSize = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetVSwitchId(v string) *ListSupabaseProjectsResponseBodyItems {
	s.VSwitchId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetVpcId(v string) *ListSupabaseProjectsResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) SetZoneId(v string) *ListSupabaseProjectsResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *ListSupabaseProjectsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
