// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceViaEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstance(v *DescribeDBInstanceViaEndpointResponseBodyDBInstance) *DescribeDBInstanceViaEndpointResponseBody
	GetDBInstance() *DescribeDBInstanceViaEndpointResponseBodyDBInstance
	SetRequestId(v string) *DescribeDBInstanceViaEndpointResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceViaEndpointResponseBody struct {
	DBInstance *DescribeDBInstanceViaEndpointResponseBodyDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" type:"Struct"`
	// example:
	//
	// 14036EBE-CF2E-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBody) GetDBInstance() *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	return s.DBInstance
}

func (s *DescribeDBInstanceViaEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceViaEndpointResponseBody) SetDBInstance(v *DescribeDBInstanceViaEndpointResponseBodyDBInstance) *DescribeDBInstanceViaEndpointResponseBody {
	s.DBInstance = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBody) SetRequestId(v string) *DescribeDBInstanceViaEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBody) Validate() error {
	if s.DBInstance != nil {
		if err := s.DBInstance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstance struct {
	// example:
	//
	// polarx.x4.xlarge.2e
	CnNodeClassCode *string `json:"CnNodeClassCode,omitempty" xml:"CnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	CnNodeCount *int32 `json:"CnNodeCount,omitempty" xml:"CnNodeCount,omitempty"`
	// example:
	//
	// drds_polarxpost_public_cn
	CommodityCode *string                                                         `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConnAddrs     []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs `json:"ConnAddrs,omitempty" xml:"ConnAddrs,omitempty" type:"Repeated"`
	// example:
	//
	// pxc-sprpx766vo****.polarx.singapore.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 2021-08-31T08:56:25.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ReadWrite
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	DBNodeClass *string `json:"DBNodeClass,omitempty" xml:"DBNodeClass,omitempty"`
	// example:
	//
	// 2
	DBNodeCount *int32                                                        `json:"DBNodeCount,omitempty" xml:"DBNodeCount,omitempty"`
	DBNodes     []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Repeated"`
	// example:
	//
	// polarx
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 5.5
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// test instance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mysql.x8.large.25
	DnNodeClassCode *string `json:"DnNodeClassCode,omitempty" xml:"DnNodeClassCode,omitempty"`
	// example:
	//
	// 2
	DnNodeCount *int32 `json:"DnNodeCount,omitempty" xml:"DnNodeCount,omitempty"`
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2022-08-31T16:00:00.000+0000
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// example:
	//
	// false
	Expired *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 18
	KindCode *int32 `json:"KindCode,omitempty" xml:"KindCode,omitempty"`
	// This parameter is required.
	LTSVersions []*string `json:"LTSVersions,omitempty" xml:"LTSVersions,omitempty" type:"Repeated"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	LatestMinorVersion *string `json:"LatestMinorVersion,omitempty" xml:"LatestMinorVersion,omitempty"`
	// example:
	//
	// Unlock
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 06:00
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// example:
	//
	// 06:00
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.11-16301083_xcluster-20210805
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// example:
	//
	// VPC
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 3306
	Port            *string   `json:"Port,omitempty" xml:"Port,omitempty"`
	ReadDBInstances []*string `json:"ReadDBInstances,omitempty" xml:"ReadDBInstances,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// false
	RightsSeparationEnabled *bool `json:"RightsSeparationEnabled,omitempty" xml:"RightsSeparationEnabled,omitempty"`
	// example:
	//
	// disabled
	RightsSeparationStatus *string `json:"RightsSeparationStatus,omitempty" xml:"RightsSeparationStatus,omitempty"`
	// example:
	//
	// enterprise
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 17042505728
	StorageUsed *int64                                                       `json:"StorageUsed,omitempty" xml:"StorageUsed,omitempty"`
	TagSet      []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty" type:"Repeated"`
	// example:
	//
	// ReadWrite
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetCnNodeClassCode() *string {
	return s.CnNodeClassCode
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetCnNodeCount() *int32 {
	return s.CnNodeCount
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetConnAddrs() []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	return s.ConnAddrs
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBNodeClass() *string {
	return s.DBNodeClass
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBNodeCount() *int32 {
	return s.DBNodeCount
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBNodes() []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	return s.DBNodes
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDnNodeClassCode() *string {
	return s.DnNodeClassCode
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetDnNodeCount() *int32 {
	return s.DnNodeCount
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetExpired() *string {
	return s.Expired
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetKindCode() *int32 {
	return s.KindCode
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetLTSVersions() []*string {
	return s.LTSVersions
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetLatestMinorVersion() *string {
	return s.LatestMinorVersion
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetMaintainEndTime() *string {
	return s.MaintainEndTime
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetMaintainStartTime() *string {
	return s.MaintainStartTime
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetNetwork() *string {
	return s.Network
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetReadDBInstances() []*string {
	return s.ReadDBInstances
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetRightsSeparationEnabled() *bool {
	return s.RightsSeparationEnabled
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetRightsSeparationStatus() *string {
	return s.RightsSeparationStatus
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetSeries() *string {
	return s.Series
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetStorageUsed() *int64 {
	return s.StorageUsed
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetTagSet() []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet {
	return s.TagSet
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCnNodeClassCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCnNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCommodityCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CommodityCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetConnAddrs(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ConnAddrs = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetConnectionString(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetCreateTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBInstanceType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodeClass(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodeClass = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBNodes(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDBVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDescription(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDnNodeClassCode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DnNodeClassCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetDnNodeCount(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.DnNodeCount = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetEngine(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetExpireDate(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ExpireDate = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetExpired(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Expired = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetKindCode(v int32) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.KindCode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLTSVersions(v []*string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LTSVersions = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLatestMinorVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LatestMinorVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetLockMode(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMaintainEndTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMaintainStartTime(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetMinorVersion(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.MinorVersion = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetNetwork(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Network = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetPayType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetPort(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetReadDBInstances(v []*string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ReadDBInstances = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRegionId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetResourceGroupId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRightsSeparationEnabled(v bool) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RightsSeparationEnabled = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetRightsSeparationStatus(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.RightsSeparationStatus = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetSeries(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Series = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetStatus(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Status = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetStorageUsed(v int64) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.StorageUsed = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetTagSet(v []*DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.TagSet = v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetVPCId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetVSwitchId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) SetZoneId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstance) Validate() error {
	if s.ConnAddrs != nil {
		for _, item := range s.ConnAddrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DBNodes != nil {
		for _, item := range s.DBNodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TagSet != nil {
		for _, item := range s.TagSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs struct {
	// example:
	//
	// polardbx-xxx.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// 3306
	Port *int64 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-xxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// pxc-zkralxpc5d****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetPort() *int64 {
	return s.Port
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetConnectionString(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetPort(v int64) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetType(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVPCId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVSwitchId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) SetVpcInstanceId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceConnAddrs) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes struct {
	// example:
	//
	// pxc-i-********
	ComputeNodeId *string `json:"ComputeNodeId,omitempty" xml:"ComputeNodeId,omitempty"`
	// example:
	//
	// pxc-xdb-xxxxxx
	DataNodeId *string `json:"DataNodeId,omitempty" xml:"DataNodeId,omitempty"`
	// example:
	//
	// pxi-*********
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// polarx.x4.large.2e
	NodeClass *string `json:"NodeClass,omitempty" xml:"NodeClass,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetComputeNodeId() *string {
	return s.ComputeNodeId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetDataNodeId() *string {
	return s.DataNodeId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetId() *string {
	return s.Id
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetNodeClass() *string {
	return s.NodeClass
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetComputeNodeId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.ComputeNodeId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetDataNodeId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.DataNodeId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.Id = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetNodeClass(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.NodeClass = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetRegionId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) SetZoneId(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet struct {
	// example:
	//
	// key2
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value2
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) SetKey(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) SetValue(v string) *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet {
	s.Value = &v
	return s
}

func (s *DescribeDBInstanceViaEndpointResponseBodyDBInstanceTagSet) Validate() error {
	return dara.Validate(s)
}
