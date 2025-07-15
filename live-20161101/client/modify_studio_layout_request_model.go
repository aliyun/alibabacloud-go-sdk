// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStudioLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBgImageConfig(v string) *ModifyStudioLayoutRequest
	GetBgImageConfig() *string
	SetCasterId(v string) *ModifyStudioLayoutRequest
	GetCasterId() *string
	SetCommonConfig(v string) *ModifyStudioLayoutRequest
	GetCommonConfig() *string
	SetLayerOrderConfigList(v string) *ModifyStudioLayoutRequest
	GetLayerOrderConfigList() *string
	SetLayoutId(v string) *ModifyStudioLayoutRequest
	GetLayoutId() *string
	SetLayoutName(v string) *ModifyStudioLayoutRequest
	GetLayoutName() *string
	SetMediaInputConfigList(v string) *ModifyStudioLayoutRequest
	GetMediaInputConfigList() *string
	SetOwnerId(v int64) *ModifyStudioLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyStudioLayoutRequest
	GetRegionId() *string
	SetScreenInputConfigList(v string) *ModifyStudioLayoutRequest
	GetScreenInputConfigList() *string
}

type ModifyStudioLayoutRequest struct {
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
	// >  The production studio must be a virtual studio that you create in advance.
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
	// {  "ChannelId":"RV01" }
	CommonConfig *string `json:"CommonConfig,omitempty" xml:"CommonConfig,omitempty"`
	// The layer sorting configurations. The value is a JSON string. For more information, see **layerOrderConfig**. You can sort layers of background and multimedia materials. The chroma key layer cannot be sorted. A layer that is in the front of the code is placed behind other layers in the layout.
	//
	// example:
	//
	// [ { "Type":"media", "Id":"k12kj31****" }, { "Type":"media", "Id":"k12kj31****" } ]
	LayerOrderConfigList *string `json:"LayerOrderConfigList,omitempty" xml:"LayerOrderConfigList,omitempty"`
	// The ID of the layout. If the layout was added by calling the [AddStudioLayout](https://help.aliyun.com/document_detail/2848062.html) operation, check the value of the response parameter LayoutId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// The name of the layout.
	//
	// example:
	//
	// The name of the layout.
	LayoutName *string `json:"LayoutName,omitempty" xml:"LayoutName,omitempty"`
	// The multimedia input configurations. The value is a JSON string. For more information, see **MediaInputConfig**.
	//
	// >  This parameter is optional and takes effect only if you set LayoutType to studio.
	//
	// example:
	//
	// [ { "Id":"k12kj31****", "Index":"1", "ChannelId":"RV01", "FillMode":"none", "PositionRefer":"topLeft", "WidthNormalized":"0.4", "HeightNormalized":"0.4", "PositionNormalized":"[0.1, 0.2]" }, { "Id":"k12kj31****", "Index":"2", "ImageMaterialId":"lkajsdfsa8fd89asd8****", "FillMode":"none", "PositionRefer":"topLeft", "WidthNormalized":"0.6", "HeightNormalized":"0.4", "PositionNormalized":"[0.1, 0.2]" } ]
	MediaInputConfigList *string `json:"MediaInputConfigList,omitempty" xml:"MediaInputConfigList,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The input configurations for chroma key. The value is a JSON string. For more information, see **ScreenInputConfig**.
	//
	// >  This parameter is required only if you set LayoutType to studio.
	//
	// example:
	//
	// [ { "Index":"1", "ChannelId":"RV01", "Color":"green", "PositionX":"0.1", "PositionY":"0.2", "HeightNormalized":"0.4" } ]
	ScreenInputConfigList *string `json:"ScreenInputConfigList,omitempty" xml:"ScreenInputConfigList,omitempty"`
}

func (s ModifyStudioLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStudioLayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyStudioLayoutRequest) GetBgImageConfig() *string {
	return s.BgImageConfig
}

func (s *ModifyStudioLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *ModifyStudioLayoutRequest) GetCommonConfig() *string {
	return s.CommonConfig
}

func (s *ModifyStudioLayoutRequest) GetLayerOrderConfigList() *string {
	return s.LayerOrderConfigList
}

func (s *ModifyStudioLayoutRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *ModifyStudioLayoutRequest) GetLayoutName() *string {
	return s.LayoutName
}

func (s *ModifyStudioLayoutRequest) GetMediaInputConfigList() *string {
	return s.MediaInputConfigList
}

func (s *ModifyStudioLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyStudioLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStudioLayoutRequest) GetScreenInputConfigList() *string {
	return s.ScreenInputConfigList
}

func (s *ModifyStudioLayoutRequest) SetBgImageConfig(v string) *ModifyStudioLayoutRequest {
	s.BgImageConfig = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetCasterId(v string) *ModifyStudioLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetCommonConfig(v string) *ModifyStudioLayoutRequest {
	s.CommonConfig = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetLayerOrderConfigList(v string) *ModifyStudioLayoutRequest {
	s.LayerOrderConfigList = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetLayoutId(v string) *ModifyStudioLayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetLayoutName(v string) *ModifyStudioLayoutRequest {
	s.LayoutName = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetMediaInputConfigList(v string) *ModifyStudioLayoutRequest {
	s.MediaInputConfigList = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetOwnerId(v int64) *ModifyStudioLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetRegionId(v string) *ModifyStudioLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStudioLayoutRequest) SetScreenInputConfigList(v string) *ModifyStudioLayoutRequest {
	s.ScreenInputConfigList = &v
	return s
}

func (s *ModifyStudioLayoutRequest) Validate() error {
	return dara.Validate(s)
}
