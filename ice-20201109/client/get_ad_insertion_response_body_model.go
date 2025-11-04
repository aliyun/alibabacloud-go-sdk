// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAdInsertionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *GetAdInsertionResponseBodyConfig) *GetAdInsertionResponseBody
	GetConfig() *GetAdInsertionResponseBodyConfig
	SetRequestId(v string) *GetAdInsertionResponseBody
	GetRequestId() *string
}

type GetAdInsertionResponseBody struct {
	// The ad insertion configuration.
	Config *GetAdInsertionResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAdInsertionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAdInsertionResponseBody) GetConfig() *GetAdInsertionResponseBodyConfig {
	return s.Config
}

func (s *GetAdInsertionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAdInsertionResponseBody) SetConfig(v *GetAdInsertionResponseBodyConfig) *GetAdInsertionResponseBody {
	s.Config = v
	return s
}

func (s *GetAdInsertionResponseBody) SetRequestId(v string) *GetAdInsertionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAdInsertionResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAdInsertionResponseBodyConfig struct {
	// Indicates whether ad marker passthrough is enabled.
	//
	// example:
	//
	// ON
	AdMarkerPassthrough *string `json:"AdMarkerPassthrough,omitempty" xml:"AdMarkerPassthrough,omitempty"`
	// The URL of the ad decision server (ADS).
	//
	// example:
	//
	// http://ads.com/ad1?param1=[palyer_params.p1]
	AdsUrl *string `json:"AdsUrl,omitempty" xml:"AdsUrl,omitempty"`
	// The CDN configurations.
	CdnConfig *GetAdInsertionResponseBodyConfigCdnConfig `json:"CdnConfig,omitempty" xml:"CdnConfig,omitempty" type:"Struct"`
	// The aliases for dynamic variable replacement.
	//
	// example:
	//
	// {
	//
	//       "player_params.p1": {
	//
	//             "1": "abc"
	//
	//       }
	//
	// }
	ConfigAliases *string `json:"ConfigAliases,omitempty" xml:"ConfigAliases,omitempty"`
	// The prefix of the source URL.
	//
	// example:
	//
	// https://source.com/
	ContentUrlPrefix *string `json:"ContentUrlPrefix,omitempty" xml:"ContentUrlPrefix,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2024-06-13T08:26:09Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the configuration was last modified.
	//
	// example:
	//
	// 2024-06-13T08:26:09Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The playback endpoint prefix for accessing manifests.
	ManifestEndpointConfig *GetAdInsertionResponseBodyConfigManifestEndpointConfig `json:"ManifestEndpointConfig,omitempty" xml:"ManifestEndpointConfig,omitempty" type:"Struct"`
	// The name of the configuration.
	//
	// example:
	//
	// my_ad
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The personalization threshold.
	//
	// example:
	//
	// 5
	PersonalizationThreshold *int32 `json:"PersonalizationThreshold,omitempty" xml:"PersonalizationThreshold,omitempty"`
	// The URL of the slate ad.
	//
	// example:
	//
	// http://storage.com/slate1.mp4
	SlateAdUrl *string `json:"SlateAdUrl,omitempty" xml:"SlateAdUrl,omitempty"`
}

func (s GetAdInsertionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetAdInsertionResponseBodyConfig) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *GetAdInsertionResponseBodyConfig) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *GetAdInsertionResponseBodyConfig) GetCdnConfig() *GetAdInsertionResponseBodyConfigCdnConfig {
	return s.CdnConfig
}

func (s *GetAdInsertionResponseBodyConfig) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *GetAdInsertionResponseBodyConfig) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *GetAdInsertionResponseBodyConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetAdInsertionResponseBodyConfig) GetLastModified() *string {
	return s.LastModified
}

func (s *GetAdInsertionResponseBodyConfig) GetManifestEndpointConfig() *GetAdInsertionResponseBodyConfigManifestEndpointConfig {
	return s.ManifestEndpointConfig
}

func (s *GetAdInsertionResponseBodyConfig) GetName() *string {
	return s.Name
}

func (s *GetAdInsertionResponseBodyConfig) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *GetAdInsertionResponseBodyConfig) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *GetAdInsertionResponseBodyConfig) SetAdMarkerPassthrough(v string) *GetAdInsertionResponseBodyConfig {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetAdsUrl(v string) *GetAdInsertionResponseBodyConfig {
	s.AdsUrl = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetCdnConfig(v *GetAdInsertionResponseBodyConfigCdnConfig) *GetAdInsertionResponseBodyConfig {
	s.CdnConfig = v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetConfigAliases(v string) *GetAdInsertionResponseBodyConfig {
	s.ConfigAliases = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetContentUrlPrefix(v string) *GetAdInsertionResponseBodyConfig {
	s.ContentUrlPrefix = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetCreateTime(v string) *GetAdInsertionResponseBodyConfig {
	s.CreateTime = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetLastModified(v string) *GetAdInsertionResponseBodyConfig {
	s.LastModified = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetManifestEndpointConfig(v *GetAdInsertionResponseBodyConfigManifestEndpointConfig) *GetAdInsertionResponseBodyConfig {
	s.ManifestEndpointConfig = v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetName(v string) *GetAdInsertionResponseBodyConfig {
	s.Name = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetPersonalizationThreshold(v int32) *GetAdInsertionResponseBodyConfig {
	s.PersonalizationThreshold = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) SetSlateAdUrl(v string) *GetAdInsertionResponseBodyConfig {
	s.SlateAdUrl = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfig) Validate() error {
	if s.CdnConfig != nil {
		if err := s.CdnConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ManifestEndpointConfig != nil {
		if err := s.ManifestEndpointConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAdInsertionResponseBodyConfigCdnConfig struct {
	// The CDN prefix for accessing ad segments.
	//
	// example:
	//
	// http://cdn.com/
	AdSegmentUrlPrefix *string `json:"AdSegmentUrlPrefix,omitempty" xml:"AdSegmentUrlPrefix,omitempty"`
	// The CDN prefix for accessing content segments.
	//
	// example:
	//
	// http://cdn.com/
	ContentSegmentUrlPrefix *string `json:"ContentSegmentUrlPrefix,omitempty" xml:"ContentSegmentUrlPrefix,omitempty"`
}

func (s GetAdInsertionResponseBodyConfigCdnConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionResponseBodyConfigCdnConfig) GoString() string {
	return s.String()
}

func (s *GetAdInsertionResponseBodyConfigCdnConfig) GetAdSegmentUrlPrefix() *string {
	return s.AdSegmentUrlPrefix
}

func (s *GetAdInsertionResponseBodyConfigCdnConfig) GetContentSegmentUrlPrefix() *string {
	return s.ContentSegmentUrlPrefix
}

func (s *GetAdInsertionResponseBodyConfigCdnConfig) SetAdSegmentUrlPrefix(v string) *GetAdInsertionResponseBodyConfigCdnConfig {
	s.AdSegmentUrlPrefix = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfigCdnConfig) SetContentSegmentUrlPrefix(v string) *GetAdInsertionResponseBodyConfigCdnConfig {
	s.ContentSegmentUrlPrefix = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfigCdnConfig) Validate() error {
	return dara.Validate(s)
}

type GetAdInsertionResponseBodyConfigManifestEndpointConfig struct {
	// DASH清单播放端点前缀
	DashPrefix *string `json:"DashPrefix,omitempty" xml:"DashPrefix,omitempty"`
	// The playback endpoint prefix for accessing HLS manifests.
	HlsPrefix *string `json:"HlsPrefix,omitempty" xml:"HlsPrefix,omitempty"`
}

func (s GetAdInsertionResponseBodyConfigManifestEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAdInsertionResponseBodyConfigManifestEndpointConfig) GoString() string {
	return s.String()
}

func (s *GetAdInsertionResponseBodyConfigManifestEndpointConfig) GetDashPrefix() *string {
	return s.DashPrefix
}

func (s *GetAdInsertionResponseBodyConfigManifestEndpointConfig) GetHlsPrefix() *string {
	return s.HlsPrefix
}

func (s *GetAdInsertionResponseBodyConfigManifestEndpointConfig) SetDashPrefix(v string) *GetAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.DashPrefix = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfigManifestEndpointConfig) SetHlsPrefix(v string) *GetAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.HlsPrefix = &v
	return s
}

func (s *GetAdInsertionResponseBodyConfigManifestEndpointConfig) Validate() error {
	return dara.Validate(s)
}
