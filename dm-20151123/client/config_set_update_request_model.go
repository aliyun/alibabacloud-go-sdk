// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetUpdateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ConfigSetUpdateRequest
	GetDescription() *string
	SetId(v string) *ConfigSetUpdateRequest
	GetId() *string
	SetIpPoolId(v string) *ConfigSetUpdateRequest
	GetIpPoolId() *string
	SetIsPublicChannelBackoff(v bool) *ConfigSetUpdateRequest
	GetIsPublicChannelBackoff() *bool
	SetName(v string) *ConfigSetUpdateRequest
	GetName() *string
	SetValidationOption(v *ConfigSetUpdateRequestValidationOption) *ConfigSetUpdateRequest
	GetValidationOption() *ConfigSetUpdateRequestValidationOption
}

type ConfigSetUpdateRequest struct {
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
	Name             *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	ValidationOption *ConfigSetUpdateRequestValidationOption `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty" type:"Struct"`
}

func (s ConfigSetUpdateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetUpdateRequest) GoString() string {
	return s.String()
}

func (s *ConfigSetUpdateRequest) GetDescription() *string {
	return s.Description
}

func (s *ConfigSetUpdateRequest) GetId() *string {
	return s.Id
}

func (s *ConfigSetUpdateRequest) GetIpPoolId() *string {
	return s.IpPoolId
}

func (s *ConfigSetUpdateRequest) GetIsPublicChannelBackoff() *bool {
	return s.IsPublicChannelBackoff
}

func (s *ConfigSetUpdateRequest) GetName() *string {
	return s.Name
}

func (s *ConfigSetUpdateRequest) GetValidationOption() *ConfigSetUpdateRequestValidationOption {
	return s.ValidationOption
}

func (s *ConfigSetUpdateRequest) SetDescription(v string) *ConfigSetUpdateRequest {
	s.Description = &v
	return s
}

func (s *ConfigSetUpdateRequest) SetId(v string) *ConfigSetUpdateRequest {
	s.Id = &v
	return s
}

func (s *ConfigSetUpdateRequest) SetIpPoolId(v string) *ConfigSetUpdateRequest {
	s.IpPoolId = &v
	return s
}

func (s *ConfigSetUpdateRequest) SetIsPublicChannelBackoff(v bool) *ConfigSetUpdateRequest {
	s.IsPublicChannelBackoff = &v
	return s
}

func (s *ConfigSetUpdateRequest) SetName(v string) *ConfigSetUpdateRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetUpdateRequest) SetValidationOption(v *ConfigSetUpdateRequestValidationOption) *ConfigSetUpdateRequest {
	s.ValidationOption = v
	return s
}

func (s *ConfigSetUpdateRequest) Validate() error {
	if s.ValidationOption != nil {
		if err := s.ValidationOption.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigSetUpdateRequestValidationOption struct {
	Enabled                *bool     `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ForbiddenStatusList    []*string `json:"ForbiddenStatusList,omitempty" xml:"ForbiddenStatusList,omitempty" type:"Repeated"`
	ForbiddenSubStatusList []*string `json:"ForbiddenSubStatusList,omitempty" xml:"ForbiddenSubStatusList,omitempty" type:"Repeated"`
}

func (s ConfigSetUpdateRequestValidationOption) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetUpdateRequestValidationOption) GoString() string {
	return s.String()
}

func (s *ConfigSetUpdateRequestValidationOption) GetEnabled() *bool {
	return s.Enabled
}

func (s *ConfigSetUpdateRequestValidationOption) GetForbiddenStatusList() []*string {
	return s.ForbiddenStatusList
}

func (s *ConfigSetUpdateRequestValidationOption) GetForbiddenSubStatusList() []*string {
	return s.ForbiddenSubStatusList
}

func (s *ConfigSetUpdateRequestValidationOption) SetEnabled(v bool) *ConfigSetUpdateRequestValidationOption {
	s.Enabled = &v
	return s
}

func (s *ConfigSetUpdateRequestValidationOption) SetForbiddenStatusList(v []*string) *ConfigSetUpdateRequestValidationOption {
	s.ForbiddenStatusList = v
	return s
}

func (s *ConfigSetUpdateRequestValidationOption) SetForbiddenSubStatusList(v []*string) *ConfigSetUpdateRequestValidationOption {
	s.ForbiddenSubStatusList = v
	return s
}

func (s *ConfigSetUpdateRequestValidationOption) Validate() error {
	return dara.Validate(s)
}
