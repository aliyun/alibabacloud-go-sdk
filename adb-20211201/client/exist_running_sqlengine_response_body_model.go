// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExistRunningSQLEngineResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetData(v bool) *ExistRunningSQLEngineResponseBody
  GetData() *bool 
  SetRequestId(v string) *ExistRunningSQLEngineResponseBody
  GetRequestId() *string 
}

type ExistRunningSQLEngineResponseBody struct {
  // Indicates whether a running SQL engine exists in the resource group.
  // 
  // Valid values:
  // 
  // 	- **True**
  // 
  // 	- **False**
  // 
  // example:
  // 
  // True
  Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
  // The request ID.
  // 
  // example:
  // 
  // FA675D68-14A4-5D9C-8820-92537D9F447E
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExistRunningSQLEngineResponseBody) String() string {
  return dara.Prettify(s)
}

func (s ExistRunningSQLEngineResponseBody) GoString() string {
  return s.String()
}

func (s *ExistRunningSQLEngineResponseBody) GetData() *bool  {
  return s.Data
}

func (s *ExistRunningSQLEngineResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *ExistRunningSQLEngineResponseBody) SetData(v bool) *ExistRunningSQLEngineResponseBody {
  s.Data = &v
  return s
}

func (s *ExistRunningSQLEngineResponseBody) SetRequestId(v string) *ExistRunningSQLEngineResponseBody {
  s.RequestId = &v
  return s
}

func (s *ExistRunningSQLEngineResponseBody) Validate() error {
  return dara.Validate(s)
}

