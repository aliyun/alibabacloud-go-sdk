// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVodPackagingAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsset(v *GetVodPackagingAssetResponseBodyAsset) *GetVodPackagingAssetResponseBody
	GetAsset() *GetVodPackagingAssetResponseBodyAsset
	SetRequestId(v string) *GetVodPackagingAssetResponseBody
	GetRequestId() *string
}

type GetVodPackagingAssetResponseBody struct {
	// The information about the asset.
	Asset *GetVodPackagingAssetResponseBodyAsset `json:"Asset,omitempty" xml:"Asset,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0622C702-41BE-467E-AF2E-883D4517962E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVodPackagingAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetResponseBody) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetResponseBody) GetAsset() *GetVodPackagingAssetResponseBodyAsset {
	return s.Asset
}

func (s *GetVodPackagingAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVodPackagingAssetResponseBody) SetAsset(v *GetVodPackagingAssetResponseBodyAsset) *GetVodPackagingAssetResponseBody {
	s.Asset = v
	return s
}

func (s *GetVodPackagingAssetResponseBody) SetRequestId(v string) *GetVodPackagingAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVodPackagingAssetResponseBody) Validate() error {
	if s.Asset != nil {
		if err := s.Asset.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVodPackagingAssetResponseBodyAsset struct {
	// The name of the asset.
	//
	// example:
	//
	// 30min_movie
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The content ID in the DRM system. The maximum length is 256 characters. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// example:
	//
	// movie
	ContentId *string `json:"ContentId,omitempty" xml:"ContentId,omitempty"`
	// The time when the asset was created. It follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-11-21T06:45:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The egress endpoints, each corresponding to a packaging configuration.
	EgressEndpoints []*GetVodPackagingAssetResponseBodyAssetEgressEndpoints `json:"EgressEndpoints,omitempty" xml:"EgressEndpoints,omitempty" type:"Repeated"`
	// The name of the packaging group.
	//
	// example:
	//
	// vod_hls
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The asset input configurations.
	Input *GetVodPackagingAssetResponseBodyAssetInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s GetVodPackagingAssetResponseBodyAsset) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetResponseBodyAsset) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetAssetName() *string {
	return s.AssetName
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetContentId() *string {
	return s.ContentId
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetEgressEndpoints() []*GetVodPackagingAssetResponseBodyAssetEgressEndpoints {
	return s.EgressEndpoints
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetGroupName() *string {
	return s.GroupName
}

func (s *GetVodPackagingAssetResponseBodyAsset) GetInput() *GetVodPackagingAssetResponseBodyAssetInput {
	return s.Input
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetAssetName(v string) *GetVodPackagingAssetResponseBodyAsset {
	s.AssetName = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetContentId(v string) *GetVodPackagingAssetResponseBodyAsset {
	s.ContentId = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetCreateTime(v string) *GetVodPackagingAssetResponseBodyAsset {
	s.CreateTime = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetEgressEndpoints(v []*GetVodPackagingAssetResponseBodyAssetEgressEndpoints) *GetVodPackagingAssetResponseBodyAsset {
	s.EgressEndpoints = v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetGroupName(v string) *GetVodPackagingAssetResponseBodyAsset {
	s.GroupName = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) SetInput(v *GetVodPackagingAssetResponseBodyAssetInput) *GetVodPackagingAssetResponseBodyAsset {
	s.Input = v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAsset) Validate() error {
	if s.EgressEndpoints != nil {
		for _, item := range s.EgressEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVodPackagingAssetResponseBodyAssetEgressEndpoints struct {
	// The name of the packaging configuration.
	//
	// example:
	//
	// hls_3s
	ConfigurationName *string `json:"ConfigurationName,omitempty" xml:"ConfigurationName,omitempty"`
	// The asset status. Valid values:
	//
	// 	- Queuing: The asset is waiting for packaging.
	//
	// 	- Playable: The asset is packaged and playable.
	//
	// 	- Failed: The asset fails to be packaged.
	//
	// example:
	//
	// Playable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The playback URL. If the asset fails to be packaged, no playback URL is returned.
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetVodPackagingAssetResponseBodyAssetEgressEndpoints) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetResponseBodyAssetEgressEndpoints) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) GetConfigurationName() *string {
	return s.ConfigurationName
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) GetStatus() *string {
	return s.Status
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) GetUrl() *string {
	return s.Url
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) SetConfigurationName(v string) *GetVodPackagingAssetResponseBodyAssetEgressEndpoints {
	s.ConfigurationName = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) SetStatus(v string) *GetVodPackagingAssetResponseBodyAssetEgressEndpoints {
	s.Status = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) SetUrl(v string) *GetVodPackagingAssetResponseBodyAssetEgressEndpoints {
	s.Url = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAssetEgressEndpoints) Validate() error {
	return dara.Validate(s)
}

type GetVodPackagingAssetResponseBodyAssetInput struct {
	// The URL of the media file. Only M3U8 files stored in OSS are supported.
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The input type. Only Object Storage Service (OSS) is supported.
	//
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetVodPackagingAssetResponseBodyAssetInput) String() string {
	return dara.Prettify(s)
}

func (s GetVodPackagingAssetResponseBodyAssetInput) GoString() string {
	return s.String()
}

func (s *GetVodPackagingAssetResponseBodyAssetInput) GetMedia() *string {
	return s.Media
}

func (s *GetVodPackagingAssetResponseBodyAssetInput) GetType() *string {
	return s.Type
}

func (s *GetVodPackagingAssetResponseBodyAssetInput) SetMedia(v string) *GetVodPackagingAssetResponseBodyAssetInput {
	s.Media = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAssetInput) SetType(v string) *GetVodPackagingAssetResponseBodyAssetInput {
	s.Type = &v
	return s
}

func (s *GetVodPackagingAssetResponseBodyAssetInput) Validate() error {
	return dara.Validate(s)
}
