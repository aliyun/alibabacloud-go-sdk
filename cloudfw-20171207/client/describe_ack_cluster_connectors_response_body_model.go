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
	// The list of ACK cluster connectors.
	AckClusterConnectors []*DescribeAckClusterConnectorsResponseBodyAckClusterConnectors `json:"AckClusterConnectors,omitempty" xml:"AckClusterConnectors,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7F333E0-7B70-54DA-A307-4B2B49DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The ACK cluster ID. You can call the following operation to obtain the value:
	//
	// - [DescribeAckClusters](~~DescribeAckClusters~~): Queries a list of ACK clusters in batches.
	//
	// example:
	//
	// f9b9815a5280****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the ACK cluster.
	//
	// example:
	//
	// TestClusterA
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The health check status of the ACK cluster connector.
	//
	// example:
	//
	// healthy
	ConnectorHealthCheckStatus *string `json:"ConnectorHealthCheckStatus,omitempty" xml:"ConnectorHealthCheckStatus,omitempty"`
	// The ID of the ACK cluster connector.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The name of the ACK cluster connector. The name must be 1 to 64 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// ack-cluster-connector-name
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// The instance status of the ACK cluster connector.
	//
	// example:
	//
	// ready
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	// The timestamp when the ACK cluster connector was created. Unit: seconds.
	//
	// example:
	//
	// 1760493347
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The list of address book UUIDs created on the ACK cluster connector.
	GroupUuids []*string `json:"GroupUuids,omitempty" xml:"GroupUuids,omitempty" type:"Repeated"`
	// The Alibaba Cloud UID of the account to which the ACK cluster resource belongs.
	//
	// example:
	//
	// 159663371500****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The primary vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Queries the list of synchronization node vSwitches in batches.
	//
	// example:
	//
	// vsw-2ze2gtlfozrab01cfo****
	PrimaryVswitchId *string `json:"PrimaryVswitchId,omitempty" xml:"PrimaryVswitchId,omitempty"`
	// The IP address of the primary vSwitch of the ACK cluster connector.
	//
	// example:
	//
	// 10.100.2.XXX
	PrimaryVswitchIp *string `json:"PrimaryVswitchIp,omitempty" xml:"PrimaryVswitchIp,omitempty"`
	// The zone of the primary vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Queries the list of synchronization node vSwitch zones in batches.
	//
	// example:
	//
	// cn-beijing-g
	PrimaryVswitchZoneId *string `json:"PrimaryVswitchZoneId,omitempty" xml:"PrimaryVswitchZoneId,omitempty"`
	// The region ID of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceRegionList](~~DescribeAccessInstanceRegionList~~): Queries the list of synchronization node regions.
	//
	// > For more information about the regions supported by ACK cluster connectors in Cloud Firewall, see [ACK cluster synchronization nodes](https://help.aliyun.com/document_detail/2865120.html).
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The standby vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Queries the list of synchronization node vSwitches in batches.
	//
	// example:
	//
	// vsw-2zerfbbje7dvnbii2****
	StandbyVswitchId *string `json:"StandbyVswitchId,omitempty" xml:"StandbyVswitchId,omitempty"`
	// The IP address of the standby vSwitch of the ACK cluster connector.
	//
	// example:
	//
	// 10.100.1.XXX
	StandbyVswitchIp *string `json:"StandbyVswitchIp,omitempty" xml:"StandbyVswitchIp,omitempty"`
	// The zone of the standby vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Queries the list of synchronization node vSwitch zones in batches.
	//
	// example:
	//
	// cn-beijing-h
	StandbyVswitchZoneId *string `json:"StandbyVswitchZoneId,omitempty" xml:"StandbyVswitchZoneId,omitempty"`
	// The container synchronization cycle of the ACK cluster connector.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The reason why the ACK cluster connector is unhealthy.
	//
	// example:
	//
	// The ACK cluster status is unavailable.
	UnhealthyReason *string `json:"UnhealthyReason,omitempty" xml:"UnhealthyReason,omitempty"`
	// The instance ID of the VPC-connected instance to which the ACK cluster belongs.
	//
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
