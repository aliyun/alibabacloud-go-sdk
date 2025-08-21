// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEsIPWhitelist(v []*string) *UpdateWhiteIpsRequest
	GetEsIPWhitelist() []*string
	SetWhiteIpGroup(v *UpdateWhiteIpsRequestWhiteIpGroup) *UpdateWhiteIpsRequest
	GetWhiteIpGroup() *UpdateWhiteIpsRequestWhiteIpGroup
	SetClientToken(v string) *UpdateWhiteIpsRequest
	GetClientToken() *string
	SetModifyMode(v string) *UpdateWhiteIpsRequest
	GetModifyMode() *string
}

type UpdateWhiteIpsRequest struct {
	// The name of the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	EsIPWhitelist []*string `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	// The IP addresses in the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	WhiteIpGroup *UpdateWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The IP addresses in the whitelist. This parameter is available if the whiteIpGroup parameter is left empty. The default IP address whitelist is updated based on the value of this parameter.
	//
	// >  You cannot configure both the esIPWhitelist and whiteIpGroup parameters.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdateWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsRequest) GetEsIPWhitelist() []*string {
	return s.EsIPWhitelist
}

func (s *UpdateWhiteIpsRequest) GetWhiteIpGroup() *UpdateWhiteIpsRequestWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdateWhiteIpsRequest) SetEsIPWhitelist(v []*string) *UpdateWhiteIpsRequest {
	s.EsIPWhitelist = v
	return s
}

func (s *UpdateWhiteIpsRequest) SetWhiteIpGroup(v *UpdateWhiteIpsRequestWhiteIpGroup) *UpdateWhiteIpsRequest {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateWhiteIpsRequest) SetClientToken(v string) *UpdateWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWhiteIpsRequest) SetModifyMode(v string) *UpdateWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *UpdateWhiteIpsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateWhiteIpsRequestWhiteIpGroup struct {
	// The type of the whitelist. Set the value to **PRIVATE_ES**. This value indicates a private IP address whitelist.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The returned result.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// The returned result.
	//
	// example:
	//
	// PRIVATE_ES
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateWhiteIpsRequestWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsRequestWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetGroupName(v string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetIps(v []*string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) SetWhiteIpType(v string) *UpdateWhiteIpsRequestWhiteIpGroup {
	s.WhiteIpType = &v
	return s
}

func (s *UpdateWhiteIpsRequestWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
