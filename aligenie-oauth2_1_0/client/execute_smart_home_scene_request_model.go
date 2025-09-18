// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSmartHomeSceneRequest interface {
  dara.Model
  String() string
  GoString() string
  SetFamilyId(v string) *ExecuteSmartHomeSceneRequest
  GetFamilyId() *string 
  SetSceneId(v string) *ExecuteSmartHomeSceneRequest
  GetSceneId() *string 
}

type ExecuteSmartHomeSceneRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 2iS1AH5eo8qrw1PYBL/Ulq==
  FamilyId *string `json:"FamilyId,omitempty" xml:"FamilyId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // Y1lCALepjYmTEouxsTrkjB==
  SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ExecuteSmartHomeSceneRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSmartHomeSceneRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSmartHomeSceneRequest) GetFamilyId() *string  {
  return s.FamilyId
}

func (s *ExecuteSmartHomeSceneRequest) GetSceneId() *string  {
  return s.SceneId
}

func (s *ExecuteSmartHomeSceneRequest) SetFamilyId(v string) *ExecuteSmartHomeSceneRequest {
  s.FamilyId = &v
  return s
}

func (s *ExecuteSmartHomeSceneRequest) SetSceneId(v string) *ExecuteSmartHomeSceneRequest {
  s.SceneId = &v
  return s
}

func (s *ExecuteSmartHomeSceneRequest) Validate() error {
  return dara.Validate(s)
}

