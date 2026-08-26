// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *UpdateCasterSceneConfigRequest
	GetCasterId() *string
	SetComponentId(v []*string) *UpdateCasterSceneConfigRequest
	GetComponentId() []*string
	SetLayoutId(v string) *UpdateCasterSceneConfigRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *UpdateCasterSceneConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateCasterSceneConfigRequest
	GetRegionId() *string
	SetSceneId(v string) *UpdateCasterSceneConfigRequest
	GetSceneId() *string
}

type UpdateCasterSceneConfigRequest struct {
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value from the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, find the ID on the **Cloud Production Studio*	- page. Navigate to this page by choosing **Production Studio*	- > **Cloud Production Studio*	- in the ApsaraVideo Live console.
	//
	// > The name of the production studio in the list is its ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// A list of component IDs. The components in the array are layered from bottom to top.
	//
	// > N indicates the sequence number. For example, ComponentId.1 is the ID of the first component and ComponentId.2 is the ID of the second component.
	//
	// example:
	//
	// ["98778372-c30f-4442-85ba-2e3e4e3d****"]
	ComponentId []*string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty" type:"Repeated"`
	// The ID of the layout. If you query the layout list for the production studio by calling the [DescribeCasterLayouts](https://help.aliyun.com/document_detail/2848028.html) operation, use the LayoutId value from the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// eeab74fb-379d-4599-a93d-86d16a05****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
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

func (s UpdateCasterSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *UpdateCasterSceneConfigRequest) GetComponentId() []*string {
	return s.ComponentId
}

func (s *UpdateCasterSceneConfigRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *UpdateCasterSceneConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateCasterSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateCasterSceneConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateCasterSceneConfigRequest) SetCasterId(v string) *UpdateCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetComponentId(v []*string) *UpdateCasterSceneConfigRequest {
	s.ComponentId = v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetLayoutId(v string) *UpdateCasterSceneConfigRequest {
	s.LayoutId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetOwnerId(v int64) *UpdateCasterSceneConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetRegionId(v string) *UpdateCasterSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) SetSceneId(v string) *UpdateCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateCasterSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
