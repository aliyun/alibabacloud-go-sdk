// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupRequest
	GetIsAIWorkspaceDataEnabled() *bool
	SetTag(v []*GetResourceGroupRequestTag) *GetResourceGroupRequest
	GetTag() []*GetResourceGroupRequestTag
}

type GetResourceGroupRequest struct {
	// example:
	//
	// true
	IsAIWorkspaceDataEnabled *bool                         `json:"IsAIWorkspaceDataEnabled,omitempty" xml:"IsAIWorkspaceDataEnabled,omitempty"`
	Tag                      []*GetResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequest) GetIsAIWorkspaceDataEnabled() *bool {
	return s.IsAIWorkspaceDataEnabled
}

func (s *GetResourceGroupRequest) GetTag() []*GetResourceGroupRequestTag {
	return s.Tag
}

func (s *GetResourceGroupRequest) SetIsAIWorkspaceDataEnabled(v bool) *GetResourceGroupRequest {
	s.IsAIWorkspaceDataEnabled = &v
	return s
}

func (s *GetResourceGroupRequest) SetTag(v []*GetResourceGroupRequestTag) *GetResourceGroupRequest {
	s.Tag = v
	return s
}

func (s *GetResourceGroupRequest) Validate() error {
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

type GetResourceGroupRequestTag struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *GetResourceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *GetResourceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *GetResourceGroupRequestTag) SetKey(v string) *GetResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *GetResourceGroupRequestTag) SetValue(v string) *GetResourceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *GetResourceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
