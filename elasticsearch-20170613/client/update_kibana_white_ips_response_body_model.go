// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaWhiteIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateKibanaWhiteIpsResponseBody
	GetRequestId() *string
	SetResult(v *UpdateKibanaWhiteIpsResponseBodyResult) *UpdateKibanaWhiteIpsResponseBody
	GetResult() *UpdateKibanaWhiteIpsResponseBodyResult
}

type UpdateKibanaWhiteIpsResponseBody struct {
	// The details of the Elasticsearch cluster.
	//
	// example:
	//
	// E5EF11F1-DBAE-4020-AC24-DFA6C4345CAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The private IP address whitelists for access to the Kibana console of the cluster.
	Result *UpdateKibanaWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateKibanaWhiteIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateKibanaWhiteIpsResponseBody) GetResult() *UpdateKibanaWhiteIpsResponseBodyResult {
	return s.Result
}

func (s *UpdateKibanaWhiteIpsResponseBody) SetRequestId(v string) *UpdateKibanaWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBody) SetResult(v *UpdateKibanaWhiteIpsResponseBodyResult) *UpdateKibanaWhiteIpsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKibanaWhiteIpsResponseBodyResult struct {
	// The public IP address whitelists for access to the Kibana console of the cluster.
	KibanaIPWhitelist []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	// The private IP address whitelists for access to the Kibana console of the cluster.
	KibanaPrivateIPWhitelist []*string `json:"kibanaPrivateIPWhitelist,omitempty" xml:"kibanaPrivateIPWhitelist,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC).
	NetworkConfig *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
}

func (s UpdateKibanaWhiteIpsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) GetKibanaIPWhitelist() []*string {
	return s.KibanaIPWhitelist
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) GetKibanaPrivateIPWhitelist() []*string {
	return s.KibanaPrivateIPWhitelist
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) GetNetworkConfig() *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) SetKibanaIPWhitelist(v []*string) *UpdateKibanaWhiteIpsResponseBodyResult {
	s.KibanaIPWhitelist = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) SetKibanaPrivateIPWhitelist(v []*string) *UpdateKibanaWhiteIpsResponseBodyResult {
	s.KibanaPrivateIPWhitelist = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) SetNetworkConfig(v *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) *UpdateKibanaWhiteIpsResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResult) Validate() error {
	if s.NetworkConfig != nil {
		if err := s.NetworkConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig struct {
	// The IP address whitelists.
	//
	// example:
	//
	// vpc
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vpc-bp1jy348ibzulk6hn****
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The network type.
	//
	// example:
	//
	// cn-hangzhou-h
	VsArea *string `json:"vsArea,omitempty" xml:"vsArea,omitempty"`
	// The region ID.
	//
	// example:
	//
	// vsw-bp1a0mifpletdd1da****
	VswitchId *string `json:"vswitchId,omitempty" xml:"vswitchId,omitempty"`
	// The IP address whitelists.
	WhiteIpGroupList []*UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GetType() *string {
	return s.Type
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GetVsArea() *string {
	return s.VsArea
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GetVswitchId() *string {
	return s.VswitchId
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) GetWhiteIpGroupList() []*UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	return s.WhiteIpGroupList
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) SetType(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	s.Type = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) SetVpcId(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	s.VpcId = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) SetVsArea(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	s.VsArea = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) SetVswitchId(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	s.VswitchId = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfig) Validate() error {
	if s.WhiteIpGroupList != nil {
		for _, item := range s.WhiteIpGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// test_group_name
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The IP addresses in the whitelist.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// example:
	//
	// PUBLIC_KIBANA
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetIps() []*string {
	return s.Ips
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

func (s *UpdateKibanaWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) Validate() error {
	return dara.Validate(s)
}
