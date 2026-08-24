// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetUpdateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ConfigSetUpdateShrinkRequest
	GetDescription() *string
	SetId(v string) *ConfigSetUpdateShrinkRequest
	GetId() *string
	SetIpPoolId(v string) *ConfigSetUpdateShrinkRequest
	GetIpPoolId() *string
	SetIsPublicChannelBackoff(v bool) *ConfigSetUpdateShrinkRequest
	GetIsPublicChannelBackoff() *bool
	SetName(v string) *ConfigSetUpdateShrinkRequest
	GetName() *string
	SetValidationOptionShrink(v string) *ConfigSetUpdateShrinkRequest
	GetValidationOptionShrink() *string
}

type ConfigSetUpdateShrinkRequest struct {
	// The description. Maximum length: 50 characters.
	//
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The configuration set ID. This parameter is required.
	//
	// example:
	//
	// XXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The associated IP pool ID. This parameter is optional.
	//
	// example:
	//
	// XXX
	IpPoolId               *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	IsPublicChannelBackoff *bool   `json:"IsPublicChannelBackoff,omitempty" xml:"IsPublicChannelBackoff,omitempty"`
	// The configuration name. This parameter is required. Maximum length: 50 characters. The name must be unique.
	//
	// example:
	//
	// XXX
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ValidationOptionShrink *string `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty"`
}

func (s ConfigSetUpdateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetUpdateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetUpdateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetUpdateShrinkRequest) GetId() *string {
	return s.Id
}

func (s *ConfigSetUpdateShrinkRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetUpdateShrinkRequest) GetIsPublicChannelBackoff() *bool {
	return s.IsPublicChannelBackoff
}

func (s *ConfigSetUpdateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ConfigSetUpdateShrinkRequest) GetValidationOptionShrink() *string {
	return s.ValidationOptionShrink
}

func (s *ConfigSetUpdateShrinkRequest) SetDescription(v string) *ConfigSetUpdateShrinkRequest {
	s.Description = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) SetId(v string) *ConfigSetUpdateShrinkRequest {
	s.Id = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) SetIpPoolId(v string) *ConfigSetUpdateShrinkRequest {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) SetIsPublicChannelBackoff(v bool) *ConfigSetUpdateShrinkRequest {
	s.IsPublicChannelBackoff = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) SetName(v string) *ConfigSetUpdateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) SetValidationOptionShrink(v string) *ConfigSetUpdateShrinkRequest {
	s.ValidationOptionShrink = &v
	return s
}

func (s *ConfigSetUpdateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
