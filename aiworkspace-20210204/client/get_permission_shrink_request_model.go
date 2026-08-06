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
	SetCallerAccessKeyId(v string) *GetPermissionShrinkRequest
	GetCallerAccessKeyId() *string
	SetCallerSecurityToken(v string) *GetPermissionShrinkRequest
	GetCallerSecurityToken() *string
	SetCallerType(v string) *GetPermissionShrinkRequest
	GetCallerType() *string
	SetCallerUid(v string) *GetPermissionShrinkRequest
	GetCallerUid() *string
	SetCreator(v string) *GetPermissionShrinkRequest
	GetCreator() *string
	SetLabelsShrink(v string) *GetPermissionShrinkRequest
	GetLabelsShrink() *string
	SetOption(v string) *GetPermissionShrinkRequest
	GetOption() *string
	SetResource(v string) *GetPermissionShrinkRequest
	GetResource() *string
	SetSecurityToken(v string) *GetPermissionShrinkRequest
	GetSecurityToken() *string
}

type GetPermissionShrinkRequest struct {
	// The access type. Valid values:
	//
	// - PUBLIC: All members in the current workspace can access the instance.
	//
	// - PRIVATE: Only the creator can access the instance.
	//
	// example:
	//
	// PUBLIC
	Accessibility       *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	CallerAccessKeyId   *string `json:"CallerAccessKeyId,omitempty" xml:"CallerAccessKeyId,omitempty"`
	CallerSecurityToken *string `json:"CallerSecurityToken,omitempty" xml:"CallerSecurityToken,omitempty"`
	CallerType          *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid           *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The Alibaba Cloud account UID of the workspace permission creator.
	//
	// example:
	//
	// 17915******4216
	Creator      *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The optional configurations. Separate multiple configurations with commas (,). Valid values:
	//
	// - ResourceEmpty: The resource is empty. The resource is empty if Resource is not specified.
	//
	// - DisableRam: RAM authentication is not performed.
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
	Resource      *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *GetPermissionShrinkRequest) GetCallerAccessKeyId() *string {
	return s.CallerAccessKeyId
}

func (s *GetPermissionShrinkRequest) GetCallerSecurityToken() *string {
	return s.CallerSecurityToken
}

func (s *GetPermissionShrinkRequest) GetCallerType() *string {
	return s.CallerType
}

func (s *GetPermissionShrinkRequest) GetCallerUid() *string {
	return s.CallerUid
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

func (s *GetPermissionShrinkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetPermissionShrinkRequest) SetAccessibility(v string) *GetPermissionShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetCallerAccessKeyId(v string) *GetPermissionShrinkRequest {
	s.CallerAccessKeyId = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetCallerSecurityToken(v string) *GetPermissionShrinkRequest {
	s.CallerSecurityToken = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetCallerType(v string) *GetPermissionShrinkRequest {
	s.CallerType = &v
	return s
}

func (s *GetPermissionShrinkRequest) SetCallerUid(v string) *GetPermissionShrinkRequest {
	s.CallerUid = &v
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

func (s *GetPermissionShrinkRequest) SetSecurityToken(v string) *GetPermissionShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetPermissionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
