// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *SetCasterSceneConfigRequest
	GetCasterId() *string
	SetComponentId(v []*string) *SetCasterSceneConfigRequest
	GetComponentId() []*string
	SetLayoutId(v string) *SetCasterSceneConfigRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *SetCasterSceneConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetCasterSceneConfigRequest
	GetRegionId() *string
	SetSceneId(v string) *SetCasterSceneConfigRequest
	GetSceneId() *string
}

type SetCasterSceneConfigRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **China Cloud Production Studio*	- to view the ID.
	//
	// > The name of the production studio in the production studio list on the China Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The list of component IDs. The components are arranged in bottom-to-top order within the array.
	//
	// >N indicates the sequence number. For example:<br>ComponentId.1 indicates the first component ID.<br>ComponentId.2 indicates the second component ID.
	//
	// example:
	//
	// [ "a2b8e671-2fe5-4642-a2ec-bf931826****", "a2b8e671-2fe5-4642-a2ec-28374657****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The layout ID. If you call the [DescribeCasterLayouts](https://help.aliyun.com/document_detail/2848028.html) operation to query the layout list of a production studio, check the LayoutId parameter returned by the DescribeCasterLayouts operation.
	//
	// example:
	//
	// 0c6da077-f037-49e8-8440-3be13393****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scene ID. The scene must have been started by calling StartCasterScene. Otherwise, the IncorrectSceneStatus error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 242b4e2c-c30f-4442-85ba-2e3e4e3d****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s SetCasterSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *SetCasterSceneConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *SetCasterSceneConfigRequest) GetComponentId() []*string {
	return s.ComponentId
}

func (s *SetCasterSceneConfigRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *SetCasterSceneConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetCasterSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetCasterSceneConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *SetCasterSceneConfigRequest) SetCasterId(v string) *SetCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetComponentId(v []*string) *SetCasterSceneConfigRequest {
	s.ComponentId = v
	return s
}

func (s *SetCasterSceneConfigRequest) SetLayoutId(v string) *SetCasterSceneConfigRequest {
	s.LayoutId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetOwnerId(v int64) *SetCasterSceneConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetRegionId(v string) *SetCasterSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) SetSceneId(v string) *SetCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *SetCasterSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
