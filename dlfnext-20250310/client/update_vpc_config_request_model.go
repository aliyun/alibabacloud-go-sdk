// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRemovals(v []*string) *UpdateVpcConfigRequest
	GetRemovals() []*string
	SetUpdates(v []*UpdateVpcConfigRequestUpdates) *UpdateVpcConfigRequest
	GetUpdates() []*UpdateVpcConfigRequestUpdates
}

type UpdateVpcConfigRequest struct {
	// The list of VPC IDs to delete.
	Removals []*string `json:"removals,omitempty" xml:"removals,omitempty" type:"Repeated"`
	// The list of VPCs to update.
	Updates []*UpdateVpcConfigRequestUpdates `json:"updates,omitempty" xml:"updates,omitempty" type:"Repeated"`
}

func (s UpdateVpcConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcConfigRequest) GetRemovals() []*string {
	return s.Removals
}

func (s *UpdateVpcConfigRequest) GetUpdates() []*UpdateVpcConfigRequestUpdates {
	return s.Updates
}

func (s *UpdateVpcConfigRequest) SetRemovals(v []*string) *UpdateVpcConfigRequest {
	s.Removals = v
	return s
}

func (s *UpdateVpcConfigRequest) SetUpdates(v []*UpdateVpcConfigRequestUpdates) *UpdateVpcConfigRequest {
	s.Updates = v
	return s
}

func (s *UpdateVpcConfigRequest) Validate() error {
	if s.Updates != nil {
		for _, item := range s.Updates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateVpcConfigRequestUpdates struct {
	// The list of configuration items.
	ExtendedOptions map[string]*string `json:"extendedOptions,omitempty" xml:"extendedOptions,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf67xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateVpcConfigRequestUpdates) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcConfigRequestUpdates) GoString() string {
	return s.String()
}

func (s *UpdateVpcConfigRequestUpdates) GetExtendedOptions() map[string]*string {
	return s.ExtendedOptions
}

func (s *UpdateVpcConfigRequestUpdates) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateVpcConfigRequestUpdates) SetExtendedOptions(v map[string]*string) *UpdateVpcConfigRequestUpdates {
	s.ExtendedOptions = v
	return s
}

func (s *UpdateVpcConfigRequestUpdates) SetVpcId(v string) *UpdateVpcConfigRequestUpdates {
	s.VpcId = &v
	return s
}

func (s *UpdateVpcConfigRequestUpdates) Validate() error {
	return dara.Validate(s)
}
