// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressGroups(v []*AddressGroup) *CreatePrivateAccessApplicationShrinkRequest
	GetAddressGroups() []*AddressGroup
	SetAddresses(v []*string) *CreatePrivateAccessApplicationShrinkRequest
	GetAddresses() []*string
	SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetBrowserAccessStatus() *string
	SetConfigMode(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetConfigMode() *string
	SetDescription(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetDescription() *string
	SetHealthCheckConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetHealthCheckConfigShrink() *string
	SetHealthCheckMode(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetHealthCheckMode() *string
	SetL7ConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetL7ConfigShrink() *string
	SetL7ProxyDomainAutomaticPrefix(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetL7ProxyDomainAutomaticPrefix() *string
	SetL7ProxyDomainCustom(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetL7ProxyDomainCustom() *string
	SetName(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetName() *string
	SetPortRanges(v []*CreatePrivateAccessApplicationShrinkRequestPortRanges) *CreatePrivateAccessApplicationShrinkRequest
	GetPortRanges() []*CreatePrivateAccessApplicationShrinkRequestPortRanges
	SetProtocol(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetProtocol() *string
	SetStatus(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetStatus() *string
	SetTagIds(v []*string) *CreatePrivateAccessApplicationShrinkRequest
	GetTagIds() []*string
	SetUnauthorizedAccessConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetUnauthorizedAccessConfigShrink() *string
}

type CreatePrivateAccessApplicationShrinkRequest struct {
	// The application address groups. This parameter is required when ConfigMode is set to Precise, and cannot be specified when ConfigMode is an empty string.
	AddressGroups []*AddressGroup `json:"AddressGroups,omitempty" xml:"AddressGroups,omitempty" type:"Repeated"`
	// The collection of private access application addresses. You can enter up to 1,000 private access application addresses.
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The browser access status of the private access application. After you enable this feature, you can access private access applications without a client. Default value: **Disabled**. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Disabled
	BrowserAccessStatus *string `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	// The configuration mode. Valid values:
	//
	// 	- Empty string: Default mode.
	//
	// 	- Precise: Precise mode.
	//
	// example:
	//
	// Precise
	ConfigMode *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The description of the private access application. The description must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// example:
	//
	// This is a private access application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The health check configuration. This parameter is required when HealthCheckMode is set to Continuous, and cannot be specified when HealthCheckMode is set to OnAccess.
	HealthCheckConfigShrink *string `json:"HealthCheckConfig,omitempty" xml:"HealthCheckConfig,omitempty"`
	// The health check mode. Valid values:
	//
	// 	- OnAccess: Checks upon access without proactive probing.
	//
	// 	- Continuous: Performs continuous checks and proactive probing based on the configurations in HealthCheckConfig.
	//
	// If this parameter is not specified, the default value is OnAccess. When the application addresses include wildcard domain names or CIDR blocks, or when the number of application addresses exceeds 10, only OnAccess is supported.
	//
	// example:
	//
	// OnAccess
	HealthCheckMode *string `json:"HealthCheckMode,omitempty" xml:"HealthCheckMode,omitempty"`
	// The application configuration in browser access mode.
	L7ConfigShrink *string `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
	// The prefix of the mapped proxy domain name in browser access mode. The prefix must be 3 to 20 characters in length, and can contain lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// app-sample
	L7ProxyDomainAutomaticPrefix *string `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	// The custom proxy domain name in browser access mode. The value must be a valid domain name.
	//
	// example:
	//
	// app1.example.com
	L7ProxyDomainCustom *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// The name of the private access application. The name must be 1 to 128 characters in length, and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of port ranges for the private access application. You can enter up to 65,535 port ranges. Multiple port ranges cannot be duplicate or overlap.
	PortRanges []*CreatePrivateAccessApplicationShrinkRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The protocol of the private access application. Valid values:
	//
	// - **All**: All protocols.
	//
	// - **TCP**
	//
	// - **UDP**
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The status of the private access application. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of private access tag IDs. You can associate up to six custom private access tags with a single private access application.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// The configuration for unauthorized application access requests.
	UnauthorizedAccessConfigShrink *string `json:"UnauthorizedAccessConfig,omitempty" xml:"UnauthorizedAccessConfig,omitempty"`
}

func (s CreatePrivateAccessApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetAddressGroups() []*AddressGroup {
	return s.AddressGroups
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetBrowserAccessStatus() *string {
	return s.BrowserAccessStatus
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetHealthCheckConfigShrink() *string {
	return s.HealthCheckConfigShrink
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetHealthCheckMode() *string {
	return s.HealthCheckMode
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetL7ConfigShrink() *string {
	return s.L7ConfigShrink
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetL7ProxyDomainAutomaticPrefix() *string {
	return s.L7ProxyDomainAutomaticPrefix
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetL7ProxyDomainCustom() *string {
	return s.L7ProxyDomainCustom
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetPortRanges() []*CreatePrivateAccessApplicationShrinkRequestPortRanges {
	return s.PortRanges
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetUnauthorizedAccessConfigShrink() *string {
	return s.UnauthorizedAccessConfigShrink
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetAddressGroups(v []*AddressGroup) *CreatePrivateAccessApplicationShrinkRequest {
	s.AddressGroups = v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetAddresses(v []*string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Addresses = v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.BrowserAccessStatus = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetConfigMode(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.ConfigMode = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetDescription(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetHealthCheckConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.HealthCheckConfigShrink = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetHealthCheckMode(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.HealthCheckMode = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetL7ConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.L7ConfigShrink = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetL7ProxyDomainAutomaticPrefix(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetL7ProxyDomainCustom(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetName(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetPortRanges(v []*CreatePrivateAccessApplicationShrinkRequestPortRanges) *CreatePrivateAccessApplicationShrinkRequest {
	s.PortRanges = v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetProtocol(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetStatus(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetTagIds(v []*string) *CreatePrivateAccessApplicationShrinkRequest {
	s.TagIds = v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetUnauthorizedAccessConfigShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.UnauthorizedAccessConfigShrink = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) Validate() error {
	if s.AddressGroups != nil {
		for _, item := range s.AddressGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreatePrivateAccessApplicationShrinkRequestPortRanges struct {
	// The start port. The value must be less than or equal to the end port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The end port. The value must be greater than or equal to the start port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s CreatePrivateAccessApplicationShrinkRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationShrinkRequestPortRanges) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationShrinkRequestPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *CreatePrivateAccessApplicationShrinkRequestPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *CreatePrivateAccessApplicationShrinkRequestPortRanges) SetBegin(v int32) *CreatePrivateAccessApplicationShrinkRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequestPortRanges) SetEnd(v int32) *CreatePrivateAccessApplicationShrinkRequestPortRanges {
	s.End = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequestPortRanges) Validate() error {
	return dara.Validate(s)
}
