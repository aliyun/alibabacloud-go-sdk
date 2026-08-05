// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfflineTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetwork(v *ModifyOfflineTaskLogRequestNetwork) *ModifyOfflineTaskLogRequest
	GetNetwork() *ModifyOfflineTaskLogRequestNetwork
	SetRegionId(v string) *ModifyOfflineTaskLogRequest
	GetRegionId() *string
}

type ModifyOfflineTaskLogRequest struct {
	// The network configuration for enabling or disabling network access.
	Network *ModifyOfflineTaskLogRequestNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ModifyOfflineTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequest) GetNetwork() *ModifyOfflineTaskLogRequestNetwork {
	return s.Network
}

func (s *ModifyOfflineTaskLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfflineTaskLogRequest) SetNetwork(v *ModifyOfflineTaskLogRequestNetwork) *ModifyOfflineTaskLogRequest {
	s.Network = v
	return s
}

func (s *ModifyOfflineTaskLogRequest) SetRegionId(v string) *ModifyOfflineTaskLogRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfflineTaskLogRequest) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOfflineTaskLogRequestNetwork struct {
	// The ES private network information.
	PrivateEs *ModifyOfflineTaskLogRequestNetworkPrivateEs `json:"privateEs,omitempty" xml:"privateEs,omitempty" type:"Struct"`
	// **The ES public network information.**
	PublicEs *ModifyOfflineTaskLogRequestNetworkPublicEs `json:"publicEs,omitempty" xml:"publicEs,omitempty" type:"Struct"`
}

func (s ModifyOfflineTaskLogRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequestNetwork) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequestNetwork) GetPrivateEs() *ModifyOfflineTaskLogRequestNetworkPrivateEs {
	return s.PrivateEs
}

func (s *ModifyOfflineTaskLogRequestNetwork) GetPublicEs() *ModifyOfflineTaskLogRequestNetworkPublicEs {
	return s.PublicEs
}

func (s *ModifyOfflineTaskLogRequestNetwork) SetPrivateEs(v *ModifyOfflineTaskLogRequestNetworkPrivateEs) *ModifyOfflineTaskLogRequestNetwork {
	s.PrivateEs = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetwork) SetPublicEs(v *ModifyOfflineTaskLogRequestNetworkPublicEs) *ModifyOfflineTaskLogRequestNetwork {
	s.PublicEs = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetwork) Validate() error {
	if s.PrivateEs != nil {
		if err := s.PrivateEs.Validate(); err != nil {
			return err
		}
	}
	if s.PublicEs != nil {
		if err := s.PublicEs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyOfflineTaskLogRequestNetworkPrivateEs struct {
	// Specifies whether to enable or disable private network access.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The IP whitelist groups.
	WhiteIpGroup []*ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s ModifyOfflineTaskLogRequestNetworkPrivateEs) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequestNetworkPrivateEs) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEs) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEs) GetWhiteIpGroup() []*ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEs) SetEnabled(v bool) *ModifyOfflineTaskLogRequestNetworkPrivateEs {
	s.Enabled = &v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEs) SetWhiteIpGroup(v []*ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) *ModifyOfflineTaskLogRequestNetworkPrivateEs {
	s.WhiteIpGroup = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEs) Validate() error {
	if s.WhiteIpGroup != nil {
		for _, item := range s.WhiteIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup struct {
	// The name of the IP whitelist group.
	//
	// example:
	//
	// kevintest
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// The IP whitelist.
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) SetGroupName(v string) *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) SetIps(v []*string) *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPrivateEsWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}

type ModifyOfflineTaskLogRequestNetworkPublicEs struct {
	// **Specifies whether to enable or disable public network access.**
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// **The IP whitelist group information.**
	WhiteIpGroup []*ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup `json:"whiteIpGroup,omitempty" xml:"whiteIpGroup,omitempty" type:"Repeated"`
}

func (s ModifyOfflineTaskLogRequestNetworkPublicEs) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequestNetworkPublicEs) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEs) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEs) GetWhiteIpGroup() []*ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup {
	return s.WhiteIpGroup
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEs) SetEnabled(v bool) *ModifyOfflineTaskLogRequestNetworkPublicEs {
	s.Enabled = &v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEs) SetWhiteIpGroup(v []*ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) *ModifyOfflineTaskLogRequestNetworkPublicEs {
	s.WhiteIpGroup = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEs) Validate() error {
	if s.WhiteIpGroup != nil {
		for _, item := range s.WhiteIpGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup struct {
	// **The name of the IP whitelist group.**
	//
	// example:
	//
	// kevintest
	GroupName *string `json:"groupName,omitempty" xml:"groupName,omitempty"`
	// **The IP whitelist.**
	Ips []*string `json:"ips,omitempty" xml:"ips,omitempty" type:"Repeated"`
}

func (s ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) GoString() string {
	return s.String()
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) GetIps() []*string {
	return s.Ips
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) SetGroupName(v string) *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup {
	s.GroupName = &v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) SetIps(v []*string) *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup {
	s.Ips = v
	return s
}

func (s *ModifyOfflineTaskLogRequestNetworkPublicEsWhiteIpGroup) Validate() error {
	return dara.Validate(s)
}
