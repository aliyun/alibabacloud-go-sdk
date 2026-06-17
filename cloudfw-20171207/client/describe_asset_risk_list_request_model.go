// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetRiskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpAddrList(v []*string) *DescribeAssetRiskListRequest
	GetIpAddrList() []*string
	SetIpVersion(v int32) *DescribeAssetRiskListRequest
	GetIpVersion() *int32
	SetLang(v string) *DescribeAssetRiskListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeAssetRiskListRequest
	GetSourceIp() *string
}

type DescribeAssetRiskListRequest struct {
	// The IP addresses to query. Separate multiple IP addresses with commas (,). You can query a maximum of 20 IP addresses at a time.
	//
	// > - Example of an IPv4 address: 47.97.XX.XX.
	//
	// >
	//
	// > - Example of an IPv6 address: 2001:db8:ffff:ffff:ffff:XXXX:ffff.
	IpAddrList []*string `json:"IpAddrList,omitempty" xml:"IpAddrList,omitempty" type:"Repeated"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// - **4*	- (default): IPv4
	//
	// - **6**: IPv6
	//
	// This parameter is required.
	//
	// example:
	//
	// 4
	IpVersion *int32 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the visitor.
	//
	// example:
	//
	// 47.100.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeAssetRiskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetRiskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListRequest) GetIpAddrList() []*string {
	return s.IpAddrList
}

func (s *DescribeAssetRiskListRequest) GetIpVersion() *int32 {
	return s.IpVersion
}

func (s *DescribeAssetRiskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAssetRiskListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeAssetRiskListRequest) SetIpAddrList(v []*string) *DescribeAssetRiskListRequest {
	s.IpAddrList = v
	return s
}

func (s *DescribeAssetRiskListRequest) SetIpVersion(v int32) *DescribeAssetRiskListRequest {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetRiskListRequest) SetLang(v string) *DescribeAssetRiskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAssetRiskListRequest) SetSourceIp(v string) *DescribeAssetRiskListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeAssetRiskListRequest) Validate() error {
	return dara.Validate(s)
}
