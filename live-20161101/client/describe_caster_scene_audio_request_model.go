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
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, obtain the value of CasterId from the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, view the ID on the **Production Studio*	- > **Cloud Production Studio*	- page.
	//
	// > The name of a production studio in the list on the Cloud Production Studio page is the ID of the production studio.
	//
	// This parameter is required.
	//
	// example:
	//
	// 97df6b7f-3490-47d2-ac50-88338765****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scenario.
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
