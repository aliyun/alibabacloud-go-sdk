// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterSceneAudioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterSceneAudioRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterSceneAudioRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterSceneAudioRequest
	GetRegionId() *string
	SetSceneId(v string) *DescribeCasterSceneAudioRequest
	GetSceneId() *string
}

type DescribeCasterSceneAudioRequest struct {
	// The ID of the production studio.
	//
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97df6b7f-3490-47d2-ac50-88338765****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97df6b7f-3490-47d2-ac50-88339087****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeCasterSceneAudioRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterSceneAudioRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterSceneAudioRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterSceneAudioRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterSceneAudioRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterSceneAudioRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeCasterSceneAudioRequest) SetCasterId(v string) *DescribeCasterSceneAudioRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterSceneAudioRequest) SetOwnerId(v int64) *DescribeCasterSceneAudioRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterSceneAudioRequest) SetRegionId(v string) *DescribeCasterSceneAudioRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterSceneAudioRequest) SetSceneId(v string) *DescribeCasterSceneAudioRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterSceneAudioRequest) Validate() error {
	return dara.Validate(s)
}
