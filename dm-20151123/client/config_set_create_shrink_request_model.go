// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCreateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ConfigSetCreateShrinkRequest
	GetDescription() *string
	SetIpPoolId(v string) *ConfigSetCreateShrinkRequest
	GetIpPoolId() *string
	SetIsPublicChannelBackoff(v bool) *ConfigSetCreateShrinkRequest
	GetIsPublicChannelBackoff() *bool
	SetName(v string) *ConfigSetCreateShrinkRequest
	GetName() *string
	SetValidationOptionShrink(v string) *ConfigSetCreateShrinkRequest
	GetValidationOptionShrink() *string
}

type ConfigSetCreateShrinkRequest struct {
	// The description. The description can be up to 50 characters in length.
	//
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the associated IP pool. This parameter is optional.
	//
	// example:
	//
	// XXX
	IpPoolId               *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	IsPublicChannelBackoff *bool   `json:"IsPublicChannelBackoff,omitempty" xml:"IsPublicChannelBackoff,omitempty"`
	// The configuration name. This parameter is required. The name can be up to 50 characters in length and must be unique.
	//
	// example:
	//
	// XXX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// if can be null:
	// false
	ValidationOptionShrink *string `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty"`
}

func (s ConfigSetCreateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCreateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetCreateShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetCreateShrinkRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetCreateShrinkRequest) GetIsPublicChannelBackoff() *bool {
	return s.IsPublicChannelBackoff
}

func (s *ConfigSetCreateShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ConfigSetCreateShrinkRequest) GetValidationOptionShrink() *string {
	return s.ValidationOptionShrink
}

func (s *ConfigSetCreateShrinkRequest) SetDescription(v string) *ConfigSetCreateShrinkRequest {
	s.Description = &v
	return s
}

func (s *ConfigSetCreateShrinkRequest) SetIpPoolId(v string) *ConfigSetCreateShrinkRequest {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetCreateShrinkRequest) SetIsPublicChannelBackoff(v bool) *ConfigSetCreateShrinkRequest {
	s.IsPublicChannelBackoff = &v
	return s
}

func (s *ConfigSetCreateShrinkRequest) SetName(v string) *ConfigSetCreateShrinkRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetCreateShrinkRequest) SetValidationOptionShrink(v string) *ConfigSetCreateShrinkRequest {
	s.ValidationOptionShrink = &v
	return s
}

func (s *ConfigSetCreateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
