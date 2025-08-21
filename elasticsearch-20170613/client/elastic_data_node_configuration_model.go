// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iElasticDataNodeConfiguration interface {
  dara.Model
  String() string
  GoString() string
  SetAmount(v int64) *ElasticDataNodeConfiguration
  GetAmount() *int64 
  SetDisk(v int64) *ElasticDataNodeConfiguration
  GetDisk() *int64 
  SetDiskEncryption(v bool) *ElasticDataNodeConfiguration
  GetDiskEncryption() *bool 
  SetDiskType(v string) *ElasticDataNodeConfiguration
  GetDiskType() *string 
  SetPerformanceLevel(v string) *ElasticDataNodeConfiguration
  GetPerformanceLevel() *string 
  SetSpec(v string) *ElasticDataNodeConfiguration
  GetSpec() *string 
}

type ElasticDataNodeConfiguration struct {
  Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
  Disk *int64 `json:"disk,omitempty" xml:"disk,omitempty"`
  DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
  DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
  PerformanceLevel *string `json:"performanceLevel,omitempty" xml:"performanceLevel,omitempty"`
  // This parameter is required.
  Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ElasticDataNodeConfiguration) String() string {
  return dara.Prettify(s)
}

func (s ElasticDataNodeConfiguration) GoString() string {
  return s.String()
}

func (s *ElasticDataNodeConfiguration) GetAmount() *int64  {
  return s.Amount
}

func (s *ElasticDataNodeConfiguration) GetDisk() *int64  {
  return s.Disk
}

func (s *ElasticDataNodeConfiguration) GetDiskEncryption() *bool  {
  return s.DiskEncryption
}

func (s *ElasticDataNodeConfiguration) GetDiskType() *string  {
  return s.DiskType
}

func (s *ElasticDataNodeConfiguration) GetPerformanceLevel() *string  {
  return s.PerformanceLevel
}

func (s *ElasticDataNodeConfiguration) GetSpec() *string  {
  return s.Spec
}

func (s *ElasticDataNodeConfiguration) SetAmount(v int64) *ElasticDataNodeConfiguration {
  s.Amount = &v
  return s
}

func (s *ElasticDataNodeConfiguration) SetDisk(v int64) *ElasticDataNodeConfiguration {
  s.Disk = &v
  return s
}

func (s *ElasticDataNodeConfiguration) SetDiskEncryption(v bool) *ElasticDataNodeConfiguration {
  s.DiskEncryption = &v
  return s
}

func (s *ElasticDataNodeConfiguration) SetDiskType(v string) *ElasticDataNodeConfiguration {
  s.DiskType = &v
  return s
}

func (s *ElasticDataNodeConfiguration) SetPerformanceLevel(v string) *ElasticDataNodeConfiguration {
  s.PerformanceLevel = &v
  return s
}

func (s *ElasticDataNodeConfiguration) SetSpec(v string) *ElasticDataNodeConfiguration {
  s.Spec = &v
  return s
}

func (s *ElasticDataNodeConfiguration) Validate() error {
  return dara.Validate(s)
}

