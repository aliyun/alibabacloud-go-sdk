// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAckClusterConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateAckClusterConnectorRequest
	GetClusterId() *string
	SetConnectorName(v string) *CreateAckClusterConnectorRequest
	GetConnectorName() *string
	SetMemberUid(v string) *CreateAckClusterConnectorRequest
	GetMemberUid() *string
	SetPrimaryVswitchId(v string) *CreateAckClusterConnectorRequest
	GetPrimaryVswitchId() *string
	SetPrimaryVswitchIp(v string) *CreateAckClusterConnectorRequest
	GetPrimaryVswitchIp() *string
	SetRegionNo(v string) *CreateAckClusterConnectorRequest
	GetRegionNo() *string
	SetStandbyVswitchId(v string) *CreateAckClusterConnectorRequest
	GetStandbyVswitchId() *string
	SetStandbyVswitchIp(v string) *CreateAckClusterConnectorRequest
	GetStandbyVswitchIp() *string
	SetTtl(v string) *CreateAckClusterConnectorRequest
	GetTtl() *string
}

type CreateAckClusterConnectorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cb0f5640b1b2d404cad6ba21509d7847b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ack-cluster-connector-name
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-2zerfbbje7dvnbii2****
	PrimaryVswitchId *string `json:"PrimaryVswitchId,omitempty" xml:"PrimaryVswitchId,omitempty"`
	// example:
	//
	// 10.100.1.1
	PrimaryVswitchIp *string `json:"PrimaryVswitchIp,omitempty" xml:"PrimaryVswitchIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// vsw-2ze2gtlfozrab01cfo****
	StandbyVswitchId *string `json:"StandbyVswitchId,omitempty" xml:"StandbyVswitchId,omitempty"`
	// example:
	//
	// 10.100.2.1
	StandbyVswitchIp *string `json:"StandbyVswitchIp,omitempty" xml:"StandbyVswitchIp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *string `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateAckClusterConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAckClusterConnectorRequest) GoString() string {
	return s.String()
}

func (s *CreateAckClusterConnectorRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateAckClusterConnectorRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *CreateAckClusterConnectorRequest) GetMemberUid() *string {
	return s.MemberUid
}

func (s *CreateAckClusterConnectorRequest) GetPrimaryVswitchId() *string {
	return s.PrimaryVswitchId
}

func (s *CreateAckClusterConnectorRequest) GetPrimaryVswitchIp() *string {
	return s.PrimaryVswitchIp
}

func (s *CreateAckClusterConnectorRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *CreateAckClusterConnectorRequest) GetStandbyVswitchId() *string {
	return s.StandbyVswitchId
}

func (s *CreateAckClusterConnectorRequest) GetStandbyVswitchIp() *string {
	return s.StandbyVswitchIp
}

func (s *CreateAckClusterConnectorRequest) GetTtl() *string {
	return s.Ttl
}

func (s *CreateAckClusterConnectorRequest) SetClusterId(v string) *CreateAckClusterConnectorRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetConnectorName(v string) *CreateAckClusterConnectorRequest {
	s.ConnectorName = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetMemberUid(v string) *CreateAckClusterConnectorRequest {
	s.MemberUid = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetPrimaryVswitchId(v string) *CreateAckClusterConnectorRequest {
	s.PrimaryVswitchId = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetPrimaryVswitchIp(v string) *CreateAckClusterConnectorRequest {
	s.PrimaryVswitchIp = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetRegionNo(v string) *CreateAckClusterConnectorRequest {
	s.RegionNo = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetStandbyVswitchId(v string) *CreateAckClusterConnectorRequest {
	s.StandbyVswitchId = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetStandbyVswitchIp(v string) *CreateAckClusterConnectorRequest {
	s.StandbyVswitchIp = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) SetTtl(v string) *CreateAckClusterConnectorRequest {
	s.Ttl = &v
	return s
}

func (s *CreateAckClusterConnectorRequest) Validate() error {
	return dara.Validate(s)
}
