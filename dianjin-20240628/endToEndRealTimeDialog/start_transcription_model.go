// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iStartTranscription interface {
  dara.Model
  String() string
  GoString() string
  SetPlayCode(v string) *StartTranscription
  GetPlayCode() *string 
  SetMetaData(v map[string]interface{}) *StartTranscription
  GetMetaData() map[string]interface{} 
  SetSelfDirected(v bool) *StartTranscription
  GetSelfDirected() *bool 
}

type StartTranscription struct {
  // This parameter is required.
  PlayCode *string `json:"playCode,omitempty" xml:"playCode,omitempty"`
  MetaData map[string]interface{} `json:"metaData,omitempty" xml:"metaData,omitempty"`
  SelfDirected *bool `json:"selfDirected,omitempty" xml:"selfDirected,omitempty"`
}

func (s StartTranscription) String() string {
  return dara.Prettify(s)
}

func (s StartTranscription) GoString() string {
  return s.String()
}

func (s *StartTranscription) GetPlayCode() *string  {
  return s.PlayCode
}

func (s *StartTranscription) GetMetaData() map[string]interface{}  {
  return s.MetaData
}

func (s *StartTranscription) GetSelfDirected() *bool  {
  return s.SelfDirected
}

func (s *StartTranscription) SetPlayCode(v string) *StartTranscription {
  s.PlayCode = &v
  return s
}

func (s *StartTranscription) SetMetaData(v map[string]interface{}) *StartTranscription {
  s.MetaData = v
  return s
}

func (s *StartTranscription) SetSelfDirected(v bool) *StartTranscription {
  s.SelfDirected = &v
  return s
}

func (s *StartTranscription) Validate() error {
  return dara.Validate(s)
}

