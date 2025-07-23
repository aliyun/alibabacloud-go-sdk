// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteOperationSyncResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *ExecuteOperationSyncResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *ExecuteOperationSyncResponse
  GetStatusCode() *int32 
  SetBody(v *ExecuteOperationSyncResponseBody) *ExecuteOperationSyncResponse
  GetBody() *ExecuteOperationSyncResponseBody 
}

type ExecuteOperationSyncResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *ExecuteOperationSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteOperationSyncResponse) String() string {
  return dara.Prettify(s)
}

func (s ExecuteOperationSyncResponse) GoString() string {
  return s.String()
}

func (s *ExecuteOperationSyncResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *ExecuteOperationSyncResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *ExecuteOperationSyncResponse) GetBody() *ExecuteOperationSyncResponseBody  {
  return s.Body
}

func (s *ExecuteOperationSyncResponse) SetHeaders(v map[string]*string) *ExecuteOperationSyncResponse {
  s.Headers = v
  return s
}

func (s *ExecuteOperationSyncResponse) SetStatusCode(v int32) *ExecuteOperationSyncResponse {
  s.StatusCode = &v
  return s
}

func (s *ExecuteOperationSyncResponse) SetBody(v *ExecuteOperationSyncResponseBody) *ExecuteOperationSyncResponse {
  s.Body = v
  return s
}

func (s *ExecuteOperationSyncResponse) Validate() error {
  return dara.Validate(s)
}

