// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetAccessInstanceId() *string
	SetAccessInstanceName(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetAccessInstanceName() *string
	SetAliUid(v int64) *DescribePrivateDnsEndpointDetailResponseBody
	GetAliUid() *int64
	SetEndpointId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetEndpointId() *string
	SetFirewallType(v []*string) *DescribePrivateDnsEndpointDetailResponseBody
	GetFirewallType() []*string
	SetGmtCreate(v int64) *DescribePrivateDnsEndpointDetailResponseBody
	GetGmtCreate() *int64
	SetIpProtocol(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetIpProtocol() *string
	SetMemberUid(v int64) *DescribePrivateDnsEndpointDetailResponseBody
	GetMemberUid() *int64
	SetPort(v int32) *DescribePrivateDnsEndpointDetailResponseBody
	GetPort() *int32
	SetPrimaryDns(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetPrimaryDns() *string
	SetPrimaryVSwitchId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetPrimaryVSwitchId() *string
	SetPrimaryVSwitchIp(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetPrimaryVSwitchIp() *string
	SetPrimaryZoneId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetPrimaryZoneId() *string
	SetPrivateDnsType(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetPrivateDnsType() *string
	SetRegionNo(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetRegionNo() *string
	SetRequestId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetRequestId() *string
	SetStandbyDns(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetStandbyDns() *string
	SetStandbyVSwitchId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetStandbyVSwitchId() *string
	SetStandbyVSwitchIp(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetStandbyVSwitchIp() *string
	SetStandbyZoneId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetStandbyZoneId() *string
	SetStatus(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetStatus() *string
	SetTaskId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetTaskId() *string
	SetVpcId(v string) *DescribePrivateDnsEndpointDetailResponseBody
	GetVpcId() *string
}

type DescribePrivateDnsEndpointDetailResponseBody struct {
	// example:
	//
	// pd-12345
	AccessInstanceId   *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	AccessInstanceName *string `json:"AccessInstanceName,omitempty" xml:"AccessInstanceName,omitempty"`
	// example:
	//
	// 119898001566xxxx
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// ep-1nmi412c28c374****
	EndpointId   *string   `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	FirewallType []*string `json:"FirewallType,omitempty" xml:"FirewallType,omitempty" type:"Repeated"`
	// example:
	//
	// 1715075765
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// example:
	//
	// 258039427902****
	MemberUid *int64 `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 1.1.1.1
	PrimaryDns *string `json:"PrimaryDns,omitempty" xml:"PrimaryDns,omitempty"`
	// example:
	//
	// vsw-8vbno9zxz8j9qiot****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// example:
	//
	// 10.1.1.1
	PrimaryVSwitchIp *string `json:"PrimaryVSwitchIp,omitempty" xml:"PrimaryVSwitchIp,omitempty"`
	// example:
	//
	// cn-shenzhen-d
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// example:
	//
	// Custom
	PrivateDnsType *string `json:"PrivateDnsType,omitempty" xml:"PrivateDnsType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 4E7F94C7-781F-5192-86CF-DB0850****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1.1.1.2
	StandbyDns *string `json:"StandbyDns,omitempty" xml:"StandbyDns,omitempty"`
	// example:
	//
	// vsw-8vb6jk75wfcwnuq****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// example:
	//
	// 10.1.1.2
	StandbyVSwitchIp *string `json:"StandbyVSwitchIp,omitempty" xml:"StandbyVSwitchIp,omitempty"`
	// example:
	//
	// cn-shenzhen-e
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 132
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// vpc-uf6b5lyul0x******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribePrivateDnsEndpointDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetAccessInstanceName() *string {
	return s.AccessInstanceName
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetFirewallType() []*string {
	return s.FirewallType
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetMemberUid() *int64 {
	return s.MemberUid
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPort() *int32 {
	return s.Port
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPrimaryDns() *string {
	return s.PrimaryDns
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPrimaryVSwitchId() *string {
	return s.PrimaryVSwitchId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPrimaryVSwitchIp() *string {
	return s.PrimaryVSwitchIp
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetPrivateDnsType() *string {
	return s.PrivateDnsType
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetStandbyDns() *string {
	return s.StandbyDns
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetStandbyVSwitchIp() *string {
	return s.StandbyVSwitchIp
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetAccessInstanceId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetAccessInstanceName(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.AccessInstanceName = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetAliUid(v int64) *DescribePrivateDnsEndpointDetailResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetEndpointId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.EndpointId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetFirewallType(v []*string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.FirewallType = v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetGmtCreate(v int64) *DescribePrivateDnsEndpointDetailResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetIpProtocol(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.IpProtocol = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetMemberUid(v int64) *DescribePrivateDnsEndpointDetailResponseBody {
	s.MemberUid = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPort(v int32) *DescribePrivateDnsEndpointDetailResponseBody {
	s.Port = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPrimaryDns(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.PrimaryDns = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPrimaryVSwitchId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPrimaryVSwitchIp(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.PrimaryVSwitchIp = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPrimaryZoneId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetPrivateDnsType(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.PrivateDnsType = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetRegionNo(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetRequestId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetStandbyDns(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.StandbyDns = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetStandbyVSwitchId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetStandbyVSwitchIp(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.StandbyVSwitchIp = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetStandbyZoneId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetStatus(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetTaskId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) SetVpcId(v string) *DescribePrivateDnsEndpointDetailResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
