// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportGeneratedContentRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgentKey(v string) *ExportGeneratedContentRequest
  GetAgentKey() *string 
  SetId(v int64) *ExportGeneratedContentRequest
  GetId() *int64 
}

type ExportGeneratedContentRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // xxxxx_p_efm
  AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ExportGeneratedContentRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportGeneratedContentRequest) GoString() string {
  return s.String()
}

func (s *ExportGeneratedContentRequest) GetAgentKey() *string  {
  return s.AgentKey
}

func (s *ExportGeneratedContentRequest) GetId() *int64  {
  return s.Id
}

func (s *ExportGeneratedContentRequest) SetAgentKey(v string) *ExportGeneratedContentRequest {
  s.AgentKey = &v
  return s
}

func (s *ExportGeneratedContentRequest) SetId(v int64) *ExportGeneratedContentRequest {
  s.Id = &v
  return s
}

func (s *ExportGeneratedContentRequest) Validate() error {
  return dara.Validate(s)
}

