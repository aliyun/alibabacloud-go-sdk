// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyConfig interface {
	dara.Model
	String() string
	GoString() string
	SetClassName(v string) *AttachPolicyConfig
	GetClassName() *string
	SetConfig(v string) *AttachPolicyConfig
	GetConfig() *string
	SetName(v string) *AttachPolicyConfig
	GetName() *string
}

type AttachPolicyConfig struct {
	ClassName *string `json:"className,omitempty" xml:"className,omitempty"`
	Config    *string `json:"config,omitempty" xml:"config,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AttachPolicyConfig) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyConfig) GoString() string {
	return s.String()
}

func (s *AttachPolicyConfig) GetClassName() *string {
	return s.ClassName
}

func (s *AttachPolicyConfig) GetConfig() *string {
	return s.Config
}

func (s *AttachPolicyConfig) GetName() *string {
	return s.Name
}

func (s *AttachPolicyConfig) SetClassName(v string) *AttachPolicyConfig {
	s.ClassName = &v
	return s
}

func (s *AttachPolicyConfig) SetConfig(v string) *AttachPolicyConfig {
	s.Config = &v
	return s
}

func (s *AttachPolicyConfig) SetName(v string) *AttachPolicyConfig {
	s.Name = &v
	return s
}

func (s *AttachPolicyConfig) Validate() error {
	return dara.Validate(s)
}
