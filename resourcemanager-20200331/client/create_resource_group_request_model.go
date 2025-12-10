// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *CreateResourceGroupRequest
	GetDisplayName() *string
	SetName(v string) *CreateResourceGroupRequest
	GetName() *string
	SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest
	GetTag() []*CreateResourceGroupRequestTag
}

type CreateResourceGroupRequest struct {
	// The display name of the resource group.
	//
	// It must be 1 to 50 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier of the resource group.
	//
	// It must be 2 to 50 characters in length and can contain letters, digits, and hyphens (-). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of tags.
	Tag []*CreateResourceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateResourceGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceGroupRequest) GetTag() []*CreateResourceGroupRequestTag {
	return s.Tag
}

func (s *CreateResourceGroupRequest) SetDisplayName(v string) *CreateResourceGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceGroupRequest) SetName(v string) *CreateResourceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupRequest) SetTag(v []*CreateResourceGroupRequestTag) *CreateResourceGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateResourceGroupRequest) Validate() error {
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

type CreateResourceGroupRequestTag struct {
	// The key of the tag.
	//
	// The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// The tag value can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateResourceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateResourceGroupRequestTag) SetKey(v string) *CreateResourceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateResourceGroupRequestTag) SetValue(v string) *CreateResourceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateResourceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
