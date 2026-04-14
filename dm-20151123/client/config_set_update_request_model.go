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
	SetName(v string) *ConfigSetUpdateRequest
	GetName() *string
}

type ConfigSetUpdateRequest struct {
	// example:
	//
	// XXX
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// XXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// XXX
	IpPoolId *string `json:"IpPoolId,omitempty" xml:"IpPoolId,omitempty"`
	// example:
	//
	// XXX
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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

func (s *ConfigSetUpdateRequest) GetName() *string {
	return s.Name
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

func (s *ConfigSetUpdateRequest) SetName(v string) *ConfigSetUpdateRequest {
	s.Name = &v
	return s
}

func (s *ConfigSetUpdateRequest) Validate() error {
	return dara.Validate(s)
}
