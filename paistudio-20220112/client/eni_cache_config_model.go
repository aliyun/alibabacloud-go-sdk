// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEniCacheConfig interface {
  dara.Model
  String() string
  GoString() string
  SetCachePoolSize(v int32) *EniCacheConfig
  GetCachePoolSize() *int32 
  SetEnabled(v bool) *EniCacheConfig
  GetEnabled() *bool 
}

type EniCacheConfig struct {
  CachePoolSize *int32 `json:"CachePoolSize,omitempty" xml:"CachePoolSize,omitempty"`
  Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s EniCacheConfig) String() string {
  return dara.Prettify(s)
}

func (s EniCacheConfig) GoString() string {
  return s.String()
}

func (s *EniCacheConfig) GetCachePoolSize() *int32  {
  return s.CachePoolSize
}

func (s *EniCacheConfig) GetEnabled() *bool  {
  return s.Enabled
}

func (s *EniCacheConfig) SetCachePoolSize(v int32) *EniCacheConfig {
  s.CachePoolSize = &v
  return s
}

func (s *EniCacheConfig) SetEnabled(v bool) *EniCacheConfig {
  s.Enabled = &v
  return s
}

func (s *EniCacheConfig) Validate() error {
  return dara.Validate(s)
}

