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
  // - If you create a production studio using the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value from the response.
  // 
  // - If you create a production studio in the ApsaraVideo Live console, find the ID on the **Production Studio*	- > **Cloud Production Studio*	- page.
  // 
  // > The name of the production studio in the list is the ID.
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
  // The ID of the resource. To get this ID, call the [DescribeCasterChannels](https://help.aliyun.com/document_detail/2848046.html) operation and check the ResourceId value in the response.
  // 
  // This parameter is required.
  // 
  // example:
  // 
  // f096e8d6-0319-4c96-82bc-ecbc79cf****
  ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
  // The ID of the scenario.
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

