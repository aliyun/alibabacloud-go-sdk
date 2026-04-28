// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExternalMultiFileRevisionConfig interface {
  dara.Model
  String() string
  GoString() string
  SetRevisionCount(v int64) *ExternalMultiFileRevisionConfig
  GetRevisionCount() *int64 
  SetRevisionMergeEnabled(v bool) *ExternalMultiFileRevisionConfig
  GetRevisionMergeEnabled() *bool 
  SetRevisionRecyclePeriod(v int64) *ExternalMultiFileRevisionConfig
  GetRevisionRecyclePeriod() *int64 
}

type ExternalMultiFileRevisionConfig struct {
  RevisionCount *int64 `json:"revision_count,omitempty" xml:"revision_count,omitempty"`
  RevisionMergeEnabled *bool `json:"revision_merge_enabled,omitempty" xml:"revision_merge_enabled,omitempty"`
  RevisionRecyclePeriod *int64 `json:"revision_recycle_period,omitempty" xml:"revision_recycle_period,omitempty"`
}

func (s ExternalMultiFileRevisionConfig) String() string {
  return dara.Prettify(s)
}

func (s ExternalMultiFileRevisionConfig) GoString() string {
  return s.String()
}

func (s *ExternalMultiFileRevisionConfig) GetRevisionCount() *int64  {
  return s.RevisionCount
}

func (s *ExternalMultiFileRevisionConfig) GetRevisionMergeEnabled() *bool  {
  return s.RevisionMergeEnabled
}

func (s *ExternalMultiFileRevisionConfig) GetRevisionRecyclePeriod() *int64  {
  return s.RevisionRecyclePeriod
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionCount(v int64) *ExternalMultiFileRevisionConfig {
  s.RevisionCount = &v
  return s
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionMergeEnabled(v bool) *ExternalMultiFileRevisionConfig {
  s.RevisionMergeEnabled = &v
  return s
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionRecyclePeriod(v int64) *ExternalMultiFileRevisionConfig {
  s.RevisionRecyclePeriod = &v
  return s
}

func (s *ExternalMultiFileRevisionConfig) Validate() error {
  return dara.Validate(s)
}

