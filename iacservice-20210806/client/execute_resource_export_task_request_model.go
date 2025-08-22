// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteResourceExportTaskRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClientToken(v string) *ExecuteResourceExportTaskRequest
  GetClientToken() *string 
}

type ExecuteResourceExportTaskRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // a65451293e64979ba7a4b573950217fe
  ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s ExecuteResourceExportTaskRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteResourceExportTaskRequest) GoString() string {
  return s.String()
}

func (s *ExecuteResourceExportTaskRequest) GetClientToken() *string  {
  return s.ClientToken
}

func (s *ExecuteResourceExportTaskRequest) SetClientToken(v string) *ExecuteResourceExportTaskRequest {
  s.ClientToken = &v
  return s
}

func (s *ExecuteResourceExportTaskRequest) Validate() error {
  return dara.Validate(s)
}

