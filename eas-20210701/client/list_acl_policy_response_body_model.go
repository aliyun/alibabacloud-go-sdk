// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAclPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListAclPolicyResponseBody
	GetGatewayId() *string
	SetInternetAclPolicyList(v []*ListAclPolicyResponseBodyInternetAclPolicyList) *ListAclPolicyResponseBody
	GetInternetAclPolicyList() []*ListAclPolicyResponseBodyInternetAclPolicyList
	SetIntranetVpcAclPolicyList(v []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList) *ListAclPolicyResponseBody
	GetIntranetVpcAclPolicyList() []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList
	SetRequestId(v string) *ListAclPolicyResponseBody
	GetRequestId() *string
}

type ListAclPolicyResponseBody struct {
	// The private gateway ID.
	//
	// example:
	//
	// gw-1uhcqmsc7x22******
	GatewayId *string `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The access control policies of the private gateway over the Internet.
	InternetAclPolicyList []*ListAclPolicyResponseBodyInternetAclPolicyList `json:"InternetAclPolicyList,omitempty" xml:"InternetAclPolicyList,omitempty" type:"Repeated"`
	// The access control policies of the private gateway over the internal network.
	IntranetVpcAclPolicyList []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList `json:"IntranetVpcAclPolicyList,omitempty" xml:"IntranetVpcAclPolicyList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAclPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBody) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAclPolicyResponseBody) GetInternetAclPolicyList() []*ListAclPolicyResponseBodyInternetAclPolicyList {
	return s.InternetAclPolicyList
}

func (s *ListAclPolicyResponseBody) GetIntranetVpcAclPolicyList() []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList {
	return s.IntranetVpcAclPolicyList
}

func (s *ListAclPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAclPolicyResponseBody) SetGatewayId(v string) *ListAclPolicyResponseBody {
	s.GatewayId = &v
	return s
}

func (s *ListAclPolicyResponseBody) SetInternetAclPolicyList(v []*ListAclPolicyResponseBodyInternetAclPolicyList) *ListAclPolicyResponseBody {
	s.InternetAclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBody) SetIntranetVpcAclPolicyList(v []*ListAclPolicyResponseBodyIntranetVpcAclPolicyList) *ListAclPolicyResponseBody {
	s.IntranetVpcAclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBody) SetRequestId(v string) *ListAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAclPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAclPolicyResponseBodyInternetAclPolicyList struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway over the Internet.
	AclPolicyList []*ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
}

func (s ListAclPolicyResponseBodyInternetAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponseBodyInternetAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyList) GetAclPolicyList() []*ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList {
	return s.AclPolicyList
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyList) SetAclPolicyList(v []*ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) *ListAclPolicyResponseBodyInternetAclPolicyList {
	s.AclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyList) Validate() error {
	return dara.Validate(s)
}

type ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway over the Internet.
	//
	// example:
	//
	// default
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway over the Internet.
	//
	// example:
	//
	// 10.23.XX.XX/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) GetComment() *string {
	return s.Comment
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) GetEntry() *string {
	return s.Entry
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) SetComment(v string) *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList {
	s.Comment = &v
	return s
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) SetEntry(v string) *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList {
	s.Entry = &v
	return s
}

func (s *ListAclPolicyResponseBodyInternetAclPolicyListAclPolicyList) Validate() error {
	return dara.Validate(s)
}

type ListAclPolicyResponseBodyIntranetVpcAclPolicyList struct {
	// The whitelisted IP CIDR blocks in the VPC that can access the private gateway over the internal network.
	AclPolicyList []*ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList `json:"AclPolicyList,omitempty" xml:"AclPolicyList,omitempty" type:"Repeated"`
	// The VPC ID. For more information about how to obtain the VPC ID, see DescribeVpcs.
	//
	// example:
	//
	// vpc-uf66uio7md****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) GetAclPolicyList() []*ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList {
	return s.AclPolicyList
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) SetAclPolicyList(v []*ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) *ListAclPolicyResponseBodyIntranetVpcAclPolicyList {
	s.AclPolicyList = v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) SetVpcId(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyList {
	s.VpcId = &v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyList) Validate() error {
	return dara.Validate(s)
}

type ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList struct {
	// The comment on the IP CIDR block in the VPC that can access the private gateway over the internal network.
	//
	// example:
	//
	// Test Entry
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IP CIDR block in the VPC that can access the private gateway over the internal network.
	//
	// example:
	//
	// 192.168.XX.XX/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) String() string {
	return dara.Prettify(s)
}

func (s ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) GoString() string {
	return s.String()
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) GetComment() *string {
	return s.Comment
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) GetEntry() *string {
	return s.Entry
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) SetComment(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList {
	s.Comment = &v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) SetEntry(v string) *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList {
	s.Entry = &v
	return s
}

func (s *ListAclPolicyResponseBodyIntranetVpcAclPolicyListAclPolicyList) Validate() error {
	return dara.Validate(s)
}
