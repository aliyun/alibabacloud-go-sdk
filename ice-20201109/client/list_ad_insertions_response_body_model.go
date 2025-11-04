// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAdInsertionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListAdInsertionsResponseBodyConfigs) *ListAdInsertionsResponseBody
	GetConfigs() []*ListAdInsertionsResponseBodyConfigs
	SetMaxResults(v int32) *ListAdInsertionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAdInsertionsResponseBody
	GetNextToken() *string
	SetPageNo(v int64) *ListAdInsertionsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListAdInsertionsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAdInsertionsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListAdInsertionsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListAdInsertionsResponseBody
	GetTotalCount() *int64
}

type ListAdInsertionsResponseBody struct {
	// Array
	Configs []*ListAdInsertionsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The maximum number of entries to retrieve in a subsequent request. If this parameter is used, the pagination parameters become invalid.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used in the next request to retrieve a new page of results. If this parameter is used, the pagination parameters become invalid.
	//
	// example:
	//
	// ******8EqYpQbZ6Eh7+Zz8DxVYoQ*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sorting order of the configurations by creation time. asc: ascending. desc: descending.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAdInsertionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsResponseBody) GetConfigs() []*ListAdInsertionsResponseBodyConfigs {
	return s.Configs
}

func (s *ListAdInsertionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAdInsertionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAdInsertionsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListAdInsertionsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAdInsertionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAdInsertionsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListAdInsertionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAdInsertionsResponseBody) SetConfigs(v []*ListAdInsertionsResponseBodyConfigs) *ListAdInsertionsResponseBody {
	s.Configs = v
	return s
}

func (s *ListAdInsertionsResponseBody) SetMaxResults(v int32) *ListAdInsertionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetNextToken(v string) *ListAdInsertionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetPageNo(v int64) *ListAdInsertionsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetPageSize(v int64) *ListAdInsertionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetRequestId(v string) *ListAdInsertionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetSortBy(v string) *ListAdInsertionsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListAdInsertionsResponseBody) SetTotalCount(v int64) *ListAdInsertionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAdInsertionsResponseBody) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAdInsertionsResponseBodyConfigs struct {
	// Indicates whether ad marker passthrough is enabled.
	//
	// example:
	//
	// ON
	AdMarkerPassthrough *string `json:"AdMarkerPassthrough,omitempty" xml:"AdMarkerPassthrough,omitempty"`
	// The request URL of the ad decision server (ADS).
	//
	// example:
	//
	// http://ads.com/ad1?param1=[palyer_params.p1]
	AdsUrl *string `json:"AdsUrl,omitempty" xml:"AdsUrl,omitempty"`
	// The CDN configurations.
	CdnConfig *ListAdInsertionsResponseBodyConfigsCdnConfig `json:"CdnConfig,omitempty" xml:"CdnConfig,omitempty" type:"Struct"`
	// The player parameter variables and aliases.
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
	ManifestEndpointConfig *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig `json:"ManifestEndpointConfig,omitempty" xml:"ManifestEndpointConfig,omitempty" type:"Struct"`
	// The name of the ad insertion configuration.
	//
	// example:
	//
	// my_ad
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The personalization threshold that defines the maximum duration of underfilled time allowed in an ad break.
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

func (s ListAdInsertionsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsResponseBodyConfigs) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *ListAdInsertionsResponseBodyConfigs) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *ListAdInsertionsResponseBodyConfigs) GetCdnConfig() *ListAdInsertionsResponseBodyConfigsCdnConfig {
	return s.CdnConfig
}

func (s *ListAdInsertionsResponseBodyConfigs) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *ListAdInsertionsResponseBodyConfigs) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *ListAdInsertionsResponseBodyConfigs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAdInsertionsResponseBodyConfigs) GetLastModified() *string {
	return s.LastModified
}

func (s *ListAdInsertionsResponseBodyConfigs) GetManifestEndpointConfig() *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig {
	return s.ManifestEndpointConfig
}

func (s *ListAdInsertionsResponseBodyConfigs) GetName() *string {
	return s.Name
}

func (s *ListAdInsertionsResponseBodyConfigs) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *ListAdInsertionsResponseBodyConfigs) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *ListAdInsertionsResponseBodyConfigs) SetAdMarkerPassthrough(v string) *ListAdInsertionsResponseBodyConfigs {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetAdsUrl(v string) *ListAdInsertionsResponseBodyConfigs {
	s.AdsUrl = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetCdnConfig(v *ListAdInsertionsResponseBodyConfigsCdnConfig) *ListAdInsertionsResponseBodyConfigs {
	s.CdnConfig = v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetConfigAliases(v string) *ListAdInsertionsResponseBodyConfigs {
	s.ConfigAliases = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetContentUrlPrefix(v string) *ListAdInsertionsResponseBodyConfigs {
	s.ContentUrlPrefix = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetCreateTime(v string) *ListAdInsertionsResponseBodyConfigs {
	s.CreateTime = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetLastModified(v string) *ListAdInsertionsResponseBodyConfigs {
	s.LastModified = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetManifestEndpointConfig(v *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) *ListAdInsertionsResponseBodyConfigs {
	s.ManifestEndpointConfig = v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetName(v string) *ListAdInsertionsResponseBodyConfigs {
	s.Name = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetPersonalizationThreshold(v int32) *ListAdInsertionsResponseBodyConfigs {
	s.PersonalizationThreshold = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) SetSlateAdUrl(v string) *ListAdInsertionsResponseBodyConfigs {
	s.SlateAdUrl = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigs) Validate() error {
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

type ListAdInsertionsResponseBodyConfigsCdnConfig struct {
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

func (s ListAdInsertionsResponseBodyConfigsCdnConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsResponseBodyConfigsCdnConfig) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsResponseBodyConfigsCdnConfig) GetAdSegmentUrlPrefix() *string {
	return s.AdSegmentUrlPrefix
}

func (s *ListAdInsertionsResponseBodyConfigsCdnConfig) GetContentSegmentUrlPrefix() *string {
	return s.ContentSegmentUrlPrefix
}

func (s *ListAdInsertionsResponseBodyConfigsCdnConfig) SetAdSegmentUrlPrefix(v string) *ListAdInsertionsResponseBodyConfigsCdnConfig {
	s.AdSegmentUrlPrefix = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigsCdnConfig) SetContentSegmentUrlPrefix(v string) *ListAdInsertionsResponseBodyConfigsCdnConfig {
	s.ContentSegmentUrlPrefix = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigsCdnConfig) Validate() error {
	return dara.Validate(s)
}

type ListAdInsertionsResponseBodyConfigsManifestEndpointConfig struct {
	// DASH清单播放端点前缀
	DashPrefix *string `json:"DashPrefix,omitempty" xml:"DashPrefix,omitempty"`
	// The prefix of the playback endpoint for HLS manifests.
	HlsPrefix *string `json:"HlsPrefix,omitempty" xml:"HlsPrefix,omitempty"`
}

func (s ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) GoString() string {
	return s.String()
}

func (s *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) GetDashPrefix() *string {
	return s.DashPrefix
}

func (s *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) GetHlsPrefix() *string {
	return s.HlsPrefix
}

func (s *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) SetDashPrefix(v string) *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig {
	s.DashPrefix = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) SetHlsPrefix(v string) *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig {
	s.HlsPrefix = &v
	return s
}

func (s *ListAdInsertionsResponseBodyConfigsManifestEndpointConfig) Validate() error {
	return dara.Validate(s)
}
