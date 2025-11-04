// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdInsertionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateAdInsertionResponseBodyConfig) *UpdateAdInsertionResponseBody
	GetConfig() *UpdateAdInsertionResponseBodyConfig
	SetRequestId(v string) *UpdateAdInsertionResponseBody
	GetRequestId() *string
}

type UpdateAdInsertionResponseBody struct {
	// The ad insertion configuration.
	Config *UpdateAdInsertionResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAdInsertionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionResponseBody) GetConfig() *UpdateAdInsertionResponseBodyConfig {
	return s.Config
}

func (s *UpdateAdInsertionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAdInsertionResponseBody) SetConfig(v *UpdateAdInsertionResponseBodyConfig) *UpdateAdInsertionResponseBody {
	s.Config = v
	return s
}

func (s *UpdateAdInsertionResponseBody) SetRequestId(v string) *UpdateAdInsertionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAdInsertionResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAdInsertionResponseBodyConfig struct {
	// Indicates whether ad marker passthrough is enabled.
	//
	// example:
	//
	// ON
	AdMarkerPassthrough *string `json:"AdMarkerPassthrough,omitempty" xml:"AdMarkerPassthrough,omitempty"`
	// The request URL of ADS.
	//
	// example:
	//
	// http://ads.com/ad1?param1=[palyer_params.p1]
	AdsUrl *string `json:"AdsUrl,omitempty" xml:"AdsUrl,omitempty"`
	// The CDN configurations.
	CdnConfig *UpdateAdInsertionResponseBodyConfigCdnConfig `json:"CdnConfig,omitempty" xml:"CdnConfig,omitempty" type:"Struct"`
	// The player parameter variables and aliases.
	//
	// example:
	//
	// { "player_params.p1": { "1": "abc" } }
	ConfigAliases *string `json:"ConfigAliases,omitempty" xml:"ConfigAliases,omitempty"`
	// The URL prefix for the source content.
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
	// The playback endpoint configuration.
	ManifestEndpointConfig *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig `json:"ManifestEndpointConfig,omitempty" xml:"ManifestEndpointConfig,omitempty" type:"Struct"`
	// The name of the ad insertion configuration.
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

func (s UpdateAdInsertionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionResponseBodyConfig) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *UpdateAdInsertionResponseBodyConfig) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *UpdateAdInsertionResponseBodyConfig) GetCdnConfig() *UpdateAdInsertionResponseBodyConfigCdnConfig {
	return s.CdnConfig
}

func (s *UpdateAdInsertionResponseBodyConfig) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *UpdateAdInsertionResponseBodyConfig) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *UpdateAdInsertionResponseBodyConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateAdInsertionResponseBodyConfig) GetLastModified() *string {
	return s.LastModified
}

func (s *UpdateAdInsertionResponseBodyConfig) GetManifestEndpointConfig() *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig {
	return s.ManifestEndpointConfig
}

func (s *UpdateAdInsertionResponseBodyConfig) GetName() *string {
	return s.Name
}

func (s *UpdateAdInsertionResponseBodyConfig) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *UpdateAdInsertionResponseBodyConfig) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *UpdateAdInsertionResponseBodyConfig) SetAdMarkerPassthrough(v string) *UpdateAdInsertionResponseBodyConfig {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetAdsUrl(v string) *UpdateAdInsertionResponseBodyConfig {
	s.AdsUrl = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetCdnConfig(v *UpdateAdInsertionResponseBodyConfigCdnConfig) *UpdateAdInsertionResponseBodyConfig {
	s.CdnConfig = v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetConfigAliases(v string) *UpdateAdInsertionResponseBodyConfig {
	s.ConfigAliases = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetContentUrlPrefix(v string) *UpdateAdInsertionResponseBodyConfig {
	s.ContentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetCreateTime(v string) *UpdateAdInsertionResponseBodyConfig {
	s.CreateTime = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetLastModified(v string) *UpdateAdInsertionResponseBodyConfig {
	s.LastModified = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetManifestEndpointConfig(v *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) *UpdateAdInsertionResponseBodyConfig {
	s.ManifestEndpointConfig = v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetName(v string) *UpdateAdInsertionResponseBodyConfig {
	s.Name = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetPersonalizationThreshold(v int32) *UpdateAdInsertionResponseBodyConfig {
	s.PersonalizationThreshold = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) SetSlateAdUrl(v string) *UpdateAdInsertionResponseBodyConfig {
	s.SlateAdUrl = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfig) Validate() error {
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

type UpdateAdInsertionResponseBodyConfigCdnConfig struct {
	// The CDN prefix for ad segments.
	//
	// example:
	//
	// http://cdn.com/
	AdSegmentUrlPrefix *string `json:"AdSegmentUrlPrefix,omitempty" xml:"AdSegmentUrlPrefix,omitempty"`
	// The CDN prefix for content segments.
	//
	// example:
	//
	// http://cdn.com/
	ContentSegmentUrlPrefix *string `json:"ContentSegmentUrlPrefix,omitempty" xml:"ContentSegmentUrlPrefix,omitempty"`
}

func (s UpdateAdInsertionResponseBodyConfigCdnConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionResponseBodyConfigCdnConfig) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionResponseBodyConfigCdnConfig) GetAdSegmentUrlPrefix() *string {
	return s.AdSegmentUrlPrefix
}

func (s *UpdateAdInsertionResponseBodyConfigCdnConfig) GetContentSegmentUrlPrefix() *string {
	return s.ContentSegmentUrlPrefix
}

func (s *UpdateAdInsertionResponseBodyConfigCdnConfig) SetAdSegmentUrlPrefix(v string) *UpdateAdInsertionResponseBodyConfigCdnConfig {
	s.AdSegmentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfigCdnConfig) SetContentSegmentUrlPrefix(v string) *UpdateAdInsertionResponseBodyConfigCdnConfig {
	s.ContentSegmentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfigCdnConfig) Validate() error {
	return dara.Validate(s)
}

type UpdateAdInsertionResponseBodyConfigManifestEndpointConfig struct {
	// DASH清单播放端点前缀
	DashPrefix *string `json:"DashPrefix,omitempty" xml:"DashPrefix,omitempty"`
	// The prefix of the playback endpoint for HLS manifests.
	HlsPrefix *string `json:"HlsPrefix,omitempty" xml:"HlsPrefix,omitempty"`
}

func (s UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) GetDashPrefix() *string {
	return s.DashPrefix
}

func (s *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) GetHlsPrefix() *string {
	return s.HlsPrefix
}

func (s *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) SetDashPrefix(v string) *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.DashPrefix = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) SetHlsPrefix(v string) *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.HlsPrefix = &v
	return s
}

func (s *UpdateAdInsertionResponseBodyConfigManifestEndpointConfig) Validate() error {
	return dara.Validate(s)
}
