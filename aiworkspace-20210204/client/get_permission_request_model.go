// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetPermissionRequest
	GetAccessibility() *string
	SetCallerType(v string) *GetPermissionRequest
	GetCallerType() *string
	SetCallerUid(v string) *GetPermissionRequest
	GetCallerUid() *string
	SetCreator(v string) *GetPermissionRequest
	GetCreator() *string
	SetLabels(v map[string]interface{}) *GetPermissionRequest
	GetLabels() map[string]interface{}
	SetOption(v string) *GetPermissionRequest
	GetOption() *string
	SetResource(v string) *GetPermissionRequest
	GetResource() *string
	SetSecurityToken(v string) *GetPermissionRequest
	GetSecurityToken() *string
}

type GetPermissionRequest struct {
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
	CallerType    *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid     *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// The UID of the Alibaba Cloud account that is used to create the workspace.
	//
	// example:
	//
	// 17915******4216
	Creator *string                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Labels  map[string]interface{} `json:"Labels,omitempty" xml:"Labels,omitempty"`
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
	Resource      *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPermissionRequest) GoString() string {
	return s.String()
}

func (s *GetPermissionRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetPermissionRequest) GetCallerType() *string {
	return s.CallerType
}

func (s *GetPermissionRequest) GetCallerUid() *string {
	return s.CallerUid
}

func (s *GetPermissionRequest) GetCreator() *string {
	return s.Creator
}

func (s *GetPermissionRequest) GetLabels() map[string]interface{} {
	return s.Labels
}

func (s *GetPermissionRequest) GetOption() *string {
	return s.Option
}

func (s *GetPermissionRequest) GetResource() *string {
	return s.Resource
}

func (s *GetPermissionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetPermissionRequest) SetAccessibility(v string) *GetPermissionRequest {
	s.Accessibility = &v
	return s
}

func (s *GetPermissionRequest) SetCallerType(v string) *GetPermissionRequest {
	s.CallerType = &v
	return s
}

func (s *GetPermissionRequest) SetCallerUid(v string) *GetPermissionRequest {
	s.CallerUid = &v
	return s
}

func (s *GetPermissionRequest) SetCreator(v string) *GetPermissionRequest {
	s.Creator = &v
	return s
}

func (s *GetPermissionRequest) SetLabels(v map[string]interface{}) *GetPermissionRequest {
	s.Labels = v
	return s
}

func (s *GetPermissionRequest) SetOption(v string) *GetPermissionRequest {
	s.Option = &v
	return s
}

func (s *GetPermissionRequest) SetResource(v string) *GetPermissionRequest {
	s.Resource = &v
	return s
}

func (s *GetPermissionRequest) SetSecurityToken(v string) *GetPermissionRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetPermissionRequest) Validate() error {
	return dara.Validate(s)
}
