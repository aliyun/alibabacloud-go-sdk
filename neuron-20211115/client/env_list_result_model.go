// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvListResult interface {
  dara.Model
  String() string
  GoString() string
  SetEnvList(v []*string) *EnvListResult
  GetEnvList() []*string 
  SetRequestId(v string) *EnvListResult
  GetRequestId() *string 
}

type EnvListResult struct {
  EnvList []*string `json:"envList,omitempty" xml:"envList,omitempty" type:"Repeated"`
  // example:
  // 
  // 3239281273464326823
  RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s EnvListResult) String() string {
  return dara.Prettify(s)
}

func (s EnvListResult) GoString() string {
  return s.String()
}

func (s *EnvListResult) GetEnvList() []*string  {
  return s.EnvList
}

func (s *EnvListResult) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnvListResult) SetEnvList(v []*string) *EnvListResult {
  s.EnvList = v
  return s
}

func (s *EnvListResult) SetRequestId(v string) *EnvListResult {
  s.RequestId = &v
  return s
}

func (s *EnvListResult) Validate() error {
  return dara.Validate(s)
}

