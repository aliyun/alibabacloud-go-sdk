// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterScenesRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterScenesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterScenesRequest
	GetRegionId() *string
	SetSceneId(v string) *DescribeCasterScenesRequest
	GetSceneId() *string
}

type DescribeCasterScenesRequest struct {
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
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// b5f8c837-ceeb-424f-b30b-68e94e86****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DescribeCasterScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterScenesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterScenesRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterScenesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterScenesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterScenesRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DescribeCasterScenesRequest) SetCasterId(v string) *DescribeCasterScenesRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterScenesRequest) SetOwnerId(v int64) *DescribeCasterScenesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterScenesRequest) SetRegionId(v string) *DescribeCasterScenesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterScenesRequest) SetSceneId(v string) *DescribeCasterScenesRequest {
	s.SceneId = &v
	return s
}

func (s *DescribeCasterScenesRequest) Validate() error {
	return dara.Validate(s)
}
