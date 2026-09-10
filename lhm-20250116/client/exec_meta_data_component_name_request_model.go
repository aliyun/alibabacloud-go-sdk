// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecMetaDataComponentNameRequest interface {
  dara.Model
  String() string
  GoString() string
  SetDsName(v string) *ExecMetaDataComponentNameRequest
  GetDsName() *string 
}

type ExecMetaDataComponentNameRequest struct {
  // The datasource name to check. The system performs an exact match against non-deleted datasources under the current tenant.
  // 
  // example:
  // 
  // test_ds318_hangzhou_0428
  DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
}

func (s ExecMetaDataComponentNameRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecMetaDataComponentNameRequest) GoString() string {
  return s.String()
}

func (s *ExecMetaDataComponentNameRequest) GetDsName() *string  {
  return s.DsName
}

func (s *ExecMetaDataComponentNameRequest) SetDsName(v string) *ExecMetaDataComponentNameRequest {
  s.DsName = &v
  return s
}

func (s *ExecMetaDataComponentNameRequest) Validate() error {
  return dara.Validate(s)
}

