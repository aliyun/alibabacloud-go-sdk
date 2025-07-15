// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterVideoResourceRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCasterId(v string) *EffectCasterVideoResourceRequest
  GetCasterId() *string 
  SetOwnerId(v int64) *EffectCasterVideoResourceRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EffectCasterVideoResourceRequest
  GetRegionId() *string 
  SetResourceId(v string) *EffectCasterVideoResourceRequest
  GetResourceId() *string 
  SetSceneId(v string) *EffectCasterVideoResourceRequest
  GetSceneId() *string 
}

type EffectCasterVideoResourceRequest struct {
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
  // 80787064-1c94-4dc1-85ce-9409960a****
  CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The resource ID. If you call the [DescribeCasterChannels](https://help.aliyun.com/document_detail/2848046.html) operation to query the channels of the production studio, you can obtain the ID of the resource in a specific channel from the ResourceId parameter in the response.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // f096e8d6-0319-4c96-82bc-ecbc79cf****
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
  // The ID of the scene. You can call the [DescribeCasterScenes](~~60262#doc-api-live-DescribeCasterScenes~~ "Queries information about the scenes of a production studio.") operation to get the ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 05ab713c-676e-49c0-96ce-cc408da1****
  SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s EffectCasterVideoResourceRequest) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterVideoResourceRequest) GoString() string {
  return s.String()
}

func (s *EffectCasterVideoResourceRequest) GetCasterId() *string  {
  return s.CasterId
}

func (s *EffectCasterVideoResourceRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EffectCasterVideoResourceRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EffectCasterVideoResourceRequest) GetResourceId() *string  {
  return s.ResourceId
}

func (s *EffectCasterVideoResourceRequest) GetSceneId() *string  {
  return s.SceneId
}

func (s *EffectCasterVideoResourceRequest) SetCasterId(v string) *EffectCasterVideoResourceRequest {
  s.CasterId = &v
  return s
}

func (s *EffectCasterVideoResourceRequest) SetOwnerId(v int64) *EffectCasterVideoResourceRequest {
  s.OwnerId = &v
  return s
}

func (s *EffectCasterVideoResourceRequest) SetRegionId(v string) *EffectCasterVideoResourceRequest {
  s.RegionId = &v
  return s
}

func (s *EffectCasterVideoResourceRequest) SetResourceId(v string) *EffectCasterVideoResourceRequest {
  s.ResourceId = &v
  return s
}

func (s *EffectCasterVideoResourceRequest) SetSceneId(v string) *EffectCasterVideoResourceRequest {
  s.SceneId = &v
  return s
}

func (s *EffectCasterVideoResourceRequest) Validate() error {
  return dara.Validate(s)
}

