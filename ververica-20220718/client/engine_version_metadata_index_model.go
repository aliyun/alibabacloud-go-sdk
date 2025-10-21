// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEngineVersionMetadataIndex interface {
  dara.Model
  String() string
  GoString() string
  SetDefaultEngineVersion(v string) *EngineVersionMetadataIndex
  GetDefaultEngineVersion() *string 
  SetEngineVersionMetadata(v []*EngineVersionMetadata) *EngineVersionMetadataIndex
  GetEngineVersionMetadata() []*EngineVersionMetadata 
}

type EngineVersionMetadataIndex struct {
  // example:
  // 
  // vvr-6.0.1-flink-1.15
  DefaultEngineVersion *string `json:"defaultEngineVersion,omitempty" xml:"defaultEngineVersion,omitempty"`
  EngineVersionMetadata []*EngineVersionMetadata `json:"engineVersionMetadata,omitempty" xml:"engineVersionMetadata,omitempty" type:"Repeated"`
}

func (s EngineVersionMetadataIndex) String() string {
  return dara.Prettify(s)
}

func (s EngineVersionMetadataIndex) GoString() string {
  return s.String()
}

func (s *EngineVersionMetadataIndex) GetDefaultEngineVersion() *string  {
  return s.DefaultEngineVersion
}

func (s *EngineVersionMetadataIndex) GetEngineVersionMetadata() []*EngineVersionMetadata  {
  return s.EngineVersionMetadata
}

func (s *EngineVersionMetadataIndex) SetDefaultEngineVersion(v string) *EngineVersionMetadataIndex {
  s.DefaultEngineVersion = &v
  return s
}

func (s *EngineVersionMetadataIndex) SetEngineVersionMetadata(v []*EngineVersionMetadata) *EngineVersionMetadataIndex {
  s.EngineVersionMetadata = v
  return s
}

func (s *EngineVersionMetadataIndex) Validate() error {
  if s.EngineVersionMetadata != nil {
    for _, item := range s.EngineVersionMetadata {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

