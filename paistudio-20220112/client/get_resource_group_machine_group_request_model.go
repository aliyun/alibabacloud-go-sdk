// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupMachineGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTag(v []*GetResourceGroupMachineGroupRequestTag) *GetResourceGroupMachineGroupRequest
	GetTag() []*GetResourceGroupMachineGroupRequestTag
}

type GetResourceGroupMachineGroupRequest struct {
	Tag []*GetResourceGroupMachineGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupMachineGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupRequest) GetTag() []*GetResourceGroupMachineGroupRequestTag {
	return s.Tag
}

func (s *GetResourceGroupMachineGroupRequest) SetTag(v []*GetResourceGroupMachineGroupRequestTag) *GetResourceGroupMachineGroupRequest {
	s.Tag = v
	return s
}

func (s *GetResourceGroupMachineGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceGroupMachineGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceGroupMachineGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupMachineGroupRequestTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupMachineGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetResourceGroupMachineGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetResourceGroupMachineGroupRequestTag) SetKey(v string) *GetResourceGroupMachineGroupRequestTag {
	s.Key = &v
	return s
}

func (s *GetResourceGroupMachineGroupRequestTag) SetValue(v string) *GetResourceGroupMachineGroupRequestTag {
	s.Value = &v
	return s
}

func (s *GetResourceGroupMachineGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
