// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetShowListBackgroundRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *SetShowListBackgroundRequest
	GetCasterId() *string
	SetMaterialId(v string) *SetShowListBackgroundRequest
	GetMaterialId() *string
	SetOwnerId(v int64) *SetShowListBackgroundRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetShowListBackgroundRequest
	GetRegionId() *string
	SetResourceType(v string) *SetShowListBackgroundRequest
	GetResourceType() *string
	SetResourceUrl(v string) *SetShowListBackgroundRequest
	GetResourceUrl() *string
}

type SetShowListBackgroundRequest struct {
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
	// The ID of the material in ApsaraVideo VOD.
	//
	// >  Specify either this parameter or the ResourceUrl parameter.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type. Valid values:
	//
	// 	- LIVE: live stream. You can add a live stream from ApsaraVideo Live or by using a third-party URL.
	//
	// 	- VOD: on-demand video. You can add an on-demand video from ApsaraVideo VOD or by using a third-party URL.
	//
	// 	- PIC: image. You can add an image from ApsaraVideo VOD or by using a third-party URL.
	//
	// >  Set this parameter to one of the preceding values, or leave this parameter empty.
	//
	// example:
	//
	// VOD
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The URL of the third-party material.
	ResourceUrl *string `json:"ResourceUrl,omitempty" xml:"ResourceUrl,omitempty"`
}

func (s SetShowListBackgroundRequest) String() string {
	return dara.Prettify(s)
}

func (s SetShowListBackgroundRequest) GoString() string {
	return s.String()
}

func (s *SetShowListBackgroundRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *SetShowListBackgroundRequest) GetMaterialId() *string {
	return s.MaterialId
}

func (s *SetShowListBackgroundRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetShowListBackgroundRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetShowListBackgroundRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *SetShowListBackgroundRequest) GetResourceUrl() *string {
	return s.ResourceUrl
}

func (s *SetShowListBackgroundRequest) SetCasterId(v string) *SetShowListBackgroundRequest {
	s.CasterId = &v
	return s
}

func (s *SetShowListBackgroundRequest) SetMaterialId(v string) *SetShowListBackgroundRequest {
	s.MaterialId = &v
	return s
}

func (s *SetShowListBackgroundRequest) SetOwnerId(v int64) *SetShowListBackgroundRequest {
	s.OwnerId = &v
	return s
}

func (s *SetShowListBackgroundRequest) SetRegionId(v string) *SetShowListBackgroundRequest {
	s.RegionId = &v
	return s
}

func (s *SetShowListBackgroundRequest) SetResourceType(v string) *SetShowListBackgroundRequest {
	s.ResourceType = &v
	return s
}

func (s *SetShowListBackgroundRequest) SetResourceUrl(v string) *SetShowListBackgroundRequest {
	s.ResourceUrl = &v
	return s
}

func (s *SetShowListBackgroundRequest) Validate() error {
	return dara.Validate(s)
}
