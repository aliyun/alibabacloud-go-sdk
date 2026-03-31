// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateProfileRequest
	GetAppId() *string
	SetAttributesOperations(v []*UpdateProfileRequestAttributesOperations) *UpdateProfileRequest
	GetAttributesOperations() []*UpdateProfileRequestAttributesOperations
	SetDescription(v string) *UpdateProfileRequest
	GetDescription() *string
	SetName(v string) *UpdateProfileRequest
	GetName() *string
	SetUserDefinedId(v string) *UpdateProfileRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *UpdateProfileRequest
	GetWorkspaceId() *string
}

type UpdateProfileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId                *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AttributesOperations []*UpdateProfileRequestAttributesOperations `json:"AttributesOperations,omitempty" xml:"AttributesOperations,omitempty" type:"Repeated"`
	Description          *string                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10b6d435
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateProfileRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateProfileRequest) GetAttributesOperations() []*UpdateProfileRequestAttributesOperations {
	return s.AttributesOperations
}

func (s *UpdateProfileRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProfileRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProfileRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *UpdateProfileRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateProfileRequest) SetAppId(v string) *UpdateProfileRequest {
	s.AppId = &v
	return s
}

func (s *UpdateProfileRequest) SetAttributesOperations(v []*UpdateProfileRequestAttributesOperations) *UpdateProfileRequest {
	s.AttributesOperations = v
	return s
}

func (s *UpdateProfileRequest) SetDescription(v string) *UpdateProfileRequest {
	s.Description = &v
	return s
}

func (s *UpdateProfileRequest) SetName(v string) *UpdateProfileRequest {
	s.Name = &v
	return s
}

func (s *UpdateProfileRequest) SetUserDefinedId(v string) *UpdateProfileRequest {
	s.UserDefinedId = &v
	return s
}

func (s *UpdateProfileRequest) SetWorkspaceId(v string) *UpdateProfileRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateProfileRequest) Validate() error {
	if s.AttributesOperations != nil {
		for _, item := range s.AttributesOperations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProfileRequestAttributesOperations struct {
	// example:
	//
	// 48944
	AttributeId *string `json:"AttributeId,omitempty" xml:"AttributeId,omitempty"`
	// example:
	//
	// 18
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// add
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
}

func (s UpdateProfileRequestAttributesOperations) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileRequestAttributesOperations) GoString() string {
	return s.String()
}

func (s *UpdateProfileRequestAttributesOperations) GetAttributeId() *string {
	return s.AttributeId
}

func (s *UpdateProfileRequestAttributesOperations) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *UpdateProfileRequestAttributesOperations) GetDescription() *string {
	return s.Description
}

func (s *UpdateProfileRequestAttributesOperations) GetName() *string {
	return s.Name
}

func (s *UpdateProfileRequestAttributesOperations) GetOp() *string {
	return s.Op
}

func (s *UpdateProfileRequestAttributesOperations) SetAttributeId(v string) *UpdateProfileRequestAttributesOperations {
	s.AttributeId = &v
	return s
}

func (s *UpdateProfileRequestAttributesOperations) SetDefaultValue(v string) *UpdateProfileRequestAttributesOperations {
	s.DefaultValue = &v
	return s
}

func (s *UpdateProfileRequestAttributesOperations) SetDescription(v string) *UpdateProfileRequestAttributesOperations {
	s.Description = &v
	return s
}

func (s *UpdateProfileRequestAttributesOperations) SetName(v string) *UpdateProfileRequestAttributesOperations {
	s.Name = &v
	return s
}

func (s *UpdateProfileRequestAttributesOperations) SetOp(v string) *UpdateProfileRequestAttributesOperations {
	s.Op = &v
	return s
}

func (s *UpdateProfileRequestAttributesOperations) Validate() error {
	return dara.Validate(s)
}
