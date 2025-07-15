// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportDoNotCallNumbersRequest interface {
  dara.Model
  String() string
  GoString() string
  SetInstanceId(v string) *ExportDoNotCallNumbersRequest
  GetInstanceId() *string 
  SetScope(v string) *ExportDoNotCallNumbersRequest
  GetScope() *string 
  SetSearchPattern(v string) *ExportDoNotCallNumbersRequest
  GetSearchPattern() *string 
}

type ExportDoNotCallNumbersRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // example:
  // 
  // INSTANCE
  Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
  // example:
  // 
  // RemarkA
  SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ExportDoNotCallNumbersRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportDoNotCallNumbersRequest) GoString() string {
  return s.String()
}

func (s *ExportDoNotCallNumbersRequest) GetInstanceId() *string  {
  return s.InstanceId
}

func (s *ExportDoNotCallNumbersRequest) GetScope() *string  {
  return s.Scope
}

func (s *ExportDoNotCallNumbersRequest) GetSearchPattern() *string  {
  return s.SearchPattern
}

func (s *ExportDoNotCallNumbersRequest) SetInstanceId(v string) *ExportDoNotCallNumbersRequest {
  s.InstanceId = &v
  return s
}

func (s *ExportDoNotCallNumbersRequest) SetScope(v string) *ExportDoNotCallNumbersRequest {
  s.Scope = &v
  return s
}

func (s *ExportDoNotCallNumbersRequest) SetSearchPattern(v string) *ExportDoNotCallNumbersRequest {
  s.SearchPattern = &v
  return s
}

func (s *ExportDoNotCallNumbersRequest) Validate() error {
  return dara.Validate(s)
}

