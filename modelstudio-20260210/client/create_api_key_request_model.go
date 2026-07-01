// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v *CreateApiKeyRequestAuth) *CreateApiKeyRequest
	GetAuth() *CreateApiKeyRequestAuth
	SetDescription(v string) *CreateApiKeyRequest
	GetDescription() *string
	SetWorkspaceId(v string) *CreateApiKeyRequest
	GetWorkspaceId() *string
}

type CreateApiKeyRequest struct {
	// The API key permission settings.
	Auth *CreateApiKeyRequestAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The workspace ID.
	//
	// >
	//
	// > - If you leave this parameter empty, the API key is automatically assigned to the default workspace.
	//
	// example:
	//
	// ws-8af73c50f5596193
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequest) GetAuth() *CreateApiKeyRequestAuth {
	return s.Auth
}

func (s *CreateApiKeyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiKeyRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateApiKeyRequest) SetAuth(v *CreateApiKeyRequestAuth) *CreateApiKeyRequest {
	s.Auth = v
	return s
}

func (s *CreateApiKeyRequest) SetDescription(v string) *CreateApiKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateApiKeyRequest) SetWorkspaceId(v string) *CreateApiKeyRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateApiKeyRequest) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApiKeyRequestAuth struct {
	// The IP access whitelist.
	//
	// >
	//
	// > - When you set custom permissions and leave the IP access whitelist empty, the server sets the default values to IPv4 (0.0.0.0/0) and IPv6 (::/0), which allows all traffic.
	AccessIps []*string `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	// The model access scope.
	ModelAccessScope *CreateApiKeyRequestAuthModelAccessScope `json:"modelAccessScope,omitempty" xml:"modelAccessScope,omitempty" type:"Struct"`
	// Valid values:
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

func (s CreateApiKeyRequestAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequestAuth) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequestAuth) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *CreateApiKeyRequestAuth) GetModelAccessScope() *CreateApiKeyRequestAuthModelAccessScope {
	return s.ModelAccessScope
}

func (s *CreateApiKeyRequestAuth) GetType() *string {
	return s.Type
}

func (s *CreateApiKeyRequestAuth) SetAccessIps(v []*string) *CreateApiKeyRequestAuth {
	s.AccessIps = v
	return s
}

func (s *CreateApiKeyRequestAuth) SetModelAccessScope(v *CreateApiKeyRequestAuthModelAccessScope) *CreateApiKeyRequestAuth {
	s.ModelAccessScope = v
	return s
}

func (s *CreateApiKeyRequestAuth) SetType(v string) *CreateApiKeyRequestAuth {
	s.Type = &v
	return s
}

func (s *CreateApiKeyRequestAuth) Validate() error {
	if s.ModelAccessScope != nil {
		if err := s.ModelAccessScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApiKeyRequestAuthModelAccessScope struct {
	// The list of accessible models.
	//
	// 	Notice: This parameter takes effect only when allowAllModels is set to false.
	AccessibleModels []*string `json:"accessibleModels,omitempty" xml:"accessibleModels,omitempty" type:"Repeated"`
	// Specifies whether all models with granted inference permissions in the workspace are accessible. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	AllowAllModels *bool `json:"allowAllModels,omitempty" xml:"allowAllModels,omitempty"`
}

func (s CreateApiKeyRequestAuthModelAccessScope) String() string {
	return dara.Prettify(s)
}

func (s CreateApiKeyRequestAuthModelAccessScope) GoString() string {
	return s.String()
}

func (s *CreateApiKeyRequestAuthModelAccessScope) GetAccessibleModels() []*string {
	return s.AccessibleModels
}

func (s *CreateApiKeyRequestAuthModelAccessScope) GetAllowAllModels() *bool {
	return s.AllowAllModels
}

func (s *CreateApiKeyRequestAuthModelAccessScope) SetAccessibleModels(v []*string) *CreateApiKeyRequestAuthModelAccessScope {
	s.AccessibleModels = v
	return s
}

func (s *CreateApiKeyRequestAuthModelAccessScope) SetAllowAllModels(v bool) *CreateApiKeyRequestAuthModelAccessScope {
	s.AllowAllModels = &v
	return s
}

func (s *CreateApiKeyRequestAuthModelAccessScope) Validate() error {
	return dara.Validate(s)
}
