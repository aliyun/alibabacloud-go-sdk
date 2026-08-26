// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStudioLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgImageConfig(v string) *AddStudioLayoutRequest
	GetBgImageConfig() *string
	SetCasterId(v string) *AddStudioLayoutRequest
	GetCasterId() *string
	SetCommonConfig(v string) *AddStudioLayoutRequest
	GetCommonConfig() *string
	SetLayerOrderConfigList(v string) *AddStudioLayoutRequest
	GetLayerOrderConfigList() *string
	SetLayoutName(v string) *AddStudioLayoutRequest
	GetLayoutName() *string
	SetLayoutType(v string) *AddStudioLayoutRequest
	GetLayoutType() *string
	SetMediaInputConfigList(v string) *AddStudioLayoutRequest
	GetMediaInputConfigList() *string
	SetOwnerId(v int64) *AddStudioLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddStudioLayoutRequest
	GetRegionId() *string
	SetScreenInputConfigList(v string) *AddStudioLayoutRequest
	GetScreenInputConfigList() *string
}

type AddStudioLayoutRequest struct {
	// The configuration of the background resource. The value is a JSON string. For more information, see **BgImageConfig**.
	//
	// 	Notice:
	//
	// This parameter is required only when you set LayoutType to studio.
	//
	// example:
	//
	// { "Id":"k12kj31****", "MaterialId":"f080575eb5f4427684fc0715159a****" }
	BgImageConfig *string `json:"BgImageConfig,omitempty" xml:"BgImageConfig,omitempty"`
	// The ID of the production studio.
	//
	// 	Notice:
	//
	// Create a virtual production studio in advance. You can create a production studio in the console or by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) API operation. The production studio must be a virtual production studio.
	//
	//
	//
	// - If you call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) API operation to create a production studio, use the returned CasterId value.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- page. The name of the production studio in the list is its ID.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The configuration of the common layout. The value is a JSON string. For more information, see **CommonConfig**.
	//
	// 	Notice:
	//
	// This parameter is required only when you set LayoutType to common.
	//
	// example:
	//
	// {"ChannelId":"RV01" }
	CommonConfig *string `json:"CommonConfig,omitempty" xml:"CommonConfig,omitempty"`
	// The layer order settings. The value is a JSON string. For more information, see **LayerOrderConfig**. You can sort background materials and multimedia materials. Chroma keying layers are not supported. The earlier a material appears in the list, the lower its layer.
	//
	// example:
	//
	// [ 	{ 	"Type":"media", 	"Id":"k12kj31****" 	}, 	{ 	"Type":"media", 	"Id":"k12kj31****" 	} ]
	LayerOrderConfigList *string `json:"LayerOrderConfigList,omitempty" xml:"LayerOrderConfigList,omitempty"`
	// The name of the studio layout.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test layout
	LayoutName *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	// The type of the studio layout. Valid values:
	//
	// - **common**: A common layout. If you set LayoutType to common, you must also specify CommonConfig.
	//
	// - **studio**: A studio layout. If you set LayoutType to studio, you must also specify BgImageConfig and ScreenInputConfigList. The MediaInputConfigList parameter is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// studio
	LayoutType *string `json:"LayoutType,omitempty" xml:"LayoutType,omitempty"`
	// The settings for the multimedia input resource. The value is a JSON string. For more information, see **MediaInputConfig**.
	//
	// 	Notice:
	//
	// This parameter is valid and optional only when you set LayoutType to studio.
	//
	// example:
	//
	// [ 	{ 	"Id":"k12kj31****", 	"Index":"1", 	"ChannelId":"RV01", 	"FillMode":"none", 	"PositionRefer":"topLeft", 	"WidthNormalized":"0.4", 	"HeightNormalized":"0.4", 	"PositionNormalized":"[0.1, 0.2]" 	},   { 	"Id":"k12kj31****", 	"Index":"2", 	"ImageMaterialId":"lkajsdfsa8fd89asd8****", 	"FillMode":"none", 	"PositionRefer":"topLeft", 	"WidthNormalized":"0.6", 	"HeightNormalized":"0.4", 	"PositionNormalized":"[0.1, 0.2]" 	} ]
	MediaInputConfigList *string `json:"MediaInputConfigList,omitempty" xml:"MediaInputConfigList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The settings for the chroma keying input. The value is a JSON string. For more information, see **ScreenInputConfig**.
	//
	// 	Notice:
	//
	// This parameter is required only when you set LayoutType to studio.
	//
	// example:
	//
	// [ 	{ 	"Index":"1", 	"ChannelId":"RV01", 	"Color":"green", 	"PositionX":"0.1", 	"PositionY":"0.2", 	"HeightNormalized":"0.4" 	} ]
	ScreenInputConfigList *string `json:"ScreenInputConfigList,omitempty" xml:"ScreenInputConfigList,omitempty"`
}

func (s AddStudioLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStudioLayoutRequest) GoString() string {
	return s.String()
}

func (s *AddStudioLayoutRequest) GetBgImageConfig() *string {
	return s.BgImageConfig
}

func (s *AddStudioLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *AddStudioLayoutRequest) GetCommonConfig() *string {
	return s.CommonConfig
}

func (s *AddStudioLayoutRequest) GetLayerOrderConfigList() *string {
	return s.LayerOrderConfigList
}

func (s *AddStudioLayoutRequest) GetLayoutName() *string {
	return s.LayoutName
}

func (s *AddStudioLayoutRequest) GetLayoutType() *string {
	return s.LayoutType
}

func (s *AddStudioLayoutRequest) GetMediaInputConfigList() *string {
	return s.MediaInputConfigList
}

func (s *AddStudioLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddStudioLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddStudioLayoutRequest) GetScreenInputConfigList() *string {
	return s.ScreenInputConfigList
}

func (s *AddStudioLayoutRequest) SetBgImageConfig(v string) *AddStudioLayoutRequest {
	s.BgImageConfig = &v
	return s
}

func (s *AddStudioLayoutRequest) SetCasterId(v string) *AddStudioLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *AddStudioLayoutRequest) SetCommonConfig(v string) *AddStudioLayoutRequest {
	s.CommonConfig = &v
	return s
}

func (s *AddStudioLayoutRequest) SetLayerOrderConfigList(v string) *AddStudioLayoutRequest {
	s.LayerOrderConfigList = &v
	return s
}

func (s *AddStudioLayoutRequest) SetLayoutName(v string) *AddStudioLayoutRequest {
	s.LayoutName = &v
	return s
}

func (s *AddStudioLayoutRequest) SetLayoutType(v string) *AddStudioLayoutRequest {
	s.LayoutType = &v
	return s
}

func (s *AddStudioLayoutRequest) SetMediaInputConfigList(v string) *AddStudioLayoutRequest {
	s.MediaInputConfigList = &v
	return s
}

func (s *AddStudioLayoutRequest) SetOwnerId(v int64) *AddStudioLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *AddStudioLayoutRequest) SetRegionId(v string) *AddStudioLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *AddStudioLayoutRequest) SetScreenInputConfigList(v string) *AddStudioLayoutRequest {
	s.ScreenInputConfigList = &v
	return s
}

func (s *AddStudioLayoutRequest) Validate() error {
	return dara.Validate(s)
}
