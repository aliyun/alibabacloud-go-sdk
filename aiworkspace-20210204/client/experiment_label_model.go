// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExperimentLabel interface {
  dara.Model
  String() string
  GoString() string
  SetExperimentId(v string) *ExperimentLabel
  GetExperimentId() *string 
  SetGmtCreateTime(v string) *ExperimentLabel
  GetGmtCreateTime() *string 
  SetGmtModifiedTime(v string) *ExperimentLabel
  GetGmtModifiedTime() *string 
  SetKey(v string) *ExperimentLabel
  GetKey() *string 
  SetValue(v string) *ExperimentLabel
  GetValue() *string 
}

type ExperimentLabel struct {
  // example:
  // 
  // exp-890waerw09a0f
  ExperimentId *string `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
  // example:
  // 
  // 2023-12-27T03:30:04Z
  GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
  // example:
  // 
  // 2023-12-27T03:30:04Z
  GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
  // example:
  // 
  // key
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // example:
  // 
  // value
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExperimentLabel) String() string {
  return dara.Prettify(s)
}

func (s ExperimentLabel) GoString() string {
  return s.String()
}

func (s *ExperimentLabel) GetExperimentId() *string  {
  return s.ExperimentId
}

func (s *ExperimentLabel) GetGmtCreateTime() *string  {
  return s.GmtCreateTime
}

func (s *ExperimentLabel) GetGmtModifiedTime() *string  {
  return s.GmtModifiedTime
}

func (s *ExperimentLabel) GetKey() *string  {
  return s.Key
}

func (s *ExperimentLabel) GetValue() *string  {
  return s.Value
}

func (s *ExperimentLabel) SetExperimentId(v string) *ExperimentLabel {
  s.ExperimentId = &v
  return s
}

func (s *ExperimentLabel) SetGmtCreateTime(v string) *ExperimentLabel {
  s.GmtCreateTime = &v
  return s
}

func (s *ExperimentLabel) SetGmtModifiedTime(v string) *ExperimentLabel {
  s.GmtModifiedTime = &v
  return s
}

func (s *ExperimentLabel) SetKey(v string) *ExperimentLabel {
  s.Key = &v
  return s
}

func (s *ExperimentLabel) SetValue(v string) *ExperimentLabel {
  s.Value = &v
  return s
}

func (s *ExperimentLabel) Validate() error {
  return dara.Validate(s)
}

