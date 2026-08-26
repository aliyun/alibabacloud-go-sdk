// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCasterSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *StopCasterSceneRequest
	GetCasterId() *string
	SetOwnerId(v int64) *StopCasterSceneRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopCasterSceneRequest
	GetRegionId() *string
	SetSceneId(v string) *StopCasterSceneRequest
	GetSceneId() *string
}

type StopCasterSceneRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
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
	// This operation is available only for preview (PVW) scenes. For more information about scene types, see [Query production studio scenes](https://help.aliyun.com/document_detail/2848102.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 242b4e2c-c30f-4442-85ba-2e3e4e3d****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s StopCasterSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s StopCasterSceneRequest) GoString() string {
	return s.String()
}

func (s *StopCasterSceneRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *StopCasterSceneRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopCasterSceneRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopCasterSceneRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *StopCasterSceneRequest) SetCasterId(v string) *StopCasterSceneRequest {
	s.CasterId = &v
	return s
}

func (s *StopCasterSceneRequest) SetOwnerId(v int64) *StopCasterSceneRequest {
	s.OwnerId = &v
	return s
}

func (s *StopCasterSceneRequest) SetRegionId(v string) *StopCasterSceneRequest {
	s.RegionId = &v
	return s
}

func (s *StopCasterSceneRequest) SetSceneId(v string) *StopCasterSceneRequest {
	s.SceneId = &v
	return s
}

func (s *StopCasterSceneRequest) Validate() error {
	return dara.Validate(s)
}
