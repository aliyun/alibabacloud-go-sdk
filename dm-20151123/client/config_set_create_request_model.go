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
	SetName(v string) *ConfigSetCreateRequest
	GetName() *string
}

type ConfigSetCreateRequest struct {
	// The description of the configuration set. The description can be up to 50 characters long.
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
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// The name of the configuration set. This parameter is required. The name must be unique and can be up to 50 characters long.
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

func (s *ConfigSetCreateRequest) SetName(v string) *ConfigSetCreateRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetCreateRequest) Validate() error {
	return dara.Validate(s)
}
