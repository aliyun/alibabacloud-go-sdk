// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressGroups(v []*AddressGroup) *UpdatePrivateAccessApplicationRequest
	GetAddressGroups() []*AddressGroup
	SetAddresses(v []*string) *UpdatePrivateAccessApplicationRequest
	GetAddresses() []*string
	SetApplicationId(v string) *UpdatePrivateAccessApplicationRequest
	GetApplicationId() *string
	SetConfigMode(v string) *UpdatePrivateAccessApplicationRequest
	GetConfigMode() *string
	SetDescription(v string) *UpdatePrivateAccessApplicationRequest
	GetDescription() *string
	SetL7Config(v *PAL7Config) *UpdatePrivateAccessApplicationRequest
	GetL7Config() *PAL7Config
	SetL7ProxyDomainAutomaticPrefix(v string) *UpdatePrivateAccessApplicationRequest
	GetL7ProxyDomainAutomaticPrefix() *string
	SetL7ProxyDomainCustom(v string) *UpdatePrivateAccessApplicationRequest
	GetL7ProxyDomainCustom() *string
	SetL7ProxyDomainPrivate(v string) *UpdatePrivateAccessApplicationRequest
	GetL7ProxyDomainPrivate() *string
	SetModifyType(v string) *UpdatePrivateAccessApplicationRequest
	GetModifyType() *string
	SetName(v string) *UpdatePrivateAccessApplicationRequest
	GetName() *string
	SetPortRanges(v []*UpdatePrivateAccessApplicationRequestPortRanges) *UpdatePrivateAccessApplicationRequest
	GetPortRanges() []*UpdatePrivateAccessApplicationRequestPortRanges
	SetProtocol(v string) *UpdatePrivateAccessApplicationRequest
	GetProtocol() *string
	SetStatus(v string) *UpdatePrivateAccessApplicationRequest
	GetStatus() *string
	SetTagIds(v []*string) *UpdatePrivateAccessApplicationRequest
	GetTagIds() []*string
}

type UpdatePrivateAccessApplicationRequest struct {
	AddressGroups []*AddressGroup `json:"AddressGroups,omitempty" xml:"AddressGroups,omitempty" type:"Repeated"`
	// The addresses of the office applications. You can enter up to 1,000 addresses of office applications.
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The ID of the office application. You can obtain the value by calling the following operations:
	//
	// 	- [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): queries office applications.
	//
	// 	- [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): creates an office application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ConfigMode    *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The description of the office application. The value must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The browser access mode parameter. The parameter specifies the configurations of Layer 7 applications.
	L7Config *PAL7Config `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
	// The browser access mode parameter. The parameter specifies the prefix of the domain name that the proxy gateway uses. The prefix must be 3 to 20 characters in length, and can contain lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// app1-xxx
	L7ProxyDomainAutomaticPrefix *string `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	// The browser access mode parameter. The parameter specifies the custom domain name of the proxy gateway.
	//
	// example:
	//
	// app1.example.com
	L7ProxyDomainCustom *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// Deprecated
	//
	// 浏览器访问模式参数：私有代理域名。
	//
	// example:
	//
	// app1.example.com
	L7ProxyDomainPrivate *string `json:"L7ProxyDomainPrivate,omitempty" xml:"L7ProxyDomainPrivate,omitempty"`
	// The modification type of the office application. Valid values:
	//
	// 	- **Cover**: uses the values of the **Addresses**, **PortRanges**, and **TagIds*	- parameters to overwrite the original addresses, port ranges, and tag IDs. This is the default value.
	//
	// 	- **Append**: adds the values of the **Addresses**, **PortRanges**, and **TagIds*	- parameters respectively to the original addresses, port ranges, and tag IDs.
	//
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port ranges of the office applications. You can enter up to 65,535 port ranges. Multiple port ranges cannot be duplicated or overlapped.
	PortRanges []*UpdatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The protocol that is used by the office application. Valid values:
	//
	// 	- **All**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The status of the office application. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of the tags for the office applications. You can add up to six custom tags to an office application.
	//
	// if can be null:
	// true
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UpdatePrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationRequest) GetAddressGroups() []*AddressGroup {
	return s.AddressGroups
}

func (s *UpdatePrivateAccessApplicationRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdatePrivateAccessApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePrivateAccessApplicationRequest) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *UpdatePrivateAccessApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePrivateAccessApplicationRequest) GetL7Config() *PAL7Config {
	return s.L7Config
}

func (s *UpdatePrivateAccessApplicationRequest) GetL7ProxyDomainAutomaticPrefix() *string {
	return s.L7ProxyDomainAutomaticPrefix
}

func (s *UpdatePrivateAccessApplicationRequest) GetL7ProxyDomainCustom() *string {
	return s.L7ProxyDomainCustom
}

func (s *UpdatePrivateAccessApplicationRequest) GetL7ProxyDomainPrivate() *string {
	return s.L7ProxyDomainPrivate
}

func (s *UpdatePrivateAccessApplicationRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdatePrivateAccessApplicationRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePrivateAccessApplicationRequest) GetPortRanges() []*UpdatePrivateAccessApplicationRequestPortRanges {
	return s.PortRanges
}

func (s *UpdatePrivateAccessApplicationRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdatePrivateAccessApplicationRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrivateAccessApplicationRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdatePrivateAccessApplicationRequest) SetAddressGroups(v []*AddressGroup) *UpdatePrivateAccessApplicationRequest {
	s.AddressGroups = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetAddresses(v []*string) *UpdatePrivateAccessApplicationRequest {
	s.Addresses = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetApplicationId(v string) *UpdatePrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetConfigMode(v string) *UpdatePrivateAccessApplicationRequest {
	s.ConfigMode = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetDescription(v string) *UpdatePrivateAccessApplicationRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7Config(v *PAL7Config) *UpdatePrivateAccessApplicationRequest {
	s.L7Config = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainAutomaticPrefix(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainCustom(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainPrivate(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainPrivate = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetModifyType(v string) *UpdatePrivateAccessApplicationRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetName(v string) *UpdatePrivateAccessApplicationRequest {
	s.Name = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetPortRanges(v []*UpdatePrivateAccessApplicationRequestPortRanges) *UpdatePrivateAccessApplicationRequest {
	s.PortRanges = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetProtocol(v string) *UpdatePrivateAccessApplicationRequest {
	s.Protocol = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetStatus(v string) *UpdatePrivateAccessApplicationRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetTagIds(v []*string) *UpdatePrivateAccessApplicationRequest {
	s.TagIds = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) Validate() error {
	if s.AddressGroups != nil {
		for _, item := range s.AddressGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.L7Config != nil {
		if err := s.L7Config.Validate(); err != nil {
			return err
		}
	}
	if s.PortRanges != nil {
		for _, item := range s.PortRanges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrivateAccessApplicationRequestPortRanges struct {
	// The start port. The start port must be less than or equal to the end port.
	//
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The end port. The end port must be greater than or equal to the start port.
	//
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdatePrivateAccessApplicationRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) SetBegin(v int32) *UpdatePrivateAccessApplicationRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) SetEnd(v int32) *UpdatePrivateAccessApplicationRequestPortRanges {
	s.End = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) Validate() error {
	return dara.Validate(s)
}
