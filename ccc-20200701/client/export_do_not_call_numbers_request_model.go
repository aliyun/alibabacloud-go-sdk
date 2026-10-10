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
  // The instance ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // ccc-test
  InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
  // The application scope. Valid values: SYSTEM and INSTANCE. SYSTEM indicates system-level do-not-call, and INSTANCE indicates customer-defined do-not-call. SYSTEM is associated with the Alibaba Cloud account to which the instance belongs, and INSTANCE is associated only with the current instance. This parameter is optional. Default value: INSTANCE.
  // 
  // example:
  // 
  // INSTANCE
  Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
  // Specifies the keyword to perform a fuzzy match based on the phone number or remark. This parameter is optional. Default value: empty. An empty value indicates that no filtering is applied.
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

