// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateProfileRequest
	GetAppId() *string
	SetAttributes(v []*CreateProfileRequestAttributes) *CreateProfileRequest
	GetAttributes() []*CreateProfileRequestAttributes
	SetDescription(v string) *CreateProfileRequest
	GetDescription() *string
	SetName(v string) *CreateProfileRequest
	GetName() *string
	SetUserDefinedId(v string) *CreateProfileRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *CreateProfileRequest
	GetWorkspaceId() *string
}

type CreateProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Attributes []*CreateProfileRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 6e18191727f747ec9de06a2
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateProfileRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateProfileRequest) GetAttributes() []*CreateProfileRequestAttributes {
	return s.Attributes
}

func (s *CreateProfileRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProfileRequest) GetName() *string {
	return s.Name
}

func (s *CreateProfileRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *CreateProfileRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateProfileRequest) SetAppId(v string) *CreateProfileRequest {
	s.AppId = &v
	return s
}

func (s *CreateProfileRequest) SetAttributes(v []*CreateProfileRequestAttributes) *CreateProfileRequest {
	s.Attributes = v
	return s
}

func (s *CreateProfileRequest) SetDescription(v string) *CreateProfileRequest {
	s.Description = &v
	return s
}

func (s *CreateProfileRequest) SetName(v string) *CreateProfileRequest {
	s.Name = &v
	return s
}

func (s *CreateProfileRequest) SetUserDefinedId(v string) *CreateProfileRequest {
	s.UserDefinedId = &v
	return s
}

func (s *CreateProfileRequest) SetWorkspaceId(v string) *CreateProfileRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateProfileRequest) Validate() error {
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProfileRequestAttributes struct {
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// False
	Immutable *bool `json:"Immutable,omitempty" xml:"Immutable,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateProfileRequestAttributes) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileRequestAttributes) GoString() string {
	return s.String()
}

func (s *CreateProfileRequestAttributes) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *CreateProfileRequestAttributes) GetDescription() *string {
	return s.Description
}

func (s *CreateProfileRequestAttributes) GetImmutable() *bool {
	return s.Immutable
}

func (s *CreateProfileRequestAttributes) GetName() *string {
	return s.Name
}

func (s *CreateProfileRequestAttributes) SetDefaultValue(v string) *CreateProfileRequestAttributes {
	s.DefaultValue = &v
	return s
}

func (s *CreateProfileRequestAttributes) SetDescription(v string) *CreateProfileRequestAttributes {
	s.Description = &v
	return s
}

func (s *CreateProfileRequestAttributes) SetImmutable(v bool) *CreateProfileRequestAttributes {
	s.Immutable = &v
	return s
}

func (s *CreateProfileRequestAttributes) SetName(v string) *CreateProfileRequestAttributes {
	s.Name = &v
	return s
}

func (s *CreateProfileRequestAttributes) Validate() error {
	return dara.Validate(s)
}
