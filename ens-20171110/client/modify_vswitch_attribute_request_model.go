// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVSwitchAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyVSwitchAttributeRequest
	GetDescription() *string
	SetVSwitchId(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchId() *string
	SetVSwitchName(v string) *ModifyVSwitchAttributeRequest
	GetVSwitchName() *string
}

type ModifyVSwitchAttributeRequest struct {
	// The description of the listener.
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// 	- It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// this is my first network
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The name of the vSwitch.
	//
	// 	- The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// 	- It must start with a letter but cannot start with http:// or https://.
	//
	// example:
	//
	// Test-switch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
}

func (s ModifyVSwitchAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVSwitchAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyVSwitchAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyVSwitchAttributeRequest) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *ModifyVSwitchAttributeRequest) SetDescription(v string) *ModifyVSwitchAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchId(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) SetVSwitchName(v string) *ModifyVSwitchAttributeRequest {
	s.VSwitchName = &v
	return s
}

func (s *ModifyVSwitchAttributeRequest) Validate() error {
	return dara.Validate(s)
}
