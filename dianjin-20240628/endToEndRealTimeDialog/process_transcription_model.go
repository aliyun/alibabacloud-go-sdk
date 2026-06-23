// This file is auto-generated, don't edit it. Thanks.
package endToEndRealTimeDialog

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iProcessTranscription interface {
  dara.Model
  String() string
  GoString() string
  SetDataSourceType(v string) *ProcessTranscription
  GetDataSourceType() *string 
  SetData(v []*int64) *ProcessTranscription
  GetData() []*int64 
}

type ProcessTranscription struct {
  DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
  // This parameter is required.
  Data []*int64 `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s ProcessTranscription) String() string {
  return dara.Prettify(s)
}

func (s ProcessTranscription) GoString() string {
  return s.String()
}

func (s *ProcessTranscription) GetDataSourceType() *string  {
  return s.DataSourceType
}

func (s *ProcessTranscription) GetData() []*int64  {
  return s.Data
}

func (s *ProcessTranscription) SetDataSourceType(v string) *ProcessTranscription {
  s.DataSourceType = &v
  return s
}

func (s *ProcessTranscription) SetData(v []*int64) *ProcessTranscription {
  s.Data = v
  return s
}

func (s *ProcessTranscription) Validate() error {
  return dara.Validate(s)
}

