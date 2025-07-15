// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExportIntervenesRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAgentKey(v string) *ExportIntervenesRequest
  GetAgentKey() *string 
}

type ExportIntervenesRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // fed6555ec9e24b92aeecc34be484b887_p_efm
  AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
}

func (s ExportIntervenesRequest) String() string {
  return dara.Prettify(s)
}

func (s ExportIntervenesRequest) GoString() string {
  return s.String()
}

func (s *ExportIntervenesRequest) GetAgentKey() *string  {
  return s.AgentKey
}

func (s *ExportIntervenesRequest) SetAgentKey(v string) *ExportIntervenesRequest {
  s.AgentKey = &v
  return s
}

func (s *ExportIntervenesRequest) Validate() error {
  return dara.Validate(s)
}

