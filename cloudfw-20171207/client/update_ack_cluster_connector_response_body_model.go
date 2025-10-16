// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAckClusterConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAckClusterConnector(v *UpdateAckClusterConnectorResponseBodyAckClusterConnector) *UpdateAckClusterConnectorResponseBody
	GetAckClusterConnector() *UpdateAckClusterConnectorResponseBodyAckClusterConnector
	SetRequestId(v string) *UpdateAckClusterConnectorResponseBody
	GetRequestId() *string
}

type UpdateAckClusterConnectorResponseBody struct {
	AckClusterConnector *UpdateAckClusterConnectorResponseBodyAckClusterConnector `json:"AckClusterConnector,omitempty" xml:"AckClusterConnector,omitempty" type:"Struct"`
	// example:
	//
	// 5D16AADE-DA2E-5CAB-AA3B-AA197D97****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAckClusterConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAckClusterConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAckClusterConnectorResponseBody) GetAckClusterConnector() *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	return s.AckClusterConnector
}

func (s *UpdateAckClusterConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAckClusterConnectorResponseBody) SetAckClusterConnector(v *UpdateAckClusterConnectorResponseBodyAckClusterConnector) *UpdateAckClusterConnectorResponseBody {
	s.AckClusterConnector = v
	return s
}

func (s *UpdateAckClusterConnectorResponseBody) SetRequestId(v string) *UpdateAckClusterConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBody) Validate() error {
	if s.AckClusterConnector != nil {
		if err := s.AckClusterConnector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAckClusterConnectorResponseBodyAckClusterConnector struct {
	// example:
	//
	// c57ecf39ff32c415e8549a7df27a7e947
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ack-cluster-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// healthy
	ConnectorHealthCheckStatus *string `json:"ConnectorHealthCheckStatus,omitempty" xml:"ConnectorHealthCheckStatus,omitempty"`
	// example:
	//
	// ac-7c1bad6c3cc84c33baab
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// ack-cluster-connector-name
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// ready
	ConnectorStatus *string `json:"ConnectorStatus,omitempty" xml:"ConnectorStatus,omitempty"`
	// example:
	//
	// 1724982259
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// vsw-2zerfbbje7dvnbii2****
	PrimaryVswitchId *string `json:"PrimaryVswitchId,omitempty" xml:"PrimaryVswitchId,omitempty"`
	// example:
	//
	// 10.100.1.1
	PrimaryVswitchIp *string `json:"PrimaryVswitchIp,omitempty" xml:"PrimaryVswitchIp,omitempty"`
	// example:
	//
	// cn-beijing-g
	PrimaryVswitchZoneId *string `json:"PrimaryVswitchZoneId,omitempty" xml:"PrimaryVswitchZoneId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vsw-2ze2gtlfozrab01cfo****
	StandbyVswitchId *string `json:"StandbyVswitchId,omitempty" xml:"StandbyVswitchId,omitempty"`
	// example:
	//
	// 10.100.2.1
	StandbyVswitchIp *string `json:"StandbyVswitchIp,omitempty" xml:"StandbyVswitchIp,omitempty"`
	// example:
	//
	// cn-beijing-h
	StandbyVswitchZoneId *string `json:"StandbyVswitchZoneId,omitempty" xml:"StandbyVswitchZoneId,omitempty"`
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// vpc-j6cvhdscntzuvr0x****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s UpdateAckClusterConnectorResponseBodyAckClusterConnector) String() string {
	return dara.Prettify(s)
}

func (s UpdateAckClusterConnectorResponseBodyAckClusterConnector) GoString() string {
	return s.String()
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorHealthCheckStatus() *string {
	return s.ConnectorHealthCheckStatus
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetConnectorStatus() *string {
	return s.ConnectorStatus
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetMemberUid() *string {
	return s.MemberUid
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchId() *string {
	return s.PrimaryVswitchId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchIp() *string {
	return s.PrimaryVswitchIp
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetPrimaryVswitchZoneId() *string {
	return s.PrimaryVswitchZoneId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetRegionNo() *string {
	return s.RegionNo
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchId() *string {
	return s.StandbyVswitchId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchIp() *string {
	return s.StandbyVswitchIp
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetStandbyVswitchZoneId() *string {
	return s.StandbyVswitchZoneId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetClusterId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetClusterName(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ClusterName = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorHealthCheckStatus(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorHealthCheckStatus = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorName(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorName = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetConnectorStatus(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.ConnectorStatus = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetCreateTime(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.CreateTime = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetMemberUid(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.MemberUid = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchIp(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchIp = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetPrimaryVswitchZoneId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.PrimaryVswitchZoneId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetRegionNo(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.RegionNo = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchIp(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchIp = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetStandbyVswitchZoneId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.StandbyVswitchZoneId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetTtl(v int32) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.Ttl = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) SetVpcId(v string) *UpdateAckClusterConnectorResponseBodyAckClusterConnector {
	s.VpcId = &v
	return s
}

func (s *UpdateAckClusterConnectorResponseBodyAckClusterConnector) Validate() error {
	return dara.Validate(s)
}
