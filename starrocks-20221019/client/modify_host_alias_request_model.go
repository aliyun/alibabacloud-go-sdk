// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHostAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyHostAliasRequest
	GetInstanceId() *string
	SetHostAliases(v []*ModifyHostAliasRequestHostAliases) *ModifyHostAliasRequest
	GetHostAliases() []*ModifyHostAliasRequestHostAliases
}

type ModifyHostAliasRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId  *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	HostAliases []*ModifyHostAliasRequestHostAliases `json:"hostAliases,omitempty" xml:"hostAliases,omitempty" type:"Repeated"`
}

func (s ModifyHostAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAliasRequest) GoString() string {
	return s.String()
}

func (s *ModifyHostAliasRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyHostAliasRequest) GetHostAliases() []*ModifyHostAliasRequestHostAliases {
	return s.HostAliases
}

func (s *ModifyHostAliasRequest) SetInstanceId(v string) *ModifyHostAliasRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyHostAliasRequest) SetHostAliases(v []*ModifyHostAliasRequestHostAliases) *ModifyHostAliasRequest {
	s.HostAliases = v
	return s
}

func (s *ModifyHostAliasRequest) Validate() error {
	if s.HostAliases != nil {
		for _, item := range s.HostAliases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyHostAliasRequestHostAliases struct {
	Hostnames []*string `json:"hostnames,omitempty" xml:"hostnames,omitempty" type:"Repeated"`
	// example:
	//
	// 26.15.54.221
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s ModifyHostAliasRequestHostAliases) String() string {
	return dara.Prettify(s)
}

func (s ModifyHostAliasRequestHostAliases) GoString() string {
	return s.String()
}

func (s *ModifyHostAliasRequestHostAliases) GetHostnames() []*string {
	return s.Hostnames
}

func (s *ModifyHostAliasRequestHostAliases) GetIp() *string {
	return s.Ip
}

func (s *ModifyHostAliasRequestHostAliases) SetHostnames(v []*string) *ModifyHostAliasRequestHostAliases {
	s.Hostnames = v
	return s
}

func (s *ModifyHostAliasRequestHostAliases) SetIp(v string) *ModifyHostAliasRequestHostAliases {
	s.Ip = &v
	return s
}

func (s *ModifyHostAliasRequestHostAliases) Validate() error {
	return dara.Validate(s)
}
