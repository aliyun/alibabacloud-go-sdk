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
}

type ConfigSetCreateRequest struct {
	// A description of the configuration set (up to 50 characters).
	//
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of an IP pool to associate with the configuration set. This parameter is optional.
	//
	// example:
	//
	// XXX
	IpPoolId               *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	IsPublicChannelBackoff *bool   `json:"IsPublicChannelBackoff,omitempty" xml:"IsPublicChannelBackoff,omitempty"`
	// The name of the configuration set. The name must be unique and can contain up to 50 characters. This parameter is required.
	//
	// example:
	//
	// XXX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *ConfigSetCreateRequest) Validate() error {
	return dara.Validate(s)
}
