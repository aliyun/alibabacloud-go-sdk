// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDesktopOversoldUserGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *AddDesktopOversoldUserGroupRequest
	GetImageId() *string
	SetName(v string) *AddDesktopOversoldUserGroupRequest
	GetName() *string
	SetOversoldGroupId(v string) *AddDesktopOversoldUserGroupRequest
	GetOversoldGroupId() *string
	SetPolicyGroupId(v string) *AddDesktopOversoldUserGroupRequest
	GetPolicyGroupId() *string
	SetTag(v []*AddDesktopOversoldUserGroupRequestTag) *AddDesktopOversoldUserGroupRequest
	GetTag() []*AddDesktopOversoldUserGroupRequestTag
}

type AddDesktopOversoldUserGroupRequest struct {
	ImageId         *string                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name            *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldGroupId *string                                  `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	PolicyGroupId   *string                                  `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	Tag             []*AddDesktopOversoldUserGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddDesktopOversoldUserGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDesktopOversoldUserGroupRequest) GoString() string {
	return s.String()
}

func (s *AddDesktopOversoldUserGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AddDesktopOversoldUserGroupRequest) GetName() *string {
	return s.Name
}

func (s *AddDesktopOversoldUserGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *AddDesktopOversoldUserGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *AddDesktopOversoldUserGroupRequest) GetTag() []*AddDesktopOversoldUserGroupRequestTag {
	return s.Tag
}

func (s *AddDesktopOversoldUserGroupRequest) SetImageId(v string) *AddDesktopOversoldUserGroupRequest {
	s.ImageId = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequest) SetName(v string) *AddDesktopOversoldUserGroupRequest {
	s.Name = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequest) SetOversoldGroupId(v string) *AddDesktopOversoldUserGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequest) SetPolicyGroupId(v string) *AddDesktopOversoldUserGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequest) SetTag(v []*AddDesktopOversoldUserGroupRequestTag) *AddDesktopOversoldUserGroupRequest {
	s.Tag = v
	return s
}

func (s *AddDesktopOversoldUserGroupRequest) Validate() error {
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

type AddDesktopOversoldUserGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddDesktopOversoldUserGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddDesktopOversoldUserGroupRequestTag) GoString() string {
	return s.String()
}

func (s *AddDesktopOversoldUserGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddDesktopOversoldUserGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddDesktopOversoldUserGroupRequestTag) SetKey(v string) *AddDesktopOversoldUserGroupRequestTag {
	s.Key = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequestTag) SetValue(v string) *AddDesktopOversoldUserGroupRequestTag {
	s.Value = &v
	return s
}

func (s *AddDesktopOversoldUserGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
