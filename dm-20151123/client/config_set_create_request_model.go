// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ConfigSetCreateRequest
	GetDescription() *string
	SetIpPoolId(v string) *ConfigSetCreateRequest
	GetIpPoolId() *string
	SetIsPublicChannelBackoff(v bool) *ConfigSetCreateRequest
	GetIsPublicChannelBackoff() *bool
	SetName(v string) *ConfigSetCreateRequest
	GetName() *string
	SetValidationOption(v *ConfigSetCreateRequestValidationOption) *ConfigSetCreateRequest
	GetValidationOption() *ConfigSetCreateRequestValidationOption
}

type ConfigSetCreateRequest struct {
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
	ValidationOption *ConfigSetCreateRequestValidationOption `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty" type:"Struct"`
}

func (s ConfigSetCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCreateRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetCreateRequest) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetCreateRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetCreateRequest) GetIsPublicChannelBackoff() *bool {
	return s.IsPublicChannelBackoff
}

func (s *ConfigSetCreateRequest) GetName() *string {
	return s.Name
}

func (s *ConfigSetCreateRequest) GetValidationOption() *ConfigSetCreateRequestValidationOption {
	return s.ValidationOption
}

func (s *ConfigSetCreateRequest) SetDescription(v string) *ConfigSetCreateRequest {
	s.Description = &v
	return s
}

func (s *ConfigSetCreateRequest) SetIpPoolId(v string) *ConfigSetCreateRequest {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetCreateRequest) SetIsPublicChannelBackoff(v bool) *ConfigSetCreateRequest {
	s.IsPublicChannelBackoff = &v
	return s
}

func (s *ConfigSetCreateRequest) SetName(v string) *ConfigSetCreateRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetCreateRequest) SetValidationOption(v *ConfigSetCreateRequestValidationOption) *ConfigSetCreateRequest {
	s.ValidationOption = v
	return s
}

func (s *ConfigSetCreateRequest) Validate() error {
	if s.ValidationOption != nil {
		if err := s.ValidationOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetCreateRequestValidationOption struct {
	Enabled                *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ForbiddenStatusList    []*string `json:"ForbiddenStatusList,omitempty" xml:"ForbiddenStatusList,omitempty" type:"Repeated"`
	ForbiddenSubStatusList []*string `json:"ForbiddenSubStatusList,omitempty" xml:"ForbiddenSubStatusList,omitempty" type:"Repeated"`
}

func (s ConfigSetCreateRequestValidationOption) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetCreateRequestValidationOption) GoString() string {
	return s.String()
}

func (s *ConfigSetCreateRequestValidationOption) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConfigSetCreateRequestValidationOption) GetForbiddenStatusList() []*string {
	return s.ForbiddenStatusList
}

func (s *ConfigSetCreateRequestValidationOption) GetForbiddenSubStatusList() []*string {
	return s.ForbiddenSubStatusList
}

func (s *ConfigSetCreateRequestValidationOption) SetEnabled(v bool) *ConfigSetCreateRequestValidationOption {
	s.Enabled = &v
	return s
}

func (s *ConfigSetCreateRequestValidationOption) SetForbiddenStatusList(v []*string) *ConfigSetCreateRequestValidationOption {
	s.ForbiddenStatusList = v
	return s
}

func (s *ConfigSetCreateRequestValidationOption) SetForbiddenSubStatusList(v []*string) *ConfigSetCreateRequestValidationOption {
	s.ForbiddenSubStatusList = v
	return s
}

func (s *ConfigSetCreateRequestValidationOption) Validate() error {
	return dara.Validate(s)
}
