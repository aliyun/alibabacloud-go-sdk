// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCasterSceneConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *CopyCasterSceneConfigRequest
	GetCasterId() *string
	SetFromSceneId(v string) *CopyCasterSceneConfigRequest
	GetFromSceneId() *string
	SetOwnerId(v int64) *CopyCasterSceneConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CopyCasterSceneConfigRequest
	GetRegionId() *string
	SetToSceneId(v string) *CopyCasterSceneConfigRequest
	GetToSceneId() *string
}

type CopyCasterSceneConfigRequest struct {
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
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the source scene, which must be a PVW scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// f1a361f4-bee3-436d-ae6e-d38e6943****
	FromSceneId *string `json:"FromSceneId,omitempty" xml:"FromSceneId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the destination scene, which must be a PGM scene.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ToSceneId *string `json:"ToSceneId,omitempty" xml:"ToSceneId,omitempty"`
}

func (s CopyCasterSceneConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyCasterSceneConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyCasterSceneConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *CopyCasterSceneConfigRequest) GetFromSceneId() *string {
	return s.FromSceneId
}

func (s *CopyCasterSceneConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CopyCasterSceneConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyCasterSceneConfigRequest) GetToSceneId() *string {
	return s.ToSceneId
}

func (s *CopyCasterSceneConfigRequest) SetCasterId(v string) *CopyCasterSceneConfigRequest {
	s.CasterId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetFromSceneId(v string) *CopyCasterSceneConfigRequest {
	s.FromSceneId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetOwnerId(v int64) *CopyCasterSceneConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetRegionId(v string) *CopyCasterSceneConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) SetToSceneId(v string) *CopyCasterSceneConfigRequest {
	s.ToSceneId = &v
	return s
}

func (s *CopyCasterSceneConfigRequest) Validate() error {
	return dara.Validate(s)
}
