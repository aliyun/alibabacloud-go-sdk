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
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
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
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
	//
	// This parameter takes effect only if the scene is a PVW scene.
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
