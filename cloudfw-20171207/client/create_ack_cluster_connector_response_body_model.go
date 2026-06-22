// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAckClusterConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnector(v *CreateAckClusterConnectorResponseBodyAckClusterConnector) *CreateAckClusterConnectorResponseBody
	GetAckClusterConnector() *CreateAckClusterConnectorResponseBodyAckClusterConnector
	SetRequestId(v string) *CreateAckClusterConnectorResponseBody
	GetRequestId() *string
}

type CreateAckClusterConnectorResponseBody struct {
	// The ACK cluster connector information.
	AckClusterConnector *CreateAckClusterConnectorResponseBodyAckClusterConnector `json:"AckClusterConnector,omitempty" xml:"AckClusterConnector,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0DC783F1-B3A7-578D-8A63-*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAckClusterConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAckClusterConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAckClusterConnectorResponseBody) GetAckClusterConnector() *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	return s.AckClusterConnector
}

func (s *CreateAckClusterConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAckClusterConnectorResponseBody) SetAckClusterConnector(v *CreateAckClusterConnectorResponseBodyAckClusterConnector) *CreateAckClusterConnectorResponseBody {
	s.AckClusterConnector = v
	return s
}

func (s *CreateAckClusterConnectorResponseBody) SetRequestId(v string) *CreateAckClusterConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBody) Validate() error {
	if s.AckClusterConnector != nil {
		if err := s.AckClusterConnector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAckClusterConnectorResponseBodyAckClusterConnector struct {
	// The entry point IP address used by the ACK cluster connector to access the ACK cluster over the internal network.
	//
	// example:
	//
	// 10.40.32.240
	AckClientHostIp *string `json:"AckClientHostIp,omitempty" xml:"AckClientHostIp,omitempty"`
	// The ACK cluster ID. You can call the following operation to obtain the value:
	//
	// - [DescribeAckClusters](~~DescribeAckClusters~~): Lists ACK clusters.
	//
	// example:
	//
	// cb0f5640b1b2d404cad6ba21509d7847b
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
	// The unique ID of the ACK cluster connector.
	//
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The name of the ACK cluster connector. The name must be 1 to 64 characters in length and can contain letters, digits, Chinese characters, periods (.), underscores (_), and hyphens (-).
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
	// The primary vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Lists the vSwitches of synchronization nodes.
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
	// The zone of the primary vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Lists the zones of vSwitches for synchronization nodes.
	//
	// example:
	//
	// cn-hangzhou-g
	PrimaryVswitchZoneId *string `json:"PrimaryVswitchZoneId,omitempty" xml:"PrimaryVswitchZoneId,omitempty"`
	// The region ID of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceRegionList](~~DescribeAccessInstanceRegionList~~): Lists the regions of synchronization nodes.
	//
	// > For more information about the regions supported by ACK cluster connectors in Cloud Firewall, see [ACK cluster synchronization nodes](https://help.aliyun.com/document_detail/2865120.html).
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The standby vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceVSwitchList](~~DescribeAccessInstanceVSwitchList~~): Lists the vSwitches of synchronization nodes.
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
	// The zone of the standby vSwitch of the ACK cluster connector. You can call the following operation to obtain the value:
	//
	// - [DescribeAccessInstanceZoneList](~~DescribeAccessInstanceZoneList~~): Lists the zones of vSwitches for synchronization nodes.
	//
	// example:
	//
	// cn-hangzhou-h
	StandbyVswitchZoneId *string `json:"StandbyVswitchZoneId,omitempty" xml:"StandbyVswitchZoneId,omitempty"`
	// The unique ID of the task for creating the ACK cluster connector.
	//
	// example:
	//
	// task-c92d4544ef7b6a42
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The synchronization interval of the ACK cluster connector. Valid values: 2 to 60. Unit: seconds.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The instance ID of the VPC-connected instance to which the ACK cluster belongs.
	//
	// example:
	//
	// vpc-j6cvhdscntzuvr0x****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateAckClusterConnectorResponseBodyAckClusterConnector) String() string {
	return dara.Prettify(s)
}

func (s CreateAckClusterConnectorResponseBodyAckClusterConnector) GoString() string {
	return s.String()
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetAckClientHostIp() *string {
	return s.AckClientHostIp
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorHealthCheckStatus() *string {
	return s.ConnectorHealthCheckStatus
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchId() *string {
	return s.PrimaryVswitchId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchIp() *string {
	return s.PrimaryVswitchIp
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchZoneId() *string {
	return s.PrimaryVswitchZoneId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchId() *string {
	return s.StandbyVswitchId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchIp() *string {
	return s.StandbyVswitchIp
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchZoneId() *string {
	return s.StandbyVswitchZoneId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetAckClientHostIp(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.AckClientHostIp = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetClusterId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetClusterName(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterName = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorHealthCheckStatus(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorHealthCheckStatus = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorName(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorName = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorStatus(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorStatus = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetCreateTime(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.CreateTime = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetMemberUid(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.MemberUid = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchIp(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchIp = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchZoneId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchZoneId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetRegionNo(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.RegionNo = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchIp(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchIp = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchZoneId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchZoneId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetTaskId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.TaskId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetTtl(v int32) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.Ttl = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) SetVpcId(v string) *CreateAckClusterConnectorResponseBodyAckClusterConnector {
	s.VpcId = &v
	return s
}

func (s *CreateAckClusterConnectorResponseBodyAckClusterConnector) Validate() error {
	return dara.Validate(s)
}
