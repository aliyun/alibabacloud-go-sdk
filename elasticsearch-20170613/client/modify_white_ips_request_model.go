// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModifyMode(v string) *ModifyWhiteIpsRequest
	GetModifyMode() *string
	SetNetworkType(v string) *ModifyWhiteIpsRequest
	GetNetworkType() *string
	SetNodeType(v string) *ModifyWhiteIpsRequest
	GetNodeType() *string
	SetWhiteIpGroup(v *ModifyWhiteIpsRequestWhiteIpGroup) *ModifyWhiteIpsRequest
	GetWhiteIpGroup() *ModifyWhiteIpsRequestWhiteIpGroup
	SetWhiteIpList(v []*string) *ModifyWhiteIpsRequest
	GetWhiteIpList() []*string
	SetClientToken(v string) *ModifyWhiteIpsRequest
	GetClientToken() *string
}

type ModifyWhiteIpsRequest struct {
	// The information about the IP address whitelist that you want to update. You can specify only one whitelist.
	//
	// > You cannot configure both the whiteIpList and whiteIpGroup parameters.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
	// The IP addresses in the whitelist. This parameter is available if the whiteIpGroup parameter is left empty. The default IP address whitelist is updated based on the value of this parameter.
	//
	// > You cannot configure both the whiteIpList and whiteIpGroup parameters.
	//
	// example:
	//
	// PUBLIC
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The IP addresses in the whitelist. This parameter is available if the whiteIpGroup parameter is left empty. The default IP address whitelist is updated based on the value of this parameter.
	//
	// example:
	//
	// WORKER
	NodeType *string `json:"nodeType,omitempty" xml:"nodeType,omitempty"`
	// The IP addresses in the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	WhiteIpGroup *ModifyWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	// The name of the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	WhiteIpList []*string `json:"whiteIpList,omitempty" xml:"whiteIpList,omitempty" type:"Repeated"`
	// The network type. This parameter is required if you configure the whiteIpList parameter. Valid values:
	//
	// 	- PRIVATE
	//
	// 	- PUBLIC
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ModifyWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyWhiteIpsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ModifyWhiteIpsRequest) GetNodeType() *string {
	return s.NodeType
}

func (s *ModifyWhiteIpsRequest) GetWhiteIpGroup() *ModifyWhiteIpsRequestWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *ModifyWhiteIpsRequest) GetWhiteIpList() []*string {
	return s.WhiteIpList
}

func (s *ModifyWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyWhiteIpsRequest) SetModifyMode(v string) *ModifyWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetNetworkType(v string) *ModifyWhiteIpsRequest {
	s.NetworkType = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetNodeType(v string) *ModifyWhiteIpsRequest {
	s.NodeType = &v
	return s
}

func (s *ModifyWhiteIpsRequest) SetWhiteIpGroup(v *ModifyWhiteIpsRequestWhiteIpGroup) *ModifyWhiteIpsRequest {
	s.WhiteIpGroup = v
	return s
}

func (s *ModifyWhiteIpsRequest) SetWhiteIpList(v []*string) *ModifyWhiteIpsRequest {
	s.WhiteIpList = v
	return s
}

func (s *ModifyWhiteIpsRequest) SetClientToken(v string) *ModifyWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyWhiteIpsRequest) Validate() error {
	if s.WhiteIpGroup != nil {
		if err := s.WhiteIpGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyWhiteIpsRequestWhiteIpGroup struct {
	// The type of the IP address whitelist. Valid values:
	//
	// 	- PRIVATE_KIBANA
	//
	// 	- PRIVATE_ES
	//
	// 	- PUBLIC_ES
	//
	// 	- PUBLIC_KIBANA
	//
	// example:
	//
	// test_group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The returned result.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// PRIVATE_ES
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s ModifyWhiteIpsRequestWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyWhiteIpsRequestWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) SetGroupName(v string) *ModifyWhiteIpsRequestWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) SetIps(v []*string) *ModifyWhiteIpsRequestWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) SetWhiteIpType(v string) *ModifyWhiteIpsRequestWhiteIpGroup {
	s.WhiteIpType = &v
	return s
}

func (s *ModifyWhiteIpsRequestWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
