// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProfileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateProfileShrinkRequest
	GetAppId() *string
	SetAttributesShrink(v string) *CreateProfileShrinkRequest
	GetAttributesShrink() *string
	SetDescription(v string) *CreateProfileShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateProfileShrinkRequest
	GetName() *string
	SetUserDefinedId(v string) *CreateProfileShrinkRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *CreateProfileShrinkRequest
	GetWorkspaceId() *string
}

type CreateProfileShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
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

func (s CreateProfileShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProfileShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProfileShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateProfileShrinkRequest) GetAttributesShrink() *string {
	return s.AttributesShrink
}

func (s *CreateProfileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProfileShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProfileShrinkRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *CreateProfileShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateProfileShrinkRequest) SetAppId(v string) *CreateProfileShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateProfileShrinkRequest) SetAttributesShrink(v string) *CreateProfileShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *CreateProfileShrinkRequest) SetDescription(v string) *CreateProfileShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProfileShrinkRequest) SetName(v string) *CreateProfileShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProfileShrinkRequest) SetUserDefinedId(v string) *CreateProfileShrinkRequest {
	s.UserDefinedId = &v
	return s
}

func (s *CreateProfileShrinkRequest) SetWorkspaceId(v string) *CreateProfileShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateProfileShrinkRequest) Validate() error {
	return dara.Validate(s)
}
