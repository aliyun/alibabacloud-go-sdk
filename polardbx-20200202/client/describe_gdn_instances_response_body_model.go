// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGdnInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeGdnInstancesResponseBodyData) *DescribeGdnInstancesResponseBody
	GetData() *DescribeGdnInstancesResponseBodyData
	SetMessage(v string) *DescribeGdnInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeGdnInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGdnInstancesResponseBody
	GetSuccess() *bool
}

type DescribeGdnInstancesResponseBody struct {
	Data *DescribeGdnInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7B044BD1-6402-5DE9-9AED-63D15A994E37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGdnInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBody) GetData() *DescribeGdnInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeGdnInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGdnInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGdnInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGdnInstancesResponseBody) SetData(v *DescribeGdnInstancesResponseBodyData) *DescribeGdnInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetMessage(v string) *DescribeGdnInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetRequestId(v string) *DescribeGdnInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGdnInstancesResponseBody) SetSuccess(v bool) *DescribeGdnInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGdnInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGdnInstancesResponseBodyData struct {
	GdnInstanceList []*DescribeGdnInstancesResponseBodyDataGdnInstanceList `json:"GdnInstanceList,omitempty" xml:"GdnInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 130
	TotalNumber *string `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyData) GetGdnInstanceList() []*DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	return s.GdnInstanceList
}

func (s *DescribeGdnInstancesResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeGdnInstancesResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeGdnInstancesResponseBodyData) GetTotalNumber() *string {
	return s.TotalNumber
}

func (s *DescribeGdnInstancesResponseBodyData) SetGdnInstanceList(v []*DescribeGdnInstancesResponseBodyDataGdnInstanceList) *DescribeGdnInstancesResponseBodyData {
	s.GdnInstanceList = v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetPageNumber(v string) *DescribeGdnInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetPageSize(v string) *DescribeGdnInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) SetTotalNumber(v string) *DescribeGdnInstancesResponseBodyData {
	s.TotalNumber = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceList struct {
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// gdn-***
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string                                                          `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	MemberList []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// example:
	//
	// 5.7
	MysqlVersion *string `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ""
	SwitchHistory *string `json:"SwitchHistory,omitempty" xml:"SwitchHistory,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetDescription() *string {
	return s.Description
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetGdnInstanceName() *string {
	return s.GdnInstanceName
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetMemberList() []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	return s.MemberList
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetMysqlVersion() *string {
	return s.MysqlVersion
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetStatus() *string {
	return s.Status
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetSwitchHistory() *string {
	return s.SwitchHistory
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetDescription(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.Description = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetGdnInstanceName(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.GdnInstanceName = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetGmtCreated(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetMemberList(v []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.MemberList = v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetMysqlVersion(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.MysqlVersion = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.Status = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetSwitchHistory(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.SwitchHistory = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) Validate() error {
	return dara.Validate(s)
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList struct {
	// example:
	//
	// polarx.x4.medium.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// example:
	//
	// polarx.x4.medium.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount *string `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *string `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// pxc-***
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	PrimaryZone *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// primary、
	//
	// standby
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// example:
	//
	// 1s
	SecondsBehindMaster *string `json:"SecondsBehindMaster,omitempty" xml:"SecondsBehindMaster,omitempty"`
	// example:
	//
	// Creating
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// example:
	//
	// cn-zhangjiakou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) String() string {
	return dara.Prettify(s)
}

func (s DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GoString() string {
	return s.String()
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetClassCode() *string {
	return s.ClassCode
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetCnNodeClassCode() *string {
	return s.CnNodeClassCode
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetCnNodeCount() *string {
	return s.CnNodeCount
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetDnNodeClassCode() *string {
	return s.DnNodeClassCode
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetDnNodeCount() *string {
	return s.DnNodeCount
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetMemberName() *string {
	return s.MemberName
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetPayType() *string {
	return s.PayType
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetPrimaryZone() *string {
	return s.PrimaryZone
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetRole() *string {
	return s.Role
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetSecondaryZone() *string {
	return s.SecondaryZone
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetSecondsBehindMaster() *string {
	return s.SecondsBehindMaster
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetStatus() *string {
	return s.Status
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetTertiaryZone() *string {
	return s.TertiaryZone
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCnNodeClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCnNodeCount(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetCommodityCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.CommodityCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetDnNodeClassCode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetDnNodeCount(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetExpireTime(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetGmtCreated(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetMemberName(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.MemberName = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetPayType(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.PayType = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetPrimaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetRegionId(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.RegionId = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetRole(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.Role = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetSecondaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.SecondaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetSecondsBehindMaster(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.SecondsBehindMaster = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.Status = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetTaskStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.TaskStatus = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetTertiaryZone(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.TertiaryZone = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetZoneId(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ZoneId = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) Validate() error {
	return dara.Validate(s)
}
