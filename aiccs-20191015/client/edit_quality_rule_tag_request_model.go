// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditQualityRuleTagRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAnalysisTypes(v []*EditQualityRuleTagRequestAnalysisTypes) *EditQualityRuleTagRequest
  GetAnalysisTypes() []*EditQualityRuleTagRequestAnalysisTypes 
  SetInstanceId(v string) *EditQualityRuleTagRequest
  GetInstanceId() *string 
}

type EditQualityRuleTagRequest struct {
  // This parameter is required.
  AnalysisTypes []*EditQualityRuleTagRequestAnalysisTypes `json:"AnalysisTypes,omitempty" xml:"AnalysisTypes,omitempty" type:"Repeated"`
  // This parameter is required.
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s EditQualityRuleTagRequest) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleTagRequest) GoString() string {
  return s.String()
}

func (s *EditQualityRuleTagRequest) GetAnalysisTypes() []*EditQualityRuleTagRequestAnalysisTypes  {
  return s.AnalysisTypes
}

func (s *EditQualityRuleTagRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *EditQualityRuleTagRequest) SetAnalysisTypes(v []*EditQualityRuleTagRequestAnalysisTypes) *EditQualityRuleTagRequest {
  s.AnalysisTypes = v
  return s
}

func (s *EditQualityRuleTagRequest) SetInstanceId(v string) *EditQualityRuleTagRequest {
  s.InstanceId = &v
  return s
}

func (s *EditQualityRuleTagRequest) Validate() error {
  if s.AnalysisTypes != nil {
    for _, item := range s.AnalysisTypes {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EditQualityRuleTagRequestAnalysisTypes struct {
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s EditQualityRuleTagRequestAnalysisTypes) String() string {
  return dara.Prettify(s)
}

func (s EditQualityRuleTagRequestAnalysisTypes) GoString() string {
  return s.String()
}

func (s *EditQualityRuleTagRequestAnalysisTypes) GetId() *int64  {
  return s.Id
}

func (s *EditQualityRuleTagRequestAnalysisTypes) GetName() *string  {
  return s.Name
}

func (s *EditQualityRuleTagRequestAnalysisTypes) SetId(v int64) *EditQualityRuleTagRequestAnalysisTypes {
  s.Id = &v
  return s
}

func (s *EditQualityRuleTagRequestAnalysisTypes) SetName(v string) *EditQualityRuleTagRequestAnalysisTypes {
  s.Name = &v
  return s
}

func (s *EditQualityRuleTagRequestAnalysisTypes) Validate() error {
  return dara.Validate(s)
}

