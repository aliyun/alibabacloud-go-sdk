// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressGroups(v []*AddressGroup) *CreatePrivateAccessApplicationRequest
	GetAddressGroups() []*AddressGroup
	SetAddresses(v []*string) *CreatePrivateAccessApplicationRequest
	GetAddresses() []*string
	SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationRequest
	GetBrowserAccessStatus() *string
	SetConfigMode(v string) *CreatePrivateAccessApplicationRequest
	GetConfigMode() *string
	SetDescription(v string) *CreatePrivateAccessApplicationRequest
	GetDescription() *string
	SetL7Config(v *PAL7Config) *CreatePrivateAccessApplicationRequest
	GetL7Config() *PAL7Config
	SetL7ProxyDomainAutomaticPrefix(v string) *CreatePrivateAccessApplicationRequest
	GetL7ProxyDomainAutomaticPrefix() *string
	SetL7ProxyDomainCustom(v string) *CreatePrivateAccessApplicationRequest
	GetL7ProxyDomainCustom() *string
	SetName(v string) *CreatePrivateAccessApplicationRequest
	GetName() *string
	SetPortRanges(v []*CreatePrivateAccessApplicationRequestPortRanges) *CreatePrivateAccessApplicationRequest
	GetPortRanges() []*CreatePrivateAccessApplicationRequestPortRanges
	SetProtocol(v string) *CreatePrivateAccessApplicationRequest
	GetProtocol() *string
	SetStatus(v string) *CreatePrivateAccessApplicationRequest
	GetStatus() *string
	SetTagIds(v []*string) *CreatePrivateAccessApplicationRequest
	GetTagIds() []*string
	SetUnauthorizedAccessConfig(v *PAApplicationUnauthorizedAccessConfig) *CreatePrivateAccessApplicationRequest
	GetUnauthorizedAccessConfig() *PAApplicationUnauthorizedAccessConfig
}

type CreatePrivateAccessApplicationRequest struct {
	AddressGroups []*AddressGroup `json:"AddressGroups,omitempty" xml:"AddressGroups,omitempty" type:"Repeated"`
	// The addresses of the internal-facing access application. You can specify up to 1,000 addresses.
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The browser access status of the internal-facing access application. After this feature is enabled, you can access internal applications without a client. Default value: **Disabled**. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Disabled
	BrowserAccessStatus *string `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	ConfigMode          *string `json:"ConfigMode,omitempty" xml:"ConfigMode,omitempty"`
	// The description of the internal-facing access application. The description must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces. Chinese characters are supported.
	//
	// example:
	//
	// 这是一条内网访问应用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The browser access mode parameter: the application configuration.
	L7Config *PAL7Config `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
	// The browser access mode parameter: the prefix of the mapped proxy domain name. The prefix must be 3 to 20 characters in length and can contain lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// app-sample
	L7ProxyDomainAutomaticPrefix *string `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	// The browser access mode parameter: the custom proxy domain name. The value must be a valid domain name.
	//
	// example:
	//
	// app1.example.com
	L7ProxyDomainCustom *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// The name of the internal-facing access application. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). Chinese characters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port ranges of the internal-facing access application. You can specify up to 65,535 port ranges. Port ranges cannot be duplicate or overlapping.
	PortRanges []*CreatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The protocol of the internal-facing access application. Valid values:
	//
	// - **All**: all protocols.
	//
	// - **TCP**
	//
	// - **UDP**.
	//
	// This parameter is required.
	//
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The status of the internal-facing access application. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of internal-facing access tags. You can associate up to 6 custom internal-facing access tags with each internal-facing access application.
	TagIds                   []*string                              `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UnauthorizedAccessConfig *PAApplicationUnauthorizedAccessConfig `json:"UnauthorizedAccessConfig,omitempty" xml:"UnauthorizedAccessConfig,omitempty"`
}

func (s CreatePrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationRequest) GetAddressGroups() []*AddressGroup {
	return s.AddressGroups
}

func (s *CreatePrivateAccessApplicationRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreatePrivateAccessApplicationRequest) GetBrowserAccessStatus() *string {
	return s.BrowserAccessStatus
}

func (s *CreatePrivateAccessApplicationRequest) GetConfigMode() *string {
	return s.ConfigMode
}

func (s *CreatePrivateAccessApplicationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePrivateAccessApplicationRequest) GetL7Config() *PAL7Config {
	return s.L7Config
}

func (s *CreatePrivateAccessApplicationRequest) GetL7ProxyDomainAutomaticPrefix() *string {
	return s.L7ProxyDomainAutomaticPrefix
}

func (s *CreatePrivateAccessApplicationRequest) GetL7ProxyDomainCustom() *string {
	return s.L7ProxyDomainCustom
}

func (s *CreatePrivateAccessApplicationRequest) GetName() *string {
	return s.Name
}

func (s *CreatePrivateAccessApplicationRequest) GetPortRanges() []*CreatePrivateAccessApplicationRequestPortRanges {
	return s.PortRanges
}

func (s *CreatePrivateAccessApplicationRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreatePrivateAccessApplicationRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrivateAccessApplicationRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *CreatePrivateAccessApplicationRequest) GetUnauthorizedAccessConfig() *PAApplicationUnauthorizedAccessConfig {
	return s.UnauthorizedAccessConfig
}

func (s *CreatePrivateAccessApplicationRequest) SetAddressGroups(v []*AddressGroup) *CreatePrivateAccessApplicationRequest {
	s.AddressGroups = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetAddresses(v []*string) *CreatePrivateAccessApplicationRequest {
	s.Addresses = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationRequest {
	s.BrowserAccessStatus = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetConfigMode(v string) *CreatePrivateAccessApplicationRequest {
	s.ConfigMode = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetDescription(v string) *CreatePrivateAccessApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetL7Config(v *PAL7Config) *CreatePrivateAccessApplicationRequest {
	s.L7Config = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetL7ProxyDomainAutomaticPrefix(v string) *CreatePrivateAccessApplicationRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetL7ProxyDomainCustom(v string) *CreatePrivateAccessApplicationRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetName(v string) *CreatePrivateAccessApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetPortRanges(v []*CreatePrivateAccessApplicationRequestPortRanges) *CreatePrivateAccessApplicationRequest {
	s.PortRanges = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetProtocol(v string) *CreatePrivateAccessApplicationRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetStatus(v string) *CreatePrivateAccessApplicationRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetTagIds(v []*string) *CreatePrivateAccessApplicationRequest {
	s.TagIds = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetUnauthorizedAccessConfig(v *PAApplicationUnauthorizedAccessConfig) *CreatePrivateAccessApplicationRequest {
	s.UnauthorizedAccessConfig = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) Validate() error {
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
	if s.UnauthorizedAccessConfig != nil {
		if err := s.UnauthorizedAccessConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreatePrivateAccessApplicationRequestPortRanges struct {
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

func (s CreatePrivateAccessApplicationRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationRequestPortRanges) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) SetBegin(v int32) *CreatePrivateAccessApplicationRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) SetEnd(v int32) *CreatePrivateAccessApplicationRequestPortRanges {
	s.End = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) Validate() error {
	return dara.Validate(s)
}
