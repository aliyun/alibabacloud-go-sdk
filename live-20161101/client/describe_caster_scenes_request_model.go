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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId parameter value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, go to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the ID.
	//
	// - The production studio specified by CasterId must have a DomainName configured through SetCasterConfig. Otherwise, the error InvalidDomainName.NotFound is returned.
	//
	// > The name of the production studio in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80787064-1c94-4dc1-85ce-9409960a****
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
