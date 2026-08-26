// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterSceneConfigRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DeleteCasterSceneConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterSceneConfigRequest
	GetRegionId() *string
	SetSceneId(v string) *DeleteCasterSceneConfigRequest
	GetSceneId() *string
	SetType(v string) *DeleteCasterSceneConfigRequest
	GetType() *string
}

type DeleteCasterSceneConfigRequest struct {
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **China Cloud-based China Production Studio*	- to view the ID.
	//
	// > The name of the production studio in the production studio list on the China Cloud-based Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b4810848-bcf9-4aef-bd4a-e6bba2d9****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The scene ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene configuration type. Valid values:
	//
	// - **Component**: component configuration.
	//
	// - **Layout**: layout configuration.
	//
	// - **All**: component and layout configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// Component
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeleteCasterSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterSceneConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterSceneConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterSceneConfigRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DeleteCasterSceneConfigRequest) GetType() *string {
	return s.Type
}

func (s *DeleteCasterSceneConfigRequest) SetCasterId(v string) *DeleteCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetOwnerId(v int64) *DeleteCasterSceneConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetRegionId(v string) *DeleteCasterSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetSceneId(v string) *DeleteCasterSceneConfigRequest {
	s.SceneId = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) SetType(v string) *DeleteCasterSceneConfigRequest {
	s.Type = &v
	return s
}

func (s *DeleteCasterSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
