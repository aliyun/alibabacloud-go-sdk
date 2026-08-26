// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectCasterUrgentRequest interface {
  dara.Model
  String() string
  GoString() string
  SetCasterId(v string) *EffectCasterUrgentRequest
  GetCasterId() *string 
  SetOwnerId(v int64) *EffectCasterUrgentRequest
  GetOwnerId() *int64 
  SetRegionId(v string) *EffectCasterUrgentRequest
  GetRegionId() *string 
  SetSceneId(v string) *EffectCasterUrgentRequest
  GetSceneId() *string 
}

type EffectCasterUrgentRequest struct {
  // The ID of the production studio.
  // 
  // - If you created a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the returned CasterId value.
  // 
  // - If you created a production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to obtain the ID.
  // 
  // > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 80787064-1c94-4dc1-85ce-9409960a****
  CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
  OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
  // The ID of the region.
  // 
  // example:
  // 
  // cn-shanghai
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // The ID of the scene. This parameter is valid only for PGM scenes.
  // 
  // Call the [DescribeCasterScenes](https://help.aliyun.com/document_detail/2848039.html) operation to query the scene ID of the production studio.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // 242b4e2c-c30f-4442-85ba-2e3e4e3d****
  SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s EffectCasterUrgentRequest) String() string {
  return dara.Prettify(s)
}

func (s EffectCasterUrgentRequest) GoString() string {
  return s.String()
}

func (s *EffectCasterUrgentRequest) GetCasterId() *string  {
  return s.CasterId
}

func (s *EffectCasterUrgentRequest) GetOwnerId() *int64  {
  return s.OwnerId
}

func (s *EffectCasterUrgentRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EffectCasterUrgentRequest) GetSceneId() *string  {
  return s.SceneId
}

func (s *EffectCasterUrgentRequest) SetCasterId(v string) *EffectCasterUrgentRequest {
  s.CasterId = &v
  return s
}

func (s *EffectCasterUrgentRequest) SetOwnerId(v int64) *EffectCasterUrgentRequest {
  s.OwnerId = &v
  return s
}

func (s *EffectCasterUrgentRequest) SetRegionId(v string) *EffectCasterUrgentRequest {
  s.RegionId = &v
  return s
}

func (s *EffectCasterUrgentRequest) SetSceneId(v string) *EffectCasterUrgentRequest {
  s.SceneId = &v
  return s
}

func (s *EffectCasterUrgentRequest) Validate() error {
  return dara.Validate(s)
}

