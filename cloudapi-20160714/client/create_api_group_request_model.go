// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBasePath(v string) *CreateApiGroupRequest
	GetBasePath() *string
	SetDescription(v string) *CreateApiGroupRequest
	GetDescription() *string
	SetGroupName(v string) *CreateApiGroupRequest
	GetGroupName() *string
	SetInstanceId(v string) *CreateApiGroupRequest
	GetInstanceId() *string
	SetSecurityToken(v string) *CreateApiGroupRequest
	GetSecurityToken() *string
	SetTag(v []*CreateApiGroupRequestTag) *CreateApiGroupRequest
	GetTag() []*CreateApiGroupRequestTag
}

type CreateApiGroupRequest struct {
	// example:
	//
	// /qqq
	BasePath *string `json:"BasePath,omitempty" xml:"BasePath,omitempty"`
	// example:
	//
	// The weather informations.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Weather
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// apigateway-cn-v6419k43xxxxx
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// Key， Value
	Tag []*CreateApiGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateApiGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApiGroupRequest) GetBasePath() *string {
	return s.BasePath
}

func (s *CreateApiGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateApiGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateApiGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateApiGroupRequest) GetTag() []*CreateApiGroupRequestTag {
	return s.Tag
}

func (s *CreateApiGroupRequest) SetBasePath(v string) *CreateApiGroupRequest {
	s.BasePath = &v
	return s
}

func (s *CreateApiGroupRequest) SetDescription(v string) *CreateApiGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateApiGroupRequest) SetGroupName(v string) *CreateApiGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateApiGroupRequest) SetInstanceId(v string) *CreateApiGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateApiGroupRequest) SetSecurityToken(v string) *CreateApiGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateApiGroupRequest) SetTag(v []*CreateApiGroupRequestTag) *CreateApiGroupRequest {
	s.Tag = v
	return s
}

func (s *CreateApiGroupRequest) Validate() error {
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

type CreateApiGroupRequestTag struct {
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateApiGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateApiGroupRequestTag) GoString() string {
	return s.String()
}

func (s *CreateApiGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateApiGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateApiGroupRequestTag) SetKey(v string) *CreateApiGroupRequestTag {
	s.Key = &v
	return s
}

func (s *CreateApiGroupRequestTag) SetValue(v string) *CreateApiGroupRequestTag {
	s.Value = &v
	return s
}

func (s *CreateApiGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
