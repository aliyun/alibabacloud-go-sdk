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
	// The production studio ID.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the CasterId value returned by the CreateCaster operation.
	//
	// - If you created the production studio in the ApsaraVideo Live console, navigate to **ApsaraVideo Live console*	- > **Production Studios*	- > **Cloud Production Studio*	- to view the production studio name.
	//
	// > The production studio name in the production studio list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The video-on-demand material ID.
	//
	// > Specify either this parameter or ResourceUrl.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The material type. Valid values:
	//
	// - LIVE: live streaming material. Supports live streaming materials and third-party URLs.
	//
	// - VOD: video-on-demand material. Supports video-on-demand materials and third-party URLs.
	//
	// - PIC: image material. Supports video-on-demand materials and third-party URLs.
	//
	// > Specify one of the three values or leave this parameter empty.
	//
	// example:
	//
	// VOD
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The URL of the external material.
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
