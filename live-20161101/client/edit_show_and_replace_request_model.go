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
  // The production studio ID.
  // 
  // - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter in the response.
  // 
  // - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the ID.
  // 
  // > The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 53200b81-b761-4c10-842a-a0726d97****
  CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
  // The end time of the video clip. Unit: seconds.
  // 
  // > - The valid range of the clip time is 0 to the total duration of the show.
  // 
  // > - The default value is the end time of the video-on-demand file. The value cannot exceed the total duration of the show.
  // 
  // > - For example, to clip a video-on-demand file from the 2nd second to the 5th second, set StartTime to 2.0 and EndTime to 5.0.
  // 
  // > - You must specify at least one of StartTime and EndTime.
  // 
  // example:
  // 
  // 5.0
  EndTime *float32 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The region ID.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the show to be clipped. The referenced show must be of the video-on-demand material type (ResourceInfo.ResourceType=vod with a valid resourceId).
  // 
  // > Obtain the ShowId value from the response parameters of the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) operation.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 42200b81-b761-4c10-842a-a0726d97****
  ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
  // The start time of the video clip. Unit: seconds.
  // 
  // > - The valid range of the clip time is 0 to the total duration of the show. - By default, the clip starts from the beginning of the video-on-demand file. Value: 0.0.
  // 
  // > - For example, to clip a video-on-demand file from the 2nd second to the 5th second, set StartTime to 2.0 and EndTime to 5.0.
  // 
  // > - You must specify at least one of StartTime and EndTime.
  // 
  // example:
  // 
  // 2.0
  StartTime *float32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
  // The storage information. This parameter is required. Description:
  // 
  // - **StorageLocation**: the video-on-demand storage address of the user.
  // 
  // - **FileName**: the custom file name.
  // 
  // > The video clip storage address must be a video-on-demand storage address under the same account. To obtain the video-on-demand storage address, see [Storage management](https://help.aliyun.com/document_detail/86097.html).
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

