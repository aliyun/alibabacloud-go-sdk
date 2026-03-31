// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateProfileShrinkRequest
	GetAppId() *string
	SetAttributesOperationsShrink(v string) *UpdateProfileShrinkRequest
	GetAttributesOperationsShrink() *string
	SetDescription(v string) *UpdateProfileShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateProfileShrinkRequest
	GetName() *string
	SetUserDefinedId(v string) *UpdateProfileShrinkRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *UpdateProfileShrinkRequest
	GetWorkspaceId() *string
}

type UpdateProfileShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId                      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AttributesOperationsShrink *string `json:"AttributesOperations,omitempty" xml:"AttributesOperations,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s UpdateProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateProfileShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateProfileShrinkRequest) GetAttributesOperationsShrink() *string {
	return s.AttributesOperationsShrink
}

func (s *UpdateProfileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProfileShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProfileShrinkRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *UpdateProfileShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateProfileShrinkRequest) SetAppId(v string) *UpdateProfileShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateProfileShrinkRequest) SetAttributesOperationsShrink(v string) *UpdateProfileShrinkRequest {
	s.AttributesOperationsShrink = &v
	return s
}

func (s *UpdateProfileShrinkRequest) SetDescription(v string) *UpdateProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateProfileShrinkRequest) SetName(v string) *UpdateProfileShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateProfileShrinkRequest) SetUserDefinedId(v string) *UpdateProfileShrinkRequest {
	s.UserDefinedId = &v
	return s
}

func (s *UpdateProfileShrinkRequest) SetWorkspaceId(v string) *UpdateProfileShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
