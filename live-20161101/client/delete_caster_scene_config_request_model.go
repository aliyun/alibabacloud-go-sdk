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
	// b4810848-bcf9-4aef-bd4a-e6bba2d9****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The type of the setting that you want to delete. Valid values:
	//
	// 	- **Component**: component setting
	//
	// 	- **Layout**: layout setting
	//
	// 	- **All**: component and layout settings
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
