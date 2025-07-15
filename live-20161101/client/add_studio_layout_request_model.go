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
	// The background material configurations. The value is a JSON string. For more information, see **BgImageConfig**.
	//
	// >  This parameter is required only if you set LayoutType to studio.
	//
	// example:
	//
	// { "Id":"k12kj31****", "MaterialId":"f080575eb5f4427684fc0715159a****" }
	BgImageConfig *string `json:"BgImageConfig,omitempty" xml:"BgImageConfig,omitempty"`
	// The ID of the production studio.
	//
	// >  The production studio must be a virtual studio that you create in advance. You can use the ApsaraVideo Live console or call the CreateCaster operation to create a virtual studio.
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
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The common layout configurations. The value is a JSON string. For more information, see **CommonConfig**.
	//
	// >  This parameter is required only if you set LayoutType to common.
	//
	// example:
	//
	// {"ChannelId":"RV01" }
	CommonConfig *string `json:"CommonConfig,omitempty" xml:"CommonConfig,omitempty"`
	// The layer sorting configurations. The value is a JSON string. For more information, see **layerOrderConfig**. You can sort layers of background and multimedia materials. The chroma key layer cannot be sorted. A layer that is in the front of the code is placed behind other layers in the layout.
	//
	// example:
	//
	// [ 	{ 	"Type":"media", 	"Id":"k12kj31****" 	}, 	{ 	"Type":"media", 	"Id":"k12kj31****" 	} ]
	LayerOrderConfigList *string `json:"LayerOrderConfigList,omitempty" xml:"LayerOrderConfigList,omitempty"`
	// The name of the layout.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test layout
	LayoutName *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	// The type of the layout. Valid values:
	//
	// 	- **common**: If you set this parameter to common, you must specify the CommonConfig parameter.
	//
	// 	- **studio**: If you set this parameter to studio, you must specify the BgImageConfig and ScreenInputConfigList parameters. The MediaInputConfigList parameter is optional.
	//
	// This parameter is required.
	//
	// example:
	//
	// studio
	LayoutType *string `json:"LayoutType,omitempty" xml:"LayoutType,omitempty"`
	// The multimedia input configurations. The value is a JSON string. For more information, see **MediaInputConfig**.
	//
	// >  This parameter is optional and is valid only if you set LayoutType to studio.
	//
	// example:
	//
	// [ 	{ 	"Id":"k12kj31****", 	"Index":"1", 	"ChannelId":"RV01", 	"FillMode":"none", 	"PositionRefer":"topLeft", 	"WidthNormalized":"0.4", 	"HeightNormalized":"0.4", 	"PositionNormalized":"[0.1, 0.2]" 	},   { 	"Id":"k12kj31****", 	"Index":"2", 	"ImageMaterialId":"lkajsdfsa8fd89asd8****", 	"FillMode":"none", 	"PositionRefer":"topLeft", 	"WidthNormalized":"0.6", 	"HeightNormalized":"0.4", 	"PositionNormalized":"[0.1, 0.2]" 	} ]
	MediaInputConfigList *string `json:"MediaInputConfigList,omitempty" xml:"MediaInputConfigList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The input configurations for chroma key. The value is a JSON string. For more information, see **ScreenInputConfig**.
	//
	// >  This parameter is required only if you set LayoutType to studio.
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
