// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditWorkspaceQueueResponse interface {
  dara.Model
  String() string
  GoString() string
  SetHeaders(v map[string]*string) *EditWorkspaceQueueResponse
  GetHeaders() map[string]*string 
  SetStatusCode(v int32) *EditWorkspaceQueueResponse
  GetStatusCode() *int32 
  SetBody(v *EditWorkspaceQueueResponseBody) *EditWorkspaceQueueResponse
  GetBody() *EditWorkspaceQueueResponseBody 
}

type EditWorkspaceQueueResponse struct {
  Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
  StatusCode *int32 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
  Body *EditWorkspaceQueueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EditWorkspaceQueueResponse) String() string {
  return dara.Prettify(s)
}

func (s EditWorkspaceQueueResponse) GoString() string {
  return s.String()
}

func (s *EditWorkspaceQueueResponse) GetHeaders() map[string]*string  {
  return s.Headers
}

func (s *EditWorkspaceQueueResponse) GetStatusCode() *int32  {
  return s.StatusCode
}

func (s *EditWorkspaceQueueResponse) GetBody() *EditWorkspaceQueueResponseBody  {
  return s.Body
}

func (s *EditWorkspaceQueueResponse) SetHeaders(v map[string]*string) *EditWorkspaceQueueResponse {
  s.Headers = v
  return s
}

func (s *EditWorkspaceQueueResponse) SetStatusCode(v int32) *EditWorkspaceQueueResponse {
  s.StatusCode = &v
  return s
}

func (s *EditWorkspaceQueueResponse) SetBody(v *EditWorkspaceQueueResponseBody) *EditWorkspaceQueueResponse {
  s.Body = v
  return s
}

func (s *EditWorkspaceQueueResponse) Validate() error {
  if s.Body != nil {
    if err := s.Body.Validate(); err != nil {
      return err
    }
  }
  return nil
}

