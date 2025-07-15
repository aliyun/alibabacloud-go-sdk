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
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation, check the value of the response parameter CasterId to obtain the ID.
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
	// The IDs of the components. Components in the scene are listed from the bottom to the top in an array.
	//
	// >  N indicates a sequence number. Examples:\\
	//
	// ComponentId.1 indicates the ID of the first component.\\
	//
	// ComponentId.2 indicates the ID of the second component.
	//
	// example:
	//
	// [ "a2b8e671-2fe5-4642-a2ec-bf931826****", "a2b8e671-2fe5-4642-a2ec-28374657****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The ID of the layout. If you call the [DescribeCasterLayouts](https://help.aliyun.com/document_detail/60260.html) operation to query the layouts of the production studio, check the value of the response parameter LayoutId to obtain the ID.
	//
	// example:
	//
	// 0c6da077-f037-49e8-8440-3be13393****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
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
