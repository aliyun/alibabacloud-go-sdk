// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *StartCasterSceneRequest
	GetCasterId() *string
	SetOwnerId(v int64) *StartCasterSceneRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartCasterSceneRequest
	GetRegionId() *string
	SetSceneId(v string) *StartCasterSceneRequest
	GetSceneId() *string
}

type StartCasterSceneRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The production studio ID is displayed as its name on the Cloud Production Studio page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
	//
	// This parameter is valid only for PVW scenes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 242b4e2c-c30f-4442-85ba-2e3e4e3d****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StartCasterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCasterSceneRequest) GoString() string {
	return s.String()
}

func (s *StartCasterSceneRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *StartCasterSceneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartCasterSceneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartCasterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StartCasterSceneRequest) SetCasterId(v string) *StartCasterSceneRequest {
	s.CasterId = &v
	return s
}

func (s *StartCasterSceneRequest) SetOwnerId(v int64) *StartCasterSceneRequest {
	s.OwnerId = &v
	return s
}

func (s *StartCasterSceneRequest) SetRegionId(v string) *StartCasterSceneRequest {
	s.RegionId = &v
	return s
}

func (s *StartCasterSceneRequest) SetSceneId(v string) *StartCasterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StartCasterSceneRequest) Validate() error {
	return dara.Validate(s)
}
