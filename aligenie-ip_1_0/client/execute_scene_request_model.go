// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecuteSceneRequest interface {
  dara.Model
  String() string
  GoString() string
  SetHotelId(v string) *ExecuteSceneRequest
  GetHotelId() *string 
  SetRoomNo(v string) *ExecuteSceneRequest
  GetRoomNo() *string 
  SetSceneName(v string) *ExecuteSceneRequest
  GetSceneName() *string 
}

type ExecuteSceneRequest struct {
  // This parameter is required.
  // 
  // example:
  // 
  // 80d84ea8ed9e422fbad52715c8fc56f1
  HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1211
  RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
  // This parameter is required.
  SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ExecuteSceneRequest) String() string {
  return dara.Prettify(s)
}

func (s ExecuteSceneRequest) GoString() string {
  return s.String()
}

func (s *ExecuteSceneRequest) GetHotelId() *string  {
  return s.HotelId
}

func (s *ExecuteSceneRequest) GetRoomNo() *string  {
  return s.RoomNo
}

func (s *ExecuteSceneRequest) GetSceneName() *string  {
  return s.SceneName
}

func (s *ExecuteSceneRequest) SetHotelId(v string) *ExecuteSceneRequest {
  s.HotelId = &v
  return s
}

func (s *ExecuteSceneRequest) SetRoomNo(v string) *ExecuteSceneRequest {
  s.RoomNo = &v
  return s
}

func (s *ExecuteSceneRequest) SetSceneName(v string) *ExecuteSceneRequest {
  s.SceneName = &v
  return s
}

func (s *ExecuteSceneRequest) Validate() error {
  return dara.Validate(s)
}

