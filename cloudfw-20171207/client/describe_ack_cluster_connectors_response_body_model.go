// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnectors(v []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) *DescribeAckClusterConnectorsResponseBody
	GetAckClusterConnectors() []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors
	SetPageNo(v int32) *DescribeAckClusterConnectorsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAckClusterConnectorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAckClusterConnectorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAckClusterConnectorsResponseBody
	GetTotalCount() *int32
}

type DescribeAckClusterConnectorsResponseBody struct {
	AckClusterConnectors []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors `json:"AckClusterConnectors,omitempty" xml:"AckClusterConnectors,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// E7F333E0-7B70-54DA-A307-4B2B49DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAckClusterConnectorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorsResponseBody) GetAckClusterConnectors() []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	return s.AckClusterConnectors
}

func (s *DescribeAckClusterConnectorsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAckClusterConnectorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAckClusterConnectorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckClusterConnectorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAckClusterConnectorsResponseBody) SetAckClusterConnectors(v []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) *DescribeAckClusterConnectorsResponseBody {
	s.AckClusterConnectors = v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBody) SetPageNo(v int32) *DescribeAckClusterConnectorsResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBody) SetPageSize(v int32) *DescribeAckClusterConnectorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBody) SetRequestId(v string) *DescribeAckClusterConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBody) SetTotalCount(v int32) *DescribeAckClusterConnectorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBody) Validate() error {
	if s.AckClusterConnectors != nil {
		for _, item := range s.AckClusterConnectors {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAckClusterConnectorsResponseBodyAckClusterConnectors struct {
	// example:
	//
	// f9b9815a5280****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// TestClusterA
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// healthy
	ConnectorHealthCheckStatus *string `json:"ConnectorHealthCheckStatus,omitempty" xml:"ConnectorHealthCheckStatus,omitempty"`
	// example:
	//
	// connector-7ff4df316c9a458d****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// test
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// ready
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	// example:
	//
	// 1760493347
	CreateTime *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GroupUuids []*string `json:"GroupUuids,omitempty" xml:"GroupUuids,omitempty" type:"Repeated"`
	// example:
	//
	// 159663371500****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// vsw-2ze2gtlfozrab01cfo****
	PrimaryVswitchId *string `json:"PrimaryVswitchId,omitempty" xml:"PrimaryVswitchId,omitempty"`
	// example:
	//
	// 10.100.2.XXX
	PrimaryVswitchIp *string `json:"PrimaryVswitchIp,omitempty" xml:"PrimaryVswitchIp,omitempty"`
	// example:
	//
	// cn-beijing-g
	PrimaryVswitchZoneId *string `json:"PrimaryVswitchZoneId,omitempty" xml:"PrimaryVswitchZoneId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vsw-2zerfbbje7dvnbii2****
	StandbyVswitchId *string `json:"StandbyVswitchId,omitempty" xml:"StandbyVswitchId,omitempty"`
	// example:
	//
	// 10.100.1.XXX
	StandbyVswitchIp *string `json:"StandbyVswitchIp,omitempty" xml:"StandbyVswitchIp,omitempty"`
	// example:
	//
	// cn-beijing-h
	StandbyVswitchZoneId *string `json:"StandbyVswitchZoneId,omitempty" xml:"StandbyVswitchZoneId,omitempty"`
	// example:
	//
	// 30
	Ttl             *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	UnhealthyReason *string `json:"UnhealthyReason,omitempty" xml:"UnhealthyReason,omitempty"`
	// example:
	//
	// vpc-j6cvhdscntzuvr0x****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetConnectorHealthCheckStatus() *string {
	return s.ConnectorHealthCheckStatus
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetGroupUuids() []*string {
	return s.GroupUuids
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetPrimaryVswitchId() *string {
	return s.PrimaryVswitchId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetPrimaryVswitchIp() *string {
	return s.PrimaryVswitchIp
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetPrimaryVswitchZoneId() *string {
	return s.PrimaryVswitchZoneId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetStandbyVswitchId() *string {
	return s.StandbyVswitchId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetStandbyVswitchIp() *string {
	return s.StandbyVswitchIp
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetStandbyVswitchZoneId() *string {
	return s.StandbyVswitchZoneId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetUnhealthyReason() *string {
	return s.UnhealthyReason
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetClusterId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ClusterId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetClusterName(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ClusterName = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetConnectorHealthCheckStatus(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ConnectorHealthCheckStatus = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetConnectorId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ConnectorId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetConnectorName(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ConnectorName = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetConnectorStatus(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetCreateTime(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.CreateTime = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetGroupUuids(v []*string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.GroupUuids = v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetMemberUid(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.MemberUid = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetPrimaryVswitchId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.PrimaryVswitchId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetPrimaryVswitchIp(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.PrimaryVswitchIp = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetPrimaryVswitchZoneId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.PrimaryVswitchZoneId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetRegionNo(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.RegionNo = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetStandbyVswitchId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.StandbyVswitchId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetStandbyVswitchIp(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.StandbyVswitchIp = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetStandbyVswitchZoneId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.StandbyVswitchZoneId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetTtl(v int32) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.Ttl = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetUnhealthyReason(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.UnhealthyReason = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) SetVpcId(v string) *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors {
	s.VpcId = &v
	return s
}

func (s *DescribeAckClusterConnectorsResponseBodyAckClusterConnectors) Validate() error {
	return dara.Validate(s)
}
