// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateAccessApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddresses(v []*string) *UpdatePrivateAccessApplicationShrinkRequest
	GetAddresses() []*string
	SetApplicationId(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetApplicationId() *string
	SetDescription(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetDescription() *string
	SetL7ConfigShrink(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetL7ConfigShrink() *string
	SetL7ProxyDomainAutomaticPrefix(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetL7ProxyDomainAutomaticPrefix() *string
	SetL7ProxyDomainCustom(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetL7ProxyDomainCustom() *string
	SetL7ProxyDomainPrivate(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetL7ProxyDomainPrivate() *string
	SetModifyType(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetModifyType() *string
	SetPortRanges(v []*UpdatePrivateAccessApplicationShrinkRequestPortRanges) *UpdatePrivateAccessApplicationShrinkRequest
	GetPortRanges() []*UpdatePrivateAccessApplicationShrinkRequestPortRanges
	SetProtocol(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetProtocol() *string
	SetStatus(v string) *UpdatePrivateAccessApplicationShrinkRequest
	GetStatus() *string
	SetTagIds(v []*string) *UpdatePrivateAccessApplicationShrinkRequest
	GetTagIds() []*string
}

type UpdatePrivateAccessApplicationShrinkRequest struct {
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
	// The description of the office application. The value must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces.
	//
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The browser access mode parameter. The parameter specifies the configurations of Layer 7 applications.
	L7ConfigShrink *string `json:"L7Config,omitempty" xml:"L7Config,omitempty"`
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
	// The port ranges of the office applications. You can enter up to 65,535 port ranges. Multiple port ranges cannot be duplicated or overlapped.
	PortRanges []*UpdatePrivateAccessApplicationShrinkRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
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

func (s UpdatePrivateAccessApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetAddresses() []*string {
	return s.Addresses
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetL7ConfigShrink() *string {
	return s.L7ConfigShrink
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetL7ProxyDomainAutomaticPrefix() *string {
	return s.L7ProxyDomainAutomaticPrefix
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetL7ProxyDomainCustom() *string {
	return s.L7ProxyDomainCustom
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetL7ProxyDomainPrivate() *string {
	return s.L7ProxyDomainPrivate
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetPortRanges() []*UpdatePrivateAccessApplicationShrinkRequestPortRanges {
	return s.PortRanges
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetAddresses(v []*string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Addresses = v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetApplicationId(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetDescription(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetL7ConfigShrink(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.L7ConfigShrink = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetL7ProxyDomainAutomaticPrefix(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetL7ProxyDomainCustom(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetL7ProxyDomainPrivate(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.L7ProxyDomainPrivate = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetModifyType(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetPortRanges(v []*UpdatePrivateAccessApplicationShrinkRequestPortRanges) *UpdatePrivateAccessApplicationShrinkRequest {
	s.PortRanges = v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetProtocol(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetStatus(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetTagIds(v []*string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.TagIds = v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) Validate() error {
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

type UpdatePrivateAccessApplicationShrinkRequestPortRanges struct {
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

func (s UpdatePrivateAccessApplicationShrinkRequestPortRanges) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateAccessApplicationShrinkRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationShrinkRequestPortRanges) GetBegin() *int32 {
	return s.Begin
}

func (s *UpdatePrivateAccessApplicationShrinkRequestPortRanges) GetEnd() *int32 {
	return s.End
}

func (s *UpdatePrivateAccessApplicationShrinkRequestPortRanges) SetBegin(v int32) *UpdatePrivateAccessApplicationShrinkRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequestPortRanges) SetEnd(v int32) *UpdatePrivateAccessApplicationShrinkRequestPortRanges {
	s.End = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequestPortRanges) Validate() error {
	return dara.Validate(s)
}
