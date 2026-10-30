// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v *UpdateApiKeyRequestAuth) *UpdateApiKeyRequest
	GetAuth() *UpdateApiKeyRequestAuth
	SetDescription(v string) *UpdateApiKeyRequest
	GetDescription() *string
}

type UpdateApiKeyRequest struct {
	// The API key permission settings.
	//
	// > When calling the UpdateApiKey operation, either leave this section empty or fill it out completely. Otherwise, the configuration may not match your expectations.
	Auth *UpdateApiKeyRequestAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// update description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s UpdateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequest) GetAuth() *UpdateApiKeyRequestAuth {
	return s.Auth
}

func (s *UpdateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiKeyRequest) SetAuth(v *UpdateApiKeyRequestAuth) *UpdateApiKeyRequest {
	s.Auth = v
	return s
}

func (s *UpdateApiKeyRequest) SetDescription(v string) *UpdateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiKeyRequest) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApiKeyRequestAuth struct {
	// The IP address whitelist.
	//
	// >
	//
	// > - When customizing the permission scope, if the IP address whitelist is not specified, the server sets the default values to IPv4 (0.0.0.0/0) and IPv6 (::/0), which allows all traffic.
	AccessIps []*string `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	// The model access scope.
	ModelAccessScope *UpdateApiKeyRequestAuthModelAccessScope `json:"modelAccessScope,omitempty" xml:"modelAccessScope,omitempty" type:"Struct"`
	// The permission type. Valid values:
	//
	// - All: all permissions.
	//
	// - Custom: custom permissions.
	//
	// example:
	//
	// Custom
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateApiKeyRequestAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequestAuth) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequestAuth) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *UpdateApiKeyRequestAuth) GetModelAccessScope() *UpdateApiKeyRequestAuthModelAccessScope {
	return s.ModelAccessScope
}

func (s *UpdateApiKeyRequestAuth) GetType() *string {
	return s.Type
}

func (s *UpdateApiKeyRequestAuth) SetAccessIps(v []*string) *UpdateApiKeyRequestAuth {
	s.AccessIps = v
	return s
}

func (s *UpdateApiKeyRequestAuth) SetModelAccessScope(v *UpdateApiKeyRequestAuthModelAccessScope) *UpdateApiKeyRequestAuth {
	s.ModelAccessScope = v
	return s
}

func (s *UpdateApiKeyRequestAuth) SetType(v string) *UpdateApiKeyRequestAuth {
	s.Type = &v
	return s
}

func (s *UpdateApiKeyRequestAuth) Validate() error {
	if s.ModelAccessScope != nil {
		if err := s.ModelAccessScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApiKeyRequestAuthModelAccessScope struct {
	// The list of accessible models.
	//
	// 	Notice: The specified content takes effect only when allowAllModels is set to false.
	AccessibleModels []*string `json:"accessibleModels,omitempty" xml:"accessibleModels,omitempty" type:"Repeated"`
	// Specifies whether to allow access to all models that have been granted inference permissions in the workspace. Valid values:
	//
	// - true
	//
	// - false
	AllowAllModels *bool `json:"allowAllModels,omitempty" xml:"allowAllModels,omitempty"`
}

func (s UpdateApiKeyRequestAuthModelAccessScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequestAuthModelAccessScope) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequestAuthModelAccessScope) GetAccessibleModels() []*string {
	return s.AccessibleModels
}

func (s *UpdateApiKeyRequestAuthModelAccessScope) GetAllowAllModels() *bool {
	return s.AllowAllModels
}

func (s *UpdateApiKeyRequestAuthModelAccessScope) SetAccessibleModels(v []*string) *UpdateApiKeyRequestAuthModelAccessScope {
	s.AccessibleModels = v
	return s
}

func (s *UpdateApiKeyRequestAuthModelAccessScope) SetAllowAllModels(v bool) *UpdateApiKeyRequestAuthModelAccessScope {
	s.AllowAllModels = &v
	return s
}

func (s *UpdateApiKeyRequestAuthModelAccessScope) Validate() error {
	return dara.Validate(s)
}
