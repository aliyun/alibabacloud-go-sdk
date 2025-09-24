// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetPermissionShrinkRequest
	GetAccessibility() *string
	SetCreator(v string) *GetPermissionShrinkRequest
	GetCreator() *string
	SetLabelsShrink(v string) *GetPermissionShrinkRequest
	GetLabelsShrink() *string
	SetOption(v string) *GetPermissionShrinkRequest
	GetOption() *string
	SetResource(v string) *GetPermissionShrinkRequest
	GetResource() *string
}

type GetPermissionShrinkRequest struct {
	// The accessibility. Valid values:
	//
	// 	- PUBLIC: All members in the workspace can access the workspace.
	//
	// 	- PRIVATE: Only the creator can access the workspace.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The UID of the Alibaba Cloud account that is used to create the workspace.
	//
	// example:
	//
	// 17915******4216
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The configuration. Separate multiple configurations with commas (,). Valid values:
	//
	// 	- ResourceEmpty: The Resource parameter is not configured.
	//
	// 	- DisableRam: The RAM check is not performed.
	//
	// example:
	//
	// ResourceEmpty,DisableRam
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The resource.
	//
	// example:
	//
	// job/dlc-ksd******s12
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GetPermissionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetPermissionShrinkRequest) GetCreator() *string {
	return s.Creator
}

func (s *GetPermissionShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *GetPermissionShrinkRequest) GetOption() *string {
	return s.Option
}

func (s *GetPermissionShrinkRequest) GetResource() *string {
	return s.Resource
}

func (s *GetPermissionShrinkRequest) SetAccessibility(v string) *GetPermissionShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetCreator(v string) *GetPermissionShrinkRequest {
	s.Creator = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetLabelsShrink(v string) *GetPermissionShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetOption(v string) *GetPermissionShrinkRequest {
	s.Option = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetResource(v string) *GetPermissionShrinkRequest {
	s.Resource = &v
	return s
}

func (s *GetPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
