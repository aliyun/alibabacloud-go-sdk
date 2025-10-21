// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineVersionMetadata interface {
  dara.Model
  String() string
  GoString() string
  SetEngineVersion(v string) *EngineVersionMetadata
  GetEngineVersion() *string 
  SetFeatures(v *EngineVersionSupportedFeatures) *EngineVersionMetadata
  GetFeatures() *EngineVersionSupportedFeatures 
  SetStatus(v string) *EngineVersionMetadata
  GetStatus() *string 
}

type EngineVersionMetadata struct {
  // This parameter is required.
  // 
  // example:
  // 
  // vvr-6.0.0-flink-1.15
  EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
  Features *EngineVersionSupportedFeatures `json:"features,omitempty" xml:"features,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // STABLE
  Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s EngineVersionMetadata) String() string {
  return dara.Prettify(s)
}

func (s EngineVersionMetadata) GoString() string {
  return s.String()
}

func (s *EngineVersionMetadata) GetEngineVersion() *string  {
  return s.EngineVersion
}

func (s *EngineVersionMetadata) GetFeatures() *EngineVersionSupportedFeatures  {
  return s.Features
}

func (s *EngineVersionMetadata) GetStatus() *string  {
  return s.Status
}

func (s *EngineVersionMetadata) SetEngineVersion(v string) *EngineVersionMetadata {
  s.EngineVersion = &v
  return s
}

func (s *EngineVersionMetadata) SetFeatures(v *EngineVersionSupportedFeatures) *EngineVersionMetadata {
  s.Features = v
  return s
}

func (s *EngineVersionMetadata) SetStatus(v string) *EngineVersionMetadata {
  s.Status = &v
  return s
}

func (s *EngineVersionMetadata) Validate() error {
  if s.Features != nil {
    if err := s.Features.Validate(); err != nil {
      return err
    }
  }
  return nil
}

