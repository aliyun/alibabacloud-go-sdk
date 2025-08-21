// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKibanaWhiteIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKibanaIPWhitelist(v []*string) *UpdateKibanaWhiteIpsRequest
	GetKibanaIPWhitelist() []*string
	SetWhiteIpGroup(v *UpdateKibanaWhiteIpsRequestWhiteIpGroup) *UpdateKibanaWhiteIpsRequest
	GetWhiteIpGroup() *UpdateKibanaWhiteIpsRequestWhiteIpGroup
	SetClientToken(v string) *UpdateKibanaWhiteIpsRequest
	GetClientToken() *string
	SetModifyMode(v string) *UpdateKibanaWhiteIpsRequest
	GetModifyMode() *string
}

type UpdateKibanaWhiteIpsRequest struct {
	// The IP address whitelists. This parameter is available if the whiteIpGroup parameter is left empty. The default IP address whitelist is updated based on the value of this parameter.
	//
	// You cannot configure both the kibanaIPWhitelist and whiteIpGroup parameters.
	KibanaIPWhitelist []*string `json:"kibanaIPWhitelist,omitempty" xml:"kibanaIPWhitelist,omitempty" type:"Repeated"`
	// The name of the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	WhiteIpGroup *UpdateKibanaWhiteIpsRequestWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Struct"`
	// The update mode. Valid values:
	//
	// 	- Cover: overwrites the IP addresses in the specified IP address whitelist with the IP addresses specified by using the ips parameter. This is the default value.
	//
	// 	- Append: adds the IP addresses specified by using the ips parameter to the specified IP address whitelist.
	//
	// 	- Delete: deletes the IP addresses specified by using the ips parameter from the specified IP address whitelist. At least one IP address must be retained for the whitelist.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The body of the request.
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"modifyMode,omitempty" xml:"modifyMode,omitempty"`
}

func (s UpdateKibanaWhiteIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsRequest) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsRequest) GetKibanaIPWhitelist() []*string {
	return s.KibanaIPWhitelist
}

func (s *UpdateKibanaWhiteIpsRequest) GetWhiteIpGroup() *UpdateKibanaWhiteIpsRequestWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *UpdateKibanaWhiteIpsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateKibanaWhiteIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *UpdateKibanaWhiteIpsRequest) SetKibanaIPWhitelist(v []*string) *UpdateKibanaWhiteIpsRequest {
	s.KibanaIPWhitelist = v
	return s
}

func (s *UpdateKibanaWhiteIpsRequest) SetWhiteIpGroup(v *UpdateKibanaWhiteIpsRequestWhiteIpGroup) *UpdateKibanaWhiteIpsRequest {
	s.WhiteIpGroup = v
	return s
}

func (s *UpdateKibanaWhiteIpsRequest) SetClientToken(v string) *UpdateKibanaWhiteIpsRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKibanaWhiteIpsRequest) SetModifyMode(v string) *UpdateKibanaWhiteIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *UpdateKibanaWhiteIpsRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateKibanaWhiteIpsRequestWhiteIpGroup struct {
	// The type of the whitelist. Set the value to PUBLIC_KIBANA. This value indicates a public IP address whitelist.
	//
	// example:
	//
	// test_group_name
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The IP addresses in the whitelist. This parameter is required if you configure the whiteIpGroup parameter.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// PUBLIC_KIBANA
	WhiteIpType *string `json:"whiteIpType,omitempty" xml:"whiteIpType,omitempty"`
}

func (s UpdateKibanaWhiteIpsRequestWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s UpdateKibanaWhiteIpsRequestWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) GetWhiteIpType() *string {
	return s.WhiteIpType
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) SetGroupName(v string) *UpdateKibanaWhiteIpsRequestWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) SetIps(v []*string) *UpdateKibanaWhiteIpsRequestWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) SetWhiteIpType(v string) *UpdateKibanaWhiteIpsRequestWhiteIpGroup {
	s.WhiteIpType = &v
	return s
}

func (s *UpdateKibanaWhiteIpsRequestWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
