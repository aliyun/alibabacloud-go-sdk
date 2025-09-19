// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResidentResourcePoolInput interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *UpdateResidentResourcePoolInput
	GetName() *string
	SetUseScaling(v bool) *UpdateResidentResourcePoolInput
	GetUseScaling() *bool
}

type UpdateResidentResourcePoolInput struct {
	Name       *string `json:"name,omitempty" xml:"name,omitempty"`
	UseScaling *bool   `json:"useScaling,omitempty" xml:"useScaling,omitempty"`
}

func (s UpdateResidentResourcePoolInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateResidentResourcePoolInput) GoString() string {
	return s.String()
}

func (s *UpdateResidentResourcePoolInput) GetName() *string {
	return s.Name
}

func (s *UpdateResidentResourcePoolInput) GetUseScaling() *bool {
	return s.UseScaling
}

func (s *UpdateResidentResourcePoolInput) SetName(v string) *UpdateResidentResourcePoolInput {
	s.Name = &v
	return s
}

func (s *UpdateResidentResourcePoolInput) SetUseScaling(v bool) *UpdateResidentResourcePoolInput {
	s.UseScaling = &v
	return s
}

func (s *UpdateResidentResourcePoolInput) Validate() error {
	return dara.Validate(s)
}
