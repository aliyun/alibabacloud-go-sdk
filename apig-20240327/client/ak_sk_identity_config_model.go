// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAkSkIdentityConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAk(v string) *AkSkIdentityConfig
	GetAk() *string
	SetGenerateMode(v string) *AkSkIdentityConfig
	GetGenerateMode() *string
	SetSk(v string) *AkSkIdentityConfig
	GetSk() *string
	SetType(v string) *AkSkIdentityConfig
	GetType() *string
}

type AkSkIdentityConfig struct {
	Ak           *string `json:"ak,omitempty" xml:"ak,omitempty"`
	GenerateMode *string `json:"generateMode,omitempty" xml:"generateMode,omitempty"`
	Sk           *string `json:"sk,omitempty" xml:"sk,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AkSkIdentityConfig) String() string {
	return dara.Prettify(s)
}

func (s AkSkIdentityConfig) GoString() string {
	return s.String()
}

func (s *AkSkIdentityConfig) GetAk() *string {
	return s.Ak
}

func (s *AkSkIdentityConfig) GetGenerateMode() *string {
	return s.GenerateMode
}

func (s *AkSkIdentityConfig) GetSk() *string {
	return s.Sk
}

func (s *AkSkIdentityConfig) GetType() *string {
	return s.Type
}

func (s *AkSkIdentityConfig) SetAk(v string) *AkSkIdentityConfig {
	s.Ak = &v
	return s
}

func (s *AkSkIdentityConfig) SetGenerateMode(v string) *AkSkIdentityConfig {
	s.GenerateMode = &v
	return s
}

func (s *AkSkIdentityConfig) SetSk(v string) *AkSkIdentityConfig {
	s.Sk = &v
	return s
}

func (s *AkSkIdentityConfig) SetType(v string) *AkSkIdentityConfig {
	s.Type = &v
	return s
}

func (s *AkSkIdentityConfig) Validate() error {
	return dara.Validate(s)
}
