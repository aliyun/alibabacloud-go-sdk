// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWhiteIpsResponseBody
	GetRequestId() *string
	SetResult(v *UpdateWhiteIpsResponseBodyResult) *UpdateWhiteIpsResponseBody
	GetResult() *UpdateWhiteIpsResponseBodyResult
}

type UpdateWhiteIpsResponseBody struct {
	// The updated whitelist.
	//
	// example:
	//
	// 8D58B014-BBD7-4D80-B219-00B9D5C6860C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The network configurations.
	Result *UpdateWhiteIpsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s UpdateWhiteIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWhiteIpsResponseBody) GetResult() *UpdateWhiteIpsResponseBodyResult {
	return s.Result
}

func (s *UpdateWhiteIpsResponseBody) SetRequestId(v string) *UpdateWhiteIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWhiteIpsResponseBody) SetResult(v *UpdateWhiteIpsResponseBodyResult) *UpdateWhiteIpsResponseBody {
	s.Result = v
	return s
}

func (s *UpdateWhiteIpsResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateWhiteIpsResponseBodyResult struct {
	// The list of whitelists.
	EsIPWhitelist []*string `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	// The name of the whitelist. By default, the default whitelist is included.
	NetworkConfig *UpdateWhiteIpsResponseBodyResultNetworkConfig `json:"networkConfig,omitempty" xml:"networkConfig,omitempty" type:"Struct"`
}

func (s UpdateWhiteIpsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResult) GetEsIPWhitelist() []*string {
	return s.EsIPWhitelist
}

func (s *UpdateWhiteIpsResponseBodyResult) GetNetworkConfig() *UpdateWhiteIpsResponseBodyResultNetworkConfig {
	return s.NetworkConfig
}

func (s *UpdateWhiteIpsResponseBodyResult) SetEsIPWhitelist(v []*string) *UpdateWhiteIpsResponseBodyResult {
	s.EsIPWhitelist = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResult) SetNetworkConfig(v *UpdateWhiteIpsResponseBodyResultNetworkConfig) *UpdateWhiteIpsResponseBodyResult {
	s.NetworkConfig = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type UpdateWhiteIpsResponseBodyResultNetworkConfig struct {
	// The IP addresses in the whitelist.
	WhiteIpGroupList []*UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList `json:"whiteIpGroupList,omitempty" xml:"whiteIpGroupList,omitempty" type:"Repeated"`
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfig) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfig) GetWhiteIpGroupList() []*UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	return s.WhiteIpGroupList
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfig) SetWhiteIpGroupList(v []*UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) *UpdateWhiteIpsResponseBodyResultNetworkConfig {
	s.WhiteIpGroupList = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList struct {
	// The type of the whitelist. The value of this parameter is fixed as PRIVATE_ES, which indicates a private IP address whitelist.
	//
	// example:
	//
	// test_group
	GroupName *string   `json:"groupName,omitempty" xml:"groupName,omitempty"`
	Ips       []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// example:
	//
	// PRIVATE_ES
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetIps() []*string {
	return s.Ips
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetGroupName(v string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.GroupName = &v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetIps(v []*string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.Ips = v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) SetWhiteIpType(v string) *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList {
	s.WhiteIpType = &v
	return s
}

func (s *UpdateWhiteIpsResponseBodyResultNetworkConfigWhiteIpGroupList) Validate() error {
	return dara.Validate(s)
}
