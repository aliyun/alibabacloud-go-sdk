// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallIPSWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVpcFirewallIPSWhitelistResponseBody
	GetRequestId() *string
	SetWhitelists(v []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) *DescribeVpcFirewallIPSWhitelistResponseBody
	GetWhitelists() []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists
}

type DescribeVpcFirewallIPSWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B5EE02F9-4F21-56CA-AA49-F9F8D69483C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the IPS whitelist of the VPC firewall.
	Whitelists []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallIPSWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) GetWhitelists() []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	return s.Whitelists
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) SetRequestId(v string) *DescribeVpcFirewallIPSWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) SetWhitelists(v []*DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) *DescribeVpcFirewallIPSWhitelistResponseBody {
	s.Whitelists = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists struct {
	// The type of the list. Valid values:
	//
	// 	- **1**: user-defined
	//
	// 	- **2**: address book
	//
	// example:
	//
	// 1
	ListType *int64 `json:"ListType,omitempty" xml:"ListType,omitempty"`
	// The entries in the list.
	//
	// example:
	//
	// 10.10.200.4/32,10.10.200.25/32
	ListValue *string `json:"ListValue,omitempty" xml:"ListValue,omitempty"`
	// The instance ID of the VPC firewall.
	//
	// example:
	//
	// vfw-57431e9abe424852a578
	VpcFirewallId *string `json:"VpcFirewallId,omitempty" xml:"VpcFirewallId,omitempty"`
	// An array of entries in the list.
	WhiteListValue []*string `json:"WhiteListValue,omitempty" xml:"WhiteListValue,omitempty" type:"Repeated"`
	// The type of the whitelist. Valid values:
	//
	// 	- **1**: destination
	//
	// 	- **2**: source
	//
	// example:
	//
	// 1
	WhiteType *int64 `json:"WhiteType,omitempty" xml:"WhiteType,omitempty"`
}

func (s DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GetListType() *int64 {
	return s.ListType
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GetListValue() *string {
	return s.ListValue
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GetVpcFirewallId() *string {
	return s.VpcFirewallId
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GetWhiteListValue() []*string {
	return s.WhiteListValue
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) GetWhiteType() *int64 {
	return s.WhiteType
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetListType(v int64) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.ListType = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetListValue(v string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.ListValue = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetVpcFirewallId(v string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.VpcFirewallId = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetWhiteListValue(v []*string) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.WhiteListValue = v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) SetWhiteType(v int64) *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists {
	s.WhiteType = &v
	return s
}

func (s *DescribeVpcFirewallIPSWhitelistResponseBodyWhitelists) Validate() error {
	return dara.Validate(s)
}
