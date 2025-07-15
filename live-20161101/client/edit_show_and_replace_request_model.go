// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditShowAndReplaceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCasterId(v string) *EditShowAndReplaceRequest
  GetCasterId() *string 
  SetEndTime(v float32) *EditShowAndReplaceRequest
  GetEndTime() *float32 
  SetOwnerId(v int64) *EditShowAndReplaceRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EditShowAndReplaceRequest
  GetRegionId() *string 
  SetShowId(v string) *EditShowAndReplaceRequest
  GetShowId() *string 
  SetStartTime(v float32) *EditShowAndReplaceRequest
  GetStartTime() *float32 
  SetStorageInfo(v string) *EditShowAndReplaceRequest
  GetStorageInfo() *string 
  SetUserData(v string) *EditShowAndReplaceRequest
  GetUserData() *string 
}

type EditShowAndReplaceRequest struct {
  // The ID of the production studio.
  // 
  // 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
  // 
  // 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
  // 
  // >  You can find the ID of the production studio in the Instance ID/Name column.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 53200b81-b761-4c10-842a-a0726d97****
  CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
  // The end time of the editing task. Unit: seconds.
  // 
  // > 
  // 
  // 	- The valid values range from 0 to the value indicated by the total length of the episode.
  // 
  // 	- By default, this parameter is set to the value that indicates the total length of the episode. The editing period cannot exceed the total length of the episode.
  // 
  // 	- If you want to edit a VOD file from the 2nd second to the 5th second, set the StartTime parameter to 2.0 and the EndTime parameter to 5.0.
  // 
  // example:
  // 
  // 5.0
  EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the episode to be edited.
  // 
  // >  You can obtain the ID from the response parameter ShowId of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 42200b81-b761-4c10-842a-a0726d97****
  ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
  // The start time of the editing task. Unit: seconds.
  // 
  // > 
  // 
  // 	- The valid values range from 0 to the value indicated by the total length of the episode. By default, the editing task starts from the beginning of the episode. Default value: 0.0.
  // 
  // 	- If you want to edit a VOD file from the 2nd second to the 5th second, set the StartTime parameter to 2.0 and the EndTime parameter to 5.0.
  // 
  // example:
  // 
  // 2.0
  StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
  // The storage information of the episode. The following fields are included:
  // 
  // 	- **StorageLocation**: the storage location of ApsaraVideo VOD.
  // 
  // 	- **FileName**: the custom file name.
  // 
  // >  Editing outputs must be stored in the VOD bucket within the same account that is used to access both ApsaraVideo VOD and ApsaraVideo Live. For more information about how to obtain the storage location, see [Manage VOD resources](https://help.aliyun.com/document_detail/86097.html).
  // 
  // example:
  // 
  // { "StorageLocation":"***bucket***", "FileName":"EditFile****.mp4" }
  StorageInfo *string `json:"StorageInfo,omitempty" xml:"StorageInfo,omitempty"`
  // The user information.
  // 
  // example:
  // 
  // 900a2b2r8-13c2-****-88f2-75e4a07c1ed9
  UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s EditShowAndReplaceRequest) String() string {
  return dara.Prettify(s)
}

func (s EditShowAndReplaceRequest) GoString() string {
  return s.String()
}

func (s *EditShowAndReplaceRequest) GetCasterId() *string  {
  return s.CasterId
}

func (s *EditShowAndReplaceRequest) GetEndTime() *float32  {
  return s.EndTime
}

func (s *EditShowAndReplaceRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EditShowAndReplaceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EditShowAndReplaceRequest) GetShowId() *string  {
  return s.ShowId
}

func (s *EditShowAndReplaceRequest) GetStartTime() *float32  {
  return s.StartTime
}

func (s *EditShowAndReplaceRequest) GetStorageInfo() *string  {
  return s.StorageInfo
}

func (s *EditShowAndReplaceRequest) GetUserData() *string  {
  return s.UserData
}

func (s *EditShowAndReplaceRequest) SetCasterId(v string) *EditShowAndReplaceRequest {
  s.CasterId = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetEndTime(v float32) *EditShowAndReplaceRequest {
  s.EndTime = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetOwnerId(v int64) *EditShowAndReplaceRequest {
  s.OwnerId = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetRegionId(v string) *EditShowAndReplaceRequest {
  s.RegionId = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetShowId(v string) *EditShowAndReplaceRequest {
  s.ShowId = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetStartTime(v float32) *EditShowAndReplaceRequest {
  s.StartTime = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetStorageInfo(v string) *EditShowAndReplaceRequest {
  s.StorageInfo = &v
  return s
}

func (s *EditShowAndReplaceRequest) SetUserData(v string) *EditShowAndReplaceRequest {
  s.UserData = &v
  return s
}

func (s *EditShowAndReplaceRequest) Validate() error {
  return dara.Validate(s)
}

