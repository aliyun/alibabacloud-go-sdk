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
	// The list of instance details.
	Data *DescribeGdnInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// > This parameter is empty when the request succeeds. When the request fails, an exception message is returned, such as an error code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B044BD1-6402-5DE9-9AED-63D15A994E37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGdnInstancesResponseBodyData struct {
	// The list of GDN instances.
	GdnInstanceList []*DescribeGdnInstancesResponseBodyDataGdnInstanceList `json:"GdnInstanceList,omitempty" xml:"GdnInstanceList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
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
	if s.GdnInstanceList != nil {
		for _, item := range s.GdnInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceList struct {
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The GDN instance name.
	//
	// example:
	//
	// gdn-***
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	GdnMode         *string `json:"GdnMode,omitempty" xml:"GdnMode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The list of members.
	MemberList []*DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// The MySQL version supported by the instance.
	//
	// example:
	//
	// 5.7
	MysqlVersion        *string `json:"MysqlVersion,omitempty" xml:"MysqlVersion,omitempty"`
	RplConflictStrategy *string `json:"RplConflictStrategy,omitempty" xml:"RplConflictStrategy,omitempty"`
	RplDmlStrategy      *string `json:"RplDmlStrategy,omitempty" xml:"RplDmlStrategy,omitempty"`
	RplSyncDdl          *bool   `json:"RplSyncDdl,omitempty" xml:"RplSyncDdl,omitempty"`
	// The status.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The switchover log.
	//
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetGdnMode() *string {
	return s.GdnMode
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetRplConflictStrategy() *string {
	return s.RplConflictStrategy
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetRplDmlStrategy() *string {
	return s.RplDmlStrategy
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) GetRplSyncDdl() *bool {
	return s.RplSyncDdl
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetGdnMode(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.GdnMode = &v
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetRplConflictStrategy(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.RplConflictStrategy = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetRplDmlStrategy(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.RplDmlStrategy = &v
	return s
}

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceList) SetRplSyncDdl(v bool) *DescribeGdnInstancesResponseBodyDataGdnInstanceList {
	s.RplSyncDdl = &v
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
	if s.MemberList != nil {
		for _, item := range s.MemberList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList struct {
	// The instance type.
	//
	// example:
	//
	// polarx.x4.medium.2e
	ClassCode *string `json:"ClassCode,omitempty" xml:"ClassCode,omitempty"`
	// The CN node specifications. Valid values:
	//
	// - **polarx.x4.medium.2e**: 2 cores, 8 GB
	//
	// - **polarx.x4.large.2e**: 4 cores, 16 GB
	//
	// - **polarx.x8.large.2e**: 4 cores, 32 GB
	//
	// - **polarx.x4.xlarge.2e**: 8 cores, 32 GB
	//
	// - **polarx.x8.xlarge.2e**: 8 cores, 64 GB
	//
	// - **polarx.x4.2xlarge.2e**: 16 cores, 64 GB
	//
	// - **polarx.x8.2xlarge.2e**: 16 cores, 128 GB
	//
	// - **polarx.x4.4xlarge.2e**: 32 cores, 128 GB
	//
	// - **polarx.x8.4xlarge.2e**: 32 cores, 256 GB
	//
	// - **polarx.st.8xlarge.2e**: 60 cores, 470 GB
	//
	// - **polarx.st.12xlarge.2e**: 90 cores, 720 GB.
	//
	// example:
	//
	// polarx.x4.medium.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// The number of CN nodes.
	//
	// example:
	//
	// 2
	CnNodeCount *string `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// drds_polarxpre_public_cn
	CommodityCode  *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	DataSyncStatus *string `json:"DataSyncStatus,omitempty" xml:"DataSyncStatus,omitempty"`
	// The DN node specifications. Valid values:
	//
	// - **mysql.n2.medium.25**: 2 cores, 4 GB
	//
	// - **mysql.n4.medium.25**: 2 cores, 8 GB
	//
	// - **mysql.x8.medium.25**: 2 cores, 16 GB
	//
	// - **mysql.n2.large.25**: 4 cores, 8 GB
	//
	// - **mysql.n4.large.25**: 4 cores, 16 GB
	//
	// - **mysql.x8.large.25**: 4 cores, 32 GB
	//
	// - **mysql.n2.xlarge.25**: 8 cores, 16 GB
	//
	// - **mysql.n4.xlarge.25**: 8 cores, 32 GB
	//
	// - **mysql.x8.xlarge.25**: 8 cores, 64 GB
	//
	// - **mysql.n4.2xlarge.25**: 16 cores, 64 GB
	//
	// - **mysql.x8.2xlarge.25**: 16 cores, 128 GB
	//
	// - **mysql.x4.4xlarge.25**: 32 cores, 128 GB
	//
	// - **mysql.x8.4xlarge.25**: 32 cores, 256 GB
	//
	// - **mysql.st.8xlarge.25**: 60 cores, 470 GB
	//
	// - **mysql.st.12xlarge.25**: 90 cores, 720 GB.
	//
	// example:
	//
	// mysql.n4.medium.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// The number of DN nodes.
	//
	// example:
	//
	// 2
	DnNodeCount *string `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-01-02T13:11:10.000+0000
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The member name (PolarDB-X instance name).
	//
	// example:
	//
	// pxc-***
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go.
	//
	// - **Prepaid**: subscription.
	//
	// example:
	//
	// Prepaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-zhangjiakou-a
	PrimaryZone     *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	ReadWriteStatus *string `json:"ReadWriteStatus,omitempty" xml:"ReadWriteStatus,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The member role.
	//
	// example:
	//
	// primary、
	//
	// standby
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The secondary zone.
	//
	// example:
	//
	// cn-zhangjiakou-a
	SecondaryZone *string `json:"SecondaryZone,omitempty" xml:"SecondaryZone,omitempty"`
	// The data latency.
	//
	// example:
	//
	// 1s
	SecondsBehindMaster *string `json:"SecondsBehindMaster,omitempty" xml:"SecondsBehindMaster,omitempty"`
	// The member status.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The switchover task status.
	//
	// example:
	//
	// prepared：参数初始化完毕
	//
	// set_old_primary_readonly：原主实例已禁写
	//
	// set_new_primary_read_write：已切换
	//
	// timeout：任务超时
	//
	// rollback：已回滚
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The zone for Three-zone deployment. This zone is active only when three-zone deployment is enabled.
	//
	// example:
	//
	// cn-zhangjiakou-a
	TertiaryZone *string `json:"TertiaryZone,omitempty" xml:"TertiaryZone,omitempty"`
	// The zone ID.
	//
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetDataSyncStatus() *string {
	return s.DataSyncStatus
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) GetReadWriteStatus() *string {
	return s.ReadWriteStatus
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetDataSyncStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.DataSyncStatus = &v
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

func (s *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList) SetReadWriteStatus(v string) *DescribeGdnInstancesResponseBodyDataGdnInstanceListMemberList {
	s.ReadWriteStatus = &v
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
