// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *ModelConfig
	GetBucket() *string
	SetFramework(v string) *ModelConfig
	GetFramework() *string
	SetModel(v string) *ModelConfig
	GetModel() *string
	SetMultiModelConfig(v []*ModelConfig) *ModelConfig
	GetMultiModelConfig() []*ModelConfig
	SetPath(v string) *ModelConfig
	GetPath() *string
	SetPrefix(v string) *ModelConfig
	GetPrefix() *string
	SetRegion(v string) *ModelConfig
	GetRegion() *string
	SetReversion(v string) *ModelConfig
	GetReversion() *string
	SetToken(v string) *ModelConfig
	GetToken() *string
	SetType(v string) *ModelConfig
	GetType() *string
}

type ModelConfig struct {
	Bucket    *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Framework *string `json:"framework,omitempty" xml:"framework,omitempty"`
	Model     *string `json:"model,omitempty" xml:"model,omitempty"`
	// if can be null:
	// true
	MultiModelConfig []*ModelConfig `json:"multiModelConfig,omitempty" xml:"multiModelConfig,omitempty" type:"Repeated"`
	Path             *string        `json:"path,omitempty" xml:"path,omitempty"`
	Prefix           *string        `json:"prefix,omitempty" xml:"prefix,omitempty"`
	Region           *string        `json:"region,omitempty" xml:"region,omitempty"`
	Reversion        *string        `json:"reversion,omitempty" xml:"reversion,omitempty"`
	Token            *string        `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// modelscope
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelConfig) String() string {
	return dara.Prettify(s)
}

func (s ModelConfig) GoString() string {
	return s.String()
}

func (s *ModelConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ModelConfig) GetFramework() *string {
	return s.Framework
}

func (s *ModelConfig) GetModel() *string {
	return s.Model
}

func (s *ModelConfig) GetMultiModelConfig() []*ModelConfig {
	return s.MultiModelConfig
}

func (s *ModelConfig) GetPath() *string {
	return s.Path
}

func (s *ModelConfig) GetPrefix() *string {
	return s.Prefix
}

func (s *ModelConfig) GetRegion() *string {
	return s.Region
}

func (s *ModelConfig) GetReversion() *string {
	return s.Reversion
}

func (s *ModelConfig) GetToken() *string {
	return s.Token
}

func (s *ModelConfig) GetType() *string {
	return s.Type
}

func (s *ModelConfig) SetBucket(v string) *ModelConfig {
	s.Bucket = &v
	return s
}

func (s *ModelConfig) SetFramework(v string) *ModelConfig {
	s.Framework = &v
	return s
}

func (s *ModelConfig) SetModel(v string) *ModelConfig {
	s.Model = &v
	return s
}

func (s *ModelConfig) SetMultiModelConfig(v []*ModelConfig) *ModelConfig {
	s.MultiModelConfig = v
	return s
}

func (s *ModelConfig) SetPath(v string) *ModelConfig {
	s.Path = &v
	return s
}

func (s *ModelConfig) SetPrefix(v string) *ModelConfig {
	s.Prefix = &v
	return s
}

func (s *ModelConfig) SetRegion(v string) *ModelConfig {
	s.Region = &v
	return s
}

func (s *ModelConfig) SetReversion(v string) *ModelConfig {
	s.Reversion = &v
	return s
}

func (s *ModelConfig) SetToken(v string) *ModelConfig {
	s.Token = &v
	return s
}

func (s *ModelConfig) SetType(v string) *ModelConfig {
	s.Type = &v
	return s
}

func (s *ModelConfig) Validate() error {
	return dara.Validate(s)
}
