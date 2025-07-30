// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAuth(v string) *ImageConfig
	GetAuth() *string
	SetDockerRegistry(v string) *ImageConfig
	GetDockerRegistry() *string
	SetPassword(v string) *ImageConfig
	GetPassword() *string
	SetUsername(v string) *ImageConfig
	GetUsername() *string
}

type ImageConfig struct {
	Auth *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com
	DockerRegistry *string `json:"DockerRegistry,omitempty" xml:"DockerRegistry,omitempty"`
	Password       *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username       *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ImageConfig) String() string {
	return dara.Prettify(s)
}

func (s ImageConfig) GoString() string {
	return s.String()
}

func (s *ImageConfig) GetAuth() *string {
	return s.Auth
}

func (s *ImageConfig) GetDockerRegistry() *string {
	return s.DockerRegistry
}

func (s *ImageConfig) GetPassword() *string {
	return s.Password
}

func (s *ImageConfig) GetUsername() *string {
	return s.Username
}

func (s *ImageConfig) SetAuth(v string) *ImageConfig {
	s.Auth = &v
	return s
}

func (s *ImageConfig) SetDockerRegistry(v string) *ImageConfig {
	s.DockerRegistry = &v
	return s
}

func (s *ImageConfig) SetPassword(v string) *ImageConfig {
	s.Password = &v
	return s
}

func (s *ImageConfig) SetUsername(v string) *ImageConfig {
	s.Username = &v
	return s
}

func (s *ImageConfig) Validate() error {
	return dara.Validate(s)
}
