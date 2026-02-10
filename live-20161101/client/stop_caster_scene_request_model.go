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
	// If you create a production studio through the [CreateCaster](~~69338#doc-api-live-CreateCaster~~ "Creates a production studio.") interface, check the value of the CasterId parameter in the response.
	//
	// If you create a production studio through the ApsaraVideo Live Console, log in to the console, then check the ID of the production studio through the following path:
	//
	// Production Studios > Production Studio Management
	//
	// >  The CasterId is reflected in the Name column on the Production Studio Management page.
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
	// This operation is available only when the scene is a preview scene. For more information about the scene types, see [Query the scenes of production studios](https://help.aliyun.com/document_detail/60262.html).
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
