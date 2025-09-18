// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneRequest interface {
  dara.Model
  String() string
  GoString() string
  SetSceneId(v string) *ExecuteSceneRequest
  GetSceneId() *string 
}

type ExecuteSceneRequest struct {
  // example:
  // 
  // a84a55aa410e460a9ac753570c76fecc
  SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ExecuteSceneRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSceneRequest) GetSceneId() *string  {
  return s.SceneId
}

func (s *ExecuteSceneRequest) SetSceneId(v string) *ExecuteSceneRequest {
  s.SceneId = &v
  return s
}

func (s *ExecuteSceneRequest) Validate() error {
  return dara.Validate(s)
}

