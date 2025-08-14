// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAdInsertionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *AddAdInsertionResponseBodyConfig) *AddAdInsertionResponseBody
	GetConfig() *AddAdInsertionResponseBodyConfig
	SetRequestId(v string) *AddAdInsertionResponseBody
	GetRequestId() *string
}

type AddAdInsertionResponseBody struct {
	// The ad insertion configuration.
	Config *AddAdInsertionResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAdInsertionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionResponseBody) GoString() string {
	return s.String()
}

func (s *AddAdInsertionResponseBody) GetConfig() *AddAdInsertionResponseBodyConfig {
	return s.Config
}

func (s *AddAdInsertionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAdInsertionResponseBody) SetConfig(v *AddAdInsertionResponseBodyConfig) *AddAdInsertionResponseBody {
	s.Config = v
	return s
}

func (s *AddAdInsertionResponseBody) SetRequestId(v string) *AddAdInsertionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAdInsertionResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddAdInsertionResponseBodyConfig struct {
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
	CdnConfig *AddAdInsertionResponseBodyConfigCdnConfig `json:"CdnConfig,omitempty" xml:"CdnConfig,omitempty" type:"Struct"`
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
	ManifestEndpointConfig *AddAdInsertionResponseBodyConfigManifestEndpointConfig `json:"ManifestEndpointConfig,omitempty" xml:"ManifestEndpointConfig,omitempty" type:"Struct"`
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

func (s AddAdInsertionResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *AddAdInsertionResponseBodyConfig) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *AddAdInsertionResponseBodyConfig) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *AddAdInsertionResponseBodyConfig) GetCdnConfig() *AddAdInsertionResponseBodyConfigCdnConfig {
	return s.CdnConfig
}

func (s *AddAdInsertionResponseBodyConfig) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *AddAdInsertionResponseBodyConfig) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *AddAdInsertionResponseBodyConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddAdInsertionResponseBodyConfig) GetLastModified() *string {
	return s.LastModified
}

func (s *AddAdInsertionResponseBodyConfig) GetManifestEndpointConfig() *AddAdInsertionResponseBodyConfigManifestEndpointConfig {
	return s.ManifestEndpointConfig
}

func (s *AddAdInsertionResponseBodyConfig) GetName() *string {
	return s.Name
}

func (s *AddAdInsertionResponseBodyConfig) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *AddAdInsertionResponseBodyConfig) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *AddAdInsertionResponseBodyConfig) SetAdMarkerPassthrough(v string) *AddAdInsertionResponseBodyConfig {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetAdsUrl(v string) *AddAdInsertionResponseBodyConfig {
	s.AdsUrl = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetCdnConfig(v *AddAdInsertionResponseBodyConfigCdnConfig) *AddAdInsertionResponseBodyConfig {
	s.CdnConfig = v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetConfigAliases(v string) *AddAdInsertionResponseBodyConfig {
	s.ConfigAliases = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetContentUrlPrefix(v string) *AddAdInsertionResponseBodyConfig {
	s.ContentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetCreateTime(v string) *AddAdInsertionResponseBodyConfig {
	s.CreateTime = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetLastModified(v string) *AddAdInsertionResponseBodyConfig {
	s.LastModified = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetManifestEndpointConfig(v *AddAdInsertionResponseBodyConfigManifestEndpointConfig) *AddAdInsertionResponseBodyConfig {
	s.ManifestEndpointConfig = v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetName(v string) *AddAdInsertionResponseBodyConfig {
	s.Name = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetPersonalizationThreshold(v int32) *AddAdInsertionResponseBodyConfig {
	s.PersonalizationThreshold = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) SetSlateAdUrl(v string) *AddAdInsertionResponseBodyConfig {
	s.SlateAdUrl = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}

type AddAdInsertionResponseBodyConfigCdnConfig struct {
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

func (s AddAdInsertionResponseBodyConfigCdnConfig) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionResponseBodyConfigCdnConfig) GoString() string {
	return s.String()
}

func (s *AddAdInsertionResponseBodyConfigCdnConfig) GetAdSegmentUrlPrefix() *string {
	return s.AdSegmentUrlPrefix
}

func (s *AddAdInsertionResponseBodyConfigCdnConfig) GetContentSegmentUrlPrefix() *string {
	return s.ContentSegmentUrlPrefix
}

func (s *AddAdInsertionResponseBodyConfigCdnConfig) SetAdSegmentUrlPrefix(v string) *AddAdInsertionResponseBodyConfigCdnConfig {
	s.AdSegmentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfigCdnConfig) SetContentSegmentUrlPrefix(v string) *AddAdInsertionResponseBodyConfigCdnConfig {
	s.ContentSegmentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfigCdnConfig) Validate() error {
	return dara.Validate(s)
}

type AddAdInsertionResponseBodyConfigManifestEndpointConfig struct {
	// DASH清单播放端点前缀
	DashPrefix *string `json:"DashPrefix,omitempty" xml:"DashPrefix,omitempty"`
	// The prefix of the playback endpoint for HLS manifests.
	HlsPrefix *string `json:"HlsPrefix,omitempty" xml:"HlsPrefix,omitempty"`
}

func (s AddAdInsertionResponseBodyConfigManifestEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionResponseBodyConfigManifestEndpointConfig) GoString() string {
	return s.String()
}

func (s *AddAdInsertionResponseBodyConfigManifestEndpointConfig) GetDashPrefix() *string {
	return s.DashPrefix
}

func (s *AddAdInsertionResponseBodyConfigManifestEndpointConfig) GetHlsPrefix() *string {
	return s.HlsPrefix
}

func (s *AddAdInsertionResponseBodyConfigManifestEndpointConfig) SetDashPrefix(v string) *AddAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.DashPrefix = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfigManifestEndpointConfig) SetHlsPrefix(v string) *AddAdInsertionResponseBodyConfigManifestEndpointConfig {
	s.HlsPrefix = &v
	return s
}

func (s *AddAdInsertionResponseBodyConfigManifestEndpointConfig) Validate() error {
	return dara.Validate(s)
}
