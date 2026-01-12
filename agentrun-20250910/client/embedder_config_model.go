// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmbedderConfig interface {
  dara.Model
  String() string
  GoString() string
  SetConfig(v *EmbedderConfigConfig) *EmbedderConfig
  GetConfig() *EmbedderConfigConfig 
  SetModelServiceName(v string) *EmbedderConfig
  GetModelServiceName() *string 
}

type EmbedderConfig struct {
  Config *EmbedderConfigConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
  ModelServiceName *string `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
}

func (s EmbedderConfig) String() string {
  return dara.Prettify(s)
}

func (s EmbedderConfig) GoString() string {
  return s.String()
}

func (s *EmbedderConfig) GetConfig() *EmbedderConfigConfig  {
  return s.Config
}

func (s *EmbedderConfig) GetModelServiceName() *string  {
  return s.ModelServiceName
}

func (s *EmbedderConfig) SetConfig(v *EmbedderConfigConfig) *EmbedderConfig {
  s.Config = v
  return s
}

func (s *EmbedderConfig) SetModelServiceName(v string) *EmbedderConfig {
  s.ModelServiceName = &v
  return s
}

func (s *EmbedderConfig) Validate() error {
  if s.Config != nil {
    if err := s.Config.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EmbedderConfigConfig struct {
  Model *string `json:"model,omitempty" xml:"model,omitempty"`
}

func (s EmbedderConfigConfig) String() string {
  return dara.Prettify(s)
}

func (s EmbedderConfigConfig) GoString() string {
  return s.String()
}

func (s *EmbedderConfigConfig) GetModel() *string  {
  return s.Model
}

func (s *EmbedderConfigConfig) SetModel(v string) *EmbedderConfigConfig {
  s.Model = &v
  return s
}

func (s *EmbedderConfigConfig) Validate() error {
  return dara.Validate(s)
}

