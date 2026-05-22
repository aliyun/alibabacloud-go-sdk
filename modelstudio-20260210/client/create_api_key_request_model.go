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
	Auth *CreateApiKeyRequestAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
	AccessIps []*string `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
	// example:
	//
	// All
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

func (s *CreateApiKeyRequestAuth) GetType() *string {
	return s.Type
}

func (s *CreateApiKeyRequestAuth) SetAccessIps(v []*string) *CreateApiKeyRequestAuth {
	s.AccessIps = v
	return s
}

func (s *CreateApiKeyRequestAuth) SetType(v string) *CreateApiKeyRequestAuth {
	s.Type = &v
	return s
}

func (s *CreateApiKeyRequestAuth) Validate() error {
	return dara.Validate(s)
}
