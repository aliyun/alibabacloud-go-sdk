// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *CreatePrivateAccessApplicationRequest
	GetAddresses() []*string
	SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationRequest
	GetBrowserAccessStatus() *string
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
}

type CreatePrivateAccessApplicationRequest struct {
	// The addresses of the office applications. You can enter up to 1,000 addresses of office applications.
	//
	// This parameter is required.
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// Specifies whether to allow access from a browser. Default value: **Disabled**. Valid values:
	//
	// 	- **Enabled**
	//
	// 	- **Disabled**
	//
	// example:
	//
	// Disabled
	BrowserAccessStatus *string `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	// The description of the office application. The value must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The browser access mode parameter. The parameter specifies the configurations of Layer 7 applications.
	L7Config *PAL7Config `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
	// The browser access mode parameter. The parameter specifies the prefix of the domain name that the proxy gateway uses. The prefix must be 3 to 20 characters in length, and can contain lowercase letters, digits, and hyphens (-).
	//
	// example:
	//
	// app-sample
	L7ProxyDomainAutomaticPrefix *string `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	// The browser access mode parameter. The parameter specifies the custom domain name of the proxy gateway. Enter a valid domain name.
	//
	// example:
	//
	// app1.example.com
	L7ProxyDomainCustom *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// The name of the office application. The value must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port ranges of the office applications. You can enter up to 65,535 port ranges. Multiple port ranges cannot be duplicated or overlapped.
	//
	// This parameter is required.
	PortRanges []*CreatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// The protocol that is used by the office application. Valid values:
	//
	// 	- **All**
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// This parameter is required.
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
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of the tags for the office applications. You can add up to six custom tags to an office application.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s CreatePrivateAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreatePrivateAccessApplicationRequest) GetBrowserAccessStatus() *string {
	return s.BrowserAccessStatus
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

func (s *CreatePrivateAccessApplicationRequest) SetAddresses(v []*string) *CreatePrivateAccessApplicationRequest {
	s.Addresses = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationRequest {
	s.BrowserAccessStatus = &v
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

func (s *CreatePrivateAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePrivateAccessApplicationRequestPortRanges struct {
	// The start port. The start port must be less than or equal to the end port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The end port. The end port must be greater than or equal to the start port.
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
