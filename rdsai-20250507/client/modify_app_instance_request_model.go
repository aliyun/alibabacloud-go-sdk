// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyAppInstanceRequest
	GetClientToken() *string
	SetComponents(v []*ModifyAppInstanceRequestComponents) *ModifyAppInstanceRequest
	GetComponents() []*ModifyAppInstanceRequestComponents
	SetInstanceName(v string) *ModifyAppInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyAppInstanceRequest
	GetRegionId() *string
}

type ModifyAppInstanceRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string                               `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Components  []*ModifyAppInstanceRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// example:
	//
	// ra-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAppInstanceRequest) GetComponents() []*ModifyAppInstanceRequestComponents {
	return s.Components
}

func (s *ModifyAppInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyAppInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAppInstanceRequest) SetClientToken(v string) *ModifyAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAppInstanceRequest) SetComponents(v []*ModifyAppInstanceRequestComponents) *ModifyAppInstanceRequest {
	s.Components = v
	return s
}

func (s *ModifyAppInstanceRequest) SetInstanceName(v string) *ModifyAppInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyAppInstanceRequest) SetRegionId(v string) *ModifyAppInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAppInstanceRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAppInstanceRequestComponents struct {
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// supabase
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyAppInstanceRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppInstanceRequestComponents) GoString() string {
	return s.String()
}

func (s *ModifyAppInstanceRequestComponents) GetStatus() *string {
	return s.Status
}

func (s *ModifyAppInstanceRequestComponents) GetType() *string {
	return s.Type
}

func (s *ModifyAppInstanceRequestComponents) SetStatus(v string) *ModifyAppInstanceRequestComponents {
	s.Status = &v
	return s
}

func (s *ModifyAppInstanceRequestComponents) SetType(v string) *ModifyAppInstanceRequestComponents {
	s.Type = &v
	return s
}

func (s *ModifyAppInstanceRequestComponents) Validate() error {
	return dara.Validate(s)
}
