// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *CreatePrivateAccessApplicationShrinkRequest
	GetAddresses() []*string
	SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetBrowserAccessStatus() *string
	SetDescription(v string) *CreatePrivateAccessApplicationShrinkRequest
	GetDescription() *string
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
}

type CreatePrivateAccessApplicationShrinkRequest struct {
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
	L7ConfigShrink *string `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
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
	PortRanges []*CreatePrivateAccessApplicationShrinkRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
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

func (s CreatePrivateAccessApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetBrowserAccessStatus() *string {
	return s.BrowserAccessStatus
}

func (s *CreatePrivateAccessApplicationShrinkRequest) GetDescription() *string {
	return s.Description
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

func (s *CreatePrivateAccessApplicationShrinkRequest) SetAddresses(v []*string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Addresses = v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.BrowserAccessStatus = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetDescription(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Description = &v
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

func (s *CreatePrivateAccessApplicationShrinkRequest) Validate() error {
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
