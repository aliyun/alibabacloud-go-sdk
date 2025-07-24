// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredential interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *Credential
	GetAccessId() *string
	SetDir(v string) *Credential
	GetDir() *string
	SetExpire(v string) *Credential
	GetExpire() *string
	SetHost(v string) *Credential
	GetHost() *string
	SetPolicy(v string) *Credential
	GetPolicy() *string
	SetSecurityToken(v string) *Credential
	GetSecurityToken() *string
	SetSignature(v string) *Credential
	GetSignature() *string
}

type Credential struct {
	// This parameter is required.
	AccessId *string `json:"accessId,omitempty" xml:"accessId,omitempty"`
	// This parameter is required.
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// This parameter is required.
	Expire *string `json:"expire,omitempty" xml:"expire,omitempty"`
	// This parameter is required.
	Host *string `json:"host,omitempty" xml:"host,omitempty"`
	// This parameter is required.
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// This parameter is required.
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// This parameter is required.
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s Credential) String() string {
	return dara.Prettify(s)
}

func (s Credential) GoString() string {
	return s.String()
}

func (s *Credential) GetAccessId() *string {
	return s.AccessId
}

func (s *Credential) GetDir() *string {
	return s.Dir
}

func (s *Credential) GetExpire() *string {
	return s.Expire
}

func (s *Credential) GetHost() *string {
	return s.Host
}

func (s *Credential) GetPolicy() *string {
	return s.Policy
}

func (s *Credential) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *Credential) GetSignature() *string {
	return s.Signature
}

func (s *Credential) SetAccessId(v string) *Credential {
	s.AccessId = &v
	return s
}

func (s *Credential) SetDir(v string) *Credential {
	s.Dir = &v
	return s
}

func (s *Credential) SetExpire(v string) *Credential {
	s.Expire = &v
	return s
}

func (s *Credential) SetHost(v string) *Credential {
	s.Host = &v
	return s
}

func (s *Credential) SetPolicy(v string) *Credential {
	s.Policy = &v
	return s
}

func (s *Credential) SetSecurityToken(v string) *Credential {
	s.SecurityToken = &v
	return s
}

func (s *Credential) SetSignature(v string) *Credential {
	s.Signature = &v
	return s
}

func (s *Credential) Validate() error {
	return dara.Validate(s)
}
