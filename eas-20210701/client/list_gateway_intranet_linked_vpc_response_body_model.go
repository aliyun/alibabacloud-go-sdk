// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIntranetLinkedVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListGatewayIntranetLinkedVpcResponseBody
	GetGatewayId() *string
	SetIntranetLinkedVpcList(v []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) *ListGatewayIntranetLinkedVpcResponseBody
	GetIntranetLinkedVpcList() []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList
	SetRequestId(v string) *ListGatewayIntranetLinkedVpcResponseBody
	GetRequestId() *string
}

type ListGatewayIntranetLinkedVpcResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The internal endpoints.
	IntranetLinkedVpcList []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList `json:"IntranetLinkedVpcList,omitempty" xml:"IntranetLinkedVpcList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) GetIntranetLinkedVpcList() []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	return s.IntranetLinkedVpcList
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetGatewayId(v string) *ListGatewayIntranetLinkedVpcResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetIntranetLinkedVpcList(v []*ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) *ListGatewayIntranetLinkedVpcResponseBody {
	s.IntranetLinkedVpcList = v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) SetRequestId(v string) *ListGatewayIntranetLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList struct {
	AccountId               *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AuthoritativeDnsEnabled *bool   `json:"AuthoritativeDnsEnabled,omitempty" xml:"AuthoritativeDnsEnabled,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.10.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze4pgstgszvgq******
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The state of the private gateway.
	//
	// Valid values:
	//
	// 	- Creating
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The private gateway is being created.
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     The private gateway is running.
	//
	//     <!-- -->
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-8vb2qjoiio6m9pg******
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// example:
	//
	// vpc-2zetuli9ws0qgjd******
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GoString() string {
	return s.String()
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetAccountId() *string {
	return s.AccountId
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetAuthoritativeDnsEnabled() *bool {
	return s.AuthoritativeDnsEnabled
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetIp() *string {
	return s.Ip
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetStatus() *string {
	return s.Status
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetAccountId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.AccountId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetAuthoritativeDnsEnabled(v bool) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.AuthoritativeDnsEnabled = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetIp(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.Ip = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetSecurityGroupId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.SecurityGroupId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetStatus(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.Status = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetVSwitchId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.VSwitchId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) SetVpcId(v string) *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList {
	s.VpcId = &v
	return s
}

func (s *ListGatewayIntranetLinkedVpcResponseBodyIntranetLinkedVpcList) Validate() error {
	return dara.Validate(s)
}
