// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClusterConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnector(v *DescribeAckClusterConnectorResponseBodyAckClusterConnector) *DescribeAckClusterConnectorResponseBody
	GetAckClusterConnector() *DescribeAckClusterConnectorResponseBodyAckClusterConnector
	SetRequestId(v string) *DescribeAckClusterConnectorResponseBody
	GetRequestId() *string
}

type DescribeAckClusterConnectorResponseBody struct {
	// The ACK cluster connector.
	AckClusterConnector *DescribeAckClusterConnectorResponseBodyAckClusterConnector `json:"AckClusterConnector,omitempty" xml:"AckClusterConnector,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 45E2E720-D2B4-506F-B682-1FCBE971****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAckClusterConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorResponseBody) GetAckClusterConnector() *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	return s.AckClusterConnector
}

func (s *DescribeAckClusterConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckClusterConnectorResponseBody) SetAckClusterConnector(v *DescribeAckClusterConnectorResponseBodyAckClusterConnector) *DescribeAckClusterConnectorResponseBody {
	s.AckClusterConnector = v
	return s
}

func (s *DescribeAckClusterConnectorResponseBody) SetRequestId(v string) *DescribeAckClusterConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBody) Validate() error {
	if s.AckClusterConnector != nil {
		if err := s.AckClusterConnector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAckClusterConnectorResponseBodyAckClusterConnector struct {
	// The ID of the ACK cluster. You can obtain the value from:
	//
	// - [DescribeAckClusters](~~DescribeAckClusters~~): Lists ACK clusters.
	//
	// example:
	//
	// c857d908016794125883a9ee8196cba17
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the ACK cluster.
	//
	// example:
	//
	// ack-cluster-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The health check status of the ACK cluster connector.
	//
	// example:
	//
	// healthy
	ConnectorHealthCheckStatus *string `json:"ConnectorHealthCheckStatus,omitempty" xml:"ConnectorHealthCheckStatus,omitempty"`
	// The ID of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAckClusterConnectors](~~DescribeAckClusterConnectors~~): Lists ACK cluster connectors.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The name of the ACK cluster connector. The name is 1 to 64 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
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
	// 1724982259
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud UID of the account to which the ACK cluster resource belongs.
	//
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// The primary vSwitch of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Lists vSwitches of synchronization nodes.
	//
	// example:
	//
	// vsw-2zerfbbje7dvnbii2****
	PrimaryVswitchId *string `json:"PrimaryVswitchId,omitempty" xml:"PrimaryVswitchId,omitempty"`
	// The IP address of the primary vSwitch of the ACK cluster connector.
	//
	// example:
	//
	// 10.100.1.1
	PrimaryVswitchIp *string `json:"PrimaryVswitchIp,omitempty" xml:"PrimaryVswitchIp,omitempty"`
	// The zone of the primary vSwitch of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Lists zones of synchronization node vSwitches.
	//
	// example:
	//
	// cn-beijing-g
	PrimaryVswitchZoneId *string `json:"PrimaryVswitchZoneId,omitempty" xml:"PrimaryVswitchZoneId,omitempty"`
	// The region ID of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAccessInstanceRegionList](~~DescribeAccessInstanceRegionList~~): Queries the list of synchronization node regions.
	//
	// > For more information about the regions supported by ACK cluster connectors in Cloud Firewall, see [ACK cluster synchronization nodes](https://help.aliyun.com/document_detail/2865120.html).
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The standby vSwitch of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Lists vSwitches of synchronization nodes.
	//
	// example:
	//
	// vsw-2ze2gtlfozrab01cfo****
	StandbyVswitchId *string `json:"StandbyVswitchId,omitempty" xml:"StandbyVswitchId,omitempty"`
	// The IP address of the standby vSwitch of the ACK cluster connector.
	//
	// example:
	//
	// 10.100.2.1
	StandbyVswitchIp *string `json:"StandbyVswitchIp,omitempty" xml:"StandbyVswitchIp,omitempty"`
	// The zone of the standby vSwitch of the ACK cluster connector. You can obtain the value from:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Lists zones of synchronization node vSwitches.
	//
	// example:
	//
	// cn-beijing-h
	StandbyVswitchZoneId *string `json:"StandbyVswitchZoneId,omitempty" xml:"StandbyVswitchZoneId,omitempty"`
	// The synchronization interval of the ACK cluster connector. Valid values: 2 to 60. Unit: seconds.
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

func (s DescribeAckClusterConnectorResponseBodyAckClusterConnector) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClusterConnectorResponseBodyAckClusterConnector) GoString() string {
	return s.String()
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorHealthCheckStatus() *string {
	return s.ConnectorHealthCheckStatus
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchId() *string {
	return s.PrimaryVswitchId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchIp() *string {
	return s.PrimaryVswitchIp
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchZoneId() *string {
	return s.PrimaryVswitchZoneId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchId() *string {
	return s.StandbyVswitchId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchIp() *string {
	return s.StandbyVswitchIp
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchZoneId() *string {
	return s.StandbyVswitchZoneId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetUnhealthyReason() *string {
	return s.UnhealthyReason
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetClusterId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetClusterName(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterName = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorHealthCheckStatus(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorHealthCheckStatus = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorName(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorName = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorStatus(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorStatus = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetCreateTime(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.CreateTime = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetMemberUid(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.MemberUid = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchIp(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchIp = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchZoneId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchZoneId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetRegionNo(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.RegionNo = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchIp(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchIp = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchZoneId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchZoneId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetTtl(v int32) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.Ttl = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetUnhealthyReason(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.UnhealthyReason = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) SetVpcId(v string) *DescribeAckClusterConnectorResponseBodyAckClusterConnector {
	s.VpcId = &v
	return s
}

func (s *DescribeAckClusterConnectorResponseBodyAckClusterConnector) Validate() error {
	return dara.Validate(s)
}
