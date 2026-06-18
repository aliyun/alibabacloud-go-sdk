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
	// The IP access whitelist.
	//
	// >
	//
	// > - When you set custom permissions and do not specify the IP access whitelist, the server sets the whitelist to IPv4 (0.0.0.0/0) and IPv6 (::/0) by default, which allows all traffic.
	AccessIps []*string `json:"accessIps,omitempty" xml:"accessIps,omitempty" type:"Repeated"`
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

func (s UpdateApiKeyRequestAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiKeyRequestAuth) GoString() string {
	return s.String()
}

func (s *UpdateApiKeyRequestAuth) GetAccessIps() []*string {
	return s.AccessIps
}

func (s *UpdateApiKeyRequestAuth) GetType() *string {
	return s.Type
}

func (s *UpdateApiKeyRequestAuth) SetAccessIps(v []*string) *UpdateApiKeyRequestAuth {
	s.AccessIps = v
	return s
}

func (s *UpdateApiKeyRequestAuth) SetType(v string) *UpdateApiKeyRequestAuth {
	s.Type = &v
	return s
}

func (s *UpdateApiKeyRequestAuth) Validate() error {
	return dara.Validate(s)
}
