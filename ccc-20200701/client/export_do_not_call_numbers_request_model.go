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
  // The ID of the instance.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // Specifies the scope of the do-not-call numbers. A value of SYSTEM applies to your entire Alibaba Cloud account, while INSTANCE applies only to the current instance. The default value is INSTANCE.
  // 
  // example:
  // 
  // INSTANCE
  Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
  // The keyword for a fuzzy search of phone numbers or remarks. If this parameter is left empty, no keyword-based filtering is applied.
  // 
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

