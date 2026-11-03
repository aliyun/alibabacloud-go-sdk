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
	// The IP address information in the whitelist.
	EsIPWhitelist []*string `json:"esIPWhitelist,omitempty" xml:"esIPWhitelist,omitempty" type:"Repeated"`
	// Updates the whitelist configuration of the instance by using a whitelist group. You can update only one whitelist group.
	//
	// 	Notice:  You cannot configure esIPWhitelist and whiteIpGroup at the same time.
	WhiteIpGroup *UpdateWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The modification method. Valid values:
	//
	// - Cover (default): Overwrites the IP addresses in the original whitelist with the value of the ips parameter.
	//
	// - Append: Adds the IP addresses specified by the ips parameter to the original whitelist.
	//
	// - Delete: Removes the IP addresses specified by the ips parameter from the original whitelist. You must retain at least one IP address.
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
	if s.WhiteIpGroup != nil {
		if err := s.WhiteIpGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWhiteIpsRequestWhiteIpGroup struct {
	// The name of the whitelist group. This parameter is required if the whiteIpGroup parameter is specified.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The list of IP addresses in the whitelist group.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// The type of the IP address whitelist. Set the value to **PRIVATE_ES**, which indicates a private network access whitelist.
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
