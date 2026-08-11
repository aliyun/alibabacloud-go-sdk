// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthInfoConfigValue interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *AuthInfoConfigValue
	GetAuthToken() *string
	SetPrivateDomain(v string) *AuthInfoConfigValue
	GetPrivateDomain() *string
	SetProject(v string) *AuthInfoConfigValue
	GetProject() *string
	SetPublicDomain(v string) *AuthInfoConfigValue
	GetPublicDomain() *string
}

type AuthInfoConfigValue struct {
	// The credential.
	//
	// example:
	//
	// token-xxx
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// The private domain name.
	//
	// example:
	//
	// https://xxx
	PrivateDomain *string `json:"PrivateDomain,omitempty" xml:"PrivateDomain,omitempty"`
	// The project space.
	//
	// example:
	//
	// proj-xxx
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The public domain name.
	//
	// example:
	//
	// https://xxx
	PublicDomain *string `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
}

func (s AuthInfoConfigValue) String() string {
	return dara.Prettify(s)
}

func (s AuthInfoConfigValue) GoString() string {
	return s.String()
}

func (s *AuthInfoConfigValue) GetAuthToken() *string {
	return s.AuthToken
}

func (s *AuthInfoConfigValue) GetPrivateDomain() *string {
	return s.PrivateDomain
}

func (s *AuthInfoConfigValue) GetProject() *string {
	return s.Project
}

func (s *AuthInfoConfigValue) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *AuthInfoConfigValue) SetAuthToken(v string) *AuthInfoConfigValue {
	s.AuthToken = &v
	return s
}

func (s *AuthInfoConfigValue) SetPrivateDomain(v string) *AuthInfoConfigValue {
	s.PrivateDomain = &v
	return s
}

func (s *AuthInfoConfigValue) SetProject(v string) *AuthInfoConfigValue {
	s.Project = &v
	return s
}

func (s *AuthInfoConfigValue) SetPublicDomain(v string) *AuthInfoConfigValue {
	s.PublicDomain = &v
	return s
}

func (s *AuthInfoConfigValue) Validate() error {
	return dara.Validate(s)
}
