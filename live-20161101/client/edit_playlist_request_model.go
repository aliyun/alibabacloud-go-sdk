// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditPlaylistRequest interface {
  dara.Model
  String() string
  GoString() string
  SetOwnerId(v int64) *EditPlaylistRequest
  GetOwnerId() *int64 
  SetProgramConfig(v string) *EditPlaylistRequest
  GetProgramConfig() *string 
  SetProgramId(v string) *EditPlaylistRequest
  GetProgramId() *string 
  SetProgramItems(v string) *EditPlaylistRequest
  GetProgramItems() *string 
  SetRegionId(v string) *EditPlaylistRequest
  GetRegionId() *string 
}

type EditPlaylistRequest struct {
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The configurations of the episode list. For more information, see the **ProgramConfig*	- section of this topic.
  // 
  // example:
  // 
  // [{"RepeatNumber":"0","ProgramName":"my program"}]
  ProgramConfig *string `json:"ProgramConfig,omitempty" xml:"ProgramConfig,omitempty"`
  // The ID of the episode list. If the episode list was created by calling the [AddPlaylistItems](https://help.aliyun.com/document_detail/2848078.html) operation, check the value of the response parameter ProgramId to obtain the ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 445409ec-7eaa-461d-8f29-4bec2eb9****
  ProgramId *string `json:"ProgramId,omitempty" xml:"ProgramId,omitempty"`
  // The episodes that you want to add to the production studio. The value is a JSON string. For more information, see the **InputProgramItem*	- section of this topic.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // [{"ItemName":"item1","ResourceType":"vod","ResourceValue":"5f8809f2-3352-4d1f-a8f7-86f9429f****"}, {"ItemName": "item2","ResourceType": "vod","ResourceValue": "e7411c0b-dd98-4c61-a545-f8bfba6c****"}]
  ProgramItems *string `json:"ProgramItems,omitempty" xml:"ProgramItems,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s EditPlaylistRequest) String() string {
  return dara.Prettify(s)
}

func (s EditPlaylistRequest) GoString() string {
  return s.String()
}

func (s *EditPlaylistRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EditPlaylistRequest) GetProgramConfig() *string  {
  return s.ProgramConfig
}

func (s *EditPlaylistRequest) GetProgramId() *string  {
  return s.ProgramId
}

func (s *EditPlaylistRequest) GetProgramItems() *string  {
  return s.ProgramItems
}

func (s *EditPlaylistRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EditPlaylistRequest) SetOwnerId(v int64) *EditPlaylistRequest {
  s.OwnerId = &v
  return s
}

func (s *EditPlaylistRequest) SetProgramConfig(v string) *EditPlaylistRequest {
  s.ProgramConfig = &v
  return s
}

func (s *EditPlaylistRequest) SetProgramId(v string) *EditPlaylistRequest {
  s.ProgramId = &v
  return s
}

func (s *EditPlaylistRequest) SetProgramItems(v string) *EditPlaylistRequest {
  s.ProgramItems = &v
  return s
}

func (s *EditPlaylistRequest) SetRegionId(v string) *EditPlaylistRequest {
  s.RegionId = &v
  return s
}

func (s *EditPlaylistRequest) Validate() error {
  return dara.Validate(s)
}

